// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/dailyroutine"
	"routine/internal/ent/image"
	"routine/internal/ent/predicate"
	"routine/internal/ent/program"
	"routine/internal/ent/programversion"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramVersionUpdate is the builder for updating ProgramVersion entities.
type ProgramVersionUpdate struct {
	config
	hooks    []Hook
	mutation *ProgramVersionMutation
}

// Where appends a list predicates to the ProgramVersionUpdate builder.
func (pvu *ProgramVersionUpdate) Where(ps ...predicate.ProgramVersion) *ProgramVersionUpdate {
	pvu.mutation.Where(ps...)
	return pvu
}

// SetProgramID sets the "program" edge to the Program entity by ID.
func (pvu *ProgramVersionUpdate) SetProgramID(id uint64) *ProgramVersionUpdate {
	pvu.mutation.SetProgramID(id)
	return pvu
}

// SetProgram sets the "program" edge to the Program entity.
func (pvu *ProgramVersionUpdate) SetProgram(p *Program) *ProgramVersionUpdate {
	return pvu.SetProgramID(p.ID)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (pvu *ProgramVersionUpdate) AddImageIDs(ids ...uint64) *ProgramVersionUpdate {
	pvu.mutation.AddImageIDs(ids...)
	return pvu
}

// AddImages adds the "images" edges to the Image entity.
func (pvu *ProgramVersionUpdate) AddImages(i ...*Image) *ProgramVersionUpdate {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pvu.AddImageIDs(ids...)
}

// AddDailyRoutineIDs adds the "daily_routines" edge to the DailyRoutine entity by IDs.
func (pvu *ProgramVersionUpdate) AddDailyRoutineIDs(ids ...uint64) *ProgramVersionUpdate {
	pvu.mutation.AddDailyRoutineIDs(ids...)
	return pvu
}

// AddDailyRoutines adds the "daily_routines" edges to the DailyRoutine entity.
func (pvu *ProgramVersionUpdate) AddDailyRoutines(d ...*DailyRoutine) *ProgramVersionUpdate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pvu.AddDailyRoutineIDs(ids...)
}

// Mutation returns the ProgramVersionMutation object of the builder.
func (pvu *ProgramVersionUpdate) Mutation() *ProgramVersionMutation {
	return pvu.mutation
}

// ClearProgram clears the "program" edge to the Program entity.
func (pvu *ProgramVersionUpdate) ClearProgram() *ProgramVersionUpdate {
	pvu.mutation.ClearProgram()
	return pvu
}

// ClearImages clears all "images" edges to the Image entity.
func (pvu *ProgramVersionUpdate) ClearImages() *ProgramVersionUpdate {
	pvu.mutation.ClearImages()
	return pvu
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (pvu *ProgramVersionUpdate) RemoveImageIDs(ids ...uint64) *ProgramVersionUpdate {
	pvu.mutation.RemoveImageIDs(ids...)
	return pvu
}

// RemoveImages removes "images" edges to Image entities.
func (pvu *ProgramVersionUpdate) RemoveImages(i ...*Image) *ProgramVersionUpdate {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pvu.RemoveImageIDs(ids...)
}

// ClearDailyRoutines clears all "daily_routines" edges to the DailyRoutine entity.
func (pvu *ProgramVersionUpdate) ClearDailyRoutines() *ProgramVersionUpdate {
	pvu.mutation.ClearDailyRoutines()
	return pvu
}

// RemoveDailyRoutineIDs removes the "daily_routines" edge to DailyRoutine entities by IDs.
func (pvu *ProgramVersionUpdate) RemoveDailyRoutineIDs(ids ...uint64) *ProgramVersionUpdate {
	pvu.mutation.RemoveDailyRoutineIDs(ids...)
	return pvu
}

// RemoveDailyRoutines removes "daily_routines" edges to DailyRoutine entities.
func (pvu *ProgramVersionUpdate) RemoveDailyRoutines(d ...*DailyRoutine) *ProgramVersionUpdate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pvu.RemoveDailyRoutineIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pvu *ProgramVersionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pvu.sqlSave, pvu.mutation, pvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pvu *ProgramVersionUpdate) SaveX(ctx context.Context) int {
	affected, err := pvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pvu *ProgramVersionUpdate) Exec(ctx context.Context) error {
	_, err := pvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvu *ProgramVersionUpdate) ExecX(ctx context.Context) {
	if err := pvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvu *ProgramVersionUpdate) check() error {
	if _, ok := pvu.mutation.ProgramID(); pvu.mutation.ProgramCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProgramVersion.program"`)
	}
	return nil
}

