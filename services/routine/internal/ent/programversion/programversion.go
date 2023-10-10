// Code generated by ent, DO NOT EDIT.

package programversion

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the programversion type in the database.
	Label = "program_version"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldProgramCode holds the string denoting the program_code field in the database.
	FieldProgramCode = "program_code"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// EdgeProgram holds the string denoting the program edge name in mutations.
	EdgeProgram = "program"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// EdgeDailyRoutines holds the string denoting the daily_routines edge name in mutations.
	EdgeDailyRoutines = "daily_routines"
	// EdgeProgramImages holds the string denoting the program_images edge name in mutations.
	EdgeProgramImages = "program_images"
	// Table holds the table name of the programversion in the database.
	Table = "program_versions"
	// ProgramTable is the table that holds the program relation/edge.
	ProgramTable = "program_versions"
	// ProgramInverseTable is the table name for the Program entity.
	// It exists in this package in order to avoid circular dependency with the "program" package.
	ProgramInverseTable = "programs"
	// ProgramColumn is the table column denoting the program relation/edge.
	ProgramColumn = "program_program_versions"
	// ImagesTable is the table that holds the images relation/edge. The primary key declared below.
	ImagesTable = "program_images"
	// ImagesInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImagesInverseTable = "images"
	// DailyRoutinesTable is the table that holds the daily_routines relation/edge.
	DailyRoutinesTable = "daily_routines"
	// DailyRoutinesInverseTable is the table name for the DailyRoutine entity.
	// It exists in this package in order to avoid circular dependency with the "dailyroutine" package.
	DailyRoutinesInverseTable = "daily_routines"
	// DailyRoutinesColumn is the table column denoting the daily_routines relation/edge.
	DailyRoutinesColumn = "program_version_daily_routines"
	// ProgramImagesTable is the table that holds the program_images relation/edge.
	ProgramImagesTable = "program_images"
	// ProgramImagesInverseTable is the table name for the ProgramImage entity.
	// It exists in this package in order to avoid circular dependency with the "programimage" package.
	ProgramImagesInverseTable = "program_images"
	// ProgramImagesColumn is the table column denoting the program_images relation/edge.
	ProgramImagesColumn = "program_version_id"
)

// Columns holds all SQL columns for programversion fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldProgramCode,
	FieldVersion,
	FieldCreatedAt,
	FieldText,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "program_versions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"program_program_versions",
}

var (
	// ImagesPrimaryKey and ImagesColumn2 are the table columns denoting the
	// primary key for the images relation (M2M).
	ImagesPrimaryKey = []string{"program_version_id", "image_id"}
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
	// ProgramCodeValidator is a validator for the "program_code" field. It is called by the builders before save.
	ProgramCodeValidator func(string) error
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
)

// OrderOption defines the ordering options for the ProgramVersion queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByProgramCode orders the results by the program_code field.
func ByProgramCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProgramCode, opts...).ToFunc()
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

// ByProgramField orders the results by program field.
func ByProgramField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramStep(), sql.OrderByField(field, opts...))
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

// ByDailyRoutinesCount orders the results by daily_routines count.
func ByDailyRoutinesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDailyRoutinesStep(), opts...)
	}
}

// ByDailyRoutines orders the results by daily_routines terms.
func ByDailyRoutines(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDailyRoutinesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProgramImagesCount orders the results by program_images count.
func ByProgramImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProgramImagesStep(), opts...)
	}
}

// ByProgramImages orders the results by program_images terms.
func ByProgramImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProgramStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProgramTable, ProgramColumn),
	)
}
func newImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ImagesTable, ImagesPrimaryKey...),
	)
}
func newDailyRoutinesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DailyRoutinesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DailyRoutinesTable, DailyRoutinesColumn),
	)
}
func newProgramImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ProgramImagesTable, ProgramImagesColumn),
	)
}
