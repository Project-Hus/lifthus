package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the LifthusSession entity.
type Session struct {
	ent.Schema
}

func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Options: "ENGINE=MEMORY",
		},
	}
}

// Fields of the LifthusSession.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StructTag(`json:"sid,omitempty"`).Default(uuid.New).Unique(), // sid
		field.UUID("tid", uuid.UUID{}).Default(uuid.New),                                           // tid for rotation

		field.UUID("hsid", uuid.UUID{}).Optional().Nillable(), // Hus session ID
		field.Time("connected_at").Default(time.Now),          // connected at

		field.Uint64("uid").Optional().Nillable(),
		field.Time("signed_at").Optional().Nillable(), // signed at
	}
}

// Edges of the LifthusSession.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("sessions").Unique().Field("uid"),
	}
}
