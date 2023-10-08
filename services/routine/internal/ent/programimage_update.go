// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/image"
	"routine/internal/ent/predicate"
	"routine/internal/ent/programimage"
	"routine/internal/ent/programversion"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramImageUpdate is the builder for updating ProgramImage entities.
type ProgramImageUpdate struct {
	config
	hooks    []Hook
	mutation *ProgramImageMutation
}

// Where appends a list predicates to the ProgramImageUpdate builder.
func (piu *ProgramImageUpdate) Where(ps ...predicate.ProgramImage) *ProgramImageUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetProgramVersionID sets the "program_version_id" field.
func (piu *ProgramImageUpdate) SetProgramVersionID(u uint64) *ProgramImageUpdate {
	piu.mutation.SetProgramVersionID(u)
	return piu
}

// SetImageID sets the "image_id" field.
func (piu *ProgramImageUpdate) SetImageID(u uint64) *ProgramImageUpdate {
	piu.mutation.SetImageID(u)
	return piu
}

// SetProgramVersion sets the "program_version" edge to the ProgramVersion entity.
func (piu *ProgramImageUpdate) SetProgramVersion(p *ProgramVersion) *ProgramImageUpdate {
	return piu.SetProgramVersionID(p.ID)
}

// SetImage sets the "image" edge to the Image entity.
func (piu *ProgramImageUpdate) SetImage(i *Image) *ProgramImageUpdate {
	return piu.SetImageID(i.ID)
}

// Mutation returns the ProgramImageMutation object of the builder.
func (piu *ProgramImageUpdate) Mutation() *ProgramImageMutation {
	return piu.mutation
}

// ClearProgramVersion clears the "program_version" edge to the ProgramVersion entity.
func (piu *ProgramImageUpdate) ClearProgramVersion() *ProgramImageUpdate {
	piu.mutation.ClearProgramVersion()
	return piu
}

// ClearImage clears the "image" edge to the Image entity.
func (piu *ProgramImageUpdate) ClearImage() *ProgramImageUpdate {
	piu.mutation.ClearImage()
	return piu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *ProgramImageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, piu.sqlSave, piu.mutation, piu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piu *ProgramImageUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *ProgramImageUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *ProgramImageUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piu *ProgramImageUpdate) check() error {
	if _, ok := piu.mutation.ProgramVersionID(); piu.mutation.ProgramVersionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProgramImage.program_version"`)
	}
	if _, ok := piu.mutation.ImageID(); piu.mutation.ImageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProgramImage.image"`)
	}
	return nil
}

func (piu *ProgramImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := piu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(programimage.Table, programimage.Columns, sqlgraph.NewFieldSpec(programimage.FieldID, field.TypeUint64))
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if piu.mutation.ProgramVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programimage.ProgramVersionTable,
			Columns: []string{programimage.ProgramVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.ProgramVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programimage.ProgramVersionTable,
			Columns: []string{programimage.ProgramVersionColumn},
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
	if piu.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programimage.ImageTable,
			Columns: []string{programimage.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programimage.ImageTable,
			Columns: []string{programimage.ImageColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{programimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	piu.mutation.done = true
	return n, nil
}

// ProgramImageUpdateOne is the builder for updating a single ProgramImage entity.
type ProgramImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProgramImageMutation
}

// SetProgramVersionID sets the "program_version_id" field.
func (piuo *ProgramImageUpdateOne) SetProgramVersionID(u uint64) *ProgramImageUpdateOne {
	piuo.mutation.SetProgramVersionID(u)
	return piuo
}

// SetImageID sets the "image_id" field.
func (piuo *ProgramImageUpdateOne) SetImageID(u uint64) *ProgramImageUpdateOne {
	piuo.mutation.SetImageID(u)
	return piuo
}

// SetProgramVersion sets the "program_version" edge to the ProgramVersion entity.
func (piuo *ProgramImageUpdateOne) SetProgramVersion(p *ProgramVersion) *ProgramImageUpdateOne {
	return piuo.SetProgramVersionID(p.ID)
}

// SetImage sets the "image" edge to the Image entity.
func (piuo *ProgramImageUpdateOne) SetImage(i *Image) *ProgramImageUpdateOne {
	return piuo.SetImageID(i.ID)
}

// Mutation returns the ProgramImageMutation object of the builder.
func (piuo *ProgramImageUpdateOne) Mutation() *ProgramImageMutation {
	return piuo.mutation
}

// ClearProgramVersion clears the "program_version" edge to the ProgramVersion entity.
func (piuo *ProgramImageUpdateOne) ClearProgramVersion() *ProgramImageUpdateOne {
	piuo.mutation.ClearProgramVersion()
	return piuo
}

// ClearImage clears the "image" edge to the Image entity.
func (piuo *ProgramImageUpdateOne) ClearImage() *ProgramImageUpdateOne {
	piuo.mutation.ClearImage()
	return piuo
}

// Where appends a list predicates to the ProgramImageUpdate builder.
func (piuo *ProgramImageUpdateOne) Where(ps ...predicate.ProgramImage) *ProgramImageUpdateOne {
	piuo.mutation.Where(ps...)
	return piuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *ProgramImageUpdateOne) Select(field string, fields ...string) *ProgramImageUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated ProgramImage entity.
func (piuo *ProgramImageUpdateOne) Save(ctx context.Context) (*ProgramImage, error) {
	return withHooks(ctx, piuo.sqlSave, piuo.mutation, piuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *ProgramImageUpdateOne) SaveX(ctx context.Context) *ProgramImage {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *ProgramImageUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *ProgramImageUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piuo *ProgramImageUpdateOne) check() error {
	if _, ok := piuo.mutation.ProgramVersionID(); piuo.mutation.ProgramVersionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProgramImage.program_version"`)
	}
	if _, ok := piuo.mutation.ImageID(); piuo.mutation.ImageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProgramImage.image"`)
	}
	return nil
}

func (piuo *ProgramImageUpdateOne) sqlSave(ctx context.Context) (_node *ProgramImage, err error) {
	if err := piuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(programimage.Table, programimage.Columns, sqlgraph.NewFieldSpec(programimage.FieldID, field.TypeUint64))
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProgramImage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, programimage.FieldID)
		for _, f := range fields {
			if !programimage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != programimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if piuo.mutation.ProgramVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programimage.ProgramVersionTable,
			Columns: []string{programimage.ProgramVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(programversion.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.ProgramVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programimage.ProgramVersionTable,
			Columns: []string{programimage.ProgramVersionColumn},
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
	if piuo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programimage.ImageTable,
			Columns: []string{programimage.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   programimage.ImageTable,
			Columns: []string{programimage.ImageColumn},
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
	_node = &ProgramImage{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{programimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	piuo.mutation.done = true
	return _node, nil
}
