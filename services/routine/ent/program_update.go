// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/dailyroutine"
	"routine/ent/predicate"
	"routine/ent/program"
	"routine/ent/tag"
	"routine/ent/weeklyroutine"
	"time"

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

// SetTitle sets the "title" field.
func (pu *ProgramUpdate) SetTitle(s string) *ProgramUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetType sets the "type" field.
func (pu *ProgramUpdate) SetType(pr program.Type) *ProgramUpdate {
	pu.mutation.SetType(pr)
	return pu
}

// SetAuthor sets the "author" field.
func (pu *ProgramUpdate) SetAuthor(u uint64) *ProgramUpdate {
	pu.mutation.ResetAuthor()
	pu.mutation.SetAuthor(u)
	return pu
}

// AddAuthor adds u to the "author" field.
func (pu *ProgramUpdate) AddAuthor(u int64) *ProgramUpdate {
	pu.mutation.AddAuthor(u)
	return pu
}

// SetImage sets the "image" field.
func (pu *ProgramUpdate) SetImage(s string) *ProgramUpdate {
	pu.mutation.SetImage(s)
	return pu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (pu *ProgramUpdate) SetNillableImage(s *string) *ProgramUpdate {
	if s != nil {
		pu.SetImage(*s)
	}
	return pu
}

// ClearImage clears the value of the "image" field.
func (pu *ProgramUpdate) ClearImage() *ProgramUpdate {
	pu.mutation.ClearImage()
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProgramUpdate) SetDescription(s string) *ProgramUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProgramUpdate) SetNillableDescription(s *string) *ProgramUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *ProgramUpdate) ClearDescription() *ProgramUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProgramUpdate) SetUpdatedAt(t time.Time) *ProgramUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (pu *ProgramUpdate) AddTagIDs(ids ...uint64) *ProgramUpdate {
	pu.mutation.AddTagIDs(ids...)
	return pu
}

// AddTags adds the "tags" edges to the Tag entity.
func (pu *ProgramUpdate) AddTags(t ...*Tag) *ProgramUpdate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTagIDs(ids...)
}

// AddWeeklyRoutineIDs adds the "weekly_routines" edge to the WeeklyRoutine entity by IDs.
func (pu *ProgramUpdate) AddWeeklyRoutineIDs(ids ...uint64) *ProgramUpdate {
	pu.mutation.AddWeeklyRoutineIDs(ids...)
	return pu
}

// AddWeeklyRoutines adds the "weekly_routines" edges to the WeeklyRoutine entity.
func (pu *ProgramUpdate) AddWeeklyRoutines(w ...*WeeklyRoutine) *ProgramUpdate {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pu.AddWeeklyRoutineIDs(ids...)
}

// AddDailyRoutineIDs adds the "daily_routines" edge to the DailyRoutine entity by IDs.
func (pu *ProgramUpdate) AddDailyRoutineIDs(ids ...uint64) *ProgramUpdate {
	pu.mutation.AddDailyRoutineIDs(ids...)
	return pu
}

// AddDailyRoutines adds the "daily_routines" edges to the DailyRoutine entity.
func (pu *ProgramUpdate) AddDailyRoutines(d ...*DailyRoutine) *ProgramUpdate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.AddDailyRoutineIDs(ids...)
}

// Mutation returns the ProgramMutation object of the builder.
func (pu *ProgramUpdate) Mutation() *ProgramMutation {
	return pu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (pu *ProgramUpdate) ClearTags() *ProgramUpdate {
	pu.mutation.ClearTags()
	return pu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (pu *ProgramUpdate) RemoveTagIDs(ids ...uint64) *ProgramUpdate {
	pu.mutation.RemoveTagIDs(ids...)
	return pu
}

// RemoveTags removes "tags" edges to Tag entities.
func (pu *ProgramUpdate) RemoveTags(t ...*Tag) *ProgramUpdate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTagIDs(ids...)
}

// ClearWeeklyRoutines clears all "weekly_routines" edges to the WeeklyRoutine entity.
func (pu *ProgramUpdate) ClearWeeklyRoutines() *ProgramUpdate {
	pu.mutation.ClearWeeklyRoutines()
	return pu
}

// RemoveWeeklyRoutineIDs removes the "weekly_routines" edge to WeeklyRoutine entities by IDs.
func (pu *ProgramUpdate) RemoveWeeklyRoutineIDs(ids ...uint64) *ProgramUpdate {
	pu.mutation.RemoveWeeklyRoutineIDs(ids...)
	return pu
}

// RemoveWeeklyRoutines removes "weekly_routines" edges to WeeklyRoutine entities.
func (pu *ProgramUpdate) RemoveWeeklyRoutines(w ...*WeeklyRoutine) *ProgramUpdate {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pu.RemoveWeeklyRoutineIDs(ids...)
}

// ClearDailyRoutines clears all "daily_routines" edges to the DailyRoutine entity.
func (pu *ProgramUpdate) ClearDailyRoutines() *ProgramUpdate {
	pu.mutation.ClearDailyRoutines()
	return pu
}

// RemoveDailyRoutineIDs removes the "daily_routines" edge to DailyRoutine entities by IDs.
func (pu *ProgramUpdate) RemoveDailyRoutineIDs(ids ...uint64) *ProgramUpdate {
	pu.mutation.RemoveDailyRoutineIDs(ids...)
	return pu
}

// RemoveDailyRoutines removes "daily_routines" edges to DailyRoutine entities.
func (pu *ProgramUpdate) RemoveDailyRoutines(d ...*DailyRoutine) *ProgramUpdate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.RemoveDailyRoutineIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProgramUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
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

// defaults sets the default values of the builder before save.
func (pu *ProgramUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := program.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProgramUpdate) check() error {
	if v, ok := pu.mutation.Title(); ok {
		if err := program.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Program.title": %w`, err)}
		}
	}
	if v, ok := pu.mutation.GetType(); ok {
		if err := program.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Program.type": %w`, err)}
		}
	}
	return nil
}

