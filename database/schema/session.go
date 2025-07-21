package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("refresh_token").
			NotEmpty().
			Immutable().
			Unique(),
		field.String("ip_address").
			Optional(),
		field.String("user_agent").
			Optional(),
		field.String("device_id").
			Optional(),
		field.Uint64("number_of_uses").
			Default(0).
			Max(10000000).
			Min(0),
		field.Time("expires_at").
			Immutable().
			Default(func() time.Time {
				return time.Now().Add(30 * 24 * time.Hour)
			}),
		// NotEmpty(),
		field.Time("last_used_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
		field.Time("revoked_at").
			Nillable().
			Optional(),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("sessions").
			Field("user_id").
			Required().
			Unique(),
	}
}
