// Code generated by ent, DO NOT EDIT.

package actversion

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the actversion type in the database.
	Label = "act_version"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldActCode holds the string denoting the act_code field in the database.
	FieldActCode = "act_code"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// EdgeAct holds the string denoting the act edge name in mutations.
	EdgeAct = "act"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// EdgeRoutineActs holds the string denoting the routine_acts edge name in mutations.
	EdgeRoutineActs = "routine_acts"
	// EdgeActImages holds the string denoting the act_images edge name in mutations.
	EdgeActImages = "act_images"
	// Table holds the table name of the actversion in the database.
	Table = "act_versions"
	// ActTable is the table that holds the act relation/edge.
	ActTable = "act_versions"
	// ActInverseTable is the table name for the Act entity.
	// It exists in this package in order to avoid circular dependency with the "act" package.
	ActInverseTable = "acts"
	// ActColumn is the table column denoting the act relation/edge.
	ActColumn = "act_act_versions"
	// ImagesTable is the table that holds the images relation/edge. The primary key declared below.
	ImagesTable = "act_images"
	// ImagesInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImagesInverseTable = "images"
	// RoutineActsTable is the table that holds the routine_acts relation/edge.
	RoutineActsTable = "routine_acts"
	// RoutineActsInverseTable is the table name for the RoutineAct entity.
	// It exists in this package in order to avoid circular dependency with the "routineact" package.
	RoutineActsInverseTable = "routine_acts"
	// RoutineActsColumn is the table column denoting the routine_acts relation/edge.
	RoutineActsColumn = "act_version_routine_acts"
	// ActImagesTable is the table that holds the act_images relation/edge.
	ActImagesTable = "act_images"
	// ActImagesInverseTable is the table name for the ActImage entity.
	// It exists in this package in order to avoid circular dependency with the "actimage" package.
	ActImagesInverseTable = "act_images"
	// ActImagesColumn is the table column denoting the act_images relation/edge.
	ActImagesColumn = "act_version_id"
)

// Columns holds all SQL columns for actversion fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldActCode,
	FieldVersion,
	FieldCreatedAt,
	FieldText,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "act_versions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"act_act_versions",
}

var (
	// ImagesPrimaryKey and ImagesColumn2 are the table columns denoting the
	// primary key for the images relation (M2M).
	ImagesPrimaryKey = []string{"act_version_id", "image_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// ActCodeValidator is a validator for the "act_code" field. It is called by the builders before save.
	ActCodeValidator func(string) error
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
)

// OrderOption defines the ordering options for the ActVersion queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByActCode orders the results by the act_code field.
func ByActCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActCode, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByActField orders the results by act field.
func ByActField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActStep(), sql.OrderByField(field, opts...))
	}
}

// ByImagesCount orders the results by images count.
func ByImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImagesStep(), opts...)
	}
}

// ByImages orders the results by images terms.
func ByImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByActImagesCount orders the results by act_images count.
func ByActImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newActImagesStep(), opts...)
	}
}

// ByActImages orders the results by act_images terms.
func ByActImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newActStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ActTable, ActColumn),
	)
}
func newImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ImagesTable, ImagesPrimaryKey...),
	)
}
func newRoutineActsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutineActsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutineActsTable, RoutineActsColumn),
	)
}
func newActImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ActImagesTable, ActImagesColumn),
	)
}
