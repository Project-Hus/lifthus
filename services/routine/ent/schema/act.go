package schema

import "entgo.io/ent"

// Act holds the schema definition for the Act entity.
type Act struct {
	ent.Schema
}

// Fields of the Act.
func (Act) Fields() []ent.Field {
	return nil
}

// Edges of the Act.
func (Act) Edges() []ent.Edge {
	return nil
}
