// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/predicate"
	"routine/internal/ent/program"
	"routine/internal/ent/programrelease"
	"routine/internal/ent/routine"
	"routine/internal/ent/s3programimage"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramReleaseUpdate is the builder for updating ProgramRelease entities.
type ProgramReleaseUpdate struct {
	config
	hooks    []Hook
	mutation *ProgramReleaseMutation
}

// Where appends a list predicates to the ProgramReleaseUpdate builder.
func (pru *ProgramReleaseUpdate) Where(ps ...predicate.ProgramRelease) *ProgramReleaseUpdate {
	pru.mutation.Where(ps...)
	return pru
}

// SetProgramID sets the "program" edge to the Program entity by ID.
func (pru *ProgramReleaseUpdate) SetProgramID(id int64) *ProgramReleaseUpdate {
	pru.mutation.SetProgramID(id)
	return pru
}

// SetProgram sets the "program" edge to the Program entity.
func (pru *ProgramReleaseUpdate) SetProgram(p *Program) *ProgramReleaseUpdate {
	return pru.SetProgramID(p.ID)
}

// AddS3ProgramImageIDs adds the "s3_program_images" edge to the S3ProgramImage entity by IDs.
func (pru *ProgramReleaseUpdate) AddS3ProgramImageIDs(ids ...int64) *ProgramReleaseUpdate {
	pru.mutation.AddS3ProgramImageIDs(ids...)
	return pru
}

// AddS3ProgramImages adds the "s3_program_images" edges to the S3ProgramImage entity.
func (pru *ProgramReleaseUpdate) AddS3ProgramImages(s ...*S3ProgramImage) *ProgramReleaseUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pru.AddS3ProgramImageIDs(ids...)
}

// AddRoutineIDs adds the "routines" edge to the Routine entity by IDs.
func (pru *ProgramReleaseUpdate) AddRoutineIDs(ids ...int64) *ProgramReleaseUpdate {
	pru.mutation.AddRoutineIDs(ids...)
	return pru
}

// AddRoutines adds the "routines" edges to the Routine entity.
func (pru *ProgramReleaseUpdate) AddRoutines(r ...*Routine) *ProgramReleaseUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pru.AddRoutineIDs(ids...)
}

// Mutation returns the ProgramReleaseMutation object of the builder.
func (pru *ProgramReleaseUpdate) Mutation() *ProgramReleaseMutation {
	return pru.mutation
}

// ClearProgram clears the "program" edge to the Program entity.
func (pru *ProgramReleaseUpdate) ClearProgram() *ProgramReleaseUpdate {
	pru.mutation.ClearProgram()
	return pru
}

// ClearS3ProgramImages clears all "s3_program_images" edges to the S3ProgramImage entity.
func (pru *ProgramReleaseUpdate) ClearS3ProgramImages() *ProgramReleaseUpdate {
	pru.mutation.ClearS3ProgramImages()
	return pru
}

// RemoveS3ProgramImageIDs removes the "s3_program_images" edge to S3ProgramImage entities by IDs.
func (pru *ProgramReleaseUpdate) RemoveS3ProgramImageIDs(ids ...int64) *ProgramReleaseUpdate {
	pru.mutation.RemoveS3ProgramImageIDs(ids...)
	return pru
}

// RemoveS3ProgramImages removes "s3_program_images" edges to S3ProgramImage entities.
func (pru *ProgramReleaseUpdate) RemoveS3ProgramImages(s ...*S3ProgramImage) *ProgramReleaseUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pru.RemoveS3ProgramImageIDs(ids...)
}

// ClearRoutines clears all "routines" edges to the Routine entity.
func (pru *ProgramReleaseUpdate) ClearRoutines() *ProgramReleaseUpdate {
	pru.mutation.ClearRoutines()
	return pru
}

// RemoveRoutineIDs removes the "routines" edge to Routine entities by IDs.
func (pru *ProgramReleaseUpdate) RemoveRoutineIDs(ids ...int64) *ProgramReleaseUpdate {
	pru.mutation.RemoveRoutineIDs(ids...)
	return pru
}

