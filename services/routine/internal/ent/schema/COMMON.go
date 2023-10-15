package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// enum
var ProgramType = []string{"weekly", "daily"}
var ActType = []string{"weight", "time", "simple"}
var RoutineActStage = []string{"warmup", "main", "cooldown"}

var RecStatus = []string{"history", "waiting", "proceeding", "completed", "failed", "canceled"}

func IdField() ent.Field {
	return field.Int64("id").Unique().Immutable()
}

func CodeIndex() ent.Index {
	return index.Fields("code").Unique()
}

func CodeField() ent.Field {
	return field.String("code").NotEmpty().Unique().Annotations(entsql.Annotation{Size: 20}).Immutable()
}

func CodeRef(fn string) ent.Field {
	return field.String(fn).NotEmpty().Annotations(entsql.Annotation{Size: 20}).Immutable()
}

func CodeRefNillable(fn string) ent.Field {
	return field.String(fn).NotEmpty().Annotations(entsql.Annotation{Size: 20}).Immutable().Optional().Nillable()
}

func CreatedAtField() ent.Field {
	return field.Time("created_at").Immutable()
}
