package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProgramImage holds the schema definition for the ProgramImage entity.
type S3ProgramImage struct {
	ent.Schema
}

// Fields of the ProgramImage.
func (S3ProgramImage) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.Int("order"),
		field.Int64("program_release_id"),
		field.Int64("image_id"),
	}
}

// Edges of the ProgramImage.
func (S3ProgramImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program_release", ProgramRelease.Type).Ref("s3_program_images").Unique().Field("program_release_id").Required(),
		edge.To("s3_image", S3Image.Type).
			Required().
			Unique().
			Field("image_id"),
	}
}
