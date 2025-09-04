// ent/schema/product.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Product is a compact schema designed for light-weight scrapers.
type Product struct {
	ent.Schema
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{mixin.Time{}}
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		// product core fields
		field.String("title").Optional().Comment("nama_produk / product title"),
		field.String("short_description").Optional(),
		field.String("long_description").Optional(),

		// price: store as integer in smallest currency unit.
		// For IDR you can treat this as whole IDR units (e.g. 107000).
		field.Int64("price").Optional().Comment("price in smallest currency unit (IDR -> 107000)"),
		field.String("currency").Optional(),

		// inventory & shipments
		field.Int("stock").Optional(),
		field.Int("weight_grams").Optional().Comment("weight in grams, convert if source uses different unit"),
		field.Int("package_length_mm").Optional(),
		field.Int("package_width_mm").Optional(),
		field.Int("package_height_mm").Optional(),

		field.UUID("user_id", uuid.UUID{}),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("products").
			Field("user_id").
			Unique().
			Required(),
	}
}
