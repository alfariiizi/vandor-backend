// Package sse provides a production-ready, lightweight Server-Sent Events manager
// suitable for integration into a Vandor-style Go backend. It supports:
//  - topic-based subscriptions (e.g. per-user topics like "user:123")
//  - non-blocking broadcasting with configurable backpressure handling (drops)
//  - heartbeats to keep connections alive
//  - graceful shutdown
//  - convenience helpers for common patterns
//
// Usage (high level):
//  m := sse.NewManager(30 * time.Second)
//  http.Handle("/events", m) // m implements http.Handler
//
//  // publish message to all clients:
//  _ = m.Publish(sse.Event{Event: "ping", Data: "hello"})
//
//  // publish to a user-specific topic:
//  _ = m.PublishToTopic("user:42", sse.Event{Event: "notification", Data: payload})
//
// Important production notes:
//  - In fronting proxies (nginx) disable buffering for the route (proxy_buffering off)
//    and increase proxy_read_timeout. Also set header X-Accel-Buffering: no to
//    hint certain proxies to not buffer.
//  - SSE works best with HTTP/1.1. Some HTTP/2 proxies may buffer; validate your infra.
//  - For very high throughput, consider a pub/sub (Redis, NATS) where app nodes push
//    events into the Manager (this Manager is in-memory for a single process/node).

package sse

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alfariiizi/vandor/internal/pkg/logger"
)

// Event is a single SSE payload.
// Data can be any JSON-marshalable value or a string.
type Event struct {
	ID      string `json:"id,omitempty"`
	Event   string `json:"event,omitempty"`
	Data    any    `json:"data,omitempty"`
	Retry   int    `json:"retry,omitempty"`   // milliseconds
	Comment string `json:"comment,omitempty"` // comment line (prefixed with ':' )
}

// internal client representation
type client struct {
	id     string
	topics map[string]struct{}
	send   chan Event
}

// Manager orchestrates clients and broadcasts events.
// It's safe for concurrent use.
type Manager struct {
	mu               sync.RWMutex
	clients          map[*client]struct{}
	topics           map[string]map[*client]struct{}
	register         chan *client
	unregister       chan *client
	broadcast        chan broadcastMessage
	heartbeat        time.Duration
	shutdown         chan struct{}
	wg               sync.WaitGroup
	closed           uint32 // atomic flag
	clientCounter    uint64 // atomic counter for client IDs
	dropCount        uint64 // how many messages were dropped due to backpressure
	broadcastBufSize int
	clientBufSize    int
}

type broadcastMessage struct {
	event  Event
	topics []string
}

// Default buffer sizes. Tune per app.
const (
	defaultBroadcastBuf = 256
	defaultClientBuf    = 64
)

// NewManager returns a configured Manager. heartbeat controls how often a
// lightweight heartbeat (SSE comment) is sent to keep idle connections alive.
func NewManager(heartbeat time.Duration) *Manager {
	m := &Manager{
		clients:          make(map[*client]struct{}),
		topics:           make(map[string]map[*client]struct{}),
		register:         make(chan *client),
		unregister:       make(chan *client),
		broadcast:        make(chan broadcastMessage, defaultBroadcastBuf),
		heartbeat:        heartbeat,
		shutdown:         make(chan struct{}),
		broadcastBufSize: defaultBroadcastBuf,
		clientBufSize:    defaultClientBuf,
	}
	m.wg.Add(1)
	go m.run()
	return m
}

// run is the manager's main loop.
func (m *Manager) run() {
	defer m.wg.Done()
	ticker := time.NewTicker(m.heartbeat)
	defer ticker.Stop()

	for {
		select {
		case c := <-m.register:
			m.addClient(c)
		case c := <-m.unregister:
			m.removeClient(c)
		case msg := <-m.broadcast:
			m.dispatch(msg)
		case <-ticker.C:
			// heartbeat as a comment to keep connections and proxies alive
			m.dispatch(broadcastMessage{event: Event{Comment: "heartbeat"}})
		case <-m.shutdown:
			// on shutdown we stop accepting new messages and cleanup
			m.mu.Lock()
			for c := range m.clients {
				// safe close — notify per-client writer
				close(c.send)
				delete(m.clients, c)
			}
			m.topics = make(map[string]map[*client]struct{})
			m.mu.Unlock()
			return
		}
	}
}

// addClient registers a client into the maps.
func (m *Manager) addClient(c *client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[c] = struct{}{}
	for t := range c.topics {
		subs := m.topics[t]
		if subs == nil {
			subs = make(map[*client]struct{})
			m.topics[t] = subs
		}
		subs[c] = struct{}{}
	}
	log.Printf("sse: client registered id=%s topics=%v", c.id, mapKeys(c.topics))
}