func (pvu *ProgramVersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(programversion.Table, programversion.Columns, sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64))
	if ps := pvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pvu.mutation.ProgramCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   programversion.ProgramTable,
			Columns: []string{programversion.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.ProgramIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   programversion.ProgramTable,
			Columns: []string{programversion.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   programversion.ImagesTable,
			Columns: programversion.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.RemovedImagesIDs(); len(nodes) > 0 && !pvu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   programversion.ImagesTable,
			Columns: programversion.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   programversion.ImagesTable,
			Columns: programversion.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvu.mutation.DailyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programversion.DailyRoutinesTable,
			Columns: []string{programversion.DailyRoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.RemovedDailyRoutinesIDs(); len(nodes) > 0 && !pvu.mutation.DailyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programversion.DailyRoutinesTable,
			Columns: []string{programversion.DailyRoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.DailyRoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programversion.DailyRoutinesTable,
			Columns: []string{programversion.DailyRoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{programversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pvu.mutation.done = true
	return n, nil
}

// ProgramVersionUpdateOne is the builder for updating a single ProgramVersion entity.
type ProgramVersionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProgramVersionMutation
}

// SetProgramID sets the "program" edge to the Program entity by ID.
func (pvuo *ProgramVersionUpdateOne) SetProgramID(id uint64) *ProgramVersionUpdateOne {
	pvuo.mutation.SetProgramID(id)
	return pvuo
}

// SetProgram sets the "program" edge to the Program entity.
func (pvuo *ProgramVersionUpdateOne) SetProgram(p *Program) *ProgramVersionUpdateOne {
	return pvuo.SetProgramID(p.ID)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (pvuo *ProgramVersionUpdateOne) AddImageIDs(ids ...uint64) *ProgramVersionUpdateOne {
	pvuo.mutation.AddImageIDs(ids...)
	return pvuo
}

// AddImages adds the "images" edges to the Image entity.
func (pvuo *ProgramVersionUpdateOne) AddImages(i ...*Image) *ProgramVersionUpdateOne {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pvuo.AddImageIDs(ids...)
}

// AddDailyRoutineIDs adds the "daily_routines" edge to the DailyRoutine entity by IDs.
func (pvuo *ProgramVersionUpdateOne) AddDailyRoutineIDs(ids ...uint64) *ProgramVersionUpdateOne {
	pvuo.mutation.AddDailyRoutineIDs(ids...)
	return pvuo
}

// AddDailyRoutines adds the "daily_routines" edges to the DailyRoutine entity.
func (pvuo *ProgramVersionUpdateOne) AddDailyRoutines(d ...*DailyRoutine) *ProgramVersionUpdateOne {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pvuo.AddDailyRoutineIDs(ids...)
}

// Mutation returns the ProgramVersionMutation object of the builder.
func (pvuo *ProgramVersionUpdateOne) Mutation() *ProgramVersionMutation {
	return pvuo.mutation
}

// ClearProgram clears the "program" edge to the Program entity.
func (pvuo *ProgramVersionUpdateOne) ClearProgram() *ProgramVersionUpdateOne {
	pvuo.mutation.ClearProgram()
	return pvuo
}

// ClearImages clears all "images" edges to the Image entity.
func (pvuo *ProgramVersionUpdateOne) ClearImages() *ProgramVersionUpdateOne {
	pvuo.mutation.ClearImages()
	return pvuo
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (pvuo *ProgramVersionUpdateOne) RemoveImageIDs(ids ...uint64) *ProgramVersionUpdateOne {
	pvuo.mutation.RemoveImageIDs(ids...)
	return pvuo
}

// RemoveImages removes "images" edges to Image entities.
func (pvuo *ProgramVersionUpdateOne) RemoveImages(i ...*Image) *ProgramVersionUpdateOne {
	ids := make([]uint64, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pvuo.RemoveImageIDs(ids...)
}

// ClearDailyRoutines clears all "daily_routines" edges to the DailyRoutine entity.
func (pvuo *ProgramVersionUpdateOne) ClearDailyRoutines() *ProgramVersionUpdateOne {
	pvuo.mutation.ClearDailyRoutines()
	return pvuo
}

// RemoveDailyRoutineIDs removes the "daily_routines" edge to DailyRoutine entities by IDs.
func (pvuo *ProgramVersionUpdateOne) RemoveDailyRoutineIDs(ids ...uint64) *ProgramVersionUpdateOne {
	pvuo.mutation.RemoveDailyRoutineIDs(ids...)
	return pvuo
}

// RemoveDailyRoutines removes "daily_routines" edges to DailyRoutine entities.
func (pvuo *ProgramVersionUpdateOne) RemoveDailyRoutines(d ...*DailyRoutine) *ProgramVersionUpdateOne {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pvuo.RemoveDailyRoutineIDs(ids...)
}

// Where appends a list predicates to the ProgramVersionUpdate builder.
func (pvuo *ProgramVersionUpdateOne) Where(ps ...predicate.ProgramVersion) *ProgramVersionUpdateOne {
	pvuo.mutation.Where(ps...)
	return pvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pvuo *ProgramVersionUpdateOne) Select(field string, fields ...string) *ProgramVersionUpdateOne {
	pvuo.fields = append([]string{field}, fields...)
	return pvuo
}

// Save executes the query and returns the updated ProgramVersion entity.
func (pvuo *ProgramVersionUpdateOne) Save(ctx context.Context) (*ProgramVersion, error) {
	return withHooks(ctx, pvuo.sqlSave, pvuo.mutation, pvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pvuo *ProgramVersionUpdateOne) SaveX(ctx context.Context) *ProgramVersion {
	node, err := pvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pvuo *ProgramVersionUpdateOne) Exec(ctx context.Context) error {
	_, err := pvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvuo *ProgramVersionUpdateOne) ExecX(ctx context.Context) {
	if err := pvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvuo *ProgramVersionUpdateOne) check() error {
	if _, ok := pvuo.mutation.ProgramID(); pvuo.mutation.ProgramCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProgramVersion.program"`)
	}
	return nil
}

func (pvuo *ProgramVersionUpdateOne) sqlSave(ctx context.Context) (_node *ProgramVersion, err error) {
	if err := pvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(programversion.Table, programversion.Columns, sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64))
	id, ok := pvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProgramVersion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, programversion.FieldID)
		for _, f := range fields {
			if !programversion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != programversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pvuo.mutation.ProgramCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   programversion.ProgramTable,
			Columns: []string{programversion.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.ProgramIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   programversion.ProgramTable,
			Columns: []string{programversion.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvuo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   programversion.ImagesTable,
			Columns: programversion.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.RemovedImagesIDs(); len(nodes) > 0 && !pvuo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   programversion.ImagesTable,
			Columns: programversion.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   programversion.ImagesTable,
			Columns: programversion.ImagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvuo.mutation.DailyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programversion.DailyRoutinesTable,
			Columns: []string{programversion.DailyRoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.RemovedDailyRoutinesIDs(); len(nodes) > 0 && !pvuo.mutation.DailyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programversion.DailyRoutinesTable,
			Columns: []string{programversion.DailyRoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.DailyRoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programversion.DailyRoutinesTable,
			Columns: []string{programversion.DailyRoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProgramVersion{config: pvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{programversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pvuo.mutation.done = true
	return _node, nil
}
