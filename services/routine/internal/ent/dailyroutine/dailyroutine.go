// Code generated by ent, DO NOT EDIT.

package dailyroutine

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the dailyroutine type in the database.
	Label = "daily_routine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldProgramVersionCode holds the string denoting the program_version_code field in the database.
	FieldProgramVersionCode = "program_version_code"
	// FieldDay holds the string denoting the day field in the database.
	FieldDay = "day"
	// EdgeProgramVersion holds the string denoting the program_version edge name in mutations.
	EdgeProgramVersion = "program_version"
	// EdgeRoutineActs holds the string denoting the routine_acts edge name in mutations.
	EdgeRoutineActs = "routine_acts"
	// Table holds the table name of the dailyroutine in the database.
	Table = "daily_routines"
	// ProgramVersionTable is the table that holds the program_version relation/edge.
	ProgramVersionTable = "daily_routines"
	// ProgramVersionInverseTable is the table name for the ProgramVersion entity.
	// It exists in this package in order to avoid circular dependency with the "programversion" package.
	ProgramVersionInverseTable = "program_versions"
	// ProgramVersionColumn is the table column denoting the program_version relation/edge.
	ProgramVersionColumn = "program_version_daily_routines"
	// RoutineActsTable is the table that holds the routine_acts relation/edge.
	RoutineActsTable = "routine_acts"
	// RoutineActsInverseTable is the table name for the RoutineAct entity.
	// It exists in this package in order to avoid circular dependency with the "routineact" package.
	RoutineActsInverseTable = "routine_acts"
	// RoutineActsColumn is the table column denoting the routine_acts relation/edge.
	RoutineActsColumn = "daily_routine_routine_acts"
)

// Columns holds all SQL columns for dailyroutine fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldProgramVersionCode,
	FieldDay,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "daily_routines"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"program_version_daily_routines",
}

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
	// ProgramVersionCodeValidator is a validator for the "program_version_code" field. It is called by the builders before save.
	ProgramVersionCodeValidator func(string) error
)

// OrderOption defines the ordering options for the DailyRoutine queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByProgramVersionCode orders the results by the program_version_code field.
func ByProgramVersionCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProgramVersionCode, opts...).ToFunc()
}

// ByDay orders the results by the day field.
func ByDay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDay, opts...).ToFunc()
}

// ByProgramVersionField orders the results by program_version field.
func ByProgramVersionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramVersionStep(), sql.OrderByField(field, opts...))
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
func newProgramVersionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramVersionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProgramVersionTable, ProgramVersionColumn),
	)
}
func newRoutineActsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutineActsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutineActsTable, RoutineActsColumn),
	)
}