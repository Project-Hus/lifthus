package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProgramRec holds the schema definition for the ProgramRec entity.
type ProgramRec struct {
	ent.Schema
}

// Fields of the ProgramRec.
func (ProgramRec) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),

		field.Uint64("author"),

		field.Uint64("program_id"),

		field.Time("start_date"),
		field.Time("end_date"),
		field.Enum("status").Values(RecStatus...),

		field.String("comment").Nillable().Optional().Annotations(entsql.Annotation{Size: 1000}),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ProgramRec.
func (ProgramRec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program", Program.Type).Field("program_id").Ref("program_recs").Unique().Required(),

		edge.To("weekly_routine_recs", WeeklyRoutineRec.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("daily_routine_recs", DailyRoutineRec.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		edge.To("body_info", BodyInfo.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("one_rep_max", OneRepMax.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
