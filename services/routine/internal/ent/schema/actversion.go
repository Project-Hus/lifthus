package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ActVersion holds the schema definition for the ActVersion entity.
type ActVersion struct {
	ent.Schema
}

// Fields of the ActVersion.
func (ActVersion) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		CodeField(),
		CodeRef("act_code"),
		field.Uint("version").Immutable(),
		CreatedAtField(),
		field.Text("text").NotEmpty(),
	}
}

// Edges of the ActVersion.
func (ActVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("act", Act.Type).Ref("act_versions").Unique().Required(),

		edge.To("images", Image.Type).Through("act_images", ActImage.Type),

		edge.To("routine_acts", RoutineAct.Type),
	}
}
