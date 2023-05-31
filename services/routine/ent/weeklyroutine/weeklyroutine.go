// Code generated by ent, DO NOT EDIT.

package weeklyroutine

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the weeklyroutine type in the database.
	Label = "weekly_routine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProgramID holds the string denoting the program_id field in the database.
	FieldProgramID = "program_id"
	// FieldWeek holds the string denoting the week field in the database.
	FieldWeek = "week"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProgram holds the string denoting the program edge name in mutations.
	EdgeProgram = "program"
	// EdgeDailyRoutines holds the string denoting the daily_routines edge name in mutations.
	EdgeDailyRoutines = "daily_routines"
	// Table holds the table name of the weeklyroutine in the database.
	Table = "weekly_routines"
	// ProgramTable is the table that holds the program relation/edge. The primary key declared below.
	ProgramTable = "program_weekly_routines"
	// ProgramInverseTable is the table name for the Program entity.
	// It exists in this package in order to avoid circular dependency with the "program" package.
	ProgramInverseTable = "programs"
	// DailyRoutinesTable is the table that holds the daily_routines relation/edge. The primary key declared below.
	DailyRoutinesTable = "weekly_routine_daily_routines"
	// DailyRoutinesInverseTable is the table name for the DailyRoutine entity.
	// It exists in this package in order to avoid circular dependency with the "dailyroutine" package.
	DailyRoutinesInverseTable = "daily_routines"
)

// Columns holds all SQL columns for weeklyroutine fields.
var Columns = []string{
	FieldID,
	FieldProgramID,
	FieldWeek,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// ProgramPrimaryKey and ProgramColumn2 are the table columns denoting the
	// primary key for the program relation (M2M).
	ProgramPrimaryKey = []string{"program_id", "weekly_routine_id"}
	// DailyRoutinesPrimaryKey and DailyRoutinesColumn2 are the table columns denoting the
	// primary key for the daily_routines relation (M2M).
	DailyRoutinesPrimaryKey = []string{"weekly_routine_id", "daily_routine_id"}
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
	// WeekValidator is a validator for the "week" field. It is called by the builders before save.
	WeekValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the WeeklyRoutine queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProgramID orders the results by the program_id field.
func ByProgramID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProgramID, opts...).ToFunc()
}

// ByWeek orders the results by the week field.
func ByWeek(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeek, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByProgramCount orders the results by program count.
func ByProgramCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProgramStep(), opts...)
	}
}

// ByProgram orders the results by program terms.
func ByProgram(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newProgramStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProgramTable, ProgramPrimaryKey...),
	)
}
func newDailyRoutinesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DailyRoutinesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, DailyRoutinesTable, DailyRoutinesPrimaryKey...),
	)
}