// RemoveRoutines removes "routines" edges to Routine entities.
func (pru *ProgramReleaseUpdate) RemoveRoutines(r ...*Routine) *ProgramReleaseUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pru.RemoveRoutineIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pru *ProgramReleaseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pru.sqlSave, pru.mutation, pru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pru *ProgramReleaseUpdate) SaveX(ctx context.Context) int {
	affected, err := pru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pru *ProgramReleaseUpdate) Exec(ctx context.Context) error {
	_, err := pru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pru *ProgramReleaseUpdate) ExecX(ctx context.Context) {
	if err := pru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pru *ProgramReleaseUpdate) check() error {
	if _, ok := pru.mutation.ProgramID(); pru.mutation.ProgramCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProgramRelease.program"`)
	}
	return nil
}

func (pru *ProgramReleaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(programrelease.Table, programrelease.Columns, sqlgraph.NewFieldSpec(programrelease.FieldID, field.TypeInt64))
	if ps := pru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pru.mutation.ProgramCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   programrelease.ProgramTable,
			Columns: []string{programrelease.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.ProgramIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   programrelease.ProgramTable,
			Columns: []string{programrelease.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pru.mutation.S3ProgramImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.S3ProgramImagesTable,
			Columns: []string{programrelease.S3ProgramImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(s3programimage.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RemovedS3ProgramImagesIDs(); len(nodes) > 0 && !pru.mutation.S3ProgramImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.S3ProgramImagesTable,
			Columns: []string{programrelease.S3ProgramImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(s3programimage.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.S3ProgramImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.S3ProgramImagesTable,
			Columns: []string{programrelease.S3ProgramImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(s3programimage.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pru.mutation.RoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.RoutinesTable,
			Columns: []string{programrelease.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RemovedRoutinesIDs(); len(nodes) > 0 && !pru.mutation.RoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.RoutinesTable,
			Columns: []string{programrelease.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.RoutinesTable,
			Columns: []string{programrelease.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{programrelease.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pru.mutation.done = true
	return n, nil
}

// ProgramReleaseUpdateOne is the builder for updating a single ProgramRelease entity.
type ProgramReleaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProgramReleaseMutation
}

// SetProgramID sets the "program" edge to the Program entity by ID.
func (pruo *ProgramReleaseUpdateOne) SetProgramID(id int64) *ProgramReleaseUpdateOne {
	pruo.mutation.SetProgramID(id)
	return pruo
}

// SetProgram sets the "program" edge to the Program entity.
func (pruo *ProgramReleaseUpdateOne) SetProgram(p *Program) *ProgramReleaseUpdateOne {
	return pruo.SetProgramID(p.ID)
}

// AddS3ProgramImageIDs adds the "s3_program_images" edge to the S3ProgramImage entity by IDs.
func (pruo *ProgramReleaseUpdateOne) AddS3ProgramImageIDs(ids ...int64) *ProgramReleaseUpdateOne {
	pruo.mutation.AddS3ProgramImageIDs(ids...)
	return pruo
}

// AddS3ProgramImages adds the "s3_program_images" edges to the S3ProgramImage entity.
func (pruo *ProgramReleaseUpdateOne) AddS3ProgramImages(s ...*S3ProgramImage) *ProgramReleaseUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pruo.AddS3ProgramImageIDs(ids...)
}

// AddRoutineIDs adds the "routines" edge to the Routine entity by IDs.
func (pruo *ProgramReleaseUpdateOne) AddRoutineIDs(ids ...int64) *ProgramReleaseUpdateOne {
	pruo.mutation.AddRoutineIDs(ids...)
	return pruo
}

// AddRoutines adds the "routines" edges to the Routine entity.
func (pruo *ProgramReleaseUpdateOne) AddRoutines(r ...*Routine) *ProgramReleaseUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pruo.AddRoutineIDs(ids...)
}

// Mutation returns the ProgramReleaseMutation object of the builder.
func (pruo *ProgramReleaseUpdateOne) Mutation() *ProgramReleaseMutation {
	return pruo.mutation
}

// ClearProgram clears the "program" edge to the Program entity.
func (pruo *ProgramReleaseUpdateOne) ClearProgram() *ProgramReleaseUpdateOne {
	pruo.mutation.ClearProgram()
	return pruo
}

// ClearS3ProgramImages clears all "s3_program_images" edges to the S3ProgramImage entity.
func (pruo *ProgramReleaseUpdateOne) ClearS3ProgramImages() *ProgramReleaseUpdateOne {
	pruo.mutation.ClearS3ProgramImages()
	return pruo
}

// RemoveS3ProgramImageIDs removes the "s3_program_images" edge to S3ProgramImage entities by IDs.
func (pruo *ProgramReleaseUpdateOne) RemoveS3ProgramImageIDs(ids ...int64) *ProgramReleaseUpdateOne {
	pruo.mutation.RemoveS3ProgramImageIDs(ids...)
	return pruo
}

// RemoveS3ProgramImages removes "s3_program_images" edges to S3ProgramImage entities.
func (pruo *ProgramReleaseUpdateOne) RemoveS3ProgramImages(s ...*S3ProgramImage) *ProgramReleaseUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pruo.RemoveS3ProgramImageIDs(ids...)
}

// ClearRoutines clears all "routines" edges to the Routine entity.
func (pruo *ProgramReleaseUpdateOne) ClearRoutines() *ProgramReleaseUpdateOne {
	pruo.mutation.ClearRoutines()
	return pruo
}

// RemoveRoutineIDs removes the "routines" edge to Routine entities by IDs.
func (pruo *ProgramReleaseUpdateOne) RemoveRoutineIDs(ids ...int64) *ProgramReleaseUpdateOne {
	pruo.mutation.RemoveRoutineIDs(ids...)
	return pruo
}

// RemoveRoutines removes "routines" edges to Routine entities.
func (pruo *ProgramReleaseUpdateOne) RemoveRoutines(r ...*Routine) *ProgramReleaseUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pruo.RemoveRoutineIDs(ids...)
}

// Where appends a list predicates to the ProgramReleaseUpdate builder.
func (pruo *ProgramReleaseUpdateOne) Where(ps ...predicate.ProgramRelease) *ProgramReleaseUpdateOne {
	pruo.mutation.Where(ps...)
	return pruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pruo *ProgramReleaseUpdateOne) Select(field string, fields ...string) *ProgramReleaseUpdateOne {
	pruo.fields = append([]string{field}, fields...)
	return pruo
}

// Save executes the query and returns the updated ProgramRelease entity.
func (pruo *ProgramReleaseUpdateOne) Save(ctx context.Context) (*ProgramRelease, error) {
	return withHooks(ctx, pruo.sqlSave, pruo.mutation, pruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pruo *ProgramReleaseUpdateOne) SaveX(ctx context.Context) *ProgramRelease {
	node, err := pruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pruo *ProgramReleaseUpdateOne) Exec(ctx context.Context) error {
	_, err := pruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pruo *ProgramReleaseUpdateOne) ExecX(ctx context.Context) {
	if err := pruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pruo *ProgramReleaseUpdateOne) check() error {
	if _, ok := pruo.mutation.ProgramID(); pruo.mutation.ProgramCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProgramRelease.program"`)
	}
	return nil
}

func (pruo *ProgramReleaseUpdateOne) sqlSave(ctx context.Context) (_node *ProgramRelease, err error) {
	if err := pruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(programrelease.Table, programrelease.Columns, sqlgraph.NewFieldSpec(programrelease.FieldID, field.TypeInt64))
	id, ok := pruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProgramRelease.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, programrelease.FieldID)
		for _, f := range fields {
			if !programrelease.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != programrelease.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pruo.mutation.ProgramCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   programrelease.ProgramTable,
			Columns: []string{programrelease.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.ProgramIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   programrelease.ProgramTable,
			Columns: []string{programrelease.ProgramColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(program.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pruo.mutation.S3ProgramImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.S3ProgramImagesTable,
			Columns: []string{programrelease.S3ProgramImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(s3programimage.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RemovedS3ProgramImagesIDs(); len(nodes) > 0 && !pruo.mutation.S3ProgramImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.S3ProgramImagesTable,
			Columns: []string{programrelease.S3ProgramImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(s3programimage.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.S3ProgramImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.S3ProgramImagesTable,
			Columns: []string{programrelease.S3ProgramImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(s3programimage.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pruo.mutation.RoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.RoutinesTable,
			Columns: []string{programrelease.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RemovedRoutinesIDs(); len(nodes) > 0 && !pruo.mutation.RoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.RoutinesTable,
			Columns: []string{programrelease.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   programrelease.RoutinesTable,
			Columns: []string{programrelease.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProgramRelease{config: pruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{programrelease.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pruo.mutation.done = true
	return _node, nil
}