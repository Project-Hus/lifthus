// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/internal/ent/program"
	"routine/internal/ent/programversion"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramCreate is the builder for creating a Program entity.
type ProgramCreate struct {
	config
	mutation *ProgramMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (pc *ProgramCreate) SetCode(s string) *ProgramCreate {
	pc.mutation.SetCode(s)
	return pc
}

// SetProgramType sets the "program_type" field.
func (pc *ProgramCreate) SetProgramType(pt program.ProgramType) *ProgramCreate {
	pc.mutation.SetProgramType(pt)
	return pc
}

// SetTitle sets the "title" field.
func (pc *ProgramCreate) SetTitle(s string) *ProgramCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetAuthor sets the "author" field.
func (pc *ProgramCreate) SetAuthor(u uint64) *ProgramCreate {
	pc.mutation.SetAuthor(u)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProgramCreate) SetCreatedAt(t time.Time) *ProgramCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetVersionDerivedFrom sets the "version_derived_from" field.
func (pc *ProgramCreate) SetVersionDerivedFrom(s string) *ProgramCreate {
	pc.mutation.SetVersionDerivedFrom(s)
	return pc
}

// SetID sets the "id" field.
func (pc *ProgramCreate) SetID(u uint64) *ProgramCreate {
	pc.mutation.SetID(u)
	return pc
}

// AddProgramVersionIDs adds the "program_versions" edge to the ProgramVersion entity by IDs.
func (pc *ProgramCreate) AddProgramVersionIDs(ids ...uint64) *ProgramCreate {
	pc.mutation.AddProgramVersionIDs(ids...)
	return pc
}

// AddProgramVersions adds the "program_versions" edges to the ProgramVersion entity.
func (pc *ProgramCreate) AddProgramVersions(p ...*ProgramVersion) *ProgramCreate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProgramVersionIDs(ids...)
}

// Mutation returns the ProgramMutation object of the builder.
func (pc *ProgramCreate) Mutation() *ProgramMutation {
	return pc.mutation
}

// Save creates the Program in the database.
func (pc *ProgramCreate) Save(ctx context.Context) (*Program, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProgramCreate) SaveX(ctx context.Context) *Program {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProgramCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProgramCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProgramCreate) check() error {
	if _, ok := pc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Program.code"`)}
	}
	if v, ok := pc.mutation.Code(); ok {
		if err := program.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Program.code": %w`, err)}
		}
	}
	if _, ok := pc.mutation.ProgramType(); !ok {
		return &ValidationError{Name: "program_type", err: errors.New(`ent: missing required field "Program.program_type"`)}
	}
	if v, ok := pc.mutation.ProgramType(); ok {
		if err := program.ProgramTypeValidator(v); err != nil {
			return &ValidationError{Name: "program_type", err: fmt.Errorf(`ent: validator failed for field "Program.program_type": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Program.title"`)}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := program.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Program.title": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "Program.author"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Program.created_at"`)}
	}
	if _, ok := pc.mutation.VersionDerivedFrom(); !ok {
		return &ValidationError{Name: "version_derived_from", err: errors.New(`ent: missing required field "Program.version_derived_from"`)}
	}
	if v, ok := pc.mutation.VersionDerivedFrom(); ok {
		if err := program.VersionDerivedFromValidator(v); err != nil {
			return &ValidationError{Name: "version_derived_from", err: fmt.Errorf(`ent: validator failed for field "Program.version_derived_from": %w`, err)}
		}
	}
	return nil
}

func (pc *ProgramCreate) sqlSave(ctx context.Context) (*Program, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProgramCreate) createSpec() (*Program, *sqlgraph.CreateSpec) {
	var (
		_node = &Program{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(program.Table, sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Code(); ok {
		_spec.SetField(program.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := pc.mutation.ProgramType(); ok {
		_spec.SetField(program.FieldProgramType, field.TypeEnum, value)
		_node.ProgramType = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(program.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Author(); ok {
		_spec.SetField(program.FieldAuthor, field.TypeUint64, value)
		_node.Author = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(program.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.VersionDerivedFrom(); ok {
		_spec.SetField(program.FieldVersionDerivedFrom, field.TypeString, value)
		_node.VersionDerivedFrom = value
	}
	if nodes := pc.mutation.ProgramVersionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProgramCreateBulk is the builder for creating many Program entities in bulk.
type ProgramCreateBulk struct {
	config
	builders []*ProgramCreate
}

// Save creates the Program entities in the database.
func (pcb *ProgramCreateBulk) Save(ctx context.Context) ([]*Program, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Program, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProgramMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProgramCreateBulk) SaveX(ctx context.Context) []*Program {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProgramCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProgramCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
