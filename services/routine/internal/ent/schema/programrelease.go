package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Program holds the schema definition for the Program entity.
type ProgramRelease struct {
	ent.Schema
}

// Fields of the Program.
func (ProgramRelease) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.Int("version").Immutable(),
		CreatedAtField(),
		field.Text("text").Immutable(),
	}
}

// Edges of the Program.
func (ProgramRelease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program", Program.Type).Ref("program_releases").Unique().Required(),

		edge.To("s3_program_images", S3ProgramImage.Type),
		edge.To("routines", Routine.Type),
	}
}
