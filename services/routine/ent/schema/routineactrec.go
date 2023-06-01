package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RoutineActRec holds the schema definition for the RoutineActRec entity.
type RoutineActRec struct {
	ent.Schema
}

// Fields of the RoutineActRec.
func (RoutineActRec) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("daily_routine_rec_id"),

		field.Uint64("routine_act_id").Nillable().Optional(),
		field.Uint64("act_id").Nillable().Optional(),

		field.Int("order").Min(1),
		field.Int("reps").Min(1).Nillable().Optional(),
		field.Int("lap").Min(1).Nillable().Optional(),
		field.Int("current_reps").Min(0).Default(0),
		field.Int("current_lap").Min(0).Default(0),
		field.String("image").Nillable().Optional(),
		field.String("comment").Nillable().Optional(),

		field.Enum("status").Values(recStatus...),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the RoutineActRec.
func (RoutineActRec) Edges() []ent.Edge {
	return nil
}
