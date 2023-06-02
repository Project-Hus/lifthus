// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/act"
	"routine/ent/onerepmax"
	"routine/ent/programrec"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OneRepMaxCreate is the builder for creating a OneRepMax entity.
type OneRepMaxCreate struct {
	config
	mutation *OneRepMaxMutation
	hooks    []Hook
}

// SetAuthor sets the "author" field.
func (ormc *OneRepMaxCreate) SetAuthor(u uint64) *OneRepMaxCreate {
	ormc.mutation.SetAuthor(u)
	return ormc
}

// SetActID sets the "act_id" field.
func (ormc *OneRepMaxCreate) SetActID(u uint64) *OneRepMaxCreate {
	ormc.mutation.SetActID(u)
	return ormc
}

// SetProgramRecID sets the "program_rec_id" field.
func (ormc *OneRepMaxCreate) SetProgramRecID(u uint64) *OneRepMaxCreate {
	ormc.mutation.SetProgramRecID(u)
	return ormc
}

// SetNillableProgramRecID sets the "program_rec_id" field if the given value is not nil.
func (ormc *OneRepMaxCreate) SetNillableProgramRecID(u *uint64) *OneRepMaxCreate {
	if u != nil {
		ormc.SetProgramRecID(*u)
	}
	return ormc
}

// SetDate sets the "date" field.
func (ormc *OneRepMaxCreate) SetDate(t time.Time) *OneRepMaxCreate {
	ormc.mutation.SetDate(t)
	return ormc
}

// SetOneRepMax sets the "one_rep_max" field.
func (ormc *OneRepMaxCreate) SetOneRepMax(f float64) *OneRepMaxCreate {
	ormc.mutation.SetOneRepMax(f)
	return ormc
}

// SetNillableOneRepMax sets the "one_rep_max" field if the given value is not nil.
func (ormc *OneRepMaxCreate) SetNillableOneRepMax(f *float64) *OneRepMaxCreate {
	if f != nil {
		ormc.SetOneRepMax(*f)
	}
	return ormc
}

// SetCertified sets the "certified" field.
func (ormc *OneRepMaxCreate) SetCertified(b bool) *OneRepMaxCreate {
	ormc.mutation.SetCertified(b)
	return ormc
}

// SetNillableCertified sets the "certified" field if the given value is not nil.
func (ormc *OneRepMaxCreate) SetNillableCertified(b *bool) *OneRepMaxCreate {
	if b != nil {
		ormc.SetCertified(*b)
	}
	return ormc
}

// SetCalculated sets the "calculated" field.
func (ormc *OneRepMaxCreate) SetCalculated(b bool) *OneRepMaxCreate {
	ormc.mutation.SetCalculated(b)
	return ormc
}

// SetNillableCalculated sets the "calculated" field if the given value is not nil.
func (ormc *OneRepMaxCreate) SetNillableCalculated(b *bool) *OneRepMaxCreate {
	if b != nil {
		ormc.SetCalculated(*b)
	}
	return ormc
}

// SetID sets the "id" field.
func (ormc *OneRepMaxCreate) SetID(u uint64) *OneRepMaxCreate {
	ormc.mutation.SetID(u)
	return ormc
}

// SetAct sets the "act" edge to the Act entity.
func (ormc *OneRepMaxCreate) SetAct(a *Act) *OneRepMaxCreate {
	return ormc.SetActID(a.ID)
}

// SetProgramRec sets the "program_rec" edge to the ProgramRec entity.
func (ormc *OneRepMaxCreate) SetProgramRec(p *ProgramRec) *OneRepMaxCreate {
	return ormc.SetProgramRecID(p.ID)
}

// Mutation returns the OneRepMaxMutation object of the builder.
func (ormc *OneRepMaxCreate) Mutation() *OneRepMaxMutation {
	return ormc.mutation
}

