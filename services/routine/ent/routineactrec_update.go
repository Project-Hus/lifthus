// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"routine/ent/act"
	"routine/ent/dailyroutinerec"
	"routine/ent/predicate"
	"routine/ent/routineact"
	"routine/ent/routineactrec"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoutineActRecUpdate is the builder for updating RoutineActRec entities.
type RoutineActRecUpdate struct {
	config
	hooks    []Hook
	mutation *RoutineActRecMutation
}

// Where appends a list predicates to the RoutineActRecUpdate builder.
func (raru *RoutineActRecUpdate) Where(ps ...predicate.RoutineActRec) *RoutineActRecUpdate {
	raru.mutation.Where(ps...)
	return raru
}

// SetDailyRoutineRecID sets the "daily_routine_rec_id" field.
func (raru *RoutineActRecUpdate) SetDailyRoutineRecID(u uint64) *RoutineActRecUpdate {
	raru.mutation.SetDailyRoutineRecID(u)
	return raru
}

// SetRoutineActID sets the "routine_act_id" field.
func (raru *RoutineActRecUpdate) SetRoutineActID(u uint64) *RoutineActRecUpdate {
	raru.mutation.SetRoutineActID(u)
	return raru
}

// SetNillableRoutineActID sets the "routine_act_id" field if the given value is not nil.
func (raru *RoutineActRecUpdate) SetNillableRoutineActID(u *uint64) *RoutineActRecUpdate {
	if u != nil {
		raru.SetRoutineActID(*u)
	}
	return raru
}

// ClearRoutineActID clears the value of the "routine_act_id" field.
func (raru *RoutineActRecUpdate) ClearRoutineActID() *RoutineActRecUpdate {
	raru.mutation.ClearRoutineActID()
	return raru
}

// SetActID sets the "act_id" field.
func (raru *RoutineActRecUpdate) SetActID(u uint64) *RoutineActRecUpdate {
	raru.mutation.SetActID(u)
	return raru
}

// SetOrder sets the "order" field.
func (raru *RoutineActRecUpdate) SetOrder(i int) *RoutineActRecUpdate {
	raru.mutation.ResetOrder()
	raru.mutation.SetOrder(i)
	return raru
}

// AddOrder adds i to the "order" field.
func (raru *RoutineActRecUpdate) AddOrder(i int) *RoutineActRecUpdate {
	raru.mutation.AddOrder(i)
	return raru
}

// SetReps sets the "reps" field.
func (raru *RoutineActRecUpdate) SetReps(i int) *RoutineActRecUpdate {
	raru.mutation.ResetReps()
	raru.mutation.SetReps(i)
	return raru
}

// SetNillableReps sets the "reps" field if the given value is not nil.
func (raru *RoutineActRecUpdate) SetNillableReps(i *int) *RoutineActRecUpdate {
	if i != nil {
		raru.SetReps(*i)
	}
	return raru
}

// AddReps adds i to the "reps" field.
func (raru *RoutineActRecUpdate) AddReps(i int) *RoutineActRecUpdate {
	raru.mutation.AddReps(i)
	return raru
}

// ClearReps clears the value of the "reps" field.
func (raru *RoutineActRecUpdate) ClearReps() *RoutineActRecUpdate {
	raru.mutation.ClearReps()
	return raru
}

// SetLap sets the "lap" field.
func (raru *RoutineActRecUpdate) SetLap(i int) *RoutineActRecUpdate {
	raru.mutation.ResetLap()
	raru.mutation.SetLap(i)
	return raru
}

// SetNillableLap sets the "lap" field if the given value is not nil.
func (raru *RoutineActRecUpdate) SetNillableLap(i *int) *RoutineActRecUpdate {
	if i != nil {
		raru.SetLap(*i)
	}
	return raru
}

// AddLap adds i to the "lap" field.
func (raru *RoutineActRecUpdate) AddLap(i int) *RoutineActRecUpdate {
	raru.mutation.AddLap(i)
	return raru
}

// ClearLap clears the value of the "lap" field.
func (raru *RoutineActRecUpdate) ClearLap() *RoutineActRecUpdate {
	raru.mutation.ClearLap()
	return raru
}

