// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"routine/internal/ent/program"
	"routine/internal/ent/programversion"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ProgramVersion is the model entity for the ProgramVersion schema.
type ProgramVersion struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// ProgramCode holds the value of the "program_code" field.
	ProgramCode string `json:"program_code,omitempty"`
	// Version holds the value of the "version" field.
	Version uint `json:"version,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProgramVersionQuery when eager-loading is set.
	Edges                    ProgramVersionEdges `json:"edges"`
	program_program_versions *uint64
	selectValues             sql.SelectValues
}

// ProgramVersionEdges holds the relations/edges for other nodes in the graph.
type ProgramVersionEdges struct {
	// Program holds the value of the program edge.
	Program *Program `json:"program,omitempty"`
	// Images holds the value of the images edge.
	Images []*Image `json:"images,omitempty"`
	// DailyRoutines holds the value of the daily_routines edge.
	DailyRoutines []*DailyRoutine `json:"daily_routines,omitempty"`
	// ProgramImages holds the value of the program_images edge.
	ProgramImages []*ProgramImage `json:"program_images,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProgramOrErr returns the Program value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProgramVersionEdges) ProgramOrErr() (*Program, error) {
	if e.loadedTypes[0] {
		if e.Program == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: program.Label}
		}
		return e.Program, nil
	}
	return nil, &NotLoadedError{edge: "program"}
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramVersionEdges) ImagesOrErr() ([]*Image, error) {
	if e.loadedTypes[1] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// DailyRoutinesOrErr returns the DailyRoutines value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramVersionEdges) DailyRoutinesOrErr() ([]*DailyRoutine, error) {
	if e.loadedTypes[2] {
		return e.DailyRoutines, nil
	}
	return nil, &NotLoadedError{edge: "daily_routines"}
}

// ProgramImagesOrErr returns the ProgramImages value or an error if the edge
// was not loaded in eager-loading.
func (e ProgramVersionEdges) ProgramImagesOrErr() ([]*ProgramImage, error) {
	if e.loadedTypes[3] {
		return e.ProgramImages, nil
	}
	return nil, &NotLoadedError{edge: "program_images"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProgramVersion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case programversion.FieldID, programversion.FieldVersion:
			values[i] = new(sql.NullInt64)
		case programversion.FieldCode, programversion.FieldProgramCode, programversion.FieldText:
			values[i] = new(sql.NullString)
		case programversion.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case programversion.ForeignKeys[0]: // program_program_versions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProgramVersion fields.
func (pv *ProgramVersion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case programversion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pv.ID = uint64(value.Int64)
		case programversion.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				pv.Code = value.String
			}
		case programversion.FieldProgramCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field program_code", values[i])
			} else if value.Valid {
				pv.ProgramCode = value.String
			}
		case programversion.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				pv.Version = uint(value.Int64)
			}
		case programversion.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pv.CreatedAt = value.Time
			}
		case programversion.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				pv.Text = value.String
			}
		case programversion.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field program_program_versions", value)
			} else if value.Valid {
				pv.program_program_versions = new(uint64)
				*pv.program_program_versions = uint64(value.Int64)
			}
		default:
			pv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProgramVersion.
// This includes values selected through modifiers, order, etc.
func (pv *ProgramVersion) Value(name string) (ent.Value, error) {
	return pv.selectValues.Get(name)
}

// QueryProgram queries the "program" edge of the ProgramVersion entity.
func (pv *ProgramVersion) QueryProgram() *ProgramQuery {
	return NewProgramVersionClient(pv.config).QueryProgram(pv)
}

// QueryImages queries the "images" edge of the ProgramVersion entity.
func (pv *ProgramVersion) QueryImages() *ImageQuery {
	return NewProgramVersionClient(pv.config).QueryImages(pv)
}

// QueryDailyRoutines queries the "daily_routines" edge of the ProgramVersion entity.
func (pv *ProgramVersion) QueryDailyRoutines() *DailyRoutineQuery {
	return NewProgramVersionClient(pv.config).QueryDailyRoutines(pv)
}

// QueryProgramImages queries the "program_images" edge of the ProgramVersion entity.
func (pv *ProgramVersion) QueryProgramImages() *ProgramImageQuery {
	return NewProgramVersionClient(pv.config).QueryProgramImages(pv)
}

// Update returns a builder for updating this ProgramVersion.
// Note that you need to call ProgramVersion.Unwrap() before calling this method if this ProgramVersion
// was returned from a transaction, and the transaction was committed or rolled back.
func (pv *ProgramVersion) Update() *ProgramVersionUpdateOne {
	return NewProgramVersionClient(pv.config).UpdateOne(pv)
}

// Unwrap unwraps the ProgramVersion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pv *ProgramVersion) Unwrap() *ProgramVersion {
	_tx, ok := pv.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProgramVersion is not a transactional entity")
	}
	pv.config.driver = _tx.drv
	return pv
}

// String implements the fmt.Stringer.
func (pv *ProgramVersion) String() string {
	var builder strings.Builder
	builder.WriteString("ProgramVersion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pv.ID))
	builder.WriteString("code=")
	builder.WriteString(pv.Code)
	builder.WriteString(", ")
	builder.WriteString("program_code=")
	builder.WriteString(pv.ProgramCode)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", pv.Version))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pv.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(pv.Text)
	builder.WriteByte(')')
	return builder.String()
}

// ProgramVersions is a parsable slice of ProgramVersion.
type ProgramVersions []*ProgramVersion