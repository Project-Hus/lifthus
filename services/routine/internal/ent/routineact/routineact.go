// Code generated by ent, DO NOT EDIT.

package routineact

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the routineact type in the database.
	Label = "routine_act"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldActCode holds the string denoting the act_code field in the database.
	FieldActCode = "act_code"
	// FieldStage holds the string denoting the stage field in the database.
	FieldStage = "stage"
	// FieldRepsOrMeters holds the string denoting the reps_or_meters field in the database.
	FieldRepsOrMeters = "reps_or_meters"
	// FieldRatioOrSecs holds the string denoting the ratio_or_secs field in the database.
	FieldRatioOrSecs = "ratio_or_secs"
	// EdgeAct holds the string denoting the act edge name in mutations.
	EdgeAct = "act"
	// EdgeDayRoutine holds the string denoting the day_routine edge name in mutations.
	EdgeDayRoutine = "day_routine"
	// Table holds the table name of the routineact in the database.
	Table = "routine_acts"
	// ActTable is the table that holds the act relation/edge.
	ActTable = "routine_acts"
	// ActInverseTable is the table name for the Act entity.
	// It exists in this package in order to avoid circular dependency with the "act" package.
	ActInverseTable = "acts"
	// ActColumn is the table column denoting the act relation/edge.
	ActColumn = "act_routine_acts"
	// DayRoutineTable is the table that holds the day_routine relation/edge.
	DayRoutineTable = "routine_acts"
	// DayRoutineInverseTable is the table name for the DayRoutine entity.
	// It exists in this package in order to avoid circular dependency with the "dayroutine" package.
	DayRoutineInverseTable = "day_routines"
	// DayRoutineColumn is the table column denoting the day_routine relation/edge.
	DayRoutineColumn = "day_routine_routine_acts"
)

// Columns holds all SQL columns for routineact fields.
var Columns = []string{
	FieldID,
	FieldOrder,
	FieldActCode,
	FieldStage,
	FieldRepsOrMeters,
	FieldRatioOrSecs,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "routine_acts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"act_routine_acts",
	"day_routine_routine_acts",
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
	// ActCodeValidator is a validator for the "act_code" field. It is called by the builders before save.
	ActCodeValidator func(string) error
)

// Stage defines the type for the "stage" enum field.
type Stage string

// Stage values.
const (
	StageWarmup   Stage = "warmup"
	StageMain     Stage = "main"
	StageCooldown Stage = "cooldown"
)

func (s Stage) String() string {
	return string(s)
}

// StageValidator is a validator for the "stage" field enum values. It is called by the builders before save.
func StageValidator(s Stage) error {
	switch s {
	case StageWarmup, StageMain, StageCooldown:
		return nil
	default:
		return fmt.Errorf("routineact: invalid enum value for stage field: %q", s)
	}
}

// OrderOption defines the ordering options for the RoutineAct queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByActCode orders the results by the act_code field.
func ByActCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActCode, opts...).ToFunc()
}

// ByStage orders the results by the stage field.
func ByStage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStage, opts...).ToFunc()
}

// ByRepsOrMeters orders the results by the reps_or_meters field.
func ByRepsOrMeters(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRepsOrMeters, opts...).ToFunc()
}

// ByRatioOrSecs orders the results by the ratio_or_secs field.
func ByRatioOrSecs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRatioOrSecs, opts...).ToFunc()
}

// ByActField orders the results by act field.
func ByActField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActStep(), sql.OrderByField(field, opts...))
	}
}

// ByDayRoutineField orders the results by day_routine field.
func ByDayRoutineField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDayRoutineStep(), sql.OrderByField(field, opts...))
	}
}
func newActStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ActTable, ActColumn),
	)
}
func newDayRoutineStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DayRoutineInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DayRoutineTable, DayRoutineColumn),
	)
}
