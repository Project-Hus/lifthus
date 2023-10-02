package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// WeeklyRoutineRec holds the schema definition for the WeeklyRoutineRec entity.
type WeeklyRoutineRec struct {
	ent.Schema
}

func (WeeklyRoutineRec) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("program_rec_id", "week").Unique(),
	}
}

// Fields of the WeeklyRoutineRec.
func (WeeklyRoutineRec) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),

		field.Uint64("program_rec_id"),
		field.Uint64("weekly_routine_id"),

		field.Int("week"),

		field.Time("start_date"),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the WeeklyRoutineRec.
func (WeeklyRoutineRec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weekly_routine", WeeklyRoutine.Type).Field("weekly_routine_id").Ref("weekly_routine_recs").Unique().Required(),

		edge.From("program_rec", ProgramRec.Type).Field("program_rec_id").Ref("weekly_routine_recs").Unique().Required(),

		edge.To("daily_routine_recs", DailyRoutineRec.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
