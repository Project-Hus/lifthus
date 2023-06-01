package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WeeklyRoutineRec holds the schema definition for the WeeklyRoutineRec entity.
type WeeklyRoutineRec struct {
	ent.Schema
}

// Fields of the WeeklyRoutineRec.
func (WeeklyRoutineRec) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),

		field.Uint64("program_rec_id"),
		field.Uint64("weekly_routine_id"),
		field.Time("start_date"),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the WeeklyRoutineRec.
func (WeeklyRoutineRec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weekly_routine", WeeklyRoutine.Type).Ref("weekly_routine_recs").Unique(),

		edge.From("program_rec", ProgramRec.Type).Ref("weekly_routine_recs").Unique(),
		edge.To("daily_routine_recs", DailyRoutineRec.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
