// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/programrelease"
	"routine/internal/ent/s3image"
	"routine/internal/ent/s3programimage"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// S3ProgramImageCreate is the builder for creating a S3ProgramImage entity.
type S3ProgramImageCreate struct {
	config
	mutation *S3ProgramImageMutation
	hooks    []Hook
}

// SetOrder sets the "order" field.
func (sic *S3ProgramImageCreate) SetOrder(i int) *S3ProgramImageCreate {
	sic.mutation.SetOrder(i)
	return sic
}

// SetProgramReleaseID sets the "program_release_id" field.
func (sic *S3ProgramImageCreate) SetProgramReleaseID(i int64) *S3ProgramImageCreate {
	sic.mutation.SetProgramReleaseID(i)
	return sic
}

// SetImageID sets the "image_id" field.
func (sic *S3ProgramImageCreate) SetImageID(i int64) *S3ProgramImageCreate {
	sic.mutation.SetImageID(i)
	return sic
}

// SetID sets the "id" field.
func (sic *S3ProgramImageCreate) SetID(i int64) *S3ProgramImageCreate {
	sic.mutation.SetID(i)
	return sic
}

// SetProgramRelease sets the "program_release" edge to the ProgramRelease entity.
func (sic *S3ProgramImageCreate) SetProgramRelease(p *ProgramRelease) *S3ProgramImageCreate {
	return sic.SetProgramReleaseID(p.ID)
}

// SetS3ImageID sets the "s3_image" edge to the S3Image entity by ID.
func (sic *S3ProgramImageCreate) SetS3ImageID(id int64) *S3ProgramImageCreate {
	sic.mutation.SetS3ImageID(id)
	return sic
}

// SetS3Image sets the "s3_image" edge to the S3Image entity.
func (sic *S3ProgramImageCreate) SetS3Image(s *S3Image) *S3ProgramImageCreate {
	return sic.SetS3ImageID(s.ID)
}

// Mutation returns the S3ProgramImageMutation object of the builder.
func (sic *S3ProgramImageCreate) Mutation() *S3ProgramImageMutation {
	return sic.mutation
}

// Save creates the S3ProgramImage in the database.
func (sic *S3ProgramImageCreate) Save(ctx context.Context) (*S3ProgramImage, error) {
	return withHooks(ctx, sic.sqlSave, sic.mutation, sic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sic *S3ProgramImageCreate) SaveX(ctx context.Context) *S3ProgramImage {
	v, err := sic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sic *S3ProgramImageCreate) Exec(ctx context.Context) error {
	_, err := sic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sic *S3ProgramImageCreate) ExecX(ctx context.Context) {
	if err := sic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sic *S3ProgramImageCreate) check() error {
	if _, ok := sic.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "S3ProgramImage.order"`)}
	}
	if _, ok := sic.mutation.ProgramReleaseID(); !ok {
		return &ValidationError{Name: "program_release_id", err: errors.New(`ent: missing required field "S3ProgramImage.program_release_id"`)}
	}
	if _, ok := sic.mutation.ImageID(); !ok {
		return &ValidationError{Name: "image_id", err: errors.New(`ent: missing required field "S3ProgramImage.image_id"`)}
	}
	if _, ok := sic.mutation.ProgramReleaseID(); !ok {
		return &ValidationError{Name: "program_release", err: errors.New(`ent: missing required edge "S3ProgramImage.program_release"`)}
	}
	if _, ok := sic.mutation.S3ImageID(); !ok {
		return &ValidationError{Name: "s3_image", err: errors.New(`ent: missing required edge "S3ProgramImage.s3_image"`)}
	}
	return nil
}

func (sic *S3ProgramImageCreate) sqlSave(ctx context.Context) (*S3ProgramImage, error) {
	if err := sic.check(); err != nil {
		return nil, err
	}
	_node, _spec := sic.createSpec()
	if err := sqlgraph.CreateNode(ctx, sic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	sic.mutation.id = &_node.ID
	sic.mutation.done = true
	return _node, nil
}

func (sic *S3ProgramImageCreate) createSpec() (*S3ProgramImage, *sqlgraph.CreateSpec) {
	var (
		_node = &S3ProgramImage{config: sic.config}
		_spec = sqlgraph.NewCreateSpec(s3programimage.Table, sqlgraph.NewFieldSpec(s3programimage.FieldID, field.TypeInt64))
	)
	if id, ok := sic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sic.mutation.Order(); ok {
		_spec.SetField(s3programimage.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if nodes := sic.mutation.ProgramReleaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   s3programimage.ProgramReleaseTable,
			Columns: []string{s3programimage.ProgramReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrelease.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProgramReleaseID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sic.mutation.S3ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   s3programimage.S3ImageTable,
			Columns: []string{s3programimage.S3ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(s3image.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ImageID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// S3ProgramImageCreateBulk is the builder for creating many S3ProgramImage entities in bulk.
type S3ProgramImageCreateBulk struct {
	config
	builders []*S3ProgramImageCreate
}

// Save creates the S3ProgramImage entities in the database.
func (sicb *S3ProgramImageCreateBulk) Save(ctx context.Context) ([]*S3ProgramImage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sicb.builders))
	nodes := make([]*S3ProgramImage, len(sicb.builders))
	mutators := make([]Mutator, len(sicb.builders))
	for i := range sicb.builders {
		func(i int, root context.Context) {
			builder := sicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*S3ProgramImageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sicb *S3ProgramImageCreateBulk) SaveX(ctx context.Context) []*S3ProgramImage {
	v, err := sicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sicb *S3ProgramImageCreateBulk) Exec(ctx context.Context) error {
	_, err := sicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sicb *S3ProgramImageCreateBulk) ExecX(ctx context.Context) {
	if err := sicb.Exec(ctx); err != nil {
		panic(err)
	}
}