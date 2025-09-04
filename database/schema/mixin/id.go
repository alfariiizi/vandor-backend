package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Optional().
			Immutable().
			Default(uuid.New).
			Comment("Unique identifier for the entity, typically a UUID."),
	}
}

func (IDMixin) Edges() []ent.Edge {
	return []ent.Edge{}
}
