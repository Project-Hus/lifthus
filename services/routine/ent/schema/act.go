package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Act holds the schema definition for the Act entity.
type Act struct {
	ent.Schema
}

// Fields of the Act.
// Which represents a single act like squat, running, or even taking rest.
func (Act) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(), // user id from Hus

		field.String("name").NotEmpty().Annotations(entsql.Annotation{Size: 50}),
		field.Enum("type").Values("rep", "lap").Nillable().Optional(),

		field.Uint64("author"),
		field.String("image").Nillable().Optional(),
		field.String("description").Nillable().Optional().Annotations(entsql.Annotation{Size: 5000}),
	}
}

// Edges of the Act.
func (Act) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tags", Tag.Type).Ref("acts"),
		edge.To("routine_acts", RoutineAct.Type),
	}
}
