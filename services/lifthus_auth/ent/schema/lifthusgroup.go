package schema

import "entgo.io/ent"

// LifthusGroup holds the schema definition for the LifthusGroup entity.
type LifthusGroup struct {
	ent.Schema
}

// Fields of the LifthusGroup.
func (LifthusGroup) Fields() []ent.Field {
	return nil
}

// Edges of the LifthusGroup.
func (LifthusGroup) Edges() []ent.Edge {
	return nil
}
