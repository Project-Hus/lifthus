package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ActImage holds the schema definition for the ActImage entity.
type S3ActImage struct {
	ent.Schema
}

// Fields of the ActImage.
func (S3ActImage) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.Int("order"),
		field.Int64("act_id"),
		field.Int64("image_id"),
	}
}

// Edges of the ActImage.
func (S3ActImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("act", Act.Type).Ref("s3_act_images").Unique().Field("act_id").Required(),

		edge.To("s3_image", S3ActImage.Type).Unique().Field("image_id").Required(),
	}
}
