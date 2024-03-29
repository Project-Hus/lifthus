// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/predicate"
	"routine/internal/ent/programrelease"
	"routine/internal/ent/routine"
	"routine/internal/ent/routineact"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoutineUpdate is the builder for updating Routine entities.
type RoutineUpdate struct {
	config
	hooks    []Hook
	mutation *RoutineMutation
}

// Where appends a list predicates to the RoutineUpdate builder.
func (ru *RoutineUpdate) Where(ps ...predicate.Routine) *RoutineUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetProgramReleaseID sets the "program_release" edge to the ProgramRelease entity by ID.
func (ru *RoutineUpdate) SetProgramReleaseID(id int64) *RoutineUpdate {
	ru.mutation.SetProgramReleaseID(id)
	return ru
}

// SetProgramRelease sets the "program_release" edge to the ProgramRelease entity.
func (ru *RoutineUpdate) SetProgramRelease(p *ProgramRelease) *RoutineUpdate {
	return ru.SetProgramReleaseID(p.ID)
}

// AddRoutineActIDs adds the "routine_acts" edge to the RoutineAct entity by IDs.
func (ru *RoutineUpdate) AddRoutineActIDs(ids ...int64) *RoutineUpdate {
	ru.mutation.AddRoutineActIDs(ids...)
	return ru
}

// AddRoutineActs adds the "routine_acts" edges to the RoutineAct entity.
func (ru *RoutineUpdate) AddRoutineActs(r ...*RoutineAct) *RoutineUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRoutineActIDs(ids...)
}

// Mutation returns the RoutineMutation object of the builder.
func (ru *RoutineUpdate) Mutation() *RoutineMutation {
	return ru.mutation
}

// ClearProgramRelease clears the "program_release" edge to the ProgramRelease entity.
func (ru *RoutineUpdate) ClearProgramRelease() *RoutineUpdate {
	ru.mutation.ClearProgramRelease()
	return ru
}

// ClearRoutineActs clears all "routine_acts" edges to the RoutineAct entity.
func (ru *RoutineUpdate) ClearRoutineActs() *RoutineUpdate {
	ru.mutation.ClearRoutineActs()
	return ru
}

// RemoveRoutineActIDs removes the "routine_acts" edge to RoutineAct entities by IDs.
func (ru *RoutineUpdate) RemoveRoutineActIDs(ids ...int64) *RoutineUpdate {
	ru.mutation.RemoveRoutineActIDs(ids...)
	return ru
}

// RemoveRoutineActs removes "routine_acts" edges to RoutineAct entities.
func (ru *RoutineUpdate) RemoveRoutineActs(r ...*RoutineAct) *RoutineUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRoutineActIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoutineUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoutineUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoutineUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoutineUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoutineUpdate) check() error {
	if _, ok := ru.mutation.ProgramReleaseID(); ru.mutation.ProgramReleaseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Routine.program_release"`)
	}
	return nil
}

func (ru *RoutineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(routine.Table, routine.Columns, sqlgraph.NewFieldSpec(routine.FieldID, field.TypeInt64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ru.mutation.ProgramReleaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routine.ProgramReleaseTable,
			Columns: []string{routine.ProgramReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrelease.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProgramReleaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routine.ProgramReleaseTable,
			Columns: []string{routine.ProgramReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrelease.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineActsTable,
			Columns: []string{routine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRoutineActsIDs(); len(nodes) > 0 && !ru.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineActsTable,
			Columns: []string{routine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RoutineActsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineActsTable,
			Columns: []string{routine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoutineUpdateOne is the builder for updating a single Routine entity.
type RoutineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoutineMutation
}

// SetProgramReleaseID sets the "program_release" edge to the ProgramRelease entity by ID.
func (ruo *RoutineUpdateOne) SetProgramReleaseID(id int64) *RoutineUpdateOne {
	ruo.mutation.SetProgramReleaseID(id)
	return ruo
}

// SetProgramRelease sets the "program_release" edge to the ProgramRelease entity.
func (ruo *RoutineUpdateOne) SetProgramRelease(p *ProgramRelease) *RoutineUpdateOne {
	return ruo.SetProgramReleaseID(p.ID)
}

// AddRoutineActIDs adds the "routine_acts" edge to the RoutineAct entity by IDs.
func (ruo *RoutineUpdateOne) AddRoutineActIDs(ids ...int64) *RoutineUpdateOne {
	ruo.mutation.AddRoutineActIDs(ids...)
	return ruo
}

// AddRoutineActs adds the "routine_acts" edges to the RoutineAct entity.
func (ruo *RoutineUpdateOne) AddRoutineActs(r ...*RoutineAct) *RoutineUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRoutineActIDs(ids...)
}

// Mutation returns the RoutineMutation object of the builder.
func (ruo *RoutineUpdateOne) Mutation() *RoutineMutation {
	return ruo.mutation
}

// ClearProgramRelease clears the "program_release" edge to the ProgramRelease entity.
func (ruo *RoutineUpdateOne) ClearProgramRelease() *RoutineUpdateOne {
	ruo.mutation.ClearProgramRelease()
	return ruo
}

// ClearRoutineActs clears all "routine_acts" edges to the RoutineAct entity.
func (ruo *RoutineUpdateOne) ClearRoutineActs() *RoutineUpdateOne {
	ruo.mutation.ClearRoutineActs()
	return ruo
}

// RemoveRoutineActIDs removes the "routine_acts" edge to RoutineAct entities by IDs.
func (ruo *RoutineUpdateOne) RemoveRoutineActIDs(ids ...int64) *RoutineUpdateOne {
	ruo.mutation.RemoveRoutineActIDs(ids...)
	return ruo
}

// RemoveRoutineActs removes "routine_acts" edges to RoutineAct entities.
func (ruo *RoutineUpdateOne) RemoveRoutineActs(r ...*RoutineAct) *RoutineUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRoutineActIDs(ids...)
}

// Where appends a list predicates to the RoutineUpdate builder.
func (ruo *RoutineUpdateOne) Where(ps ...predicate.Routine) *RoutineUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoutineUpdateOne) Select(field string, fields ...string) *RoutineUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Routine entity.
func (ruo *RoutineUpdateOne) Save(ctx context.Context) (*Routine, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoutineUpdateOne) SaveX(ctx context.Context) *Routine {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoutineUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoutineUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoutineUpdateOne) check() error {
	if _, ok := ruo.mutation.ProgramReleaseID(); ruo.mutation.ProgramReleaseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Routine.program_release"`)
	}
	return nil
}

func (ruo *RoutineUpdateOne) sqlSave(ctx context.Context) (_node *Routine, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(routine.Table, routine.Columns, sqlgraph.NewFieldSpec(routine.FieldID, field.TypeInt64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Routine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routine.FieldID)
		for _, f := range fields {
			if !routine.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != routine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ruo.mutation.ProgramReleaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routine.ProgramReleaseTable,
			Columns: []string{routine.ProgramReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrelease.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProgramReleaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routine.ProgramReleaseTable,
			Columns: []string{routine.ProgramReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programrelease.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineActsTable,
			Columns: []string{routine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRoutineActsIDs(); len(nodes) > 0 && !ruo.mutation.RoutineActsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineActsTable,
			Columns: []string{routine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RoutineActsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineActsTable,
			Columns: []string{routine.RoutineActsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Routine{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
