// Code generated by ent, DO NOT EDIT.

package bodyinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the bodyinfo type in the database.
	Label = "body_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldProgramRecID holds the string denoting the program_rec_id field in the database.
	FieldProgramRecID = "program_rec_id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldBodyWeight holds the string denoting the body_weight field in the database.
	FieldBodyWeight = "body_weight"
	// FieldBodyFatMass holds the string denoting the body_fat_mass field in the database.
	FieldBodyFatMass = "body_fat_mass"
	// FieldSkeletalMuscleMass holds the string denoting the skeletal_muscle_mass field in the database.
	FieldSkeletalMuscleMass = "skeletal_muscle_mass"
	// EdgeProgramRec holds the string denoting the program_rec edge name in mutations.
	EdgeProgramRec = "program_rec"
	// Table holds the table name of the bodyinfo in the database.
	Table = "body_infos"
	// ProgramRecTable is the table that holds the program_rec relation/edge.
	ProgramRecTable = "body_infos"
	// ProgramRecInverseTable is the table name for the ProgramRec entity.
	// It exists in this package in order to avoid circular dependency with the "programrec" package.
	ProgramRecInverseTable = "program_recs"
	// ProgramRecColumn is the table column denoting the program_rec relation/edge.
	ProgramRecColumn = "program_rec_id"
)

// Columns holds all SQL columns for bodyinfo fields.
var Columns = []string{
	FieldID,
	FieldAuthor,
	FieldProgramRecID,
	FieldDate,
	FieldHeight,
	FieldBodyWeight,
	FieldBodyFatMass,
	FieldSkeletalMuscleMass,
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
	// DefaultDate holds the default value on creation for the "date" field.
	DefaultDate func() time.Time
)

// OrderOption defines the ordering options for the BodyInfo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByProgramRecID orders the results by the program_rec_id field.
func ByProgramRecID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProgramRecID, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByBodyWeight orders the results by the body_weight field.
func ByBodyWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBodyWeight, opts...).ToFunc()
}

// ByBodyFatMass orders the results by the body_fat_mass field.
func ByBodyFatMass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBodyFatMass, opts...).ToFunc()
}

// BySkeletalMuscleMass orders the results by the skeletal_muscle_mass field.
func BySkeletalMuscleMass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSkeletalMuscleMass, opts...).ToFunc()
}

// ByProgramRecField orders the results by program_rec field.
func ByProgramRecField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramRecStep(), sql.OrderByField(field, opts...))
	}
}
func newProgramRecStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramRecInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProgramRecTable, ProgramRecColumn),
	)
}