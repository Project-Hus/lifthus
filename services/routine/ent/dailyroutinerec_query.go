// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"routine/ent/dailyroutine"
	"routine/ent/dailyroutinerec"
	"routine/ent/predicate"
	"routine/ent/programrec"
	"routine/ent/routineactrec"
	"routine/ent/weeklyroutinerec"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DailyRoutineRecQuery is the builder for querying DailyRoutineRec entities.
type DailyRoutineRecQuery struct {
	config
	ctx                  *QueryContext
	order                []dailyroutinerec.OrderOption
	inters               []Interceptor
	predicates           []predicate.DailyRoutineRec
	withDailyRoutine     *DailyRoutineQuery
	withProgramRec       *ProgramRecQuery
	withWeeklyRoutineRec *WeeklyRoutineRecQuery
	withRoutineActRecs   *RoutineActRecQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DailyRoutineRecQuery builder.
func (drrq *DailyRoutineRecQuery) Where(ps ...predicate.DailyRoutineRec) *DailyRoutineRecQuery {
	drrq.predicates = append(drrq.predicates, ps...)
	return drrq
}

// Limit the number of records to be returned by this query.
func (drrq *DailyRoutineRecQuery) Limit(limit int) *DailyRoutineRecQuery {
	drrq.ctx.Limit = &limit
	return drrq
}

// Offset to start from.
func (drrq *DailyRoutineRecQuery) Offset(offset int) *DailyRoutineRecQuery {
	drrq.ctx.Offset = &offset
	return drrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (drrq *DailyRoutineRecQuery) Unique(unique bool) *DailyRoutineRecQuery {
	drrq.ctx.Unique = &unique
	return drrq
}

// Order specifies how the records should be ordered.
func (drrq *DailyRoutineRecQuery) Order(o ...dailyroutinerec.OrderOption) *DailyRoutineRecQuery {
	drrq.order = append(drrq.order, o...)
	return drrq
}

// QueryDailyRoutine chains the current query on the "daily_routine" edge.
func (drrq *DailyRoutineRecQuery) QueryDailyRoutine() *DailyRoutineQuery {
	query := (&DailyRoutineClient{config: drrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailyroutinerec.Table, dailyroutinerec.FieldID, selector),
			sqlgraph.To(dailyroutine.Table, dailyroutine.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailyroutinerec.DailyRoutineTable, dailyroutinerec.DailyRoutineColumn),
		)
		fromU = sqlgraph.SetNeighbors(drrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProgramRec chains the current query on the "program_rec" edge.
func (drrq *DailyRoutineRecQuery) QueryProgramRec() *ProgramRecQuery {
	query := (&ProgramRecClient{config: drrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailyroutinerec.Table, dailyroutinerec.FieldID, selector),
			sqlgraph.To(programrec.Table, programrec.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailyroutinerec.ProgramRecTable, dailyroutinerec.ProgramRecColumn),
		)
		fromU = sqlgraph.SetNeighbors(drrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWeeklyRoutineRec chains the current query on the "weekly_routine_rec" edge.
func (drrq *DailyRoutineRecQuery) QueryWeeklyRoutineRec() *WeeklyRoutineRecQuery {
	query := (&WeeklyRoutineRecClient{config: drrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailyroutinerec.Table, dailyroutinerec.FieldID, selector),
			sqlgraph.To(weeklyroutinerec.Table, weeklyroutinerec.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailyroutinerec.WeeklyRoutineRecTable, dailyroutinerec.WeeklyRoutineRecColumn),
		)
		fromU = sqlgraph.SetNeighbors(drrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoutineActRecs chains the current query on the "routine_act_recs" edge.
func (drrq *DailyRoutineRecQuery) QueryRoutineActRecs() *RoutineActRecQuery {
	query := (&RoutineActRecClient{config: drrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dailyroutinerec.Table, dailyroutinerec.FieldID, selector),
			sqlgraph.To(routineactrec.Table, routineactrec.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dailyroutinerec.RoutineActRecsTable, dailyroutinerec.RoutineActRecsColumn),
		)
		fromU = sqlgraph.SetNeighbors(drrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DailyRoutineRec entity from the query.
// Returns a *NotFoundError when no DailyRoutineRec was found.
func (drrq *DailyRoutineRecQuery) First(ctx context.Context) (*DailyRoutineRec, error) {
	nodes, err := drrq.Limit(1).All(setContextOp(ctx, drrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dailyroutinerec.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (drrq *DailyRoutineRecQuery) FirstX(ctx context.Context) *DailyRoutineRec {
	node, err := drrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DailyRoutineRec ID from the query.
// Returns a *NotFoundError when no DailyRoutineRec ID was found.
func (drrq *DailyRoutineRecQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = drrq.Limit(1).IDs(setContextOp(ctx, drrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dailyroutinerec.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (drrq *DailyRoutineRecQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := drrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DailyRoutineRec entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DailyRoutineRec entity is found.
// Returns a *NotFoundError when no DailyRoutineRec entities are found.
func (drrq *DailyRoutineRecQuery) Only(ctx context.Context) (*DailyRoutineRec, error) {
	nodes, err := drrq.Limit(2).All(setContextOp(ctx, drrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dailyroutinerec.Label}
	default:
		return nil, &NotSingularError{dailyroutinerec.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (drrq *DailyRoutineRecQuery) OnlyX(ctx context.Context) *DailyRoutineRec {
	node, err := drrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DailyRoutineRec ID in the query.
// Returns a *NotSingularError when more than one DailyRoutineRec ID is found.
// Returns a *NotFoundError when no entities are found.
func (drrq *DailyRoutineRecQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = drrq.Limit(2).IDs(setContextOp(ctx, drrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dailyroutinerec.Label}
	default:
		err = &NotSingularError{dailyroutinerec.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (drrq *DailyRoutineRecQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := drrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DailyRoutineRecs.
func (drrq *DailyRoutineRecQuery) All(ctx context.Context) ([]*DailyRoutineRec, error) {
	ctx = setContextOp(ctx, drrq.ctx, "All")
	if err := drrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DailyRoutineRec, *DailyRoutineRecQuery]()
	return withInterceptors[[]*DailyRoutineRec](ctx, drrq, qr, drrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (drrq *DailyRoutineRecQuery) AllX(ctx context.Context) []*DailyRoutineRec {
	nodes, err := drrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DailyRoutineRec IDs.
func (drrq *DailyRoutineRecQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if drrq.ctx.Unique == nil && drrq.path != nil {
		drrq.Unique(true)
	}
	ctx = setContextOp(ctx, drrq.ctx, "IDs")
	if err = drrq.Select(dailyroutinerec.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (drrq *DailyRoutineRecQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := drrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (drrq *DailyRoutineRecQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, drrq.ctx, "Count")
	if err := drrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, drrq, querierCount[*DailyRoutineRecQuery](), drrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (drrq *DailyRoutineRecQuery) CountX(ctx context.Context) int {
	count, err := drrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (drrq *DailyRoutineRecQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, drrq.ctx, "Exist")
	switch _, err := drrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (drrq *DailyRoutineRecQuery) ExistX(ctx context.Context) bool {
	exist, err := drrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DailyRoutineRecQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (drrq *DailyRoutineRecQuery) Clone() *DailyRoutineRecQuery {
	if drrq == nil {
		return nil
	}
	return &DailyRoutineRecQuery{
		config:               drrq.config,
		ctx:                  drrq.ctx.Clone(),
		order:                append([]dailyroutinerec.OrderOption{}, drrq.order...),
		inters:               append([]Interceptor{}, drrq.inters...),
		predicates:           append([]predicate.DailyRoutineRec{}, drrq.predicates...),
		withDailyRoutine:     drrq.withDailyRoutine.Clone(),
		withProgramRec:       drrq.withProgramRec.Clone(),
		withWeeklyRoutineRec: drrq.withWeeklyRoutineRec.Clone(),
		withRoutineActRecs:   drrq.withRoutineActRecs.Clone(),
		// clone intermediate query.
		sql:  drrq.sql.Clone(),
		path: drrq.path,
	}
}

// WithDailyRoutine tells the query-builder to eager-load the nodes that are connected to
// the "daily_routine" edge. The optional arguments are used to configure the query builder of the edge.
func (drrq *DailyRoutineRecQuery) WithDailyRoutine(opts ...func(*DailyRoutineQuery)) *DailyRoutineRecQuery {
	query := (&DailyRoutineClient{config: drrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	drrq.withDailyRoutine = query
	return drrq
}

// WithProgramRec tells the query-builder to eager-load the nodes that are connected to
// the "program_rec" edge. The optional arguments are used to configure the query builder of the edge.
func (drrq *DailyRoutineRecQuery) WithProgramRec(opts ...func(*ProgramRecQuery)) *DailyRoutineRecQuery {
	query := (&ProgramRecClient{config: drrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	drrq.withProgramRec = query
	return drrq
}

// WithWeeklyRoutineRec tells the query-builder to eager-load the nodes that are connected to
// the "weekly_routine_rec" edge. The optional arguments are used to configure the query builder of the edge.
func (drrq *DailyRoutineRecQuery) WithWeeklyRoutineRec(opts ...func(*WeeklyRoutineRecQuery)) *DailyRoutineRecQuery {
	query := (&WeeklyRoutineRecClient{config: drrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	drrq.withWeeklyRoutineRec = query
	return drrq
}

// WithRoutineActRecs tells the query-builder to eager-load the nodes that are connected to
// the "routine_act_recs" edge. The optional arguments are used to configure the query builder of the edge.
func (drrq *DailyRoutineRecQuery) WithRoutineActRecs(opts ...func(*RoutineActRecQuery)) *DailyRoutineRecQuery {
	query := (&RoutineActRecClient{config: drrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	drrq.withRoutineActRecs = query
	return drrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Author uint64 `json:"author,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DailyRoutineRec.Query().
//		GroupBy(dailyroutinerec.FieldAuthor).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (drrq *DailyRoutineRecQuery) GroupBy(field string, fields ...string) *DailyRoutineRecGroupBy {
	drrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DailyRoutineRecGroupBy{build: drrq}
	grbuild.flds = &drrq.ctx.Fields
	grbuild.label = dailyroutinerec.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Author uint64 `json:"author,omitempty"`
//	}
//
//	client.DailyRoutineRec.Query().
//		Select(dailyroutinerec.FieldAuthor).
//		Scan(ctx, &v)
func (drrq *DailyRoutineRecQuery) Select(fields ...string) *DailyRoutineRecSelect {
	drrq.ctx.Fields = append(drrq.ctx.Fields, fields...)
	sbuild := &DailyRoutineRecSelect{DailyRoutineRecQuery: drrq}
	sbuild.label = dailyroutinerec.Label
	sbuild.flds, sbuild.scan = &drrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DailyRoutineRecSelect configured with the given aggregations.
func (drrq *DailyRoutineRecQuery) Aggregate(fns ...AggregateFunc) *DailyRoutineRecSelect {
	return drrq.Select().Aggregate(fns...)
}

func (drrq *DailyRoutineRecQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range drrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, drrq); err != nil {
				return err
			}
		}
	}
	for _, f := range drrq.ctx.Fields {
		if !dailyroutinerec.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if drrq.path != nil {
		prev, err := drrq.path(ctx)
		if err != nil {
			return err
		}
		drrq.sql = prev
	}
	return nil
}

func (drrq *DailyRoutineRecQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DailyRoutineRec, error) {
	var (
		nodes       = []*DailyRoutineRec{}
		_spec       = drrq.querySpec()
		loadedTypes = [4]bool{
			drrq.withDailyRoutine != nil,
			drrq.withProgramRec != nil,
			drrq.withWeeklyRoutineRec != nil,
			drrq.withRoutineActRecs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DailyRoutineRec).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DailyRoutineRec{config: drrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, drrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := drrq.withDailyRoutine; query != nil {
		if err := drrq.loadDailyRoutine(ctx, query, nodes, nil,
			func(n *DailyRoutineRec, e *DailyRoutine) { n.Edges.DailyRoutine = e }); err != nil {
			return nil, err
		}
	}
	if query := drrq.withProgramRec; query != nil {
		if err := drrq.loadProgramRec(ctx, query, nodes, nil,
			func(n *DailyRoutineRec, e *ProgramRec) { n.Edges.ProgramRec = e }); err != nil {
			return nil, err
		}
	}
	if query := drrq.withWeeklyRoutineRec; query != nil {
		if err := drrq.loadWeeklyRoutineRec(ctx, query, nodes, nil,
			func(n *DailyRoutineRec, e *WeeklyRoutineRec) { n.Edges.WeeklyRoutineRec = e }); err != nil {
			return nil, err
		}
	}
	if query := drrq.withRoutineActRecs; query != nil {
		if err := drrq.loadRoutineActRecs(ctx, query, nodes,
			func(n *DailyRoutineRec) { n.Edges.RoutineActRecs = []*RoutineActRec{} },
			func(n *DailyRoutineRec, e *RoutineActRec) { n.Edges.RoutineActRecs = append(n.Edges.RoutineActRecs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (drrq *DailyRoutineRecQuery) loadDailyRoutine(ctx context.Context, query *DailyRoutineQuery, nodes []*DailyRoutineRec, init func(*DailyRoutineRec), assign func(*DailyRoutineRec, *DailyRoutine)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*DailyRoutineRec)
	for i := range nodes {
		if nodes[i].DailyRoutineID == nil {
			continue
		}
		fk := *nodes[i].DailyRoutineID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(dailyroutine.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "daily_routine_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (drrq *DailyRoutineRecQuery) loadProgramRec(ctx context.Context, query *ProgramRecQuery, nodes []*DailyRoutineRec, init func(*DailyRoutineRec), assign func(*DailyRoutineRec, *ProgramRec)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*DailyRoutineRec)
	for i := range nodes {
		if nodes[i].ProgramRecID == nil {
			continue
		}
		fk := *nodes[i].ProgramRecID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(programrec.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "program_rec_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (drrq *DailyRoutineRecQuery) loadWeeklyRoutineRec(ctx context.Context, query *WeeklyRoutineRecQuery, nodes []*DailyRoutineRec, init func(*DailyRoutineRec), assign func(*DailyRoutineRec, *WeeklyRoutineRec)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*DailyRoutineRec)
	for i := range nodes {
		if nodes[i].WeeklyRoutineRecID == nil {
			continue
		}
		fk := *nodes[i].WeeklyRoutineRecID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(weeklyroutinerec.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "weekly_routine_rec_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (drrq *DailyRoutineRecQuery) loadRoutineActRecs(ctx context.Context, query *RoutineActRecQuery, nodes []*DailyRoutineRec, init func(*DailyRoutineRec), assign func(*DailyRoutineRec, *RoutineActRec)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*DailyRoutineRec)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(routineactrec.FieldDailyRoutineRecID)
	}
	query.Where(predicate.RoutineActRec(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(dailyroutinerec.RoutineActRecsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DailyRoutineRecID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "daily_routine_rec_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (drrq *DailyRoutineRecQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := drrq.querySpec()
	_spec.Node.Columns = drrq.ctx.Fields
	if len(drrq.ctx.Fields) > 0 {
		_spec.Unique = drrq.ctx.Unique != nil && *drrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, drrq.driver, _spec)
}

func (drrq *DailyRoutineRecQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dailyroutinerec.Table, dailyroutinerec.Columns, sqlgraph.NewFieldSpec(dailyroutinerec.FieldID, field.TypeUint64))
	_spec.From = drrq.sql
	if unique := drrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if drrq.path != nil {
		_spec.Unique = true
	}
	if fields := drrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dailyroutinerec.FieldID)
		for i := range fields {
			if fields[i] != dailyroutinerec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if drrq.withDailyRoutine != nil {
			_spec.Node.AddColumnOnce(dailyroutinerec.FieldDailyRoutineID)
		}
		if drrq.withProgramRec != nil {
			_spec.Node.AddColumnOnce(dailyroutinerec.FieldProgramRecID)
		}
		if drrq.withWeeklyRoutineRec != nil {
			_spec.Node.AddColumnOnce(dailyroutinerec.FieldWeeklyRoutineRecID)
		}
	}
	if ps := drrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := drrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := drrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := drrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (drrq *DailyRoutineRecQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(drrq.driver.Dialect())
	t1 := builder.Table(dailyroutinerec.Table)
	columns := drrq.ctx.Fields
	if len(columns) == 0 {
		columns = dailyroutinerec.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if drrq.sql != nil {
		selector = drrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if drrq.ctx.Unique != nil && *drrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range drrq.predicates {
		p(selector)
	}
	for _, p := range drrq.order {
		p(selector)
	}
	if offset := drrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := drrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DailyRoutineRecGroupBy is the group-by builder for DailyRoutineRec entities.
type DailyRoutineRecGroupBy struct {
	selector
	build *DailyRoutineRecQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (drrgb *DailyRoutineRecGroupBy) Aggregate(fns ...AggregateFunc) *DailyRoutineRecGroupBy {
	drrgb.fns = append(drrgb.fns, fns...)
	return drrgb
}

// Scan applies the selector query and scans the result into the given value.
func (drrgb *DailyRoutineRecGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, drrgb.build.ctx, "GroupBy")
	if err := drrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DailyRoutineRecQuery, *DailyRoutineRecGroupBy](ctx, drrgb.build, drrgb, drrgb.build.inters, v)
}

func (drrgb *DailyRoutineRecGroupBy) sqlScan(ctx context.Context, root *DailyRoutineRecQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(drrgb.fns))
	for _, fn := range drrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*drrgb.flds)+len(drrgb.fns))
		for _, f := range *drrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*drrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DailyRoutineRecSelect is the builder for selecting fields of DailyRoutineRec entities.
type DailyRoutineRecSelect struct {
	*DailyRoutineRecQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (drrs *DailyRoutineRecSelect) Aggregate(fns ...AggregateFunc) *DailyRoutineRecSelect {
	drrs.fns = append(drrs.fns, fns...)
	return drrs
}

// Scan applies the selector query and scans the result into the given value.
func (drrs *DailyRoutineRecSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, drrs.ctx, "Select")
	if err := drrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DailyRoutineRecQuery, *DailyRoutineRecSelect](ctx, drrs.DailyRoutineRecQuery, drrs, drrs.inters, v)
}

func (drrs *DailyRoutineRecSelect) sqlScan(ctx context.Context, root *DailyRoutineRecQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(drrs.fns))
	for _, fn := range drrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*drrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
