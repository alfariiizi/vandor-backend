package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lrstanley/entrest"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Annotations(
				entrest.WithExample("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			),
		field.String("email").
			NotEmpty().
			Unique().
			Annotations(
				entrest.WithExample("john@mail.com"),
				entrest.WithSortable(true),
				entrest.WithFilter(entrest.FilterGroupEqual|entrest.FilterGroupArray),
			),
		field.String("first_name"),
		field.String("last_name"),
		field.String("password_hash").
			NotEmpty().
			Sensitive(),
		field.Enum("role").
			Values("USER", "ADMIN", "SUPERADMIN"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
		edge.To("sessions", Session.Type),
		edge.To("notifications", Notification.Type),
	}
}
