package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OneRepMax holds the schema definition for the OneRepMax entity.
type OneRepMax struct {
	ent.Schema
}

// Fields of the OneRepMax.
func (OneRepMax) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("author").Unique(),

		field.Uint64("act_id").Unique(),
		field.Uint64("program_rec_id").Nillable().Optional(),

		field.Time("date"),

		field.Float("one_rep_max").Nillable().Optional(),
		field.Bool("certified").Default(false),
		field.Bool("calculated").Default(false),
	}
}

// Edges of the OneRepMax.
func (OneRepMax) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("act", Act.Type).Field("act_id").Ref("one_rep_maxes").Unique().Required(),
		edge.From("program_rec", ProgramRec.Type).Field("program_rec_id").Ref("one_rep_max").Unique(),
	}
}
