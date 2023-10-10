package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DailyRoutine holds the schema definition for the DailyRoutine entity.
type DailyRoutine struct {
	ent.Schema
}

// Fields of the DailyRoutine.
func (DailyRoutine) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		CodeField(),
		CodeRef("program_version_code"),
		field.Uint("day").Immutable(),
	}
}

// Edges of the DailyRoutine.
func (DailyRoutine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program_version", ProgramVersion.Type).Ref("daily_routines").Unique().Required(),

		edge.To("routine_acts", RoutineAct.Type),
	}
}