// removeClient removes client and closes its channel.
func (m *Manager) removeClient(c *client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.clients[c]; !ok {
		return
	}
	delete(m.clients, c)
	for t := range c.topics {
		if subs, ok := m.topics[t]; ok {
			delete(subs, c)
			if len(subs) == 0 {
				delete(m.topics, t)
			}
		}
	}
	// close send channel to notify handler to exit
	close(c.send)
	log.Printf("sse: client removed id=%s", c.id)
}

// dispatch sends an event to interested clients. Non-blocking: when a client
// channel is full the event is dropped for that client (prevents slow clients
// from blocking the system). We count drops in dropCount.
func (m *Manager) dispatch(msg broadcastMessage) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(msg.topics) == 0 {
		// broadcast to all
		for c := range m.clients {
			m.sendNonBlocking(c, msg.event)
		}
		return
	}

	// send to union of clients subscribed to any requested topic
	seen := make(map[*client]struct{})
	for _, t := range msg.topics {
		if subs, ok := m.topics[t]; ok {
			for c := range subs {
				if _, s := seen[c]; !s {
					seen[c] = struct{}{}
					m.sendNonBlocking(c, msg.event)
				}
			}
		}
	}
}

// non-blocking send. If client's buffer is full we drop and increment drop count.
func (m *Manager) sendNonBlocking(c *client, e Event) {
	select {
	case c.send <- e:
		// ok
	default:
		atomic.AddUint64(&m.dropCount, 1)
		// optional: if too many drops for a client, force unregister (not implemented)
	}
}

// Publish publishes an event to all clients (no topics). Returns error if manager
// is closed or broadcast buffer is full.
func (m *Manager) Publish(e Event) error {
	if atomic.LoadUint32(&m.closed) == 1 {
		return errors.New("sse manager closed")
	}
	msg := broadcastMessage{event: e}
	select {
	case m.broadcast <- msg:
		return nil
	default:
		atomic.AddUint64(&m.dropCount, 1)
		return errors.New("broadcast queue full")
	}
}

// PublishToTopic is a convenience method: publish to a single topic string.
func (m *Manager) PublishToTopic(topic string, e Event) error {
	if topic == "" {
		return m.Publish(e)
	}
	if atomic.LoadUint32(&m.closed) == 1 {
		return errors.New("sse manager closed")
	}
	msg := broadcastMessage{event: e, topics: []string{topic}}
	select {
	case m.broadcast <- msg:
		return nil
	default:
		atomic.AddUint64(&m.dropCount, 1)
		return errors.New("broadcast queue full")
	}
}

// PublishToUser is a helper for user-scoped events. Standard topic format is
// "user:<id>". Up to you to adopt the scheme.
func (m *Manager) PublishToUser(userID string, e Event) error {
	if strings.TrimSpace(userID) == "" {
		return errors.New("empty user id")
	}
	topic := "user:" + userID
	return m.PublishToTopic(topic, e)
}

// Close gracefully stops the manager and closes client channels. After Close
// Publish will return an error.
func (m *Manager) Close() {
	if !atomic.CompareAndSwapUint32(&m.closed, 0, 1) {
		return
	}
	close(m.shutdown)
	// wait until run returns
	m.wg.Wait()
	log.Printf("sse: manager closed (drops=%d)", atomic.LoadUint64(&m.dropCount))
}

// ServeHTTP implements http.Handler. It registers the new client and streams
// events until the request context is done or the client is unregistered.
// Query parameters:
//   - topics=comma,separated,list -> subscribe to a set of topics
//   - topic=single -> alternative (multiple ?topic= may be supported)
//   - last_event_id -> will be available in r.Header.Get("Last-Event-ID") if the
//     client reconnects; you may use it if you persist IDs and want to re-send missed events.
func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	log.Printf("sse: new client connection: %T", w)

	// required headers for SSE
	header := w.Header()
	header.Set("Content-Type", "text/event-stream")
	header.Set("Cache-Control", "no-cache")
	header.Set("Connection", "keep-alive")
	// hint to some reverse proxies
	header.Set("X-Accel-Buffering", "no")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	// parse topics from query param
	topicsMap := parseTopicsFromRequest(r)

	// create client and register
	id := m.newClientID()
	c := &client{id: id, topics: topicsMap, send: make(chan Event, m.clientBufSize)}

	// register; manager will store client and start sending messages to c.send
	m.register <- c
	defer func() {
		// ensure unregister when handler returns
		m.unregister <- c
	}()

	// Optional: send a connected comment
	fmt.Fprintf(w, ": connected\n\n")
	flusher.Flush()

	ctx := r.Context()

	// main write loop: writes events as they come; if client.send is closed we exit
	for {
		select {
		case <-ctx.Done():
			// client disconnected
			return
		case ev, ok := <-c.send:
			if !ok {
				// manager closed or client removed
				return
			}
			if err := writeEvent(w, ev); err != nil {
				// error writing -> unregister and return
				log.Printf("sse: error writing to client id=%s error=%v", c.id, err)
				return
			}
			flusher.Flush()
		}
	}
}

