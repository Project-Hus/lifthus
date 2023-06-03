// Code generated by ent, DO NOT EDIT.

package act

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the act type in the database.
	Label = "act"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldBodyweight holds the string denoting the bodyweight field in the database.
	FieldBodyweight = "bodyweight"
	// FieldCardio holds the string denoting the cardio field in the database.
	FieldCardio = "cardio"
	// FieldUpper holds the string denoting the upper field in the database.
	FieldUpper = "upper"
	// FieldLower holds the string denoting the lower field in the database.
	FieldLower = "lower"
	// FieldFull holds the string denoting the full field in the database.
	FieldFull = "full"
	// FieldArms holds the string denoting the arms field in the database.
	FieldArms = "arms"
	// FieldShoulders holds the string denoting the shoulders field in the database.
	FieldShoulders = "shoulders"
	// FieldChest holds the string denoting the chest field in the database.
	FieldChest = "chest"
	// FieldCore holds the string denoting the core field in the database.
	FieldCore = "core"
	// FieldUpperBack holds the string denoting the upper_back field in the database.
	FieldUpperBack = "upper_back"
	// FieldLowerBack holds the string denoting the lower_back field in the database.
	FieldLowerBack = "lower_back"
	// FieldGlute holds the string denoting the glute field in the database.
	FieldGlute = "glute"
	// FieldLegsFront holds the string denoting the legs_front field in the database.
	FieldLegsFront = "legs_front"
	// FieldLegsBack holds the string denoting the legs_back field in the database.
	FieldLegsBack = "legs_back"
	// FieldEtc holds the string denoting the etc field in the database.
	FieldEtc = "etc"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeRoutineActs holds the string denoting the routine_acts edge name in mutations.
	EdgeRoutineActs = "routine_acts"
	// EdgeRoutineActRecs holds the string denoting the routine_act_recs edge name in mutations.
	EdgeRoutineActRecs = "routine_act_recs"
	// EdgeOneRepMaxes holds the string denoting the one_rep_maxes edge name in mutations.
	EdgeOneRepMaxes = "one_rep_maxes"
	// Table holds the table name of the act in the database.
	Table = "acts"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "tag_acts"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// RoutineActsTable is the table that holds the routine_acts relation/edge.
	RoutineActsTable = "routine_acts"
	// RoutineActsInverseTable is the table name for the RoutineAct entity.
	// It exists in this package in order to avoid circular dependency with the "routineact" package.
	RoutineActsInverseTable = "routine_acts"
	// RoutineActsColumn is the table column denoting the routine_acts relation/edge.
	RoutineActsColumn = "act_id"
	// RoutineActRecsTable is the table that holds the routine_act_recs relation/edge.
	RoutineActRecsTable = "routine_act_recs"
	// RoutineActRecsInverseTable is the table name for the RoutineActRec entity.
	// It exists in this package in order to avoid circular dependency with the "routineactrec" package.
	RoutineActRecsInverseTable = "routine_act_recs"
	// RoutineActRecsColumn is the table column denoting the routine_act_recs relation/edge.
	RoutineActRecsColumn = "act_id"
	// OneRepMaxesTable is the table that holds the one_rep_maxes relation/edge.
	OneRepMaxesTable = "one_rep_maxes"
	// OneRepMaxesInverseTable is the table name for the OneRepMax entity.
	// It exists in this package in order to avoid circular dependency with the "onerepmax" package.
	OneRepMaxesInverseTable = "one_rep_maxes"
	// OneRepMaxesColumn is the table column denoting the one_rep_maxes relation/edge.
	OneRepMaxesColumn = "act_id"
)

