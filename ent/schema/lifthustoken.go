package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LifthusToken holds the schema definition for the LifthusToken entity.
type LifthusToken struct {
	ent.Schema
}

// Fields of the LifthusToken.
func (LifthusToken) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StructTag(`json:tid,omitempty`).Default(uuid.New),
		field.UUID("uid", uuid.UUID{}), // User ID
		field.Bool("revoked").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the LifthusToken.
func (LifthusToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("lifthus_tokens").Unique().Field("uid").Required(),
	}
}
