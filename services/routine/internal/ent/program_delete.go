// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"routine/internal/ent/predicate"
	"routine/internal/ent/program"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramDelete is the builder for deleting a Program entity.
type ProgramDelete struct {
	config
	hooks    []Hook
	mutation *ProgramMutation
}

// Where appends a list predicates to the ProgramDelete builder.
func (pd *ProgramDelete) Where(ps ...predicate.Program) *ProgramDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ProgramDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ProgramDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ProgramDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(program.Table, sqlgraph.NewFieldSpec(program.FieldID, field.TypeInt64))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// ProgramDeleteOne is the builder for deleting a single Program entity.
type ProgramDeleteOne struct {
	pd *ProgramDelete
}

// Where appends a list predicates to the ProgramDelete builder.
func (pdo *ProgramDeleteOne) Where(ps ...predicate.Program) *ProgramDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *ProgramDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{program.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ProgramDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
