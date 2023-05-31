package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RoutineAct holds the schema definition for the RoutineAct entity.
type RoutineAct struct {
	ent.Schema
}

// Fields of the RoutineAct.
func (RoutineAct) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("daily_routine_id"),
		field.Uint64("act_id"),
		field.Int("order").Min(1),
		field.Int("reps").Min(1).Nillable().Optional(),
		field.Int("lap").Min(1).Nillable().Optional(),
	}
}

// Edges of the RoutineAct.
func (RoutineAct) Edges() []ent.Edge {
	return nil
}
