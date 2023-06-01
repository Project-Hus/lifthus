package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WeeklyRoutine holds the schema definition for the WeeklyRoutine entity.
type WeeklyRoutine struct {
	ent.Schema
}

// Fields of the WeeklyRoutine.
func (WeeklyRoutine) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("program_id"),
		field.Int("week").Min(1),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the WeeklyRoutine.
func (WeeklyRoutine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program", Program.Type).Ref("weekly_routines"),
		edge.To("daily_routines", DailyRoutine.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		edge.To("weekly_routine_recs", WeeklyRoutineRec.Type),
	}
}
