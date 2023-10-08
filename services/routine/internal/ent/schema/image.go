package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

func (Image) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Unique(),
		index.Fields("src").Unique(),
	}
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
