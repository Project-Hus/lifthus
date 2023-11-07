// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"routine/internal/ent/predicate"
	"routine/internal/ent/s3image"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// S3ImageDelete is the builder for deleting a S3Image entity.
type S3ImageDelete struct {
	config
	hooks    []Hook
	mutation *S3ImageMutation
}

// Where appends a list predicates to the S3ImageDelete builder.
func (sd *S3ImageDelete) Where(ps ...predicate.S3Image) *S3ImageDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *S3ImageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *S3ImageDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *S3ImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(s3image.Table, sqlgraph.NewFieldSpec(s3image.FieldID, field.TypeInt64))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// S3ImageDeleteOne is the builder for deleting a single S3Image entity.
type S3ImageDeleteOne struct {
	sd *S3ImageDelete
}

// Where appends a list predicates to the S3ImageDelete builder.
func (sdo *S3ImageDeleteOne) Where(ps ...predicate.S3Image) *S3ImageDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *S3ImageDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{s3image.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *S3ImageDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}