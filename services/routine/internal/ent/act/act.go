// Code generated by ent, DO NOT EDIT.

package act

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the act type in the database.
	Label = "act"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldActType holds the string denoting the act_type field in the database.
	FieldActType = "act_type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldStandard holds the string denoting the standard field in the database.
	FieldStandard = "standard"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeS3ActImages holds the string denoting the s3_act_images edge name in mutations.
	EdgeS3ActImages = "s3_act_images"
	// EdgeRoutineActs holds the string denoting the routine_acts edge name in mutations.
	EdgeRoutineActs = "routine_acts"
	// Table holds the table name of the act in the database.
	Table = "acts"
	// S3ActImagesTable is the table that holds the s3_act_images relation/edge.
	S3ActImagesTable = "s3act_images"
	// S3ActImagesInverseTable is the table name for the S3ActImage entity.
	// It exists in this package in order to avoid circular dependency with the "s3actimage" package.
	S3ActImagesInverseTable = "s3act_images"
	// S3ActImagesColumn is the table column denoting the s3_act_images relation/edge.
	S3ActImagesColumn = "act_id"
	// RoutineActsTable is the table that holds the routine_acts relation/edge.
	RoutineActsTable = "routine_acts"
	// RoutineActsInverseTable is the table name for the RoutineAct entity.
	// It exists in this package in order to avoid circular dependency with the "routineact" package.
	RoutineActsInverseTable = "routine_acts"
	// RoutineActsColumn is the table column denoting the routine_acts relation/edge.
	RoutineActsColumn = "act_routine_acts"
)

// Columns holds all SQL columns for act fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldAuthor,
	FieldActType,
	FieldName,
	FieldText,
	FieldStandard,
	FieldCreatedAt,
}

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
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultStandard holds the default value on creation for the "standard" field.
	DefaultStandard bool
)

// ActType defines the type for the "act_type" enum field.
type ActType string

// ActType values.
const (
	ActTypeWeight ActType = "weight"
	ActTypeTime   ActType = "time"
	ActTypeSimple ActType = "simple"
)

func (at ActType) String() string {
	return string(at)
}

// ActTypeValidator is a validator for the "act_type" field enum values. It is called by the builders before save.
func ActTypeValidator(at ActType) error {
	switch at {
	case ActTypeWeight, ActTypeTime, ActTypeSimple:
		return nil
	default:
		return fmt.Errorf("act: invalid enum value for act_type field: %q", at)
	}
}

// OrderOption defines the ordering options for the Act queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByActType orders the results by the act_type field.
func ByActType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActType, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByStandard orders the results by the standard field.
func ByStandard(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStandard, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByS3ActImagesCount orders the results by s3_act_images count.
func ByS3ActImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newS3ActImagesStep(), opts...)
	}
}

// ByS3ActImages orders the results by s3_act_images terms.
func ByS3ActImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newS3ActImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newS3ActImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(S3ActImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, S3ActImagesTable, S3ActImagesColumn),
	)
}
func newRoutineActsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutineActsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutineActsTable, RoutineActsColumn),
	)
}