func (pu *ProgramUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(program.Table, program.Columns, sqlgraph.NewFieldSpec(program.FieldID, field.TypeUint64))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(program.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(program.FieldType, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.Author(); ok {
		_spec.SetField(program.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := pu.mutation.AddedAuthor(); ok {
		_spec.AddField(program.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := pu.mutation.Image(); ok {
		_spec.SetField(program.FieldImage, field.TypeString, value)
	}
	if pu.mutation.ImageCleared() {
		_spec.ClearField(program.FieldImage, field.TypeString)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(program.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(program.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(program.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   program.TagsTable,
			Columns: program.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !pu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   program.TagsTable,
			Columns: program.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   program.TagsTable,
			Columns: program.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.WeeklyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.WeeklyRoutinesTable,
			Columns: program.WeeklyRoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedWeeklyRoutinesIDs(); len(nodes) > 0 && !pu.mutation.WeeklyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.WeeklyRoutinesTable,
			Columns: program.WeeklyRoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.WeeklyRoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.WeeklyRoutinesTable,
			Columns: program.WeeklyRoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DailyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.DailyRoutinesTable,
			Columns: program.DailyRoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedDailyRoutinesIDs(); len(nodes) > 0 && !pu.mutation.DailyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.DailyRoutinesTable,
			Columns: program.DailyRoutinesPrimaryKey,
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
	if nodes := pu.mutation.DailyRoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.DailyRoutinesTable,
			Columns: program.DailyRoutinesPrimaryKey,
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

// SetTitle sets the "title" field.
func (puo *ProgramUpdateOne) SetTitle(s string) *ProgramUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetType sets the "type" field.
func (puo *ProgramUpdateOne) SetType(pr program.Type) *ProgramUpdateOne {
	puo.mutation.SetType(pr)
	return puo
}

// SetAuthor sets the "author" field.
func (puo *ProgramUpdateOne) SetAuthor(u uint64) *ProgramUpdateOne {
	puo.mutation.ResetAuthor()
	puo.mutation.SetAuthor(u)
	return puo
}

// AddAuthor adds u to the "author" field.
func (puo *ProgramUpdateOne) AddAuthor(u int64) *ProgramUpdateOne {
	puo.mutation.AddAuthor(u)
	return puo
}

// SetImage sets the "image" field.
func (puo *ProgramUpdateOne) SetImage(s string) *ProgramUpdateOne {
	puo.mutation.SetImage(s)
	return puo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (puo *ProgramUpdateOne) SetNillableImage(s *string) *ProgramUpdateOne {
	if s != nil {
		puo.SetImage(*s)
	}
	return puo
}

// ClearImage clears the value of the "image" field.
func (puo *ProgramUpdateOne) ClearImage() *ProgramUpdateOne {
	puo.mutation.ClearImage()
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProgramUpdateOne) SetDescription(s string) *ProgramUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProgramUpdateOne) SetNillableDescription(s *string) *ProgramUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *ProgramUpdateOne) ClearDescription() *ProgramUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProgramUpdateOne) SetUpdatedAt(t time.Time) *ProgramUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (puo *ProgramUpdateOne) AddTagIDs(ids ...uint64) *ProgramUpdateOne {
	puo.mutation.AddTagIDs(ids...)
	return puo
}

// AddTags adds the "tags" edges to the Tag entity.
func (puo *ProgramUpdateOne) AddTags(t ...*Tag) *ProgramUpdateOne {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTagIDs(ids...)
}

// AddWeeklyRoutineIDs adds the "weekly_routines" edge to the WeeklyRoutine entity by IDs.
func (puo *ProgramUpdateOne) AddWeeklyRoutineIDs(ids ...uint64) *ProgramUpdateOne {
	puo.mutation.AddWeeklyRoutineIDs(ids...)
	return puo
}

// AddWeeklyRoutines adds the "weekly_routines" edges to the WeeklyRoutine entity.
func (puo *ProgramUpdateOne) AddWeeklyRoutines(w ...*WeeklyRoutine) *ProgramUpdateOne {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return puo.AddWeeklyRoutineIDs(ids...)
}

// AddDailyRoutineIDs adds the "daily_routines" edge to the DailyRoutine entity by IDs.
func (puo *ProgramUpdateOne) AddDailyRoutineIDs(ids ...uint64) *ProgramUpdateOne {
	puo.mutation.AddDailyRoutineIDs(ids...)
	return puo
}

// AddDailyRoutines adds the "daily_routines" edges to the DailyRoutine entity.
func (puo *ProgramUpdateOne) AddDailyRoutines(d ...*DailyRoutine) *ProgramUpdateOne {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.AddDailyRoutineIDs(ids...)
}

// Mutation returns the ProgramMutation object of the builder.
func (puo *ProgramUpdateOne) Mutation() *ProgramMutation {
	return puo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (puo *ProgramUpdateOne) ClearTags() *ProgramUpdateOne {
	puo.mutation.ClearTags()
	return puo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (puo *ProgramUpdateOne) RemoveTagIDs(ids ...uint64) *ProgramUpdateOne {
	puo.mutation.RemoveTagIDs(ids...)
	return puo
}

// RemoveTags removes "tags" edges to Tag entities.
func (puo *ProgramUpdateOne) RemoveTags(t ...*Tag) *ProgramUpdateOne {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTagIDs(ids...)
}

// ClearWeeklyRoutines clears all "weekly_routines" edges to the WeeklyRoutine entity.
func (puo *ProgramUpdateOne) ClearWeeklyRoutines() *ProgramUpdateOne {
	puo.mutation.ClearWeeklyRoutines()
	return puo
}

// RemoveWeeklyRoutineIDs removes the "weekly_routines" edge to WeeklyRoutine entities by IDs.
func (puo *ProgramUpdateOne) RemoveWeeklyRoutineIDs(ids ...uint64) *ProgramUpdateOne {
	puo.mutation.RemoveWeeklyRoutineIDs(ids...)
	return puo
}

// RemoveWeeklyRoutines removes "weekly_routines" edges to WeeklyRoutine entities.
func (puo *ProgramUpdateOne) RemoveWeeklyRoutines(w ...*WeeklyRoutine) *ProgramUpdateOne {
	ids := make([]uint64, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return puo.RemoveWeeklyRoutineIDs(ids...)
}

// ClearDailyRoutines clears all "daily_routines" edges to the DailyRoutine entity.
func (puo *ProgramUpdateOne) ClearDailyRoutines() *ProgramUpdateOne {
	puo.mutation.ClearDailyRoutines()
	return puo
}

// RemoveDailyRoutineIDs removes the "daily_routines" edge to DailyRoutine entities by IDs.
func (puo *ProgramUpdateOne) RemoveDailyRoutineIDs(ids ...uint64) *ProgramUpdateOne {
	puo.mutation.RemoveDailyRoutineIDs(ids...)
	return puo
}

// RemoveDailyRoutines removes "daily_routines" edges to DailyRoutine entities.
func (puo *ProgramUpdateOne) RemoveDailyRoutines(d ...*DailyRoutine) *ProgramUpdateOne {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.RemoveDailyRoutineIDs(ids...)
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
	puo.defaults()
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

// defaults sets the default values of the builder before save.
func (puo *ProgramUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := program.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProgramUpdateOne) check() error {
	if v, ok := puo.mutation.Title(); ok {
		if err := program.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Program.title": %w`, err)}
		}
	}
	if v, ok := puo.mutation.GetType(); ok {
		if err := program.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Program.type": %w`, err)}
		}
	}
	return nil
}

func (puo *ProgramUpdateOne) sqlSave(ctx context.Context) (_node *Program, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
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
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(program.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(program.FieldType, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.Author(); ok {
		_spec.SetField(program.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := puo.mutation.AddedAuthor(); ok {
		_spec.AddField(program.FieldAuthor, field.TypeUint64, value)
	}
	if value, ok := puo.mutation.Image(); ok {
		_spec.SetField(program.FieldImage, field.TypeString, value)
	}
	if puo.mutation.ImageCleared() {
		_spec.ClearField(program.FieldImage, field.TypeString)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(program.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(program.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(program.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   program.TagsTable,
			Columns: program.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !puo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   program.TagsTable,
			Columns: program.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   program.TagsTable,
			Columns: program.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.WeeklyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.WeeklyRoutinesTable,
			Columns: program.WeeklyRoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedWeeklyRoutinesIDs(); len(nodes) > 0 && !puo.mutation.WeeklyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.WeeklyRoutinesTable,
			Columns: program.WeeklyRoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.WeeklyRoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.WeeklyRoutinesTable,
			Columns: program.WeeklyRoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weeklyroutine.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DailyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.DailyRoutinesTable,
			Columns: program.DailyRoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutine.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedDailyRoutinesIDs(); len(nodes) > 0 && !puo.mutation.DailyRoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.DailyRoutinesTable,
			Columns: program.DailyRoutinesPrimaryKey,
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
	if nodes := puo.mutation.DailyRoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   program.DailyRoutinesTable,
			Columns: program.DailyRoutinesPrimaryKey,
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
