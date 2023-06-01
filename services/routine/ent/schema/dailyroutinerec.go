package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// DailyRoutineRec holds the schema definition for the DailyRoutineRec entity.
type DailyRoutineRec struct {
	ent.Schema
}

// Fields of the DailyRoutineRec.
func (DailyRoutineRec) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("program_rec_id").Nillable().Optional(),
		field.Uint64("weekly_routine_rec_id").Nillable().Optional(),
		field.Uint64("daily_routine_id").Nillable().Optional(),
		field.Time("date"),
		field.Enum("status").Values(recStatus...),

		field.String("comment").Nillable().Optional().Annotations(entsql.Annotation{Size: 1000}),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the DailyRoutineRec.
func (DailyRoutineRec) Edges() []ent.Edge {
	return nil
}
