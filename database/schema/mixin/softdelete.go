package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type SoftDeleteMixin struct {
	mixin.Schema
}

func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional().
			Nillable().
			Comment("Timestamp when the record was soft deleted. Null if not deleted."),
	}
}

func (SoftDeleteMixin) Edges() []ent.Edge {
	return []ent.Edge{}
}