// Columns holds all SQL columns for act fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCode,
	FieldType,
	FieldAuthor,
	FieldImage,
	FieldDescription,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldWeight,
	FieldBodyweight,
	FieldCardio,
	FieldUpper,
	FieldLower,
	FieldFull,
	FieldArms,
	FieldShoulders,
	FieldChest,
	FieldCore,
	FieldUpperBack,
	FieldLowerBack,
	FieldGlute,
	FieldLegsFront,
	FieldLegsBack,
	FieldEtc,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"tag_id", "act_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCode holds the default value on creation for the "code" field.
	DefaultCode func() string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultWeight holds the default value on creation for the "weight" field.
	DefaultWeight bool
	// DefaultBodyweight holds the default value on creation for the "bodyweight" field.
	DefaultBodyweight bool
	// DefaultCardio holds the default value on creation for the "cardio" field.
	DefaultCardio bool
	// DefaultUpper holds the default value on creation for the "upper" field.
	DefaultUpper bool
	// DefaultLower holds the default value on creation for the "lower" field.
	DefaultLower bool
	// DefaultFull holds the default value on creation for the "full" field.
	DefaultFull bool
	// DefaultArms holds the default value on creation for the "arms" field.
	DefaultArms bool
	// DefaultShoulders holds the default value on creation for the "shoulders" field.
	DefaultShoulders bool
	// DefaultChest holds the default value on creation for the "chest" field.
	DefaultChest bool
	// DefaultCore holds the default value on creation for the "core" field.
	DefaultCore bool
	// DefaultUpperBack holds the default value on creation for the "upper_back" field.
	DefaultUpperBack bool
	// DefaultLowerBack holds the default value on creation for the "lower_back" field.
	DefaultLowerBack bool
	// DefaultGlute holds the default value on creation for the "glute" field.
	DefaultGlute bool
	// DefaultLegsFront holds the default value on creation for the "legs_front" field.
	DefaultLegsFront bool
	// DefaultLegsBack holds the default value on creation for the "legs_back" field.
	DefaultLegsBack bool
	// DefaultEtc holds the default value on creation for the "etc" field.
	DefaultEtc bool
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeRep    Type = "rep"
	TypeLap    Type = "lap"
	TypeSimple Type = "simple"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeRep, TypeLap, TypeSimple:
		return nil
	default:
		return fmt.Errorf("act: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Act queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByBodyweight orders the results by the bodyweight field.
func ByBodyweight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBodyweight, opts...).ToFunc()
}

// ByCardio orders the results by the cardio field.
func ByCardio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCardio, opts...).ToFunc()
}

// ByUpper orders the results by the upper field.
func ByUpper(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpper, opts...).ToFunc()
}

// ByLower orders the results by the lower field.
func ByLower(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLower, opts...).ToFunc()
}

// ByFull orders the results by the full field.
func ByFull(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFull, opts...).ToFunc()
}

// ByArms orders the results by the arms field.
func ByArms(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArms, opts...).ToFunc()
}

// ByShoulders orders the results by the shoulders field.
func ByShoulders(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShoulders, opts...).ToFunc()
}

// ByChest orders the results by the chest field.
func ByChest(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChest, opts...).ToFunc()
}

// ByCore orders the results by the core field.
func ByCore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCore, opts...).ToFunc()
}

// ByUpperBack orders the results by the upper_back field.
func ByUpperBack(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpperBack, opts...).ToFunc()
}

// ByLowerBack orders the results by the lower_back field.
func ByLowerBack(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLowerBack, opts...).ToFunc()
}

// ByGlute orders the results by the glute field.
func ByGlute(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGlute, opts...).ToFunc()
}

// ByLegsFront orders the results by the legs_front field.
func ByLegsFront(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLegsFront, opts...).ToFunc()
}

// ByLegsBack orders the results by the legs_back field.
func ByLegsBack(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLegsBack, opts...).ToFunc()
}

// ByEtc orders the results by the etc field.
func ByEtc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEtc, opts...).ToFunc()
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoutineActsCount orders the results by routine_acts count.
func ByRoutineActsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoutineActsStep(), opts...)
	}
}

// ByRoutineActs orders the results by routine_acts terms.
func ByRoutineActs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoutineActsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRoutineActRecsCount orders the results by routine_act_recs count.
func ByRoutineActRecsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoutineActRecsStep(), opts...)
	}
}

// ByRoutineActRecs orders the results by routine_act_recs terms.
func ByRoutineActRecs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoutineActRecsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOneRepMaxesCount orders the results by one_rep_maxes count.
func ByOneRepMaxesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOneRepMaxesStep(), opts...)
	}
}

// ByOneRepMaxes orders the results by one_rep_maxes terms.
func ByOneRepMaxes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOneRepMaxesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
	)
}
func newRoutineActsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutineActsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutineActsTable, RoutineActsColumn),
	)
}
func newRoutineActRecsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutineActRecsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutineActRecsTable, RoutineActRecsColumn),
	)
}
func newOneRepMaxesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OneRepMaxesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OneRepMaxesTable, OneRepMaxesColumn),
	)
}
