package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ActImage holds the schema definition for the ActImage entity.
type ActImage struct {
	ent.Schema
}

func (ActImage) Indexes() []ent.Index {
	return []ent.Index{}
}

// Fields of the ActImage.
func (ActImage) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		CodeRef("act_version_code"),
		field.Uint("order"),
		field.Text("src").NotEmpty(),
	}
}

// Edges of the ActImage.
func (ActImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("act_version", ActVersion.Type).Ref("act_images").Unique().Required(),
	}
}
