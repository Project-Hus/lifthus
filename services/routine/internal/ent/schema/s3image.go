package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Image holds the schema definition for the Image entity.
type S3Image struct {
	ent.Schema
}

func (S3Image) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Unique(),
		index.Fields("src").Unique(),
	}
}

// Fields of the Image.
func (S3Image) Fields() []ent.Field {
	return []ent.Field{
		IdField(),
		field.String("key").Unique(),
		field.String("src").Unique(),
		field.Time("created_at").Immutable(),
	}
}

// Edges of the Image.
func (S3Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("s3_act_images", S3ActImage.Type).Ref("s3_image"),
	}
}
