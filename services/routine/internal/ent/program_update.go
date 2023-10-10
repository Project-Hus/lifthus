// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/predicate"
	"routine/internal/ent/program"
	"routine/internal/ent/programversion"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramUpdate is the builder for updating Program entities.
type ProgramUpdate struct {
	config
	hooks    []Hook
	mutation *ProgramMutation
}

// Where appends a list predicates to the ProgramUpdate builder.
func (pu *ProgramUpdate) Where(ps ...predicate.Program) *ProgramUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// AddProgramVersionIDs adds the "program_versions" edge to the ProgramVersion entity by IDs.
func (pu *ProgramUpdate) AddProgramVersionIDs(ids ...uint64) *ProgramUpdate {
	pu.mutation.AddProgramVersionIDs(ids...)
	return pu
}

// AddProgramVersions adds the "program_versions" edges to the ProgramVersion entity.
func (pu *ProgramUpdate) AddProgramVersions(p ...*ProgramVersion) *ProgramUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddProgramVersionIDs(ids...)
}

// Mutation returns the ProgramMutation object of the builder.
func (pu *ProgramUpdate) Mutation() *ProgramMutation {
	return pu.mutation
}

// ClearProgramVersions clears all "program_versions" edges to the ProgramVersion entity.
func (pu *ProgramUpdate) ClearProgramVersions() *ProgramUpdate {
	pu.mutation.ClearProgramVersions()
	return pu
}

// RemoveProgramVersionIDs removes the "program_versions" edge to ProgramVersion entities by IDs.
func (pu *ProgramUpdate) RemoveProgramVersionIDs(ids ...uint64) *ProgramUpdate {
	pu.mutation.RemoveProgramVersionIDs(ids...)
	return pu
}

// RemoveProgramVersions removes "program_versions" edges to ProgramVersion entities.
func (pu *ProgramUpdate) RemoveProgramVersions(p ...*ProgramVersion) *ProgramUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveProgramVersionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProgramUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProgramUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProgramUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProgramUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProgramUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(program.Table, program.Columns, sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.VersionDerivedFromCleared() {
		_spec.ClearField(program.FieldVersionDerivedFrom, field.TypeString)
	}
	if pu.mutation.ProgramVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   program.ProgramVersionsTable,
			Columns: []string{program.ProgramVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedProgramVersionsIDs(); len(nodes) > 0 && !pu.mutation.ProgramVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   program.ProgramVersionsTable,
			Columns: []string{program.ProgramVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProgramVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   program.ProgramVersionsTable,
			Columns: []string{program.ProgramVersionsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{program.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProgramUpdateOne is the builder for updating a single Program entity.
type ProgramUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProgramMutation
}

// AddProgramVersionIDs adds the "program_versions" edge to the ProgramVersion entity by IDs.
func (puo *ProgramUpdateOne) AddProgramVersionIDs(ids ...uint64) *ProgramUpdateOne {
	puo.mutation.AddProgramVersionIDs(ids...)
	return puo
}

// AddProgramVersions adds the "program_versions" edges to the ProgramVersion entity.
func (puo *ProgramUpdateOne) AddProgramVersions(p ...*ProgramVersion) *ProgramUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddProgramVersionIDs(ids...)
}

// Mutation returns the ProgramMutation object of the builder.
func (puo *ProgramUpdateOne) Mutation() *ProgramMutation {
	return puo.mutation
}

// ClearProgramVersions clears all "program_versions" edges to the ProgramVersion entity.
func (puo *ProgramUpdateOne) ClearProgramVersions() *ProgramUpdateOne {
	puo.mutation.ClearProgramVersions()
	return puo
}

// RemoveProgramVersionIDs removes the "program_versions" edge to ProgramVersion entities by IDs.
func (puo *ProgramUpdateOne) RemoveProgramVersionIDs(ids ...uint64) *ProgramUpdateOne {
	puo.mutation.RemoveProgramVersionIDs(ids...)
	return puo
}

// RemoveProgramVersions removes "program_versions" edges to ProgramVersion entities.
func (puo *ProgramUpdateOne) RemoveProgramVersions(p ...*ProgramVersion) *ProgramUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveProgramVersionIDs(ids...)
}

// Where appends a list predicates to the ProgramUpdate builder.
func (puo *ProgramUpdateOne) Where(ps ...predicate.Program) *ProgramUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProgramUpdateOne) Select(field string, fields ...string) *ProgramUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Program entity.
func (puo *ProgramUpdateOne) Save(ctx context.Context) (*Program, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProgramUpdateOne) SaveX(ctx context.Context) *Program {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProgramUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProgramUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProgramUpdateOne) sqlSave(ctx context.Context) (_node *Program, err error) {
	_spec := sqlgraph.NewUpdateSpec(program.Table, program.Columns, sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Program.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, program.FieldID)
		for _, f := range fields {
			if !program.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != program.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if puo.mutation.VersionDerivedFromCleared() {
		_spec.ClearField(program.FieldVersionDerivedFrom, field.TypeString)
	}
	if puo.mutation.ProgramVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   program.ProgramVersionsTable,
			Columns: []string{program.ProgramVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedProgramVersionsIDs(); len(nodes) > 0 && !puo.mutation.ProgramVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   program.ProgramVersionsTable,
			Columns: []string{program.ProgramVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProgramVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   program.ProgramVersionsTable,
			Columns: []string{program.ProgramVersionsColumn},
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
	_node = &Program{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{program.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}