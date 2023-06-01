// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/dailyroutine"
	"routine/ent/program"
	"routine/ent/weeklyroutine"
	"routine/ent/weeklyroutinerec"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WeeklyRoutineCreate is the builder for creating a WeeklyRoutine entity.
type WeeklyRoutineCreate struct {
	config
	mutation *WeeklyRoutineMutation
	hooks    []Hook
}

// SetProgramID sets the "program_id" field.
func (wrc *WeeklyRoutineCreate) SetProgramID(u uint64) *WeeklyRoutineCreate {
	wrc.mutation.SetProgramID(u)
	return wrc
}

// SetWeek sets the "week" field.
func (wrc *WeeklyRoutineCreate) SetWeek(i int) *WeeklyRoutineCreate {
	wrc.mutation.SetWeek(i)
	return wrc
}

// SetCreatedAt sets the "created_at" field.
func (wrc *WeeklyRoutineCreate) SetCreatedAt(t time.Time) *WeeklyRoutineCreate {
	wrc.mutation.SetCreatedAt(t)
	return wrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wrc *WeeklyRoutineCreate) SetNillableCreatedAt(t *time.Time) *WeeklyRoutineCreate {
	if t != nil {
		wrc.SetCreatedAt(*t)
	}
	return wrc
}

// SetUpdatedAt sets the "updated_at" field.
func (wrc *WeeklyRoutineCreate) SetUpdatedAt(t time.Time) *WeeklyRoutineCreate {
	wrc.mutation.SetUpdatedAt(t)
	return wrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wrc *WeeklyRoutineCreate) SetNillableUpdatedAt(t *time.Time) *WeeklyRoutineCreate {
	if t != nil {
		wrc.SetUpdatedAt(*t)
	}
	return wrc
}

// SetID sets the "id" field.
func (wrc *WeeklyRoutineCreate) SetID(u uint64) *WeeklyRoutineCreate {
	wrc.mutation.SetID(u)
	return wrc
}

// SetProgram sets the "program" edge to the Program entity.
func (wrc *WeeklyRoutineCreate) SetProgram(p *Program) *WeeklyRoutineCreate {
	return wrc.SetProgramID(p.ID)
}

// AddDailyRoutineIDs adds the "daily_routines" edge to the DailyRoutine entity by IDs.
func (wrc *WeeklyRoutineCreate) AddDailyRoutineIDs(ids ...uint64) *WeeklyRoutineCreate {
	wrc.mutation.AddDailyRoutineIDs(ids...)
	return wrc
}

// AddDailyRoutines adds the "daily_routines" edges to the DailyRoutine entity.
func (wrc *WeeklyRoutineCreate) AddDailyRoutines(d ...*DailyRoutine) *WeeklyRoutineCreate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return wrc.AddDailyRoutineIDs(ids...)
}

// AddWeeklyRoutineRecIDs adds the "weekly_routine_recs" edge to the WeeklyRoutineRec entity by IDs.
func (wrc *WeeklyRoutineCreate) AddWeeklyRoutineRecIDs(ids ...uint64) *WeeklyRoutineCreate {
	wrc.mutation.AddWeeklyRoutineRecIDs(ids...)
	return wrc
}

// AddWeeklyRoutineRecs adds the "weekly_routine_recs" edges to the WeeklyRoutineRec entity.
func (wrc *WeeklyRoutineCreate) AddWeeklyRoutineRecs(w ...*WeeklyRoutineRec) *WeeklyRoutineCreate {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wrc.AddWeeklyRoutineRecIDs(ids...)
}

// Mutation returns the WeeklyRoutineMutation object of the builder.
func (wrc *WeeklyRoutineCreate) Mutation() *WeeklyRoutineMutation {
	return wrc.mutation
}

