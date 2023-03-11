// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"lifthus-auth/ent/lifthussession"
	"lifthus-auth/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// LifthusSession is the model entity for the LifthusSession schema.
type LifthusSession struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"sid,omitempty"`
	// UID holds the value of the "uid" field.
	UID *uuid.UUID `json:"uid,omitempty"`
	// ConnectedAt holds the value of the "connected_at" field.
	ConnectedAt time.Time `json:"connected_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LifthusSessionQuery when eager-loading is set.
	Edges LifthusSessionEdges `json:"edges"`
}

// LifthusSessionEdges holds the relations/edges for other nodes in the graph.
type LifthusSessionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LifthusSessionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LifthusSession) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lifthussession.FieldUID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case lifthussession.FieldConnectedAt:
			values[i] = new(sql.NullTime)
		case lifthussession.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LifthusSession", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LifthusSession fields.
func (ls *LifthusSession) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lifthussession.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ls.ID = *value
			}
		case lifthussession.FieldUID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				ls.UID = new(uuid.UUID)
				*ls.UID = *value.S.(*uuid.UUID)
			}
		case lifthussession.FieldConnectedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field connected_at", values[i])
			} else if value.Valid {
				ls.ConnectedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the LifthusSession entity.
func (ls *LifthusSession) QueryUser() *UserQuery {
	return NewLifthusSessionClient(ls.config).QueryUser(ls)
}

// Update returns a builder for updating this LifthusSession.
// Note that you need to call LifthusSession.Unwrap() before calling this method if this LifthusSession
// was returned from a transaction, and the transaction was committed or rolled back.
func (ls *LifthusSession) Update() *LifthusSessionUpdateOne {
	return NewLifthusSessionClient(ls.config).UpdateOne(ls)
}

// Unwrap unwraps the LifthusSession entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ls *LifthusSession) Unwrap() *LifthusSession {
	_tx, ok := ls.config.driver.(*txDriver)
	if !ok {
		panic("ent: LifthusSession is not a transactional entity")
	}
	ls.config.driver = _tx.drv
	return ls
}

// String implements the fmt.Stringer.
func (ls *LifthusSession) String() string {
	var builder strings.Builder
	builder.WriteString("LifthusSession(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ls.ID))
	if v := ls.UID; v != nil {
		builder.WriteString("uid=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("connected_at=")
	builder.WriteString(ls.ConnectedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LifthusSessions is a parsable slice of LifthusSession.
type LifthusSessions []*LifthusSession