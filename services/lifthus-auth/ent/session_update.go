// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lifthus-auth/ent/predicate"
	"lifthus-auth/ent/session"
	"lifthus-auth/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetTid sets the "tid" field.
func (su *SessionUpdate) SetTid(u uuid.UUID) *SessionUpdate {
	su.mutation.SetTid(u)
	return su
}

// SetNillableTid sets the "tid" field if the given value is not nil.
func (su *SessionUpdate) SetNillableTid(u *uuid.UUID) *SessionUpdate {
	if u != nil {
		su.SetTid(*u)
	}
	return su
}

// SetHsid sets the "hsid" field.
func (su *SessionUpdate) SetHsid(u uuid.UUID) *SessionUpdate {
	su.mutation.SetHsid(u)
	return su
}

// SetNillableHsid sets the "hsid" field if the given value is not nil.
func (su *SessionUpdate) SetNillableHsid(u *uuid.UUID) *SessionUpdate {
	if u != nil {
		su.SetHsid(*u)
	}
	return su
}

// ClearHsid clears the value of the "hsid" field.
func (su *SessionUpdate) ClearHsid() *SessionUpdate {
	su.mutation.ClearHsid()
	return su
}

// SetConnectedAt sets the "connected_at" field.
func (su *SessionUpdate) SetConnectedAt(t time.Time) *SessionUpdate {
	su.mutation.SetConnectedAt(t)
	return su
}

// SetNillableConnectedAt sets the "connected_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableConnectedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetConnectedAt(*t)
	}
	return su
}

// SetUID sets the "uid" field.
func (su *SessionUpdate) SetUID(u uint64) *SessionUpdate {
	su.mutation.SetUID(u)
	return su
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUID(u *uint64) *SessionUpdate {
	if u != nil {
		su.SetUID(*u)
	}
	return su
}

// ClearUID clears the value of the "uid" field.
func (su *SessionUpdate) ClearUID() *SessionUpdate {
	su.mutation.ClearUID()
	return su
}

// SetSignedAt sets the "signed_at" field.
func (su *SessionUpdate) SetSignedAt(t time.Time) *SessionUpdate {
	su.mutation.SetSignedAt(t)
	return su
}

// SetNillableSignedAt sets the "signed_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableSignedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetSignedAt(*t)
	}
	return su
}

// ClearSignedAt clears the value of the "signed_at" field.
func (su *SessionUpdate) ClearSignedAt() *SessionUpdate {
	su.mutation.ClearSignedAt()
	return su
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *SessionUpdate) SetUserID(id uint64) *SessionUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (su *SessionUpdate) SetNillableUserID(id *uint64) *SessionUpdate {
	if id != nil {
		su = su.SetUserID(*id)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SessionUpdate) SetUser(u *User) *SessionUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *SessionUpdate) ClearUser() *SessionUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Tid(); ok {
		_spec.SetField(session.FieldTid, field.TypeUUID, value)
	}
	if value, ok := su.mutation.Hsid(); ok {
		_spec.SetField(session.FieldHsid, field.TypeUUID, value)
	}
	if su.mutation.HsidCleared() {
		_spec.ClearField(session.FieldHsid, field.TypeUUID)
	}
	if value, ok := su.mutation.ConnectedAt(); ok {
		_spec.SetField(session.FieldConnectedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.SignedAt(); ok {
		_spec.SetField(session.FieldSignedAt, field.TypeTime, value)
	}
	if su.mutation.SignedAtCleared() {
		_spec.ClearField(session.FieldSignedAt, field.TypeTime)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetTid sets the "tid" field.
func (suo *SessionUpdateOne) SetTid(u uuid.UUID) *SessionUpdateOne {
	suo.mutation.SetTid(u)
	return suo
}

// SetNillableTid sets the "tid" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableTid(u *uuid.UUID) *SessionUpdateOne {
	if u != nil {
		suo.SetTid(*u)
	}
	return suo
}

// SetHsid sets the "hsid" field.
func (suo *SessionUpdateOne) SetHsid(u uuid.UUID) *SessionUpdateOne {
	suo.mutation.SetHsid(u)
	return suo
}

// SetNillableHsid sets the "hsid" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableHsid(u *uuid.UUID) *SessionUpdateOne {
	if u != nil {
		suo.SetHsid(*u)
	}
	return suo
}

// ClearHsid clears the value of the "hsid" field.
func (suo *SessionUpdateOne) ClearHsid() *SessionUpdateOne {
	suo.mutation.ClearHsid()
	return suo
}

// SetConnectedAt sets the "connected_at" field.
func (suo *SessionUpdateOne) SetConnectedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetConnectedAt(t)
	return suo
}

// SetNillableConnectedAt sets the "connected_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableConnectedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetConnectedAt(*t)
	}
	return suo
}

// SetUID sets the "uid" field.
func (suo *SessionUpdateOne) SetUID(u uint64) *SessionUpdateOne {
	suo.mutation.SetUID(u)
	return suo
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUID(u *uint64) *SessionUpdateOne {
	if u != nil {
		suo.SetUID(*u)
	}
	return suo
}

// ClearUID clears the value of the "uid" field.
func (suo *SessionUpdateOne) ClearUID() *SessionUpdateOne {
	suo.mutation.ClearUID()
	return suo
}

// SetSignedAt sets the "signed_at" field.
func (suo *SessionUpdateOne) SetSignedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetSignedAt(t)
	return suo
}

// SetNillableSignedAt sets the "signed_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableSignedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetSignedAt(*t)
	}
	return suo
}

// ClearSignedAt clears the value of the "signed_at" field.
func (suo *SessionUpdateOne) ClearSignedAt() *SessionUpdateOne {
	suo.mutation.ClearSignedAt()
	return suo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *SessionUpdateOne) SetUserID(id uint64) *SessionUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserID(id *uint64) *SessionUpdateOne {
	if id != nil {
		suo = suo.SetUserID(*id)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SessionUpdateOne) SetUser(u *User) *SessionUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SessionUpdateOne) ClearUser() *SessionUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Tid(); ok {
		_spec.SetField(session.FieldTid, field.TypeUUID, value)
	}
	if value, ok := suo.mutation.Hsid(); ok {
		_spec.SetField(session.FieldHsid, field.TypeUUID, value)
	}
	if suo.mutation.HsidCleared() {
		_spec.ClearField(session.FieldHsid, field.TypeUUID)
	}
	if value, ok := suo.mutation.ConnectedAt(); ok {
		_spec.SetField(session.FieldConnectedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.SignedAt(); ok {
		_spec.SetField(session.FieldSignedAt, field.TypeTime, value)
	}
	if suo.mutation.SignedAtCleared() {
		_spec.ClearField(session.FieldSignedAt, field.TypeTime)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