// SetCurrentReps sets the "current_reps" field.
func (raru *RoutineActRecUpdate) SetCurrentReps(i int) *RoutineActRecUpdate {
	raru.mutation.ResetCurrentReps()
	raru.mutation.SetCurrentReps(i)
	return raru
}

// SetNillableCurrentReps sets the "current_reps" field if the given value is not nil.
func (raru *RoutineActRecUpdate) SetNillableCurrentReps(i *int) *RoutineActRecUpdate {
	if i != nil {
		raru.SetCurrentReps(*i)
	}
	return raru
}

// AddCurrentReps adds i to the "current_reps" field.
func (raru *RoutineActRecUpdate) AddCurrentReps(i int) *RoutineActRecUpdate {
	raru.mutation.AddCurrentReps(i)
	return raru
}

// SetCurrentLap sets the "current_lap" field.
func (raru *RoutineActRecUpdate) SetCurrentLap(i int) *RoutineActRecUpdate {
	raru.mutation.ResetCurrentLap()
	raru.mutation.SetCurrentLap(i)
	return raru
}

// SetNillableCurrentLap sets the "current_lap" field if the given value is not nil.
func (raru *RoutineActRecUpdate) SetNillableCurrentLap(i *int) *RoutineActRecUpdate {
	if i != nil {
		raru.SetCurrentLap(*i)
	}
	return raru
}

// AddCurrentLap adds i to the "current_lap" field.
func (raru *RoutineActRecUpdate) AddCurrentLap(i int) *RoutineActRecUpdate {
	raru.mutation.AddCurrentLap(i)
	return raru
}