// Save creates the OneRepMax in the database.
func (ormc *OneRepMaxCreate) Save(ctx context.Context) (*OneRepMax, error) {
	ormc.defaults()
	return withHooks(ctx, ormc.sqlSave, ormc.mutation, ormc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ormc *OneRepMaxCreate) SaveX(ctx context.Context) *OneRepMax {
	v, err := ormc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ormc *OneRepMaxCreate) Exec(ctx context.Context) error {
	_, err := ormc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ormc *OneRepMaxCreate) ExecX(ctx context.Context) {
	if err := ormc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ormc *OneRepMaxCreate) defaults() {
	if _, ok := ormc.mutation.Certified(); !ok {
		v := onerepmax.DefaultCertified
		ormc.mutation.SetCertified(v)
	}
	if _, ok := ormc.mutation.Calculated(); !ok {
		v := onerepmax.DefaultCalculated
		ormc.mutation.SetCalculated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ormc *OneRepMaxCreate) check() error {
	if _, ok := ormc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "OneRepMax.author"`)}
	}
	if _, ok := ormc.mutation.ActID(); !ok {
		return &ValidationError{Name: "act_id", err: errors.New(`ent: missing required field "OneRepMax.act_id"`)}
	}
	if _, ok := ormc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "OneRepMax.date"`)}
	}
	if _, ok := ormc.mutation.Certified(); !ok {
		return &ValidationError{Name: "certified", err: errors.New(`ent: missing required field "OneRepMax.certified"`)}
	}
	if _, ok := ormc.mutation.Calculated(); !ok {
		return &ValidationError{Name: "calculated", err: errors.New(`ent: missing required field "OneRepMax.calculated"`)}
	}
	if _, ok := ormc.mutation.ActID(); !ok {
		return &ValidationError{Name: "act", err: errors.New(`ent: missing required edge "OneRepMax.act"`)}
	}
	return nil
}

func (ormc *OneRepMaxCreate) sqlSave(ctx context.Context) (*OneRepMax, error) {
	if err := ormc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ormc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ormc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ormc.mutation.id = &_node.ID
	ormc.mutation.done = true
	return _node, nil
}

func (ormc *OneRepMaxCreate) createSpec() (*OneRepMax, *sqlgraph.CreateSpec) {
	var (
		_node = &OneRepMax{config: ormc.config}
		_spec = sqlgraph.NewCreateSpec(onerepmax.Table, sqlgraph.NewFieldSpec(onerepmax.FieldID, field.TypeUint64))
	)
	if id, ok := ormc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ormc.mutation.Author(); ok {
		_spec.SetField(onerepmax.FieldAuthor, field.TypeUint64, value)
		_node.Author = value
	}
	if value, ok := ormc.mutation.Date(); ok {
		_spec.SetField(onerepmax.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := ormc.mutation.OneRepMax(); ok {
		_spec.SetField(onerepmax.FieldOneRepMax, field.TypeFloat64, value)
		_node.OneRepMax = &value
	}
	if value, ok := ormc.mutation.Certified(); ok {
		_spec.SetField(onerepmax.FieldCertified, field.TypeBool, value)
		_node.Certified = value
	}
	if value, ok := ormc.mutation.Calculated(); ok {
		_spec.SetField(onerepmax.FieldCalculated, field.TypeBool, value)
		_node.Calculated = value
	}
	if nodes := ormc.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   onerepmax.ActTable,
			Columns: []string{onerepmax.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ActID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ormc.mutation.ProgramRecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   onerepmax.ProgramRecTable,
			Columns: []string{onerepmax.ProgramRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProgramRecID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OneRepMaxCreateBulk is the builder for creating many OneRepMax entities in bulk.
type OneRepMaxCreateBulk struct {
	config
	builders []*OneRepMaxCreate
}

// Save creates the OneRepMax entities in the database.
func (ormcb *OneRepMaxCreateBulk) Save(ctx context.Context) ([]*OneRepMax, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ormcb.builders))
	nodes := make([]*OneRepMax, len(ormcb.builders))
	mutators := make([]Mutator, len(ormcb.builders))
	for i := range ormcb.builders {
		func(i int, root context.Context) {
			builder := ormcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OneRepMaxMutation)
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
					_, err = mutators[i+1].Mutate(root, ormcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ormcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ormcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ormcb *OneRepMaxCreateBulk) SaveX(ctx context.Context) []*OneRepMax {
	v, err := ormcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ormcb *OneRepMaxCreateBulk) Exec(ctx context.Context) error {
	_, err := ormcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ormcb *OneRepMaxCreateBulk) ExecX(ctx context.Context) {
	if err := ormcb.Exec(ctx); err != nil {
		panic(err)
	}
}