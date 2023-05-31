package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// Program holds the schema definition for the Program entity.
type Program struct {
	ent.Schema
}

// Fields of the Program.
func (Program) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),

		field.String("title").NotEmpty().Annotations(entsql.Annotation{Size: 50}),
		field.Enum("type").Values("weekly", "daily"),

		field.Uint64("author"),

		field.String("image").Nillable().Optional(),
		field.String("description").Nillable().Optional().Annotations(entsql.Annotation{Size: 5000}),
	}
}

// Edges of the Program.
func (Program) Edges() []ent.Edge {
	return nil
}
