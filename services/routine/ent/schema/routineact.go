package schema

import (
	"time"

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
		field.Uint64("id").Unique(),

		field.Uint64("act_id"),
		field.Uint64("daily_routine_id"),

		field.Int("order").Min(1),
		field.Int("reps").Min(1).Nillable().Optional(),
		field.Int("lap").Min(1).Nillable().Optional(),
		field.Bool("warmup").Default(false),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the RoutineAct.
func (RoutineAct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("act", Act.Type).Field("act_id").Ref("routine_acts").Required().Unique(),
		edge.From("daily_routine", DailyRoutine.Type).Field("daily_routine_id").Ref("routine_acts").Required().Unique(),

		edge.To("routine_act_recs", RoutineActRec.Type),
	}
}
