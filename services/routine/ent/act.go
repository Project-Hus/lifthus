// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"routine/ent/act"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Act is the model entity for the Act schema.
type Act struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Type holds the value of the "type" field.
	Type act.Type `json:"type,omitempty"`
	// Author holds the value of the "author" field.
	Author uint64 `json:"author,omitempty"`
	// Image holds the value of the "image" field.
	Image *string `json:"image,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight bool `json:"weight,omitempty"`
	// Bodyweight holds the value of the "bodyweight" field.
	Bodyweight bool `json:"bodyweight,omitempty"`
	// Cardio holds the value of the "cardio" field.
	Cardio bool `json:"cardio,omitempty"`
	// Upper holds the value of the "upper" field.
	Upper bool `json:"upper,omitempty"`
	// Lower holds the value of the "lower" field.
	Lower bool `json:"lower,omitempty"`
	// Full holds the value of the "full" field.
	Full bool `json:"full,omitempty"`
	// Arms holds the value of the "arms" field.
	Arms bool `json:"arms,omitempty"`
	// Shoulders holds the value of the "shoulders" field.
	Shoulders bool `json:"shoulders,omitempty"`
	// Chest holds the value of the "chest" field.
	Chest bool `json:"chest,omitempty"`
	// Core holds the value of the "core" field.
	Core bool `json:"core,omitempty"`
	// UpperBack holds the value of the "upper_back" field.
	UpperBack bool `json:"upper_back,omitempty"`
	// LowerBack holds the value of the "lower_back" field.
	LowerBack bool `json:"lower_back,omitempty"`
	// Glute holds the value of the "glute" field.
	Glute bool `json:"glute,omitempty"`
	// LegsFront holds the value of the "legs_front" field.
	LegsFront bool `json:"legs_front,omitempty"`
	// LegsBack holds the value of the "legs_back" field.
	LegsBack bool `json:"legs_back,omitempty"`
	// Etc holds the value of the "etc" field.
	Etc bool `json:"etc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActQuery when eager-loading is set.
	Edges        ActEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ActEdges holds the relations/edges for other nodes in the graph.
