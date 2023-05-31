package schema

import (
	"entgo.io/ent"
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
	}
}

// Edges of the DailyRoutine.
func (DailyRoutine) Edges() []ent.Edge {
	return nil
}
