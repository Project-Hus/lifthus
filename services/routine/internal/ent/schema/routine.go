package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DayRoutine holds the schema definition for the DailyRoutine entity.
type Routine struct {
	ent.Schema
}

// Fields of the DailyRoutine.
func (Routine) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.Int("day").Immutable(),
	}
}

// Edges of the DailyRoutine.
func (Routine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program_release", ProgramRelease.Type).Ref("routines").Unique().Required(),

		edge.To("routine_acts", RoutineAct.Type),
	}
}
