package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Act holds the schema definition for the Act entity.
type Act struct {
	ent.Schema
}

// Fields of the Act.
// Which represents a single act like squat, running, or even taking rest.
func (Act) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(), // user id from Hus

		field.String("name").NotEmpty().Annotations(entsql.Annotation{Size: 50}),
		field.String("slug").Unique(),
		field.Enum("type").Values(ActType...),

		field.Uint64("author"),
		field.String("image").Nillable().Optional(),
		field.String("description").Nillable().Optional().Annotations(entsql.Annotation{Size: 5000}),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		// weight/cardio
		field.Bool("weight").Default(false),
		field.Bool("bodyweight").Default(false),
		field.Bool("cardio").Default(false),

		// upper/lower/full body
		field.Bool("upper").Default(false),
		field.Bool("lower").Default(false),
		field.Bool("full").Default(false),

		// prime movers
		field.Bool("arms").Default(false),
		field.Bool("shoulders").Default(false),
		field.Bool("chest").Default(false),
		field.Bool("core").Default(false),
		field.Bool("upper_back").Default(false),
		field.Bool("lower_back").Default(false),
		field.Bool("glute").Default(false),
		field.Bool("legs_front").Default(false),
		field.Bool("legs_back").Default(false),
		field.Bool("etc").Default(false),
	}
}

// Edges of the Act.
func (Act) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tags", Tag.Type).Ref("acts"),
		edge.To("routine_acts", RoutineAct.Type),

		edge.To("routine_act_recs", RoutineActRec.Type),

		edge.To("one_rep_maxes", OneRepMax.Type),
	}
}
