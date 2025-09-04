package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/alfariiizi/vandor/database/schema/mixin"
	"github.com/google/uuid"
	"github.com/lrstanley/entrest"
	"github.com/ogen-go/ogen"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Annotations (optional): name the table explicitly.
func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notifications"},
	}
}

func (Notification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Notification) Fields() []ent.Field {
	return []ent.Field{
		// ID as UUID to match frontend string ids.
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		// Who the notification is for (user id as string to avoid tight coupling).
		field.UUID("user_id", uuid.UUID{}),

		field.String("title").
			NotEmpty().
			MaxLen(200),

		// Long text message
		field.String("message").
			NotEmpty().
			SchemaType(map[string]string{
				"mysql":    "TEXT",
				"postgres": "TEXT",
			}),

		field.Enum("type").
			Values("info", "success", "warning", "error").
			Default("info"),

		field.Enum("priority").
			Values("low", "normal", "high", "urgent").
			Default("normal"),

		field.Enum("channel").
			Values("in_app", "email", "push", "sms", "webhook").
			Default("in_app"),

		field.Bool("read").
			Default(false),

		field.Bool("archived").
			Default(false),

		field.Bool("sticky").
			Default(false),

		// Optional URLs / actions
		field.String("link").
			Optional().
			Nillable().
			MaxLen(1000),

		field.String("action").
			Optional().
			Nillable().
			MaxLen(100),

		// Optional reference to related resource (e.g., order, ticket)
		field.String("resource_type").
			Optional().
			Nillable().
			MaxLen(100),

		field.String("resource_id").
			Optional().
			Nillable().
			MaxLen(100),

		// Aggregation and idempotency
		field.String("group_key").
			Optional().
			Nillable().
			MaxLen(100),

		field.String("dedupe_key").
			Optional().
			Nillable().
			Unique().
			MaxLen(200),

		// Delivery & lifecycle times
		field.Time("delivered_at").
			Optional().
			Nillable(),
		field.Time("read_at").
			Optional().
			Nillable(),
		field.Time("expires_at").
			Optional().
			Nillable(),

		// Flexible metadata payload
		field.JSON("meta", map[string]any{}).
			Optional().
			Annotations(
				entrest.WithSchema(ogen.NewSchema()),
			),
	}
}

func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		// If you have a User schema with primary key string/uuid,
		// you can later add an edge like this (uncomment and adjust):
		edge.From("user", User.Type).
			Ref("notifications").
			Field("user_id").
			Unique().
			Required(),
	}
}

func (Notification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "read"),
		index.Fields("user_id", "archived"),
		index.Fields("group_key"),
		index.Fields("created_at"),
		// Unique index for idempotent creates via dedupe_key.
		index.Fields("dedupe_key").Unique(),
	}
}
