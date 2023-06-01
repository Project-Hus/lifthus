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
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeRoutineActs holds the string denoting the routine_acts edge name in mutations.
	EdgeRoutineActs = "routine_acts"
	// EdgeRoutineActRecs holds the string denoting the routine_act_recs edge name in mutations.
	EdgeRoutineActRecs = "routine_act_recs"
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
	RoutineActsColumn = "act_routine_acts"
	// RoutineActRecsTable is the table that holds the routine_act_recs relation/edge.
	RoutineActRecsTable = "routine_act_recs"
	// RoutineActRecsInverseTable is the table name for the RoutineActRec entity.
	// It exists in this package in order to avoid circular dependency with the "routineactrec" package.
	RoutineActRecsInverseTable = "routine_act_recs"
	// RoutineActRecsColumn is the table column denoting the routine_act_recs relation/edge.
	RoutineActRecsColumn = "act_routine_act_recs"
)

// Columns holds all SQL columns for act fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldAuthor,
	FieldImage,
	FieldDescription,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeRep Type = "rep"
	TypeLap Type = "lap"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeRep, TypeLap:
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