type ActEdges struct {
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// RoutineActs holds the value of the routine_acts edge.
	RoutineActs []*RoutineAct `json:"routine_acts,omitempty"`
	// RoutineActRecs holds the value of the routine_act_recs edge.
	RoutineActRecs []*RoutineActRec `json:"routine_act_recs,omitempty"`
	// OneRepMaxes holds the value of the one_rep_maxes edge.
	OneRepMaxes []*OneRepMax `json:"one_rep_maxes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ActEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// RoutineActsOrErr returns the RoutineActs value or an error if the edge
// was not loaded in eager-loading.
func (e ActEdges) RoutineActsOrErr() ([]*RoutineAct, error) {
	if e.loadedTypes[1] {
		return e.RoutineActs, nil
	}
	return nil, &NotLoadedError{edge: "routine_acts"}
}

// RoutineActRecsOrErr returns the RoutineActRecs value or an error if the edge
// was not loaded in eager-loading.
func (e ActEdges) RoutineActRecsOrErr() ([]*RoutineActRec, error) {
	if e.loadedTypes[2] {
		return e.RoutineActRecs, nil
	}
	return nil, &NotLoadedError{edge: "routine_act_recs"}
}

// OneRepMaxesOrErr returns the OneRepMaxes value or an error if the edge
// was not loaded in eager-loading.
func (e ActEdges) OneRepMaxesOrErr() ([]*OneRepMax, error) {
	if e.loadedTypes[3] {
		return e.OneRepMaxes, nil
	}
	return nil, &NotLoadedError{edge: "one_rep_maxes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Act) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case act.FieldWeight, act.FieldBodyweight, act.FieldCardio, act.FieldUpper, act.FieldLower, act.FieldFull, act.FieldArms, act.FieldShoulders, act.FieldChest, act.FieldCore, act.FieldUpperBack, act.FieldLowerBack, act.FieldGlute, act.FieldLegsFront, act.FieldLegsBack, act.FieldEtc:
			values[i] = new(sql.NullBool)
		case act.FieldID, act.FieldAuthor:
			values[i] = new(sql.NullInt64)
		case act.FieldName, act.FieldSlug, act.FieldType, act.FieldImage, act.FieldDescription:
			values[i] = new(sql.NullString)
		case act.FieldCreatedAt, act.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Act fields.
func (a *Act) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case act.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case act.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case act.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				a.Slug = value.String
			}
		case act.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = act.Type(value.String)
			}
		case act.FieldAuthor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				a.Author = uint64(value.Int64)
			}
		case act.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				a.Image = new(string)
				*a.Image = value.String
			}
		case act.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = new(string)
				*a.Description = value.String
			}
		case act.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case act.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case act.FieldWeight:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				a.Weight = value.Bool
			}
		case act.FieldBodyweight:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field bodyweight", values[i])
			} else if value.Valid {
				a.Bodyweight = value.Bool
			}
		case act.FieldCardio:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field cardio", values[i])
			} else if value.Valid {
				a.Cardio = value.Bool
			}
		case act.FieldUpper:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field upper", values[i])
			} else if value.Valid {
				a.Upper = value.Bool
			}
		case act.FieldLower:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field lower", values[i])
			} else if value.Valid {
				a.Lower = value.Bool
			}
		case act.FieldFull:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field full", values[i])
			} else if value.Valid {
				a.Full = value.Bool
			}
		case act.FieldArms:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field arms", values[i])
			} else if value.Valid {
				a.Arms = value.Bool
			}
		case act.FieldShoulders:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field shoulders", values[i])
			} else if value.Valid {
				a.Shoulders = value.Bool
			}
		case act.FieldChest:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field chest", values[i])
			} else if value.Valid {
				a.Chest = value.Bool
			}
		case act.FieldCore:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field core", values[i])
			} else if value.Valid {
				a.Core = value.Bool
			}
		case act.FieldUpperBack:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field upper_back", values[i])
			} else if value.Valid {
				a.UpperBack = value.Bool
			}
		case act.FieldLowerBack:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field lower_back", values[i])
			} else if value.Valid {
				a.LowerBack = value.Bool
			}
		case act.FieldGlute:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field glute", values[i])
			} else if value.Valid {
				a.Glute = value.Bool
			}
		case act.FieldLegsFront:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field legs_front", values[i])
			} else if value.Valid {
				a.LegsFront = value.Bool
			}
		case act.FieldLegsBack:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field legs_back", values[i])
			} else if value.Valid {
				a.LegsBack = value.Bool
			}
		case act.FieldEtc:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field etc", values[i])
			} else if value.Valid {
				a.Etc = value.Bool
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Act.
// This includes values selected through modifiers, order, etc.
func (a *Act) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryTags queries the "tags" edge of the Act entity.
func (a *Act) QueryTags() *TagQuery {
	return NewActClient(a.config).QueryTags(a)
}

// QueryRoutineActs queries the "routine_acts" edge of the Act entity.
func (a *Act) QueryRoutineActs() *RoutineActQuery {
	return NewActClient(a.config).QueryRoutineActs(a)
}

// QueryRoutineActRecs queries the "routine_act_recs" edge of the Act entity.
func (a *Act) QueryRoutineActRecs() *RoutineActRecQuery {
	return NewActClient(a.config).QueryRoutineActRecs(a)
}

// QueryOneRepMaxes queries the "one_rep_maxes" edge of the Act entity.
func (a *Act) QueryOneRepMaxes() *OneRepMaxQuery {
	return NewActClient(a.config).QueryOneRepMaxes(a)
}

// Update returns a builder for updating this Act.
// Note that you need to call Act.Unwrap() before calling this method if this Act
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Act) Update() *ActUpdateOne {
	return NewActClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Act entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Act) Unwrap() *Act {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Act is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Act) String() string {
	var builder strings.Builder
	builder.WriteString("Act(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(a.Slug)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(fmt.Sprintf("%v", a.Author))
	builder.WriteString(", ")
	if v := a.Image; v != nil {
		builder.WriteString("image=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", a.Weight))
	builder.WriteString(", ")
	builder.WriteString("bodyweight=")
	builder.WriteString(fmt.Sprintf("%v", a.Bodyweight))
	builder.WriteString(", ")
	builder.WriteString("cardio=")
	builder.WriteString(fmt.Sprintf("%v", a.Cardio))
	builder.WriteString(", ")
	builder.WriteString("upper=")
	builder.WriteString(fmt.Sprintf("%v", a.Upper))
	builder.WriteString(", ")
	builder.WriteString("lower=")
	builder.WriteString(fmt.Sprintf("%v", a.Lower))
	builder.WriteString(", ")
	builder.WriteString("full=")
	builder.WriteString(fmt.Sprintf("%v", a.Full))
	builder.WriteString(", ")
	builder.WriteString("arms=")
	builder.WriteString(fmt.Sprintf("%v", a.Arms))
	builder.WriteString(", ")
	builder.WriteString("shoulders=")
	builder.WriteString(fmt.Sprintf("%v", a.Shoulders))
	builder.WriteString(", ")
	builder.WriteString("chest=")
	builder.WriteString(fmt.Sprintf("%v", a.Chest))
	builder.WriteString(", ")
	builder.WriteString("core=")
	builder.WriteString(fmt.Sprintf("%v", a.Core))
	builder.WriteString(", ")
	builder.WriteString("upper_back=")
	builder.WriteString(fmt.Sprintf("%v", a.UpperBack))
	builder.WriteString(", ")
	builder.WriteString("lower_back=")
	builder.WriteString(fmt.Sprintf("%v", a.LowerBack))
	builder.WriteString(", ")
	builder.WriteString("glute=")
	builder.WriteString(fmt.Sprintf("%v", a.Glute))
	builder.WriteString(", ")
	builder.WriteString("legs_front=")
	builder.WriteString(fmt.Sprintf("%v", a.LegsFront))
	builder.WriteString(", ")
	builder.WriteString("legs_back=")
	builder.WriteString(fmt.Sprintf("%v", a.LegsBack))
	builder.WriteString(", ")
	builder.WriteString("etc=")
	builder.WriteString(fmt.Sprintf("%v", a.Etc))
	builder.WriteByte(')')
	return builder.String()
}

// Acts is a parsable slice of Act.
type Acts []*Act
