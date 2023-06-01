package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BodyInfo holds the schema definition for the BodyInfo entity.
type BodyInfo struct {
	ent.Schema
}

// Fields of the BodyInfo.
func (BodyInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("author").Unique(),
		field.Uint64("program_rec_id").Nillable().Optional(),

		field.Time("date").Default(time.Now),

		field.Float("height").Nillable().Optional(),
		field.Float("body_weight").Nillable().Optional(),
		field.Float("body_fat_mass").Nillable().Optional(),
		field.Float("skeletal_muscle_mass").Nillable().Optional(),
	}
}

// Edges of the BodyInfo.
func (BodyInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("program_rec", ProgramRec.Type).Field("program_rec_id").Ref("body_info").Unique(),
	}
}
