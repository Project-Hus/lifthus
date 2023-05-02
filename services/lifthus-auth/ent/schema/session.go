package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the LifthusSession entity.
type Session struct {
	ent.Schema
}

// Fields of the LifthusSession.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StructTag(`json:"sid,omitempty"`).Default(uuid.New).Unique(), // sid
		field.Uint64("uid").Optional().Nillable(),
		field.Time("connected_at").Default(time.Now),                          // connected at
		field.Time("signed_at").Optional().Nillable().UpdateDefault(time.Now), // signed at
		field.Bool("used").Default(false),                                     // used to sign Lifthus session
	}
}

// Edges of the LifthusSession.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("sessions").Unique().Field("uid"),
	}
}
