package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
// Which tags acts and programs.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("tag").NotEmpty().Annotations(entsql.Annotation{Size: 20}),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
