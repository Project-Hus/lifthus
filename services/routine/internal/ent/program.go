// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"routine/internal/ent/program"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Program is the model entity for the Program schema.
type Program struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// ProgramType holds the value of the "program_type" field.
	ProgramType program.ProgramType `json:"program_type,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Author holds the value of the "author" field.
	Author uint64 `json:"author,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// VersionDerivedFrom holds the value of the "version_derived_from" field.
	VersionDerivedFrom *string `json:"version_derived_from,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProgramQuery when eager-loading is set.
	Edges        ProgramEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProgramEdges holds the relations/edges for other nodes in the graph.
type ProgramEdges struct {
	// ProgramVersions holds the value of the program_versions edge.
	ProgramVersions []*ProgramVersion `json:"program_versions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProgramVersionsOrErr returns the ProgramVersions value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramEdges) ProgramVersionsOrErr() ([]*ProgramVersion, error) {
	if e.loadedTypes[0] {
		return e.ProgramVersions, nil
	}
	return nil, &NotLoadedError{edge: "program_versions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Program) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case program.FieldID, program.FieldAuthor:
			values[i] = new(sql.NullInt64)
		case program.FieldCode, program.FieldProgramType, program.FieldTitle, program.FieldVersionDerivedFrom:
			values[i] = new(sql.NullString)
		case program.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Program fields.
func (pr *Program) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case program.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = uint64(value.Int64)
		case program.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				pr.Code = value.String
			}
		case program.FieldProgramType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field program_type", values[i])
			} else if value.Valid {
				pr.ProgramType = program.ProgramType(value.String)
			}
		case program.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pr.Title = value.String
			}
		case program.FieldAuthor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				pr.Author = uint64(value.Int64)
			}
		case program.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case program.FieldVersionDerivedFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version_derived_from", values[i])
			} else if value.Valid {
				pr.VersionDerivedFrom = new(string)
				*pr.VersionDerivedFrom = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Program.
// This includes values selected through modifiers, order, etc.
func (pr *Program) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryProgramVersions queries the "program_versions" edge of the Program entity.
func (pr *Program) QueryProgramVersions() *ProgramVersionQuery {
	return NewProgramClient(pr.config).QueryProgramVersions(pr)
}

// Update returns a builder for updating this Program.
// Note that you need to call Program.Unwrap() before calling this method if this Program
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Program) Update() *ProgramUpdateOne {
	return NewProgramClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Program entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Program) Unwrap() *Program {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Program is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Program) String() string {
	var builder strings.Builder
	builder.WriteString("Program(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("code=")
	builder.WriteString(pr.Code)
	builder.WriteString(", ")
	builder.WriteString("program_type=")
	builder.WriteString(fmt.Sprintf("%v", pr.ProgramType))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pr.Title)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(fmt.Sprintf("%v", pr.Author))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pr.VersionDerivedFrom; v != nil {
		builder.WriteString("version_derived_from=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Programs is a parsable slice of Program.
type Programs []*Program