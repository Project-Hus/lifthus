package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProgramImage holds the schema definition for the ProgramImage entity.
type ProgramImage struct {
	ent.Schema
}

// Fields of the ProgramImage.
func (ProgramImage) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.Uint("order").Immutable(),
		field.Uint64("program_version_id"),
		field.Uint64("image_id"),
	}
}

// Edges of the ProgramImage.
func (ProgramImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("program_version", ProgramVersion.Type).
			Required().
			Unique().
			Field("program_version_id"),
		edge.To("image", Image.Type).
			Required().
			Unique().
			Field("image_id"),
	}
}
