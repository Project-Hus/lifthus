package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
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
		field.String("slug").Unique(),
		field.Enum("type").Values(ProgramType...),

		field.Uint64("author"),

		field.String("image").Nillable().Optional(),
		field.String("description").Nillable().Optional().Annotations(entsql.Annotation{Size: 5000}),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(nil).UpdateDefault(time.Now).Nillable(),
	}
}

// Edges of the Program.
func (Program) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tags", Tag.Type).Ref("programs"),

		edge.To("weekly_routines", WeeklyRoutine.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("daily_routines", DailyRoutine.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		edge.To("program_recs", ProgramRec.Type),
	}
}
