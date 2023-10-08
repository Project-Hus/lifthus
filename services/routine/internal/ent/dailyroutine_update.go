// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/dailyroutine"
	"routine/internal/ent/predicate"
	"routine/internal/ent/programversion"
	"routine/internal/ent/routineact"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DailyRoutineUpdate is the builder for updating DailyRoutine entities.
type DailyRoutineUpdate struct {
	config
	hooks    []Hook
	mutation *DailyRoutineMutation
}

// Where appends a list predicates to the DailyRoutineUpdate builder.
func (dru *DailyRoutineUpdate) Where(ps ...predicate.DailyRoutine) *DailyRoutineUpdate {
	dru.mutation.Where(ps...)
	return dru
}

// SetProgramVersionID sets the "program_version" edge to the ProgramVersion entity by ID.
func (dru *DailyRoutineUpdate) SetProgramVersionID(id uint64) *DailyRoutineUpdate {
	dru.mutation.SetProgramVersionID(id)
	return dru
}

// SetProgramVersion sets the "program_version" edge to the ProgramVersion entity.
func (dru *DailyRoutineUpdate) SetProgramVersion(p *ProgramVersion) *DailyRoutineUpdate {
	return dru.SetProgramVersionID(p.ID)
}

// AddRoutineActIDs adds the "routine_acts" edge to the RoutineAct entity by IDs.
func (dru *DailyRoutineUpdate) AddRoutineActIDs(ids ...uint64) *DailyRoutineUpdate {
	dru.mutation.AddRoutineActIDs(ids...)
	return dru
}

// AddRoutineActs adds the "routine_acts" edges to the RoutineAct entity.
func (dru *DailyRoutineUpdate) AddRoutineActs(r ...*RoutineAct) *DailyRoutineUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return dru.AddRoutineActIDs(ids...)
}

// Mutation returns the DailyRoutineMutation object of the builder.
func (dru *DailyRoutineUpdate) Mutation() *DailyRoutineMutation {
	return dru.mutation
}

// ClearProgramVersion clears the "program_version" edge to the ProgramVersion entity.
func (dru *DailyRoutineUpdate) ClearProgramVersion() *DailyRoutineUpdate {
	dru.mutation.ClearProgramVersion()
	return dru
}

// ClearRoutineActs clears all "routine_acts" edges to the RoutineAct entity.
func (dru *DailyRoutineUpdate) ClearRoutineActs() *DailyRoutineUpdate {
	dru.mutation.ClearRoutineActs()
	return dru
}

// RemoveRoutineActIDs removes the "routine_acts" edge to RoutineAct entities by IDs.
func (dru *DailyRoutineUpdate) RemoveRoutineActIDs(ids ...uint64) *DailyRoutineUpdate {
	dru.mutation.RemoveRoutineActIDs(ids...)
	return dru
}

// RemoveRoutineActs removes "routine_acts" edges to RoutineAct entities.
func (dru *DailyRoutineUpdate) RemoveRoutineActs(r ...*RoutineAct) *DailyRoutineUpdate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return dru.RemoveRoutineActIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dru *DailyRoutineUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dru.sqlSave, dru.mutation, dru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dru *DailyRoutineUpdate) SaveX(ctx context.Context) int {
	affected, err := dru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dru *DailyRoutineUpdate) Exec(ctx context.Context) error {
	_, err := dru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dru *DailyRoutineUpdate) ExecX(ctx context.Context) {
	if err := dru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dru *DailyRoutineUpdate) check() error {
	if _, ok := dru.mutation.ProgramVersionID(); dru.mutation.ProgramVersionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DailyRoutine.program_version"`)
	}
	return nil
}

func (dru *DailyRoutineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(dailyroutine.Table, dailyroutine.Columns, sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64))
	if ps := dru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if dru.mutation.ProgramVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutine.ProgramVersionTable,
			Columns: []string{dailyroutine.ProgramVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.ProgramVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutine.ProgramVersionTable,
			Columns: []string{dailyroutine.ProgramVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dru.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.RemovedRoutineActsIDs(); len(nodes) > 0 && !dru.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dru.mutation.RoutineActsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailyroutine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dru.mutation.done = true
	return n, nil
}

// DailyRoutineUpdateOne is the builder for updating a single DailyRoutine entity.
type DailyRoutineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DailyRoutineMutation
}

// SetProgramVersionID sets the "program_version" edge to the ProgramVersion entity by ID.
func (druo *DailyRoutineUpdateOne) SetProgramVersionID(id uint64) *DailyRoutineUpdateOne {
	druo.mutation.SetProgramVersionID(id)
	return druo
}

