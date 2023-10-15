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
		field.Int("order").Immutable(),
		CodeRef("act_code"),
		field.Enum("stage").Values(RoutineActStage...).Immutable(),
		field.Uint("reps_or_meters").Immutable(),
		field.Float("ratio_or_secs").Immutable(),
	}
}

// Edges of the RoutineAct.
func (RoutineAct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("act", Act.Type).Ref("routine_acts").Unique().Required(),
		edge.From("routine", Routine.Type).Ref("routine_acts").Unique().Required(),
	}
}
