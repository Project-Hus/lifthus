// Code generated by ent, DO NOT EDIT.

package program

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the program type in the database.
	Label = "program"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldProgramType holds the string denoting the program_type field in the database.
	FieldProgramType = "program_type"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldVersionDerivedFrom holds the string denoting the version_derived_from field in the database.
	FieldVersionDerivedFrom = "version_derived_from"
	// EdgeProgramVersions holds the string denoting the program_versions edge name in mutations.
	EdgeProgramVersions = "program_versions"
	// Table holds the table name of the program in the database.
	Table = "programs"
	// ProgramVersionsTable is the table that holds the program_versions relation/edge.
	ProgramVersionsTable = "program_versions"
	// ProgramVersionsInverseTable is the table name for the ProgramVersion entity.
	// It exists in this package in order to avoid circular dependency with the "programversion" package.
	ProgramVersionsInverseTable = "program_versions"
	// ProgramVersionsColumn is the table column denoting the program_versions relation/edge.
	ProgramVersionsColumn = "program_program_versions"
)

// Columns holds all SQL columns for program fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldProgramType,
	FieldTitle,
	FieldAuthor,
	FieldCreatedAt,
	FieldVersionDerivedFrom,
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// VersionDerivedFromValidator is a validator for the "version_derived_from" field. It is called by the builders before save.
	VersionDerivedFromValidator func(string) error
)

// ProgramType defines the type for the "program_type" enum field.
type ProgramType string

// ProgramType values.
const (
	ProgramTypeWeekly ProgramType = "weekly"
	ProgramTypeDaily  ProgramType = "daily"
)

func (pt ProgramType) String() string {
	return string(pt)
}

// ProgramTypeValidator is a validator for the "program_type" field enum values. It is called by the builders before save.
func ProgramTypeValidator(pt ProgramType) error {
	switch pt {
	case ProgramTypeWeekly, ProgramTypeDaily:
		return nil
	default:
		return fmt.Errorf("program: invalid enum value for program_type field: %q", pt)
	}
}

// OrderOption defines the ordering options for the Program queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByProgramType orders the results by the program_type field.
func ByProgramType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProgramType, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByVersionDerivedFrom orders the results by the version_derived_from field.
func ByVersionDerivedFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersionDerivedFrom, opts...).ToFunc()
}

// ByProgramVersionsCount orders the results by program_versions count.
func ByProgramVersionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProgramVersionsStep(), opts...)
	}
}

// ByProgramVersions orders the results by program_versions terms.
func ByProgramVersions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramVersionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProgramVersionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramVersionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProgramVersionsTable, ProgramVersionsColumn),
	)
}