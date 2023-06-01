package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Uint64("id").Unique(),
		field.Uint64("program_id").Nillable().Optional(),
		field.Uint64("week_id").Nillable().Optional(),
		field.Int("day").Min(1),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the DailyRoutine.
func (DailyRoutine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program", Program.Type).Ref("daily_routines"),
		edge.From("weekly_routine", WeeklyRoutine.Type).Ref("daily_routines"),

		edge.To("routine_acts", RoutineAct.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		edge.To("daily_routine_recs", DailyRoutineRec.Type),
	}
}
