package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Program holds the schema definition for the Program entity.
type Program struct {
	ent.Schema
}

// Fields of the Program.
func (Program) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		CodeField(),

		field.Enum("program_type").Values(ProgramType...).Immutable(),
		field.String("title").NotEmpty().Annotations(entsql.Annotation{Size: 100}).Immutable(),
		field.Int64("author").Immutable(),
		CreatedAtField(),

		CodeRefNillable("parent_program"),
		field.Int("parent_version").Optional().Nillable().Immutable(),
	}
}

// Edges of the Program.
func (Program) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("program_releases", ProgramRelease.Type),
	}
}
