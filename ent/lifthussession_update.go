// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lifthus-auth/ent/lifthussession"
	"lifthus-auth/ent/predicate"
	"lifthus-auth/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LifthusSessionUpdate is the builder for updating LifthusSession entities.
type LifthusSessionUpdate struct {
	config
	hooks    []Hook
	mutation *LifthusSessionMutation
}

// Where appends a list predicates to the LifthusSessionUpdate builder.
func (lsu *LifthusSessionUpdate) Where(ps ...predicate.LifthusSession) *LifthusSessionUpdate {
	lsu.mutation.Where(ps...)
	return lsu
}

// SetUID sets the "uid" field.
func (lsu *LifthusSessionUpdate) SetUID(u uuid.UUID) *LifthusSessionUpdate {
	lsu.mutation.SetUID(u)
	return lsu
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (lsu *LifthusSessionUpdate) SetNillableUID(u *uuid.UUID) *LifthusSessionUpdate {
	if u != nil {
		lsu.SetUID(*u)
	}
	return lsu
}

// ClearUID clears the value of the "uid" field.
func (lsu *LifthusSessionUpdate) ClearUID() *LifthusSessionUpdate {
	lsu.mutation.ClearUID()
	return lsu
}

// SetConnectedAt sets the "connected_at" field.
func (lsu *LifthusSessionUpdate) SetConnectedAt(t time.Time) *LifthusSessionUpdate {
	lsu.mutation.SetConnectedAt(t)
	return lsu
}

// SetNillableConnectedAt sets the "connected_at" field if the given value is not nil.
func (lsu *LifthusSessionUpdate) SetNillableConnectedAt(t *time.Time) *LifthusSessionUpdate {
	if t != nil {
		lsu.SetConnectedAt(*t)
	}
	return lsu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lsu *LifthusSessionUpdate) SetUserID(id uuid.UUID) *LifthusSessionUpdate {
	lsu.mutation.SetUserID(id)
	return lsu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (lsu *LifthusSessionUpdate) SetNillableUserID(id *uuid.UUID) *LifthusSessionUpdate {
	if id != nil {
		lsu = lsu.SetUserID(*id)
	}
	return lsu
}

// SetUser sets the "user" edge to the User entity.
func (lsu *LifthusSessionUpdate) SetUser(u *User) *LifthusSessionUpdate {
	return lsu.SetUserID(u.ID)
}

// Mutation returns the LifthusSessionMutation object of the builder.
func (lsu *LifthusSessionUpdate) Mutation() *LifthusSessionMutation {
	return lsu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lsu *LifthusSessionUpdate) ClearUser() *LifthusSessionUpdate {
	lsu.mutation.ClearUser()
	return lsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lsu *LifthusSessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, LifthusSessionMutation](ctx, lsu.sqlSave, lsu.mutation, lsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lsu *LifthusSessionUpdate) SaveX(ctx context.Context) int {
	affected, err := lsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lsu *LifthusSessionUpdate) Exec(ctx context.Context) error {
	_, err := lsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsu *LifthusSessionUpdate) ExecX(ctx context.Context) {
	if err := lsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lsu *LifthusSessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(lifthussession.Table, lifthussession.Columns, sqlgraph.NewFieldSpec(lifthussession.FieldID, field.TypeUUID))
	if ps := lsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lsu.mutation.ConnectedAt(); ok {
		_spec.SetField(lifthussession.FieldConnectedAt, field.TypeTime, value)
	}
	if lsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lifthussession.UserTable,
			Columns: []string{lifthussession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lifthussession.UserTable,
			Columns: []string{lifthussession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lifthussession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lsu.mutation.done = true
	return n, nil
}

// LifthusSessionUpdateOne is the builder for updating a single LifthusSession entity.
type LifthusSessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LifthusSessionMutation
}

// SetUID sets the "uid" field.
func (lsuo *LifthusSessionUpdateOne) SetUID(u uuid.UUID) *LifthusSessionUpdateOne {
	lsuo.mutation.SetUID(u)
	return lsuo
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (lsuo *LifthusSessionUpdateOne) SetNillableUID(u *uuid.UUID) *LifthusSessionUpdateOne {
	if u != nil {
		lsuo.SetUID(*u)
	}
	return lsuo
}

// ClearUID clears the value of the "uid" field.
func (lsuo *LifthusSessionUpdateOne) ClearUID() *LifthusSessionUpdateOne {
	lsuo.mutation.ClearUID()
	return lsuo
}

// SetConnectedAt sets the "connected_at" field.
func (lsuo *LifthusSessionUpdateOne) SetConnectedAt(t time.Time) *LifthusSessionUpdateOne {
	lsuo.mutation.SetConnectedAt(t)
	return lsuo
}

// SetNillableConnectedAt sets the "connected_at" field if the given value is not nil.
func (lsuo *LifthusSessionUpdateOne) SetNillableConnectedAt(t *time.Time) *LifthusSessionUpdateOne {
	if t != nil {
		lsuo.SetConnectedAt(*t)
	}
	return lsuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lsuo *LifthusSessionUpdateOne) SetUserID(id uuid.UUID) *LifthusSessionUpdateOne {
	lsuo.mutation.SetUserID(id)
	return lsuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (lsuo *LifthusSessionUpdateOne) SetNillableUserID(id *uuid.UUID) *LifthusSessionUpdateOne {
	if id != nil {
		lsuo = lsuo.SetUserID(*id)
	}
	return lsuo
}

// SetUser sets the "user" edge to the User entity.
func (lsuo *LifthusSessionUpdateOne) SetUser(u *User) *LifthusSessionUpdateOne {
	return lsuo.SetUserID(u.ID)
}

// Mutation returns the LifthusSessionMutation object of the builder.
func (lsuo *LifthusSessionUpdateOne) Mutation() *LifthusSessionMutation {
	return lsuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lsuo *LifthusSessionUpdateOne) ClearUser() *LifthusSessionUpdateOne {
	lsuo.mutation.ClearUser()
	return lsuo
}

// Where appends a list predicates to the LifthusSessionUpdate builder.
func (lsuo *LifthusSessionUpdateOne) Where(ps ...predicate.LifthusSession) *LifthusSessionUpdateOne {
	lsuo.mutation.Where(ps...)
	return lsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lsuo *LifthusSessionUpdateOne) Select(field string, fields ...string) *LifthusSessionUpdateOne {
	lsuo.fields = append([]string{field}, fields...)
	return lsuo
}

// Save executes the query and returns the updated LifthusSession entity.
func (lsuo *LifthusSessionUpdateOne) Save(ctx context.Context) (*LifthusSession, error) {
	return withHooks[*LifthusSession, LifthusSessionMutation](ctx, lsuo.sqlSave, lsuo.mutation, lsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lsuo *LifthusSessionUpdateOne) SaveX(ctx context.Context) *LifthusSession {
	node, err := lsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lsuo *LifthusSessionUpdateOne) Exec(ctx context.Context) error {
	_, err := lsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsuo *LifthusSessionUpdateOne) ExecX(ctx context.Context) {
	if err := lsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lsuo *LifthusSessionUpdateOne) sqlSave(ctx context.Context) (_node *LifthusSession, err error) {
	_spec := sqlgraph.NewUpdateSpec(lifthussession.Table, lifthussession.Columns, sqlgraph.NewFieldSpec(lifthussession.FieldID, field.TypeUUID))
	id, ok := lsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LifthusSession.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lifthussession.FieldID)
		for _, f := range fields {
			if !lifthussession.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lifthussession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lsuo.mutation.ConnectedAt(); ok {
		_spec.SetField(lifthussession.FieldConnectedAt, field.TypeTime, value)
	}
	if lsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lifthussession.UserTable,
			Columns: []string{lifthussession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lifthussession.UserTable,
			Columns: []string{lifthussession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LifthusSession{config: lsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lifthussession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lsuo.mutation.done = true
	return _node, nil
}