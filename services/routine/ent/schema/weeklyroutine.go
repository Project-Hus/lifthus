package schema

import (
	"entgo.io/ent"
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
	}
}

// Edges of the WeeklyRoutine.
func (WeeklyRoutine) Edges() []ent.Edge {
	return nil
}
