// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/act"
	"routine/internal/ent/actimage"
	"routine/internal/ent/actversion"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActVersionCreate is the builder for creating a ActVersion entity.
type ActVersionCreate struct {
	config
	mutation *ActVersionMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (avc *ActVersionCreate) SetCode(s string) *ActVersionCreate {
	avc.mutation.SetCode(s)
	return avc
}

// SetActCode sets the "act_code" field.
func (avc *ActVersionCreate) SetActCode(s string) *ActVersionCreate {
	avc.mutation.SetActCode(s)
	return avc
}

// SetVersion sets the "version" field.
func (avc *ActVersionCreate) SetVersion(u uint) *ActVersionCreate {
	avc.mutation.SetVersion(u)
	return avc
}

// SetID sets the "id" field.
func (avc *ActVersionCreate) SetID(u uint64) *ActVersionCreate {
	avc.mutation.SetID(u)
	return avc
}

// AddActImageIDs adds the "act_images" edge to the ActImage entity by IDs.
func (avc *ActVersionCreate) AddActImageIDs(ids ...uint64) *ActVersionCreate {
	avc.mutation.AddActImageIDs(ids...)
	return avc
}

// AddActImages adds the "act_images" edges to the ActImage entity.
func (avc *ActVersionCreate) AddActImages(a ...*ActImage) *ActVersionCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return avc.AddActImageIDs(ids...)
}

// SetActID sets the "act" edge to the Act entity by ID.
func (avc *ActVersionCreate) SetActID(id uint64) *ActVersionCreate {
	avc.mutation.SetActID(id)
	return avc
}

// SetAct sets the "act" edge to the Act entity.
func (avc *ActVersionCreate) SetAct(a *Act) *ActVersionCreate {
	return avc.SetActID(a.ID)
}

// Mutation returns the ActVersionMutation object of the builder.
func (avc *ActVersionCreate) Mutation() *ActVersionMutation {
	return avc.mutation
}

// Save creates the ActVersion in the database.
func (avc *ActVersionCreate) Save(ctx context.Context) (*ActVersion, error) {
	return withHooks(ctx, avc.sqlSave, avc.mutation, avc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (avc *ActVersionCreate) SaveX(ctx context.Context) *ActVersion {
	v, err := avc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (avc *ActVersionCreate) Exec(ctx context.Context) error {
	_, err := avc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avc *ActVersionCreate) ExecX(ctx context.Context) {
	if err := avc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (avc *ActVersionCreate) check() error {
	if _, ok := avc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "ActVersion.code"`)}
	}
	if v, ok := avc.mutation.Code(); ok {
		if err := actversion.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "ActVersion.code": %w`, err)}
		}
	}
	if _, ok := avc.mutation.ActCode(); !ok {
		return &ValidationError{Name: "act_code", err: errors.New(`ent: missing required field "ActVersion.act_code"`)}
	}
	if v, ok := avc.mutation.ActCode(); ok {
		if err := actversion.ActCodeValidator(v); err != nil {
			return &ValidationError{Name: "act_code", err: fmt.Errorf(`ent: validator failed for field "ActVersion.act_code": %w`, err)}
		}
	}
	if _, ok := avc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "ActVersion.version"`)}
	}
	if _, ok := avc.mutation.ActID(); !ok {
		return &ValidationError{Name: "act", err: errors.New(`ent: missing required edge "ActVersion.act"`)}
	}
	return nil
}

func (avc *ActVersionCreate) sqlSave(ctx context.Context) (*ActVersion, error) {
	if err := avc.check(); err != nil {
		return nil, err
	}
	_node, _spec := avc.createSpec()
	if err := sqlgraph.CreateNode(ctx, avc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	avc.mutation.id = &_node.ID
	avc.mutation.done = true
	return _node, nil
}

func (avc *ActVersionCreate) createSpec() (*ActVersion, *sqlgraph.CreateSpec) {
	var (
		_node = &ActVersion{config: avc.config}
		_spec = sqlgraph.NewCreateSpec(actversion.Table, sqlgraph.NewFieldSpec(actversion.FieldID, field.TypeUint64))
	)
	if id, ok := avc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := avc.mutation.Code(); ok {
		_spec.SetField(actversion.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := avc.mutation.ActCode(); ok {
		_spec.SetField(actversion.FieldActCode, field.TypeString, value)
		_node.ActCode = value
	}
	if value, ok := avc.mutation.Version(); ok {
		_spec.SetField(actversion.FieldVersion, field.TypeUint, value)
		_node.Version = value
	}
	if nodes := avc.mutation.ActImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actversion.ActImagesTable,
			Columns: []string{actversion.ActImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(actimage.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := avc.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   actversion.ActTable,
			Columns: []string{actversion.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.act_act_versions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActVersionCreateBulk is the builder for creating many ActVersion entities in bulk.
type ActVersionCreateBulk struct {
	config
	builders []*ActVersionCreate
}

// Save creates the ActVersion entities in the database.
func (avcb *ActVersionCreateBulk) Save(ctx context.Context) ([]*ActVersion, error) {
	specs := make([]*sqlgraph.CreateSpec, len(avcb.builders))
	nodes := make([]*ActVersion, len(avcb.builders))
	mutators := make([]Mutator, len(avcb.builders))
	for i := range avcb.builders {
		func(i int, root context.Context) {
			builder := avcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActVersionMutation)
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
					_, err = mutators[i+1].Mutate(root, avcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, avcb.driver, spec); err != nil {
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
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, avcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (avcb *ActVersionCreateBulk) SaveX(ctx context.Context) []*ActVersion {
	v, err := avcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (avcb *ActVersionCreateBulk) Exec(ctx context.Context) error {
	_, err := avcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avcb *ActVersionCreateBulk) ExecX(ctx context.Context) {
	if err := avcb.Exec(ctx); err != nil {
		panic(err)
	}
}