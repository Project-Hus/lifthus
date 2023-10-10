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
		field.String("title").NotEmpty().Annotations(entsql.Annotation{Size: 50}).Immutable(),
		field.Uint64("author").Immutable(),
		CreatedAtField(),

		CodeRefNillable("version_derived_from"),
	}
}

// Edges of the Program.
func (Program) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("program_versions", ProgramVersion.Type),
	}
}
