// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"lifthus-auth/ent/lifthusgroup"
	"lifthus-auth/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LifthusGroupQuery is the builder for querying LifthusGroup entities.
type LifthusGroupQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.LifthusGroup
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LifthusGroupQuery builder.
func (lgq *LifthusGroupQuery) Where(ps ...predicate.LifthusGroup) *LifthusGroupQuery {
	lgq.predicates = append(lgq.predicates, ps...)
	return lgq
}

// Limit the number of records to be returned by this query.
func (lgq *LifthusGroupQuery) Limit(limit int) *LifthusGroupQuery {
	lgq.ctx.Limit = &limit
	return lgq
}

// Offset to start from.
func (lgq *LifthusGroupQuery) Offset(offset int) *LifthusGroupQuery {
	lgq.ctx.Offset = &offset
	return lgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lgq *LifthusGroupQuery) Unique(unique bool) *LifthusGroupQuery {
	lgq.ctx.Unique = &unique
	return lgq
}

// Order specifies how the records should be ordered.
func (lgq *LifthusGroupQuery) Order(o ...OrderFunc) *LifthusGroupQuery {
	lgq.order = append(lgq.order, o...)
	return lgq
}

// First returns the first LifthusGroup entity from the query.
// Returns a *NotFoundError when no LifthusGroup was found.
func (lgq *LifthusGroupQuery) First(ctx context.Context) (*LifthusGroup, error) {
	nodes, err := lgq.Limit(1).All(setContextOp(ctx, lgq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lifthusgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lgq *LifthusGroupQuery) FirstX(ctx context.Context) *LifthusGroup {
	node, err := lgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LifthusGroup ID from the query.
// Returns a *NotFoundError when no LifthusGroup ID was found.
func (lgq *LifthusGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lgq.Limit(1).IDs(setContextOp(ctx, lgq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lifthusgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lgq *LifthusGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := lgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LifthusGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LifthusGroup entity is found.
// Returns a *NotFoundError when no LifthusGroup entities are found.
func (lgq *LifthusGroupQuery) Only(ctx context.Context) (*LifthusGroup, error) {
	nodes, err := lgq.Limit(2).All(setContextOp(ctx, lgq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lifthusgroup.Label}
	default:
		return nil, &NotSingularError{lifthusgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lgq *LifthusGroupQuery) OnlyX(ctx context.Context) *LifthusGroup {
	node, err := lgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LifthusGroup ID in the query.
// Returns a *NotSingularError when more than one LifthusGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (lgq *LifthusGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lgq.Limit(2).IDs(setContextOp(ctx, lgq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lifthusgroup.Label}
	default:
		err = &NotSingularError{lifthusgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lgq *LifthusGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := lgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LifthusGroups.
func (lgq *LifthusGroupQuery) All(ctx context.Context) ([]*LifthusGroup, error) {
	ctx = setContextOp(ctx, lgq.ctx, "All")
	if err := lgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LifthusGroup, *LifthusGroupQuery]()
	return withInterceptors[[]*LifthusGroup](ctx, lgq, qr, lgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lgq *LifthusGroupQuery) AllX(ctx context.Context) []*LifthusGroup {
	nodes, err := lgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LifthusGroup IDs.
func (lgq *LifthusGroupQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lgq.ctx.Unique == nil && lgq.path != nil {
		lgq.Unique(true)
	}
	ctx = setContextOp(ctx, lgq.ctx, "IDs")
	if err = lgq.Select(lifthusgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lgq *LifthusGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := lgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lgq *LifthusGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lgq.ctx, "Count")
	if err := lgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lgq, querierCount[*LifthusGroupQuery](), lgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lgq *LifthusGroupQuery) CountX(ctx context.Context) int {
	count, err := lgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lgq *LifthusGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lgq.ctx, "Exist")
	switch _, err := lgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lgq *LifthusGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := lgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LifthusGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lgq *LifthusGroupQuery) Clone() *LifthusGroupQuery {
	if lgq == nil {
		return nil
	}
	return &LifthusGroupQuery{
		config:     lgq.config,
		ctx:        lgq.ctx.Clone(),
		order:      append([]OrderFunc{}, lgq.order...),
		inters:     append([]Interceptor{}, lgq.inters...),
		predicates: append([]predicate.LifthusGroup{}, lgq.predicates...),
		// clone intermediate query.
		sql:  lgq.sql.Clone(),
		path: lgq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (lgq *LifthusGroupQuery) GroupBy(field string, fields ...string) *LifthusGroupGroupBy {
	lgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LifthusGroupGroupBy{build: lgq}
	grbuild.flds = &lgq.ctx.Fields
	grbuild.label = lifthusgroup.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (lgq *LifthusGroupQuery) Select(fields ...string) *LifthusGroupSelect {
	lgq.ctx.Fields = append(lgq.ctx.Fields, fields...)
	sbuild := &LifthusGroupSelect{LifthusGroupQuery: lgq}
	sbuild.label = lifthusgroup.Label
	sbuild.flds, sbuild.scan = &lgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LifthusGroupSelect configured with the given aggregations.
func (lgq *LifthusGroupQuery) Aggregate(fns ...AggregateFunc) *LifthusGroupSelect {
	return lgq.Select().Aggregate(fns...)
}

func (lgq *LifthusGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lgq); err != nil {
				return err
			}
		}
	}
	for _, f := range lgq.ctx.Fields {
		if !lifthusgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lgq.path != nil {
		prev, err := lgq.path(ctx)
		if err != nil {
			return err
		}
		lgq.sql = prev
	}
	return nil
}

func (lgq *LifthusGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LifthusGroup, error) {
	var (
		nodes = []*LifthusGroup{}
		_spec = lgq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LifthusGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LifthusGroup{config: lgq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lgq *LifthusGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lgq.querySpec()
	_spec.Node.Columns = lgq.ctx.Fields
	if len(lgq.ctx.Fields) > 0 {
		_spec.Unique = lgq.ctx.Unique != nil && *lgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lgq.driver, _spec)
}

func (lgq *LifthusGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(lifthusgroup.Table, lifthusgroup.Columns, sqlgraph.NewFieldSpec(lifthusgroup.FieldID, field.TypeInt))
	_spec.From = lgq.sql
	if unique := lgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lgq.path != nil {
		_spec.Unique = true
	}
	if fields := lgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lifthusgroup.FieldID)
		for i := range fields {
			if fields[i] != lifthusgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lgq *LifthusGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lgq.driver.Dialect())
	t1 := builder.Table(lifthusgroup.Table)
	columns := lgq.ctx.Fields
	if len(columns) == 0 {
		columns = lifthusgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lgq.sql != nil {
		selector = lgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lgq.ctx.Unique != nil && *lgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lgq.predicates {
		p(selector)
	}
	for _, p := range lgq.order {
		p(selector)
	}
	if offset := lgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LifthusGroupGroupBy is the group-by builder for LifthusGroup entities.
type LifthusGroupGroupBy struct {
	selector
	build *LifthusGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lggb *LifthusGroupGroupBy) Aggregate(fns ...AggregateFunc) *LifthusGroupGroupBy {
	lggb.fns = append(lggb.fns, fns...)
	return lggb
}

// Scan applies the selector query and scans the result into the given value.
func (lggb *LifthusGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lggb.build.ctx, "GroupBy")
	if err := lggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LifthusGroupQuery, *LifthusGroupGroupBy](ctx, lggb.build, lggb, lggb.build.inters, v)
}

func (lggb *LifthusGroupGroupBy) sqlScan(ctx context.Context, root *LifthusGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lggb.fns))
	for _, fn := range lggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lggb.flds)+len(lggb.fns))
		for _, f := range *lggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LifthusGroupSelect is the builder for selecting fields of LifthusGroup entities.
type LifthusGroupSelect struct {
	*LifthusGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lgs *LifthusGroupSelect) Aggregate(fns ...AggregateFunc) *LifthusGroupSelect {
	lgs.fns = append(lgs.fns, fns...)
	return lgs
}

// Scan applies the selector query and scans the result into the given value.
func (lgs *LifthusGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgs.ctx, "Select")
	if err := lgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LifthusGroupQuery, *LifthusGroupSelect](ctx, lgs.LifthusGroupQuery, lgs, lgs.inters, v)
}

func (lgs *LifthusGroupSelect) sqlScan(ctx context.Context, root *LifthusGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lgs.fns))
	for _, fn := range lgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}