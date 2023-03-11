package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LifthusSession holds the schema definition for the LifthusSession entity.
type LifthusSession struct {
	ent.Schema
}

// Fields of the LifthusSession.
func (LifthusSession) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StructTag(`json:"sid,omitempty"`).Default(uuid.New).Unique(), // sid
		field.UUID("uid", uuid.UUID{}).Optional().Nillable(),
		field.Time("connected_at").Default(time.Now), // connected at
	}
}

// Edges of the LifthusSession.
func (LifthusSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("lifthus_sessions").Unique().Field("uid"),
	}
}
