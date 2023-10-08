package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Program holds the schema definition for the Program entity.
type ProgramVersion struct {
	ent.Schema
}

// Fields of the Program.
func (ProgramVersion) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		CodeField(),

		CodeRef("program_code"),
		field.Uint("version").Immutable(),
		CreatedAtField(),

		field.Text("text").NotEmpty().Immutable(),
	}
}

// Edges of the Program.
func (ProgramVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program", Program.Type).Ref("program_versions").Unique().Required(),

		edge.To("images", Image.Type),
		edge.To("daily_routines", DailyRoutine.Type),
	}
}
