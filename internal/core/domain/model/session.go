package domain

import (
	"fmt"
	"time"

	domain_builder "github.com/alfariiizi/vandor/internal/core/domain/builder"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
)

type Session struct {
	*db.Session
	client *db.Client
}

func NewSessionDomain(client *db.Client) domain_builder.Domain[*db.Session, *Session] {
	return domain_builder.NewDomain(
		func(e *db.Session, c *db.Client) *Session {
			return &Session{
				Session: e,
				client:  c,
			}
		}, client)
}

// TODO: Add your domain methods here
// Example methods:

func (session *Session) String() string {
	return fmt.Sprintf("Session{ID: %d}", session.ID)
}

func (session *Session) IsAvailable() bool {
	return time.Now().Before(session.ExpiresAt) && session.RevokedAt != nil
}

// Add more business logic methods as needed
