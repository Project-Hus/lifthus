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
	// FieldDailyRoutineCode holds the string denoting the daily_routine_code field in the database.
	FieldDailyRoutineCode = "daily_routine_code"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldActVersion holds the string denoting the act_version field in the database.
	FieldActVersion = "act_version"
	// FieldStage holds the string denoting the stage field in the database.
	FieldStage = "stage"
	// FieldRepsOrMeters holds the string denoting the reps_or_meters field in the database.
	FieldRepsOrMeters = "reps_or_meters"
	// FieldRatioOrSecs holds the string denoting the ratio_or_secs field in the database.
	FieldRatioOrSecs = "ratio_or_secs"
	// EdgeDailyRoutine holds the string denoting the daily_routine edge name in mutations.
	EdgeDailyRoutine = "daily_routine"
	// Table holds the table name of the routineact in the database.
	Table = "routine_acts"
	// DailyRoutineTable is the table that holds the daily_routine relation/edge.
	DailyRoutineTable = "routine_acts"
	// DailyRoutineInverseTable is the table name for the DailyRoutine entity.
	// It exists in this package in order to avoid circular dependency with the "dailyroutine" package.
	DailyRoutineInverseTable = "daily_routines"
	// DailyRoutineColumn is the table column denoting the daily_routine relation/edge.
	DailyRoutineColumn = "daily_routine_routine_acts"
)

// Columns holds all SQL columns for routineact fields.
var Columns = []string{
	FieldID,
	FieldDailyRoutineCode,
	FieldOrder,
	FieldActVersion,
	FieldStage,
	FieldRepsOrMeters,
	FieldRatioOrSecs,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "routine_acts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"daily_routine_routine_acts",
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
	// DailyRoutineCodeValidator is a validator for the "daily_routine_code" field. It is called by the builders before save.
	DailyRoutineCodeValidator func(string) error
	// ActVersionValidator is a validator for the "act_version" field. It is called by the builders before save.
	ActVersionValidator func(string) error
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

// ByDailyRoutineCode orders the results by the daily_routine_code field.
func ByDailyRoutineCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDailyRoutineCode, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByActVersion orders the results by the act_version field.
func ByActVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActVersion, opts...).ToFunc()
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

// ByDailyRoutineField orders the results by daily_routine field.
func ByDailyRoutineField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDailyRoutineStep(), sql.OrderByField(field, opts...))
	}
}
func newDailyRoutineStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DailyRoutineInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DailyRoutineTable, DailyRoutineColumn),
	)
}
