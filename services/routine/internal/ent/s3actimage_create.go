// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/act"
	"routine/internal/ent/s3actimage"
	"routine/internal/ent/s3image"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// S3ActImageCreate is the builder for creating a S3ActImage entity.
type S3ActImageCreate struct {
	config
	mutation *S3ActImageMutation
	hooks    []Hook
}

// SetOrder sets the "order" field.
func (sic *S3ActImageCreate) SetOrder(i int) *S3ActImageCreate {
	sic.mutation.SetOrder(i)
	return sic
}

// SetActID sets the "act_id" field.
func (sic *S3ActImageCreate) SetActID(i int64) *S3ActImageCreate {
	sic.mutation.SetActID(i)
	return sic
}

// SetImageID sets the "image_id" field.
func (sic *S3ActImageCreate) SetImageID(i int64) *S3ActImageCreate {
	sic.mutation.SetImageID(i)
	return sic
}

// SetID sets the "id" field.
func (sic *S3ActImageCreate) SetID(i int64) *S3ActImageCreate {
	sic.mutation.SetID(i)
	return sic
}

// SetAct sets the "act" edge to the Act entity.
func (sic *S3ActImageCreate) SetAct(a *Act) *S3ActImageCreate {
	return sic.SetActID(a.ID)
}

// SetS3ImageID sets the "s3_image" edge to the S3Image entity by ID.
func (sic *S3ActImageCreate) SetS3ImageID(id int64) *S3ActImageCreate {
	sic.mutation.SetS3ImageID(id)
	return sic
}

// SetS3Image sets the "s3_image" edge to the S3Image entity.
func (sic *S3ActImageCreate) SetS3Image(s *S3Image) *S3ActImageCreate {
	return sic.SetS3ImageID(s.ID)
}

// Mutation returns the S3ActImageMutation object of the builder.
func (sic *S3ActImageCreate) Mutation() *S3ActImageMutation {
	return sic.mutation
}

// Save creates the S3ActImage in the database.
func (sic *S3ActImageCreate) Save(ctx context.Context) (*S3ActImage, error) {
	return withHooks(ctx, sic.sqlSave, sic.mutation, sic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sic *S3ActImageCreate) SaveX(ctx context.Context) *S3ActImage {
	v, err := sic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sic *S3ActImageCreate) Exec(ctx context.Context) error {
	_, err := sic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sic *S3ActImageCreate) ExecX(ctx context.Context) {
	if err := sic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sic *S3ActImageCreate) check() error {
	if _, ok := sic.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "S3ActImage.order"`)}
	}
	if _, ok := sic.mutation.ActID(); !ok {
		return &ValidationError{Name: "act_id", err: errors.New(`ent: missing required field "S3ActImage.act_id"`)}
	}
	if _, ok := sic.mutation.ImageID(); !ok {
		return &ValidationError{Name: "image_id", err: errors.New(`ent: missing required field "S3ActImage.image_id"`)}
	}
	if _, ok := sic.mutation.ActID(); !ok {
		return &ValidationError{Name: "act", err: errors.New(`ent: missing required edge "S3ActImage.act"`)}
	}
	if _, ok := sic.mutation.S3ImageID(); !ok {
		return &ValidationError{Name: "s3_image", err: errors.New(`ent: missing required edge "S3ActImage.s3_image"`)}
	}
	return nil
}

func (sic *S3ActImageCreate) sqlSave(ctx context.Context) (*S3ActImage, error) {
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

func (sic *S3ActImageCreate) createSpec() (*S3ActImage, *sqlgraph.CreateSpec) {
	var (
		_node = &S3ActImage{config: sic.config}
		_spec = sqlgraph.NewCreateSpec(s3actimage.Table, sqlgraph.NewFieldSpec(s3actimage.FieldID, field.TypeInt64))
	)
	if id, ok := sic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sic.mutation.Order(); ok {
		_spec.SetField(s3actimage.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if nodes := sic.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   s3actimage.ActTable,
			Columns: []string{s3actimage.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ActID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sic.mutation.S3ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   s3actimage.S3ImageTable,
			Columns: []string{s3actimage.S3ImageColumn},
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

// S3ActImageCreateBulk is the builder for creating many S3ActImage entities in bulk.
type S3ActImageCreateBulk struct {
	config
	builders []*S3ActImageCreate
}

// Save creates the S3ActImage entities in the database.
func (sicb *S3ActImageCreateBulk) Save(ctx context.Context) ([]*S3ActImage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sicb.builders))
	nodes := make([]*S3ActImage, len(sicb.builders))
	mutators := make([]Mutator, len(sicb.builders))
	for i := range sicb.builders {
		func(i int, root context.Context) {
			builder := sicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*S3ActImageMutation)
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
func (sicb *S3ActImageCreateBulk) SaveX(ctx context.Context) []*S3ActImage {
	v, err := sicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sicb *S3ActImageCreateBulk) Exec(ctx context.Context) error {
	_, err := sicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sicb *S3ActImageCreateBulk) ExecX(ctx context.Context) {
	if err := sicb.Exec(ctx); err != nil {
		panic(err)
	}
}