// SetProgramVersion sets the "program_version" edge to the ProgramVersion entity.
func (druo *DailyRoutineUpdateOne) SetProgramVersion(p *ProgramVersion) *DailyRoutineUpdateOne {
	return druo.SetProgramVersionID(p.ID)
}

// AddRoutineActIDs adds the "routine_acts" edge to the RoutineAct entity by IDs.
func (druo *DailyRoutineUpdateOne) AddRoutineActIDs(ids ...uint64) *DailyRoutineUpdateOne {
	druo.mutation.AddRoutineActIDs(ids...)
	return druo
}

// AddRoutineActs adds the "routine_acts" edges to the RoutineAct entity.
func (druo *DailyRoutineUpdateOne) AddRoutineActs(r ...*RoutineAct) *DailyRoutineUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return druo.AddRoutineActIDs(ids...)
}

// Mutation returns the DailyRoutineMutation object of the builder.
func (druo *DailyRoutineUpdateOne) Mutation() *DailyRoutineMutation {
	return druo.mutation
}

// ClearProgramVersion clears the "program_version" edge to the ProgramVersion entity.
func (druo *DailyRoutineUpdateOne) ClearProgramVersion() *DailyRoutineUpdateOne {
	druo.mutation.ClearProgramVersion()
	return druo
}

// ClearRoutineActs clears all "routine_acts" edges to the RoutineAct entity.
func (druo *DailyRoutineUpdateOne) ClearRoutineActs() *DailyRoutineUpdateOne {
	druo.mutation.ClearRoutineActs()
	return druo
}

// RemoveRoutineActIDs removes the "routine_acts" edge to RoutineAct entities by IDs.
func (druo *DailyRoutineUpdateOne) RemoveRoutineActIDs(ids ...uint64) *DailyRoutineUpdateOne {
	druo.mutation.RemoveRoutineActIDs(ids...)
	return druo
}

// RemoveRoutineActs removes "routine_acts" edges to RoutineAct entities.
func (druo *DailyRoutineUpdateOne) RemoveRoutineActs(r ...*RoutineAct) *DailyRoutineUpdateOne {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return druo.RemoveRoutineActIDs(ids...)
}

// Where appends a list predicates to the DailyRoutineUpdate builder.
func (druo *DailyRoutineUpdateOne) Where(ps ...predicate.DailyRoutine) *DailyRoutineUpdateOne {
	druo.mutation.Where(ps...)
	return druo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (druo *DailyRoutineUpdateOne) Select(field string, fields ...string) *DailyRoutineUpdateOne {
	druo.fields = append([]string{field}, fields...)
	return druo
}

// Save executes the query and returns the updated DailyRoutine entity.
func (druo *DailyRoutineUpdateOne) Save(ctx context.Context) (*DailyRoutine, error) {
	return withHooks(ctx, druo.sqlSave, druo.mutation, druo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (druo *DailyRoutineUpdateOne) SaveX(ctx context.Context) *DailyRoutine {
	node, err := druo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (druo *DailyRoutineUpdateOne) Exec(ctx context.Context) error {
	_, err := druo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (druo *DailyRoutineUpdateOne) ExecX(ctx context.Context) {
	if err := druo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (druo *DailyRoutineUpdateOne) check() error {
	if _, ok := druo.mutation.ProgramVersionID(); druo.mutation.ProgramVersionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DailyRoutine.program_version"`)
	}
	return nil
}

func (druo *DailyRoutineUpdateOne) sqlSave(ctx context.Context) (_node *DailyRoutine, err error) {
	if err := druo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(dailyroutine.Table, dailyroutine.Columns, sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64))
	id, ok := druo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DailyRoutine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := druo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailyroutine.FieldID)
		for _, f := range fields {
			if !dailyroutine.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dailyroutine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := druo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if druo.mutation.ProgramVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutine.ProgramVersionTable,
			Columns: []string{dailyroutine.ProgramVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.ProgramVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dailyroutine.ProgramVersionTable,
			Columns: []string{dailyroutine.ProgramVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if druo.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.RemovedRoutineActsIDs(); len(nodes) > 0 && !druo.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := druo.mutation.RoutineActsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dailyroutine.RoutineActsTable,
			Columns: []string{dailyroutine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DailyRoutine{config: druo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, druo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dailyroutine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	druo.mutation.done = true
	return _node, nil
}
