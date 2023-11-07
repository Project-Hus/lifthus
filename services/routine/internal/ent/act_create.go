// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/act"
	"routine/internal/ent/routineact"
	"routine/internal/ent/s3actimage"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActCreate is the builder for creating a Act entity.
type ActCreate struct {
	config
	mutation *ActMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (ac *ActCreate) SetCode(s string) *ActCreate {
	ac.mutation.SetCode(s)
	return ac
}

// SetAuthor sets the "author" field.
func (ac *ActCreate) SetAuthor(i int64) *ActCreate {
	ac.mutation.SetAuthor(i)
	return ac
}

// SetActType sets the "act_type" field.
func (ac *ActCreate) SetActType(at act.ActType) *ActCreate {
	ac.mutation.SetActType(at)
	return ac
}

// SetName sets the "name" field.
func (ac *ActCreate) SetName(s string) *ActCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetText sets the "text" field.
func (ac *ActCreate) SetText(s string) *ActCreate {
	ac.mutation.SetText(s)
	return ac
}

// SetStandard sets the "standard" field.
func (ac *ActCreate) SetStandard(b bool) *ActCreate {
	ac.mutation.SetStandard(b)
	return ac
}

// SetNillableStandard sets the "standard" field if the given value is not nil.
func (ac *ActCreate) SetNillableStandard(b *bool) *ActCreate {
	if b != nil {
		ac.SetStandard(*b)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *ActCreate) SetCreatedAt(t time.Time) *ActCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetID sets the "id" field.
func (ac *ActCreate) SetID(i int64) *ActCreate {
	ac.mutation.SetID(i)
	return ac
}

// AddS3ActImageIDs adds the "s3_act_images" edge to the S3ActImage entity by IDs.
func (ac *ActCreate) AddS3ActImageIDs(ids ...int64) *ActCreate {
	ac.mutation.AddS3ActImageIDs(ids...)
	return ac
}

// AddS3ActImages adds the "s3_act_images" edges to the S3ActImage entity.
func (ac *ActCreate) AddS3ActImages(s ...*S3ActImage) *ActCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddS3ActImageIDs(ids...)
}

// AddRoutineActIDs adds the "routine_acts" edge to the RoutineAct entity by IDs.
func (ac *ActCreate) AddRoutineActIDs(ids ...int64) *ActCreate {
	ac.mutation.AddRoutineActIDs(ids...)
	return ac
}

// AddRoutineActs adds the "routine_acts" edges to the RoutineAct entity.
func (ac *ActCreate) AddRoutineActs(r ...*RoutineAct) *ActCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddRoutineActIDs(ids...)
}

// Mutation returns the ActMutation object of the builder.
func (ac *ActCreate) Mutation() *ActMutation {
	return ac.mutation
}

// Save creates the Act in the database.
func (ac *ActCreate) Save(ctx context.Context) (*Act, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ActCreate) SaveX(ctx context.Context) *Act {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ActCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ActCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ActCreate) defaults() {
	if _, ok := ac.mutation.Standard(); !ok {
		v := act.DefaultStandard
		ac.mutation.SetStandard(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ActCreate) check() error {
	if _, ok := ac.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Act.code"`)}
	}
	if v, ok := ac.mutation.Code(); ok {
		if err := act.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Act.code": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "Act.author"`)}
	}
	if _, ok := ac.mutation.ActType(); !ok {
		return &ValidationError{Name: "act_type", err: errors.New(`ent: missing required field "Act.act_type"`)}
	}
	if v, ok := ac.mutation.ActType(); ok {
		if err := act.ActTypeValidator(v); err != nil {
			return &ValidationError{Name: "act_type", err: fmt.Errorf(`ent: validator failed for field "Act.act_type": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Act.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := act.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Act.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Act.text"`)}
	}
	if _, ok := ac.mutation.Standard(); !ok {
		return &ValidationError{Name: "standard", err: errors.New(`ent: missing required field "Act.standard"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Act.created_at"`)}
	}
	return nil
}

func (ac *ActCreate) sqlSave(ctx context.Context) (*Act, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ActCreate) createSpec() (*Act, *sqlgraph.CreateSpec) {
	var (
		_node = &Act{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(act.Table, sqlgraph.NewFieldSpec(act.FieldID, field.TypeInt64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Code(); ok {
		_spec.SetField(act.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := ac.mutation.Author(); ok {
		_spec.SetField(act.FieldAuthor, field.TypeInt64, value)
		_node.Author = value
	}
	if value, ok := ac.mutation.ActType(); ok {
		_spec.SetField(act.FieldActType, field.TypeEnum, value)
		_node.ActType = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(act.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Text(); ok {
		_spec.SetField(act.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := ac.mutation.Standard(); ok {
		_spec.SetField(act.FieldStandard, field.TypeBool, value)
		_node.Standard = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(act.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.S3ActImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   act.S3ActImagesTable,
			Columns: []string{act.S3ActImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(s3actimage.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.RoutineActsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   act.RoutineActsTable,
			Columns: []string{act.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActCreateBulk is the builder for creating many Act entities in bulk.
type ActCreateBulk struct {
	config
	builders []*ActCreate
}

// Save creates the Act entities in the database.
func (acb *ActCreateBulk) Save(ctx context.Context) ([]*Act, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Act, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ActCreateBulk) SaveX(ctx context.Context) []*Act {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ActCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ActCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
