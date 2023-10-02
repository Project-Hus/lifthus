package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Act holds the schema definition for the Act entity.
type Act struct {
	ent.Schema
}

func (Act) Indexes() []ent.Index {
	return []ent.Index{
		CodeIndex(),
		index.Fields("name"),
	}
}

// Fields of the Act.
// Which represents a single act like squat, running, or even taking rest.
func (Act) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		CodeField(),
		field.Enum("act_type").Values(ActType...).Immutable(),
		field.String("name").NotEmpty().Annotations(entsql.Annotation{Size: 50}).Immutable(),
		field.Uint64("author"),
		CreatedAtField(),
	}
}

// Edges of the Act.
func (Act) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("act_versions", ActVersion.Type),
	}
}