// Save creates the WeeklyRoutine in the database.
func (wrc *WeeklyRoutineCreate) Save(ctx context.Context) (*WeeklyRoutine, error) {
	wrc.defaults()
	return withHooks(ctx, wrc.sqlSave, wrc.mutation, wrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wrc *WeeklyRoutineCreate) SaveX(ctx context.Context) *WeeklyRoutine {
	v, err := wrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wrc *WeeklyRoutineCreate) Exec(ctx context.Context) error {
	_, err := wrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrc *WeeklyRoutineCreate) ExecX(ctx context.Context) {
	if err := wrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wrc *WeeklyRoutineCreate) defaults() {
	if _, ok := wrc.mutation.CreatedAt(); !ok {
		v := weeklyroutine.DefaultCreatedAt()
		wrc.mutation.SetCreatedAt(v)
	}
	if _, ok := wrc.mutation.UpdatedAt(); !ok {
		v := weeklyroutine.DefaultUpdatedAt()
		wrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wrc *WeeklyRoutineCreate) check() error {
	if _, ok := wrc.mutation.ProgramID(); !ok {
		return &ValidationError{Name: "program_id", err: errors.New(`ent: missing required field "WeeklyRoutine.program_id"`)}
	}
	if _, ok := wrc.mutation.Week(); !ok {
		return &ValidationError{Name: "week", err: errors.New(`ent: missing required field "WeeklyRoutine.week"`)}
	}
	if v, ok := wrc.mutation.Week(); ok {
		if err := weeklyroutine.WeekValidator(v); err != nil {
			return &ValidationError{Name: "week", err: fmt.Errorf(`ent: validator failed for field "WeeklyRoutine.week": %w`, err)}
		}
	}
	if _, ok := wrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WeeklyRoutine.created_at"`)}
	}
	if _, ok := wrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WeeklyRoutine.updated_at"`)}
	}
	if _, ok := wrc.mutation.ProgramID(); !ok {
		return &ValidationError{Name: "program", err: errors.New(`ent: missing required edge "WeeklyRoutine.program"`)}
	}
	return nil
}

func (wrc *WeeklyRoutineCreate) sqlSave(ctx context.Context) (*WeeklyRoutine, error) {
	if err := wrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	wrc.mutation.id = &_node.ID
	wrc.mutation.done = true
	return _node, nil
}

func (wrc *WeeklyRoutineCreate) createSpec() (*WeeklyRoutine, *sqlgraph.CreateSpec) {
	var (
		_node = &WeeklyRoutine{config: wrc.config}
		_spec = sqlgraph.NewCreateSpec(weeklyroutine.Table, sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64))
	)
	if id, ok := wrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wrc.mutation.Week(); ok {
		_spec.SetField(weeklyroutine.FieldWeek, field.TypeInt, value)
		_node.Week = value
	}
	if value, ok := wrc.mutation.CreatedAt(); ok {
		_spec.SetField(weeklyroutine.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wrc.mutation.UpdatedAt(); ok {
		_spec.SetField(weeklyroutine.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := wrc.mutation.ProgramIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   weeklyroutine.ProgramTable,
			Columns: []string{weeklyroutine.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProgramID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wrc.mutation.DailyRoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weeklyroutine.DailyRoutinesTable,
			Columns: []string{weeklyroutine.DailyRoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wrc.mutation.WeeklyRoutineRecsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weeklyroutine.WeeklyRoutineRecsTable,
			Columns: []string{weeklyroutine.WeeklyRoutineRecsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutinerec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WeeklyRoutineCreateBulk is the builder for creating many WeeklyRoutine entities in bulk.
type WeeklyRoutineCreateBulk struct {
	config
	builders []*WeeklyRoutineCreate
}

// Save creates the WeeklyRoutine entities in the database.
func (wrcb *WeeklyRoutineCreateBulk) Save(ctx context.Context) ([]*WeeklyRoutine, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wrcb.builders))
	nodes := make([]*WeeklyRoutine, len(wrcb.builders))
	mutators := make([]Mutator, len(wrcb.builders))
	for i := range wrcb.builders {
		func(i int, root context.Context) {
			builder := wrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WeeklyRoutineMutation)
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
					_, err = mutators[i+1].Mutate(root, wrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wrcb *WeeklyRoutineCreateBulk) SaveX(ctx context.Context) []*WeeklyRoutine {
	v, err := wrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wrcb *WeeklyRoutineCreateBulk) Exec(ctx context.Context) error {
	_, err := wrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrcb *WeeklyRoutineCreateBulk) ExecX(ctx context.Context) {
	if err := wrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
