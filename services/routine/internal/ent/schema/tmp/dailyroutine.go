package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DailyRoutine holds the schema definition for the DailyRoutine entity.
type DailyRoutine struct {
	ent.Schema
}

func (DailyRoutine) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("program_id", "weekly_routine_id", "day").Unique(),
	}
}

// Fields of the DailyRoutine.
func (DailyRoutine) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),

		field.Uint64("program_id").Nillable().Optional(),
		field.Uint64("weekly_routine_id").Nillable().Optional(),

		field.Int("day").Min(1),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the DailyRoutine.
func (DailyRoutine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program", Program.Type).Field("program_id").Ref("daily_routines").Unique(),
		edge.From("weekly_routine", WeeklyRoutine.Type).Field("weekly_routine_id").Ref("daily_routines").Unique(),

		edge.To("routine_acts", RoutineAct.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		edge.To("daily_routine_recs", DailyRoutineRec.Type),
	}
}