// parseTopicsFromRequest supports query param 'topics' (comma separated) and
// multiple 'topic' params. Return a set map[string]struct{} for quick lookup.
func parseTopicsFromRequest(r *http.Request) map[string]struct{} {
	q := r.URL.Query()
	out := make(map[string]struct{})
	if v := q.Get("topics"); v != "" {
		parts := strings.Split(v, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p != "" {
				out[p] = struct{}{}
			}
		}
	}
	if vals, ok := q["topic"]; ok {
		for _, v := range vals {
			v = strings.TrimSpace(v)
			if v != "" {
				out[v] = struct{}{}
			}
		}
	}
	// convenience: if user id is present in context under "user_id" put them on a user topic
	if uid, ok := r.Context().Value("user_id").(string); ok && uid != "" {
		out["user:"+uid] = struct{}{}
	}
	return out
}

// writeEvent writes a single Event to ResponseWriter in proper SSE format.
func writeEvent(w http.ResponseWriter, ev Event) error {
	// comment first
	if ev.Comment != "" {
		for _, line := range splitLines(ev.Comment) {
			if _, err := fmt.Fprintf(w, ": %s\n", line); err != nil {
				return err
			}
		}
		// comments are followed by a blank line
		if _, err := fmt.Fprint(w, "\n"); err != nil {
			return err
		}
	}

	if ev.Retry > 0 {
		if _, err := fmt.Fprintf(w, "retry: %d\n", ev.Retry); err != nil {
			return err
		}
	}

	if ev.ID != "" {
		if _, err := fmt.Fprintf(w, "id: %s\n", ev.ID); err != nil {
			return err
		}
	}
	if ev.Event != "" {
		if _, err := fmt.Fprintf(w, "event: %s\n", ev.Event); err != nil {
			return err
		}
	}

	// data
	if ev.Data != nil {
		switch d := ev.Data.(type) {
		case string:
			for _, line := range splitLines(d) {
				if _, err := fmt.Fprintf(w, "data: %s\n", line); err != nil {
					return err
				}
			}
		default:
			b, err := json.Marshal(d)
			if err != nil {
				return err
			}
			for _, line := range splitLines(string(b)) {
				if _, err := fmt.Fprintf(w, "data: %s\n", line); err != nil {
					return err
				}
			}
		}
		// blank line finishes the event
		if _, err := fmt.Fprint(w, "\n"); err != nil {
			return err
		}
	} else {
		// if there's no data but there was an id/event/comment, end event with \n
		if ev.ID != "" || ev.Event != "" || ev.Comment != "" || ev.Retry > 0 {
			if _, err := fmt.Fprint(w, "\n"); err != nil {
				return err
			}
		}
	}
	return nil
}

// splitLines splits text into lines without losing empty lines (preserve semantics)
func splitLines(s string) []string {
	if s == "" {
		return []string{""}
	}
	return strings.Split(s, "\n")
}

// newClientID returns a simple sequential ID. Good enough for debugging.
func (m *Manager) newClientID() string {
	n := atomic.AddUint64(&m.clientCounter, 1)
	return fmt.Sprintf("c-%d", n)
}

// mapKeys returns a slice of keys for logging helpers
func mapKeys(m map[string]struct{}) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// Stats exposes minimal internal metrics you may want to export to Prometheus.
func (m *Manager) Stats() (clients int, topics int, drops uint64) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.clients), len(m.topics), atomic.LoadUint64(&m.dropCount)
}

// --- Convenience helpers & docs ---
// Example: integrating in your http routes (net/http):
//  sseManager := sse.NewManager(30 * time.Second)
//  mux.Handle("/events", sseManager)
//  // publish:
//  sseManager.PublishToUser("42", sse.Event{Event: "new_order", Data: orderPayload})
//
// If you're using fx (go.uber.org/fx) you can add a lifecycle hook in your app
// module to close the manager on shutdown (call sseManager.Close()).
//
// If you want to broadcast across multiple nodes use a small adapter layer
// that receives messages from Redis/NSQ/NATS and calls m.Publish or
// m.PublishToTopic — the Manager itself remains in-memory for a single process.

// Optional helper: get user id from request context. This implementation
// expects middleware to set r.Context().Value("user_id") to a string.
func GetUserIDFromRequest(r *http.Request) (string, bool) {
	if v := r.Context().Value("user_id"); v != nil {
		if s, ok := v.(string); ok {
			return s, true
		}
	}
	// fallback to query param (not recommended for production auth)
	if s := r.URL.Query().Get("user_id"); s != "" {
		return s, true
	}
	return "", false
}
