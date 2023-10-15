// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/act"
	"routine/internal/ent/dayroutine"
	"routine/internal/ent/routineact"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoutineActCreate is the builder for creating a RoutineAct entity.
type RoutineActCreate struct {
	config
	mutation *RoutineActMutation
	hooks    []Hook
}

// SetOrder sets the "order" field.
func (rac *RoutineActCreate) SetOrder(i int) *RoutineActCreate {
	rac.mutation.SetOrder(i)
	return rac
}

// SetActCode sets the "act_code" field.
func (rac *RoutineActCreate) SetActCode(s string) *RoutineActCreate {
	rac.mutation.SetActCode(s)
	return rac
}

// SetStage sets the "stage" field.
func (rac *RoutineActCreate) SetStage(r routineact.Stage) *RoutineActCreate {
	rac.mutation.SetStage(r)
	return rac
}

// SetRepsOrMeters sets the "reps_or_meters" field.
func (rac *RoutineActCreate) SetRepsOrMeters(u uint) *RoutineActCreate {
	rac.mutation.SetRepsOrMeters(u)
	return rac
}

// SetRatioOrSecs sets the "ratio_or_secs" field.
func (rac *RoutineActCreate) SetRatioOrSecs(f float64) *RoutineActCreate {
	rac.mutation.SetRatioOrSecs(f)
	return rac
}

// SetID sets the "id" field.
func (rac *RoutineActCreate) SetID(i int64) *RoutineActCreate {
	rac.mutation.SetID(i)
	return rac
}

// SetActID sets the "act" edge to the Act entity by ID.
func (rac *RoutineActCreate) SetActID(id int64) *RoutineActCreate {
	rac.mutation.SetActID(id)
	return rac
}

// SetAct sets the "act" edge to the Act entity.
func (rac *RoutineActCreate) SetAct(a *Act) *RoutineActCreate {
	return rac.SetActID(a.ID)
}

// SetDayRoutineID sets the "day_routine" edge to the DayRoutine entity by ID.
func (rac *RoutineActCreate) SetDayRoutineID(id int64) *RoutineActCreate {
	rac.mutation.SetDayRoutineID(id)
	return rac
}

// SetDayRoutine sets the "day_routine" edge to the DayRoutine entity.
func (rac *RoutineActCreate) SetDayRoutine(d *DayRoutine) *RoutineActCreate {
	return rac.SetDayRoutineID(d.ID)
}

// Mutation returns the RoutineActMutation object of the builder.
func (rac *RoutineActCreate) Mutation() *RoutineActMutation {
	return rac.mutation
}

// Save creates the RoutineAct in the database.
func (rac *RoutineActCreate) Save(ctx context.Context) (*RoutineAct, error) {
	return withHooks(ctx, rac.sqlSave, rac.mutation, rac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rac *RoutineActCreate) SaveX(ctx context.Context) *RoutineAct {
	v, err := rac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rac *RoutineActCreate) Exec(ctx context.Context) error {
	_, err := rac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rac *RoutineActCreate) ExecX(ctx context.Context) {
	if err := rac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rac *RoutineActCreate) check() error {
	if _, ok := rac.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "RoutineAct.order"`)}
	}
	if _, ok := rac.mutation.ActCode(); !ok {
		return &ValidationError{Name: "act_code", err: errors.New(`ent: missing required field "RoutineAct.act_code"`)}
	}
	if v, ok := rac.mutation.ActCode(); ok {
		if err := routineact.ActCodeValidator(v); err != nil {
			return &ValidationError{Name: "act_code", err: fmt.Errorf(`ent: validator failed for field "RoutineAct.act_code": %w`, err)}
		}
	}
	if _, ok := rac.mutation.Stage(); !ok {
		return &ValidationError{Name: "stage", err: errors.New(`ent: missing required field "RoutineAct.stage"`)}
	}
	if v, ok := rac.mutation.Stage(); ok {
		if err := routineact.StageValidator(v); err != nil {
			return &ValidationError{Name: "stage", err: fmt.Errorf(`ent: validator failed for field "RoutineAct.stage": %w`, err)}
		}
	}
	if _, ok := rac.mutation.RepsOrMeters(); !ok {
		return &ValidationError{Name: "reps_or_meters", err: errors.New(`ent: missing required field "RoutineAct.reps_or_meters"`)}
	}
	if _, ok := rac.mutation.RatioOrSecs(); !ok {
		return &ValidationError{Name: "ratio_or_secs", err: errors.New(`ent: missing required field "RoutineAct.ratio_or_secs"`)}
	}
	if _, ok := rac.mutation.ActID(); !ok {
		return &ValidationError{Name: "act", err: errors.New(`ent: missing required edge "RoutineAct.act"`)}
	}
	if _, ok := rac.mutation.DayRoutineID(); !ok {
		return &ValidationError{Name: "day_routine", err: errors.New(`ent: missing required edge "RoutineAct.day_routine"`)}
	}
	return nil
}

func (rac *RoutineActCreate) sqlSave(ctx context.Context) (*RoutineAct, error) {
	if err := rac.check(); err != nil {
		return nil, err
	}
	_node, _spec := rac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rac.mutation.id = &_node.ID
	rac.mutation.done = true
	return _node, nil
}

func (rac *RoutineActCreate) createSpec() (*RoutineAct, *sqlgraph.CreateSpec) {
	var (
		_node = &RoutineAct{config: rac.config}
		_spec = sqlgraph.NewCreateSpec(routineact.Table, sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeInt64))
	)
	if id, ok := rac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rac.mutation.Order(); ok {
		_spec.SetField(routineact.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if value, ok := rac.mutation.ActCode(); ok {
		_spec.SetField(routineact.FieldActCode, field.TypeString, value)
		_node.ActCode = value
	}
	if value, ok := rac.mutation.Stage(); ok {
		_spec.SetField(routineact.FieldStage, field.TypeEnum, value)
		_node.Stage = value
	}
	if value, ok := rac.mutation.RepsOrMeters(); ok {
		_spec.SetField(routineact.FieldRepsOrMeters, field.TypeUint, value)
		_node.RepsOrMeters = value
	}
	if value, ok := rac.mutation.RatioOrSecs(); ok {
		_spec.SetField(routineact.FieldRatioOrSecs, field.TypeFloat64, value)
		_node.RatioOrSecs = value
	}
	if nodes := rac.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.ActTable,
			Columns: []string{routineact.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.act_routine_acts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rac.mutation.DayRoutineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineact.DayRoutineTable,
			Columns: []string{routineact.DayRoutineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dayroutine.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.day_routine_routine_acts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoutineActCreateBulk is the builder for creating many RoutineAct entities in bulk.
type RoutineActCreateBulk struct {
	config
	builders []*RoutineActCreate
}

// Save creates the RoutineAct entities in the database.
func (racb *RoutineActCreateBulk) Save(ctx context.Context) ([]*RoutineAct, error) {
	specs := make([]*sqlgraph.CreateSpec, len(racb.builders))
	nodes := make([]*RoutineAct, len(racb.builders))
	mutators := make([]Mutator, len(racb.builders))
	for i := range racb.builders {
		func(i int, root context.Context) {
			builder := racb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoutineActMutation)
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
					_, err = mutators[i+1].Mutate(root, racb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, racb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, racb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (racb *RoutineActCreateBulk) SaveX(ctx context.Context) []*RoutineAct {
	v, err := racb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (racb *RoutineActCreateBulk) Exec(ctx context.Context) error {
	_, err := racb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (racb *RoutineActCreateBulk) ExecX(ctx context.Context) {
	if err := racb.Exec(ctx); err != nil {
		panic(err)
	}
}
