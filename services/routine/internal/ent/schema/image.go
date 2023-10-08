package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.String("key").NotEmpty().Unique(),
		field.String("src").NotEmpty().Unique(),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("act_versions", ActVersion.Type).Ref("images").Through("act_images", ActImage.Type),
		edge.From("program_versions", ProgramVersion.Type).Ref("images").Through("program_images", ProgramImage.Type),
	}
}
