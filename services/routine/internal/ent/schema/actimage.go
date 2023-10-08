package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ActImage holds the schema definition for the ActImage entity.
type ActImage struct {
	ent.Schema
}

// Fields of the ActImage.
func (ActImage) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.Uint("order"),
		field.Uint64("act_version_id"),
		field.Uint64("image_id"),
	}
}

// Edges of the ActImage.
func (ActImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("act_version", ActVersion.Type).
			Required().
			Unique().
			Field("act_version_id"),
		edge.To("image", Image.Type).
			Required().
			Unique().
			Field("image_id"),
	}
}
