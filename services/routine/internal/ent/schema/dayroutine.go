package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DayRoutine holds the schema definition for the DailyRoutine entity.
type DayRoutine struct {
	ent.Schema
}

// Fields of the DailyRoutine.
func (DayRoutine) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.Int("day").Immutable(),
	}
}

// Edges of the DailyRoutine.
func (DayRoutine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("routine_acts", RoutineAct.Type),
	}
}
