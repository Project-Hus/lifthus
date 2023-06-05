// Code generated by ent, DO NOT EDIT.

package tag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeActs holds the string denoting the acts edge name in mutations.
	EdgeActs = "acts"
	// EdgePrograms holds the string denoting the programs edge name in mutations.
	EdgePrograms = "programs"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// ActsTable is the table that holds the acts relation/edge. The primary key declared below.
	ActsTable = "tag_acts"
	// ActsInverseTable is the table name for the Act entity.
	// It exists in this package in order to avoid circular dependency with the "act" package.
	ActsInverseTable = "acts"
	// ProgramsTable is the table that holds the programs relation/edge. The primary key declared below.
	ProgramsTable = "tag_programs"
	// ProgramsInverseTable is the table name for the Program entity.
	// It exists in this package in order to avoid circular dependency with the "program" package.
	ProgramsInverseTable = "programs"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ActsPrimaryKey and ActsColumn2 are the table columns denoting the
	// primary key for the acts relation (M2M).
	ActsPrimaryKey = []string{"tag_id", "act_id"}
	// ProgramsPrimaryKey and ProgramsColumn2 are the table columns denoting the
	// primary key for the programs relation (M2M).
	ProgramsPrimaryKey = []string{"tag_id", "program_id"}
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
)

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByActsCount orders the results by acts count.
func ByActsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newActsStep(), opts...)
	}
}

// ByActs orders the results by acts terms.
func ByActs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProgramsCount orders the results by programs count.
func ByProgramsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProgramsStep(), opts...)
	}
}

// ByPrograms orders the results by programs terms.
func ByPrograms(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newActsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ActsTable, ActsPrimaryKey...),
	)
}
func newProgramsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProgramsTable, ProgramsPrimaryKey...),
	)
}
