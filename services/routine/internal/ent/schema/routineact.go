package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RoutineAct holds the schema definition for the RoutineAct entity.
type RoutineAct struct {
	ent.Schema
}

// Fields of the RoutineAct.
func (RoutineAct) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		CodeRef("daily_routine_code"),
		field.Uint("order").Immutable(),
		CodeRef("act_version"),
		field.Enum("stage").Values(RoutineActStage...).Immutable(),
		field.Uint("reps_or_meters").Immutable(),
		field.Float("ratio_or_secs").Immutable(),
	}
}

// Edges of the RoutineAct.
func (RoutineAct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("daily_routine", DailyRoutine.Type).Ref("routine_acts").Unique().Required(),
	}
}
