// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"routine/internal/ent/act"
	"routine/internal/ent/dayroutine"
	"routine/internal/ent/routineact"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// RoutineAct is the model entity for the RoutineAct schema.
type RoutineAct struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// ActCode holds the value of the "act_code" field.
	ActCode string `json:"act_code,omitempty"`
	// Stage holds the value of the "stage" field.
	Stage routineact.Stage `json:"stage,omitempty"`
	// RepsOrMeters holds the value of the "reps_or_meters" field.
	RepsOrMeters uint `json:"reps_or_meters,omitempty"`
	// RatioOrSecs holds the value of the "ratio_or_secs" field.
	RatioOrSecs float64 `json:"ratio_or_secs,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoutineActQuery when eager-loading is set.
	Edges                    RoutineActEdges `json:"edges"`
	act_routine_acts         *int64
	day_routine_routine_acts *int64
	selectValues             sql.SelectValues
}

// RoutineActEdges holds the relations/edges for other nodes in the graph.
type RoutineActEdges struct {
	// Act holds the value of the act edge.
	Act *Act `json:"act,omitempty"`
	// DayRoutine holds the value of the day_routine edge.
	DayRoutine *DayRoutine `json:"day_routine,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ActOrErr returns the Act value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoutineActEdges) ActOrErr() (*Act, error) {
	if e.loadedTypes[0] {
		if e.Act == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: act.Label}
		}
		return e.Act, nil
	}
	return nil, &NotLoadedError{edge: "act"}
}

// DayRoutineOrErr returns the DayRoutine value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoutineActEdges) DayRoutineOrErr() (*DayRoutine, error) {
	if e.loadedTypes[1] {
		if e.DayRoutine == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: dayroutine.Label}
		}
		return e.DayRoutine, nil
	}
	return nil, &NotLoadedError{edge: "day_routine"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoutineAct) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case routineact.FieldRatioOrSecs:
			values[i] = new(sql.NullFloat64)
		case routineact.FieldID, routineact.FieldOrder, routineact.FieldRepsOrMeters:
			values[i] = new(sql.NullInt64)
		case routineact.FieldActCode, routineact.FieldStage:
			values[i] = new(sql.NullString)
		case routineact.ForeignKeys[0]: // act_routine_acts
			values[i] = new(sql.NullInt64)
		case routineact.ForeignKeys[1]: // day_routine_routine_acts
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoutineAct fields.
func (ra *RoutineAct) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case routineact.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ra.ID = int64(value.Int64)
		case routineact.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				ra.Order = int(value.Int64)
			}
		case routineact.FieldActCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field act_code", values[i])
			} else if value.Valid {
				ra.ActCode = value.String
			}
		case routineact.FieldStage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stage", values[i])
			} else if value.Valid {
				ra.Stage = routineact.Stage(value.String)
			}
		case routineact.FieldRepsOrMeters:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reps_or_meters", values[i])
			} else if value.Valid {
				ra.RepsOrMeters = uint(value.Int64)
			}
		case routineact.FieldRatioOrSecs:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ratio_or_secs", values[i])
			} else if value.Valid {
				ra.RatioOrSecs = value.Float64
			}
		case routineact.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field act_routine_acts", value)
			} else if value.Valid {
				ra.act_routine_acts = new(int64)
				*ra.act_routine_acts = int64(value.Int64)
			}
		case routineact.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field day_routine_routine_acts", value)
			} else if value.Valid {
				ra.day_routine_routine_acts = new(int64)
				*ra.day_routine_routine_acts = int64(value.Int64)
			}
		default:
			ra.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoutineAct.
// This includes values selected through modifiers, order, etc.
func (ra *RoutineAct) Value(name string) (ent.Value, error) {
	return ra.selectValues.Get(name)
}

// QueryAct queries the "act" edge of the RoutineAct entity.
func (ra *RoutineAct) QueryAct() *ActQuery {
	return NewRoutineActClient(ra.config).QueryAct(ra)
}

// QueryDayRoutine queries the "day_routine" edge of the RoutineAct entity.
func (ra *RoutineAct) QueryDayRoutine() *DayRoutineQuery {
	return NewRoutineActClient(ra.config).QueryDayRoutine(ra)
}

// Update returns a builder for updating this RoutineAct.
// Note that you need to call RoutineAct.Unwrap() before calling this method if this RoutineAct
// was returned from a transaction, and the transaction was committed or rolled back.
func (ra *RoutineAct) Update() *RoutineActUpdateOne {
	return NewRoutineActClient(ra.config).UpdateOne(ra)
}

// Unwrap unwraps the RoutineAct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ra *RoutineAct) Unwrap() *RoutineAct {
	_tx, ok := ra.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoutineAct is not a transactional entity")
	}
	ra.config.driver = _tx.drv
	return ra
}

// String implements the fmt.Stringer.
func (ra *RoutineAct) String() string {
	var builder strings.Builder
	builder.WriteString("RoutineAct(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ra.ID))
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", ra.Order))
	builder.WriteString(", ")
	builder.WriteString("act_code=")
	builder.WriteString(ra.ActCode)
	builder.WriteString(", ")
	builder.WriteString("stage=")
	builder.WriteString(fmt.Sprintf("%v", ra.Stage))
	builder.WriteString(", ")
	builder.WriteString("reps_or_meters=")
	builder.WriteString(fmt.Sprintf("%v", ra.RepsOrMeters))
	builder.WriteString(", ")
	builder.WriteString("ratio_or_secs=")
	builder.WriteString(fmt.Sprintf("%v", ra.RatioOrSecs))
	builder.WriteByte(')')
	return builder.String()
}

// RoutineActs is a parsable slice of RoutineAct.
type RoutineActs []*RoutineAct
