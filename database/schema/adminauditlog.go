package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AdminAuditLog holds the schema definition for the AdminAuditLog entity.
type AdminAuditLog struct {
	ent.Schema
}

// Fields of the AdminAuditLog.
func (AdminAuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("user_email"),
		field.String("operation"),
		field.String("model"),
		field.JSON("args", map[string]any{}),
		field.JSON("result", map[string]any{}).Optional(),
		field.Text("query").Optional(),
		field.Text("params").Optional(),
		field.String("source").Default("admin-panel"),
		field.Int("duration_ms").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the AdminAuditLog.
func (AdminAuditLog) Edges() []ent.Edge {
	return nil
}