// SetStartedAt sets the "started_at" field.
func (raru *RoutineActRecUpdate) SetStartedAt(t time.Time) *RoutineActRecUpdate {
	raru.mutation.SetStartedAt(t)
	return raru
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (raru *RoutineActRecUpdate) SetNillableStartedAt(t *time.Time) *RoutineActRecUpdate {
	if t != nil {
		raru.SetStartedAt(*t)
	}
	return raru
}

// ClearStartedAt clears the value of the "started_at" field.
func (raru *RoutineActRecUpdate) ClearStartedAt() *RoutineActRecUpdate {
	raru.mutation.ClearStartedAt()
	return raru
}

// SetImage sets the "image" field.
func (raru *RoutineActRecUpdate) SetImage(s string) *RoutineActRecUpdate {
	raru.mutation.SetImage(s)
	return raru
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (raru *RoutineActRecUpdate) SetNillableImage(s *string) *RoutineActRecUpdate {
	if s != nil {
		raru.SetImage(*s)
	}
	return raru
}

// ClearImage clears the value of the "image" field.
func (raru *RoutineActRecUpdate) ClearImage() *RoutineActRecUpdate {
	raru.mutation.ClearImage()
	return raru
}

// SetComment sets the "comment" field.
func (raru *RoutineActRecUpdate) SetComment(s string) *RoutineActRecUpdate {
	raru.mutation.SetComment(s)
	return raru
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (raru *RoutineActRecUpdate) SetNillableComment(s *string) *RoutineActRecUpdate {
	if s != nil {
		raru.SetComment(*s)
	}
	return raru
}

// ClearComment clears the value of the "comment" field.
func (raru *RoutineActRecUpdate) ClearComment() *RoutineActRecUpdate {
	raru.mutation.ClearComment()
	return raru
}

// SetStatus sets the "status" field.
func (raru *RoutineActRecUpdate) SetStatus(r routineactrec.Status) *RoutineActRecUpdate {
	raru.mutation.SetStatus(r)
	return raru
}

// SetUpdatedAt sets the "updated_at" field.
func (raru *RoutineActRecUpdate) SetUpdatedAt(t time.Time) *RoutineActRecUpdate {
	raru.mutation.SetUpdatedAt(t)
	return raru
}

// SetDailyRoutineRec sets the "daily_routine_rec" edge to the DailyRoutineRec entity.
func (raru *RoutineActRecUpdate) SetDailyRoutineRec(d *DailyRoutineRec) *RoutineActRecUpdate {
	return raru.SetDailyRoutineRecID(d.ID)
}

// SetAct sets the "act" edge to the Act entity.
func (raru *RoutineActRecUpdate) SetAct(a *Act) *RoutineActRecUpdate {
	return raru.SetActID(a.ID)
}

// SetRoutineAct sets the "routine_act" edge to the RoutineAct entity.
func (raru *RoutineActRecUpdate) SetRoutineAct(r *RoutineAct) *RoutineActRecUpdate {
	return raru.SetRoutineActID(r.ID)
}

// Mutation returns the RoutineActRecMutation object of the builder.
func (raru *RoutineActRecUpdate) Mutation() *RoutineActRecMutation {
	return raru.mutation
}

// ClearDailyRoutineRec clears the "daily_routine_rec" edge to the DailyRoutineRec entity.
func (raru *RoutineActRecUpdate) ClearDailyRoutineRec() *RoutineActRecUpdate {
	raru.mutation.ClearDailyRoutineRec()
	return raru
}

// ClearAct clears the "act" edge to the Act entity.
func (raru *RoutineActRecUpdate) ClearAct() *RoutineActRecUpdate {
	raru.mutation.ClearAct()
	return raru
}

// ClearRoutineAct clears the "routine_act" edge to the RoutineAct entity.
func (raru *RoutineActRecUpdate) ClearRoutineAct() *RoutineActRecUpdate {
	raru.mutation.ClearRoutineAct()
	return raru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (raru *RoutineActRecUpdate) Save(ctx context.Context) (int, error) {
	raru.defaults()
	return withHooks(ctx, raru.sqlSave, raru.mutation, raru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (raru *RoutineActRecUpdate) SaveX(ctx context.Context) int {
	affected, err := raru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (raru *RoutineActRecUpdate) Exec(ctx context.Context) error {
	_, err := raru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (raru *RoutineActRecUpdate) ExecX(ctx context.Context) {
	if err := raru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (raru *RoutineActRecUpdate) defaults() {
	if _, ok := raru.mutation.UpdatedAt(); !ok {
		v := routineactrec.UpdateDefaultUpdatedAt()
		raru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (raru *RoutineActRecUpdate) check() error {
	if v, ok := raru.mutation.Order(); ok {
		if err := routineactrec.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.order": %w`, err)}
		}
	}
	if v, ok := raru.mutation.Reps(); ok {
		if err := routineactrec.RepsValidator(v); err != nil {
			return &ValidationError{Name: "reps", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.reps": %w`, err)}
		}
	}
	if v, ok := raru.mutation.Lap(); ok {
		if err := routineactrec.LapValidator(v); err != nil {
			return &ValidationError{Name: "lap", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.lap": %w`, err)}
		}
	}
	if v, ok := raru.mutation.CurrentReps(); ok {
		if err := routineactrec.CurrentRepsValidator(v); err != nil {
			return &ValidationError{Name: "current_reps", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.current_reps": %w`, err)}
		}
	}
	if v, ok := raru.mutation.CurrentLap(); ok {
		if err := routineactrec.CurrentLapValidator(v); err != nil {
			return &ValidationError{Name: "current_lap", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.current_lap": %w`, err)}
		}
	}
	if v, ok := raru.mutation.Status(); ok {
		if err := routineactrec.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.status": %w`, err)}
		}
	}
	if _, ok := raru.mutation.DailyRoutineRecID(); raru.mutation.DailyRoutineRecCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineActRec.daily_routine_rec"`)
	}
	if _, ok := raru.mutation.ActID(); raru.mutation.ActCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineActRec.act"`)
	}
	return nil
}

func (raru *RoutineActRecUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := raru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(routineactrec.Table, routineactrec.Columns, sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64))
	if ps := raru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := raru.mutation.Order(); ok {
		_spec.SetField(routineactrec.FieldOrder, field.TypeInt, value)
	}
	if value, ok := raru.mutation.AddedOrder(); ok {
		_spec.AddField(routineactrec.FieldOrder, field.TypeInt, value)
	}
	if value, ok := raru.mutation.Reps(); ok {
		_spec.SetField(routineactrec.FieldReps, field.TypeInt, value)
	}
	if value, ok := raru.mutation.AddedReps(); ok {
		_spec.AddField(routineactrec.FieldReps, field.TypeInt, value)
	}
	if raru.mutation.RepsCleared() {
		_spec.ClearField(routineactrec.FieldReps, field.TypeInt)
	}
	if value, ok := raru.mutation.Lap(); ok {
		_spec.SetField(routineactrec.FieldLap, field.TypeInt, value)
	}
	if value, ok := raru.mutation.AddedLap(); ok {
		_spec.AddField(routineactrec.FieldLap, field.TypeInt, value)
	}
	if raru.mutation.LapCleared() {
		_spec.ClearField(routineactrec.FieldLap, field.TypeInt)
	}
	if value, ok := raru.mutation.CurrentReps(); ok {
		_spec.SetField(routineactrec.FieldCurrentReps, field.TypeInt, value)
	}
	if value, ok := raru.mutation.AddedCurrentReps(); ok {
		_spec.AddField(routineactrec.FieldCurrentReps, field.TypeInt, value)
	}
	if value, ok := raru.mutation.CurrentLap(); ok {
		_spec.SetField(routineactrec.FieldCurrentLap, field.TypeInt, value)
	}
	if value, ok := raru.mutation.AddedCurrentLap(); ok {
		_spec.AddField(routineactrec.FieldCurrentLap, field.TypeInt, value)
	}
	if value, ok := raru.mutation.StartedAt(); ok {
		_spec.SetField(routineactrec.FieldStartedAt, field.TypeTime, value)
	}
	if raru.mutation.StartedAtCleared() {
		_spec.ClearField(routineactrec.FieldStartedAt, field.TypeTime)
	}
	if value, ok := raru.mutation.Image(); ok {
		_spec.SetField(routineactrec.FieldImage, field.TypeString, value)
	}
	if raru.mutation.ImageCleared() {
		_spec.ClearField(routineactrec.FieldImage, field.TypeString)
	}
	if value, ok := raru.mutation.Comment(); ok {
		_spec.SetField(routineactrec.FieldComment, field.TypeString, value)
	}
	if raru.mutation.CommentCleared() {
		_spec.ClearField(routineactrec.FieldComment, field.TypeString)
	}
	if value, ok := raru.mutation.Status(); ok {
		_spec.SetField(routineactrec.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := raru.mutation.UpdatedAt(); ok {
		_spec.SetField(routineactrec.FieldUpdatedAt, field.TypeTime, value)
	}
	if raru.mutation.DailyRoutineRecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.DailyRoutineRecTable,
			Columns: []string{routineactrec.DailyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutinerec.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := raru.mutation.DailyRoutineRecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.DailyRoutineRecTable,
			Columns: []string{routineactrec.DailyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutinerec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if raru.mutation.ActCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.ActTable,
			Columns: []string{routineactrec.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := raru.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.ActTable,
			Columns: []string{routineactrec.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if raru.mutation.RoutineActCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.RoutineActTable,
			Columns: []string{routineactrec.RoutineActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := raru.mutation.RoutineActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.RoutineActTable,
			Columns: []string{routineactrec.RoutineActColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, raru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routineactrec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	raru.mutation.done = true
	return n, nil
}

// RoutineActRecUpdateOne is the builder for updating a single RoutineActRec entity.
type RoutineActRecUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoutineActRecMutation
}

// SetDailyRoutineRecID sets the "daily_routine_rec_id" field.
func (raruo *RoutineActRecUpdateOne) SetDailyRoutineRecID(u uint64) *RoutineActRecUpdateOne {
	raruo.mutation.SetDailyRoutineRecID(u)
	return raruo
}

// SetRoutineActID sets the "routine_act_id" field.
func (raruo *RoutineActRecUpdateOne) SetRoutineActID(u uint64) *RoutineActRecUpdateOne {
	raruo.mutation.SetRoutineActID(u)
	return raruo
}

// SetNillableRoutineActID sets the "routine_act_id" field if the given value is not nil.
func (raruo *RoutineActRecUpdateOne) SetNillableRoutineActID(u *uint64) *RoutineActRecUpdateOne {
	if u != nil {
		raruo.SetRoutineActID(*u)
	}
	return raruo
}

// ClearRoutineActID clears the value of the "routine_act_id" field.
func (raruo *RoutineActRecUpdateOne) ClearRoutineActID() *RoutineActRecUpdateOne {
	raruo.mutation.ClearRoutineActID()
	return raruo
}

// SetActID sets the "act_id" field.
func (raruo *RoutineActRecUpdateOne) SetActID(u uint64) *RoutineActRecUpdateOne {
	raruo.mutation.SetActID(u)
	return raruo
}

// SetOrder sets the "order" field.
func (raruo *RoutineActRecUpdateOne) SetOrder(i int) *RoutineActRecUpdateOne {
	raruo.mutation.ResetOrder()
	raruo.mutation.SetOrder(i)
	return raruo
}

// AddOrder adds i to the "order" field.
func (raruo *RoutineActRecUpdateOne) AddOrder(i int) *RoutineActRecUpdateOne {
	raruo.mutation.AddOrder(i)
	return raruo
}

// SetReps sets the "reps" field.
func (raruo *RoutineActRecUpdateOne) SetReps(i int) *RoutineActRecUpdateOne {
	raruo.mutation.ResetReps()
	raruo.mutation.SetReps(i)
	return raruo
}

// SetNillableReps sets the "reps" field if the given value is not nil.
func (raruo *RoutineActRecUpdateOne) SetNillableReps(i *int) *RoutineActRecUpdateOne {
	if i != nil {
		raruo.SetReps(*i)
	}
	return raruo
}

// AddReps adds i to the "reps" field.
func (raruo *RoutineActRecUpdateOne) AddReps(i int) *RoutineActRecUpdateOne {
	raruo.mutation.AddReps(i)
	return raruo
}

// ClearReps clears the value of the "reps" field.
func (raruo *RoutineActRecUpdateOne) ClearReps() *RoutineActRecUpdateOne {
	raruo.mutation.ClearReps()
	return raruo
}

// SetLap sets the "lap" field.
func (raruo *RoutineActRecUpdateOne) SetLap(i int) *RoutineActRecUpdateOne {
	raruo.mutation.ResetLap()
	raruo.mutation.SetLap(i)
	return raruo
}

// SetNillableLap sets the "lap" field if the given value is not nil.
func (raruo *RoutineActRecUpdateOne) SetNillableLap(i *int) *RoutineActRecUpdateOne {
	if i != nil {
		raruo.SetLap(*i)
	}
	return raruo
}

// AddLap adds i to the "lap" field.
func (raruo *RoutineActRecUpdateOne) AddLap(i int) *RoutineActRecUpdateOne {
	raruo.mutation.AddLap(i)
	return raruo
}

// ClearLap clears the value of the "lap" field.
func (raruo *RoutineActRecUpdateOne) ClearLap() *RoutineActRecUpdateOne {
	raruo.mutation.ClearLap()
	return raruo
}

// SetCurrentReps sets the "current_reps" field.
func (raruo *RoutineActRecUpdateOne) SetCurrentReps(i int) *RoutineActRecUpdateOne {
	raruo.mutation.ResetCurrentReps()
	raruo.mutation.SetCurrentReps(i)
	return raruo
}

// SetNillableCurrentReps sets the "current_reps" field if the given value is not nil.
func (raruo *RoutineActRecUpdateOne) SetNillableCurrentReps(i *int) *RoutineActRecUpdateOne {
	if i != nil {
		raruo.SetCurrentReps(*i)
	}
	return raruo
}

// AddCurrentReps adds i to the "current_reps" field.
func (raruo *RoutineActRecUpdateOne) AddCurrentReps(i int) *RoutineActRecUpdateOne {
	raruo.mutation.AddCurrentReps(i)
	return raruo
}

// SetCurrentLap sets the "current_lap" field.
func (raruo *RoutineActRecUpdateOne) SetCurrentLap(i int) *RoutineActRecUpdateOne {
	raruo.mutation.ResetCurrentLap()
	raruo.mutation.SetCurrentLap(i)
	return raruo
}

// SetNillableCurrentLap sets the "current_lap" field if the given value is not nil.
func (raruo *RoutineActRecUpdateOne) SetNillableCurrentLap(i *int) *RoutineActRecUpdateOne {
	if i != nil {
		raruo.SetCurrentLap(*i)
	}
	return raruo
}

// AddCurrentLap adds i to the "current_lap" field.
func (raruo *RoutineActRecUpdateOne) AddCurrentLap(i int) *RoutineActRecUpdateOne {
	raruo.mutation.AddCurrentLap(i)
	return raruo
}

// SetStartedAt sets the "started_at" field.
func (raruo *RoutineActRecUpdateOne) SetStartedAt(t time.Time) *RoutineActRecUpdateOne {
	raruo.mutation.SetStartedAt(t)
	return raruo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (raruo *RoutineActRecUpdateOne) SetNillableStartedAt(t *time.Time) *RoutineActRecUpdateOne {
	if t != nil {
		raruo.SetStartedAt(*t)
	}
	return raruo
}

// ClearStartedAt clears the value of the "started_at" field.
func (raruo *RoutineActRecUpdateOne) ClearStartedAt() *RoutineActRecUpdateOne {
	raruo.mutation.ClearStartedAt()
	return raruo
}

// SetImage sets the "image" field.
func (raruo *RoutineActRecUpdateOne) SetImage(s string) *RoutineActRecUpdateOne {
	raruo.mutation.SetImage(s)
	return raruo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (raruo *RoutineActRecUpdateOne) SetNillableImage(s *string) *RoutineActRecUpdateOne {
	if s != nil {
		raruo.SetImage(*s)
	}
	return raruo
}

// ClearImage clears the value of the "image" field.
func (raruo *RoutineActRecUpdateOne) ClearImage() *RoutineActRecUpdateOne {
	raruo.mutation.ClearImage()
	return raruo
}

// SetComment sets the "comment" field.
func (raruo *RoutineActRecUpdateOne) SetComment(s string) *RoutineActRecUpdateOne {
	raruo.mutation.SetComment(s)
	return raruo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (raruo *RoutineActRecUpdateOne) SetNillableComment(s *string) *RoutineActRecUpdateOne {
	if s != nil {
		raruo.SetComment(*s)
	}
	return raruo
}

// ClearComment clears the value of the "comment" field.
func (raruo *RoutineActRecUpdateOne) ClearComment() *RoutineActRecUpdateOne {
	raruo.mutation.ClearComment()
	return raruo
}

// SetStatus sets the "status" field.
func (raruo *RoutineActRecUpdateOne) SetStatus(r routineactrec.Status) *RoutineActRecUpdateOne {
	raruo.mutation.SetStatus(r)
	return raruo
}

// SetUpdatedAt sets the "updated_at" field.
func (raruo *RoutineActRecUpdateOne) SetUpdatedAt(t time.Time) *RoutineActRecUpdateOne {
	raruo.mutation.SetUpdatedAt(t)
	return raruo
}

// SetDailyRoutineRec sets the "daily_routine_rec" edge to the DailyRoutineRec entity.
func (raruo *RoutineActRecUpdateOne) SetDailyRoutineRec(d *DailyRoutineRec) *RoutineActRecUpdateOne {
	return raruo.SetDailyRoutineRecID(d.ID)
}

// SetAct sets the "act" edge to the Act entity.
func (raruo *RoutineActRecUpdateOne) SetAct(a *Act) *RoutineActRecUpdateOne {
	return raruo.SetActID(a.ID)
}

// SetRoutineAct sets the "routine_act" edge to the RoutineAct entity.
func (raruo *RoutineActRecUpdateOne) SetRoutineAct(r *RoutineAct) *RoutineActRecUpdateOne {
	return raruo.SetRoutineActID(r.ID)
}

// Mutation returns the RoutineActRecMutation object of the builder.
func (raruo *RoutineActRecUpdateOne) Mutation() *RoutineActRecMutation {
	return raruo.mutation
}

// ClearDailyRoutineRec clears the "daily_routine_rec" edge to the DailyRoutineRec entity.
func (raruo *RoutineActRecUpdateOne) ClearDailyRoutineRec() *RoutineActRecUpdateOne {
	raruo.mutation.ClearDailyRoutineRec()
	return raruo
}

// ClearAct clears the "act" edge to the Act entity.
func (raruo *RoutineActRecUpdateOne) ClearAct() *RoutineActRecUpdateOne {
	raruo.mutation.ClearAct()
	return raruo
}

// ClearRoutineAct clears the "routine_act" edge to the RoutineAct entity.
func (raruo *RoutineActRecUpdateOne) ClearRoutineAct() *RoutineActRecUpdateOne {
	raruo.mutation.ClearRoutineAct()
	return raruo
}

// Where appends a list predicates to the RoutineActRecUpdate builder.
func (raruo *RoutineActRecUpdateOne) Where(ps ...predicate.RoutineActRec) *RoutineActRecUpdateOne {
	raruo.mutation.Where(ps...)
	return raruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (raruo *RoutineActRecUpdateOne) Select(field string, fields ...string) *RoutineActRecUpdateOne {
	raruo.fields = append([]string{field}, fields...)
	return raruo
}

// Save executes the query and returns the updated RoutineActRec entity.
func (raruo *RoutineActRecUpdateOne) Save(ctx context.Context) (*RoutineActRec, error) {
	raruo.defaults()
	return withHooks(ctx, raruo.sqlSave, raruo.mutation, raruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (raruo *RoutineActRecUpdateOne) SaveX(ctx context.Context) *RoutineActRec {
	node, err := raruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (raruo *RoutineActRecUpdateOne) Exec(ctx context.Context) error {
	_, err := raruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (raruo *RoutineActRecUpdateOne) ExecX(ctx context.Context) {
	if err := raruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (raruo *RoutineActRecUpdateOne) defaults() {
	if _, ok := raruo.mutation.UpdatedAt(); !ok {
		v := routineactrec.UpdateDefaultUpdatedAt()
		raruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (raruo *RoutineActRecUpdateOne) check() error {
	if v, ok := raruo.mutation.Order(); ok {
		if err := routineactrec.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.order": %w`, err)}
		}
	}
	if v, ok := raruo.mutation.Reps(); ok {
		if err := routineactrec.RepsValidator(v); err != nil {
			return &ValidationError{Name: "reps", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.reps": %w`, err)}
		}
	}
	if v, ok := raruo.mutation.Lap(); ok {
		if err := routineactrec.LapValidator(v); err != nil {
			return &ValidationError{Name: "lap", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.lap": %w`, err)}
		}
	}
	if v, ok := raruo.mutation.CurrentReps(); ok {
		if err := routineactrec.CurrentRepsValidator(v); err != nil {
			return &ValidationError{Name: "current_reps", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.current_reps": %w`, err)}
		}
	}
	if v, ok := raruo.mutation.CurrentLap(); ok {
		if err := routineactrec.CurrentLapValidator(v); err != nil {
			return &ValidationError{Name: "current_lap", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.current_lap": %w`, err)}
		}
	}
	if v, ok := raruo.mutation.Status(); ok {
		if err := routineactrec.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "RoutineActRec.status": %w`, err)}
		}
	}
	if _, ok := raruo.mutation.DailyRoutineRecID(); raruo.mutation.DailyRoutineRecCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineActRec.daily_routine_rec"`)
	}
	if _, ok := raruo.mutation.ActID(); raruo.mutation.ActCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineActRec.act"`)
	}
	return nil
}

func (raruo *RoutineActRecUpdateOne) sqlSave(ctx context.Context) (_node *RoutineActRec, err error) {
	if err := raruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(routineactrec.Table, routineactrec.Columns, sqlgraph.NewFieldSpec(routineactrec.FieldID, field.TypeUint64))
	id, ok := raruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoutineActRec.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := raruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routineactrec.FieldID)
		for _, f := range fields {
			if !routineactrec.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != routineactrec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := raruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := raruo.mutation.Order(); ok {
		_spec.SetField(routineactrec.FieldOrder, field.TypeInt, value)
	}
	if value, ok := raruo.mutation.AddedOrder(); ok {
		_spec.AddField(routineactrec.FieldOrder, field.TypeInt, value)
	}
	if value, ok := raruo.mutation.Reps(); ok {
		_spec.SetField(routineactrec.FieldReps, field.TypeInt, value)
	}
	if value, ok := raruo.mutation.AddedReps(); ok {
		_spec.AddField(routineactrec.FieldReps, field.TypeInt, value)
	}
	if raruo.mutation.RepsCleared() {
		_spec.ClearField(routineactrec.FieldReps, field.TypeInt)
	}
	if value, ok := raruo.mutation.Lap(); ok {
		_spec.SetField(routineactrec.FieldLap, field.TypeInt, value)
	}
	if value, ok := raruo.mutation.AddedLap(); ok {
		_spec.AddField(routineactrec.FieldLap, field.TypeInt, value)
	}
	if raruo.mutation.LapCleared() {
		_spec.ClearField(routineactrec.FieldLap, field.TypeInt)
	}
	if value, ok := raruo.mutation.CurrentReps(); ok {
		_spec.SetField(routineactrec.FieldCurrentReps, field.TypeInt, value)
	}
	if value, ok := raruo.mutation.AddedCurrentReps(); ok {
		_spec.AddField(routineactrec.FieldCurrentReps, field.TypeInt, value)
	}
	if value, ok := raruo.mutation.CurrentLap(); ok {
		_spec.SetField(routineactrec.FieldCurrentLap, field.TypeInt, value)
	}
	if value, ok := raruo.mutation.AddedCurrentLap(); ok {
		_spec.AddField(routineactrec.FieldCurrentLap, field.TypeInt, value)
	}
	if value, ok := raruo.mutation.StartedAt(); ok {
		_spec.SetField(routineactrec.FieldStartedAt, field.TypeTime, value)
	}
	if raruo.mutation.StartedAtCleared() {
		_spec.ClearField(routineactrec.FieldStartedAt, field.TypeTime)
	}
	if value, ok := raruo.mutation.Image(); ok {
		_spec.SetField(routineactrec.FieldImage, field.TypeString, value)
	}
	if raruo.mutation.ImageCleared() {
		_spec.ClearField(routineactrec.FieldImage, field.TypeString)
	}
	if value, ok := raruo.mutation.Comment(); ok {
		_spec.SetField(routineactrec.FieldComment, field.TypeString, value)
	}
	if raruo.mutation.CommentCleared() {
		_spec.ClearField(routineactrec.FieldComment, field.TypeString)
	}
	if value, ok := raruo.mutation.Status(); ok {
		_spec.SetField(routineactrec.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := raruo.mutation.UpdatedAt(); ok {
		_spec.SetField(routineactrec.FieldUpdatedAt, field.TypeTime, value)
	}
	if raruo.mutation.DailyRoutineRecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.DailyRoutineRecTable,
			Columns: []string{routineactrec.DailyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutinerec.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := raruo.mutation.DailyRoutineRecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.DailyRoutineRecTable,
			Columns: []string{routineactrec.DailyRoutineRecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dailyroutinerec.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if raruo.mutation.ActCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.ActTable,
			Columns: []string{routineactrec.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := raruo.mutation.ActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.ActTable,
			Columns: []string{routineactrec.ActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(act.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if raruo.mutation.RoutineActCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.RoutineActTable,
			Columns: []string{routineactrec.RoutineActColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineact.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := raruo.mutation.RoutineActIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineactrec.RoutineActTable,
			Columns: []string{routineactrec.RoutineActColumn},
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
	_node = &RoutineActRec{config: raruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, raruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routineactrec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	raruo.mutation.done = true
	return _node, nil
}
