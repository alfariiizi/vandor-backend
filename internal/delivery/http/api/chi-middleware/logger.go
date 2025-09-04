package chimiddleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

// RegisterZerolog registers a production-ready zerolog stack for a chi.Router.
//
//   - logger: base zerolog.Logger (should already have Timestamp() if desired).
//   - r: the chi router to attach middlewares to.
//   - pathPrefix: only requests with path equal to pathPrefix or starting with pathPrefix+"/" will be considered for logging.
//     e.g. "/api" will match "/api", "/api/foo", "/api/admin/..."
//   - excludePaths: exact path strings to exclude from logging (e.g. "/api/admin/docs").
func RegisterZerolog(r chi.Router, logger zerolog.Logger, pathPrefix string, excludePaths []string) {
	// Basic recovery and request ID
	r.Use(chimw.Recoverer) // recover panics
	r.Use(chimw.RequestID) // chi RequestID middleware (adds X-Request-Id to context/response)

	// Attach base logger to request context
	r.Use(hlog.NewHandler(logger))

	// Add common useful request-level fields into the logger context
	// Note: the second arg is the field name that will be added into the log context.
	r.Use(hlog.RequestIDHandler("X-Request-Id", "request_id"))
	r.Use(hlog.RemoteAddrHandler("remote"))
	r.Use(hlog.UserAgentHandler("user_agent"))
	r.Use(hlog.MethodHandler("method"))
	r.Use(hlog.URLHandler("url"))

	// Access handler logs after response is written. We apply filtering here.
	r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		path := r.URL.Path
		if !shouldLogPath(path, pathPrefix, excludePaths) {
			return
		}

		log := hlog.FromRequest(r)

		// Attach some fields that AccessHandler doesn't provide
		event := log.With().
			Int("status", status).
			Int("bytes", size).
			Dur("duration", duration).
			Logger()

		// Choose level by status code
		switch {
		case status >= 500:
			event.Error().Msg("request completed")
		case status >= 400:
			event.Warn().Msg("request completed")
		default:
			event.Info().Msg("request completed")
		}
	}))
}

// shouldLogPath returns true if path matches the prefix rule and is not present in excludes.
// pathPrefix matching logic:
//   - path == pathPrefix  OR  strings.HasPrefix(path, pathPrefix + "/")
func shouldLogPath(path, pathPrefix string, excludePaths []string) bool {
	// if no prefix configured, assume log everything
	if pathPrefix == "" {
		// still allow excludePaths to stop some endpoints
		for _, e := range excludePaths {
			if path == e {
				return false
			}
		}
		return true
	}

	// Match prefix: either equal or starts with prefix + "/"
	if !(path == pathPrefix || strings.HasPrefix(path, pathPrefix+"/")) {
		return false
	}

	// explicit exclusions (exact match)
	for _, e := range excludePaths {
		if path == e {
			return false
		}
	}

	return true
}
