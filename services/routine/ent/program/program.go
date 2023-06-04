// Code generated by ent, DO NOT EDIT.

package program

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the program type in the database.
	Label = "program"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeWeeklyRoutines holds the string denoting the weekly_routines edge name in mutations.
	EdgeWeeklyRoutines = "weekly_routines"
	// EdgeDailyRoutines holds the string denoting the daily_routines edge name in mutations.
	EdgeDailyRoutines = "daily_routines"
	// EdgeProgramRecs holds the string denoting the program_recs edge name in mutations.
	EdgeProgramRecs = "program_recs"
	// Table holds the table name of the program in the database.
	Table = "programs"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "tag_programs"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// WeeklyRoutinesTable is the table that holds the weekly_routines relation/edge.
	WeeklyRoutinesTable = "weekly_routines"
	// WeeklyRoutinesInverseTable is the table name for the WeeklyRoutine entity.
	// It exists in this package in order to avoid circular dependency with the "weeklyroutine" package.
	WeeklyRoutinesInverseTable = "weekly_routines"
	// WeeklyRoutinesColumn is the table column denoting the weekly_routines relation/edge.
	WeeklyRoutinesColumn = "program_id"
	// DailyRoutinesTable is the table that holds the daily_routines relation/edge.
	DailyRoutinesTable = "daily_routines"
	// DailyRoutinesInverseTable is the table name for the DailyRoutine entity.
	// It exists in this package in order to avoid circular dependency with the "dailyroutine" package.
	DailyRoutinesInverseTable = "daily_routines"
	// DailyRoutinesColumn is the table column denoting the daily_routines relation/edge.
	DailyRoutinesColumn = "program_id"
	// ProgramRecsTable is the table that holds the program_recs relation/edge.
	ProgramRecsTable = "program_recs"
	// ProgramRecsInverseTable is the table name for the ProgramRec entity.
	// It exists in this package in order to avoid circular dependency with the "programrec" package.
	ProgramRecsInverseTable = "program_recs"
	// ProgramRecsColumn is the table column denoting the program_recs relation/edge.
	ProgramRecsColumn = "program_id"
)

// Columns holds all SQL columns for program fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldSlug,
	FieldType,
	FieldAuthor,
	FieldImage,
	FieldDescription,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"tag_id", "program_id"}
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeWeekly Type = "weekly"
	TypeDaily  Type = "daily"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeWeekly, TypeDaily:
		return nil
	default:
		return fmt.Errorf("program: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Program queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWeeklyRoutinesCount orders the results by weekly_routines count.
func ByWeeklyRoutinesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWeeklyRoutinesStep(), opts...)
	}
}

// ByWeeklyRoutines orders the results by weekly_routines terms.
func ByWeeklyRoutines(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWeeklyRoutinesStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByProgramRecsCount orders the results by program_recs count.
func ByProgramRecsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProgramRecsStep(), opts...)
	}
}

// ByProgramRecs orders the results by program_recs terms.
func ByProgramRecs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramRecsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
	)
}
func newWeeklyRoutinesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WeeklyRoutinesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WeeklyRoutinesTable, WeeklyRoutinesColumn),
	)
}
func newDailyRoutinesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DailyRoutinesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DailyRoutinesTable, DailyRoutinesColumn),
	)
}
func newProgramRecsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramRecsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProgramRecsTable, ProgramRecsColumn),
	)
}
