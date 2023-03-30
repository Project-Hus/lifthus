// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"lifthus-auth/ent/lifthusgroup"
	"lifthus-auth/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LifthusGroupDelete is the builder for deleting a LifthusGroup entity.
type LifthusGroupDelete struct {
	config
	hooks    []Hook
	mutation *LifthusGroupMutation
}

// Where appends a list predicates to the LifthusGroupDelete builder.
func (lgd *LifthusGroupDelete) Where(ps ...predicate.LifthusGroup) *LifthusGroupDelete {
	lgd.mutation.Where(ps...)
	return lgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lgd *LifthusGroupDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LifthusGroupMutation](ctx, lgd.sqlExec, lgd.mutation, lgd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lgd *LifthusGroupDelete) ExecX(ctx context.Context) int {
	n, err := lgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lgd *LifthusGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(lifthusgroup.Table, sqlgraph.NewFieldSpec(lifthusgroup.FieldID, field.TypeInt))
	if ps := lgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lgd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lgd.mutation.done = true
	return affected, err
}

// LifthusGroupDeleteOne is the builder for deleting a single LifthusGroup entity.
type LifthusGroupDeleteOne struct {
	lgd *LifthusGroupDelete
}

// Where appends a list predicates to the LifthusGroupDelete builder.
func (lgdo *LifthusGroupDeleteOne) Where(ps ...predicate.LifthusGroup) *LifthusGroupDeleteOne {
	lgdo.lgd.mutation.Where(ps...)
	return lgdo
}

// Exec executes the deletion query.
func (lgdo *LifthusGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := lgdo.lgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lifthusgroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lgdo *LifthusGroupDeleteOne) ExecX(ctx context.Context) {
	if err := lgdo.Exec(ctx); err != nil {
		panic(err)
	}
}