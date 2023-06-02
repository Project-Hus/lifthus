// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"routine/ent/predicate"
	"routine/ent/programrec"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramRecDelete is the builder for deleting a ProgramRec entity.
type ProgramRecDelete struct {
	config
	hooks    []Hook
	mutation *ProgramRecMutation
}

// Where appends a list predicates to the ProgramRecDelete builder.
func (prd *ProgramRecDelete) Where(ps ...predicate.ProgramRec) *ProgramRecDelete {
	prd.mutation.Where(ps...)
	return prd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prd *ProgramRecDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, prd.sqlExec, prd.mutation, prd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (prd *ProgramRecDelete) ExecX(ctx context.Context) int {
	n, err := prd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prd *ProgramRecDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(programrec.Table, sqlgraph.NewFieldSpec(programrec.FieldID, field.TypeUint64))
	if ps := prd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, prd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	prd.mutation.done = true
	return affected, err
}

// ProgramRecDeleteOne is the builder for deleting a single ProgramRec entity.
type ProgramRecDeleteOne struct {
	prd *ProgramRecDelete
}

// Where appends a list predicates to the ProgramRecDelete builder.
func (prdo *ProgramRecDeleteOne) Where(ps ...predicate.ProgramRec) *ProgramRecDeleteOne {
	prdo.prd.mutation.Where(ps...)
	return prdo
}

// Exec executes the deletion query.
func (prdo *ProgramRecDeleteOne) Exec(ctx context.Context) error {
	n, err := prdo.prd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{programrec.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prdo *ProgramRecDeleteOne) ExecX(ctx context.Context) {
	if err := prdo.Exec(ctx); err != nil {
		panic(err)
	}
}