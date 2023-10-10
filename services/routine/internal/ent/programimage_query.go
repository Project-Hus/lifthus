// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"routine/internal/ent/image"
	"routine/internal/ent/predicate"
	"routine/internal/ent/programimage"
	"routine/internal/ent/programversion"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramImageQuery is the builder for querying ProgramImage entities.
type ProgramImageQuery struct {
	config
	ctx                *QueryContext
	order              []programimage.OrderOption
	inters             []Interceptor
	predicates         []predicate.ProgramImage
	withProgramVersion *ProgramVersionQuery
	withImage          *ImageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProgramImageQuery builder.
func (piq *ProgramImageQuery) Where(ps ...predicate.ProgramImage) *ProgramImageQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit the number of records to be returned by this query.
func (piq *ProgramImageQuery) Limit(limit int) *ProgramImageQuery {
	piq.ctx.Limit = &limit
	return piq
}

// Offset to start from.
func (piq *ProgramImageQuery) Offset(offset int) *ProgramImageQuery {
	piq.ctx.Offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *ProgramImageQuery) Unique(unique bool) *ProgramImageQuery {
	piq.ctx.Unique = &unique
	return piq
}

// Order specifies how the records should be ordered.
func (piq *ProgramImageQuery) Order(o ...programimage.OrderOption) *ProgramImageQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// QueryProgramVersion chains the current query on the "program_version" edge.
func (piq *ProgramImageQuery) QueryProgramVersion() *ProgramVersionQuery {
	query := (&ProgramVersionClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(programimage.Table, programimage.FieldID, selector),
			sqlgraph.To(programversion.Table, programversion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, programimage.ProgramVersionTable, programimage.ProgramVersionColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryImage chains the current query on the "image" edge.
func (piq *ProgramImageQuery) QueryImage() *ImageQuery {
	query := (&ImageClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(programimage.Table, programimage.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, programimage.ImageTable, programimage.ImageColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProgramImage entity from the query.
// Returns a *NotFoundError when no ProgramImage was found.
func (piq *ProgramImageQuery) First(ctx context.Context) (*ProgramImage, error) {
	nodes, err := piq.Limit(1).All(setContextOp(ctx, piq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{programimage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *ProgramImageQuery) FirstX(ctx context.Context) *ProgramImage {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProgramImage ID from the query.
// Returns a *NotFoundError when no ProgramImage ID was found.
func (piq *ProgramImageQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = piq.Limit(1).IDs(setContextOp(ctx, piq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{programimage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *ProgramImageQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProgramImage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProgramImage entity is found.
// Returns a *NotFoundError when no ProgramImage entities are found.
func (piq *ProgramImageQuery) Only(ctx context.Context) (*ProgramImage, error) {
	nodes, err := piq.Limit(2).All(setContextOp(ctx, piq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{programimage.Label}
	default:
		return nil, &NotSingularError{programimage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *ProgramImageQuery) OnlyX(ctx context.Context) *ProgramImage {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProgramImage ID in the query.
// Returns a *NotSingularError when more than one ProgramImage ID is found.
// Returns a *NotFoundError when no entities are found.
func (piq *ProgramImageQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = piq.Limit(2).IDs(setContextOp(ctx, piq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{programimage.Label}
	default:
		err = &NotSingularError{programimage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *ProgramImageQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProgramImages.
func (piq *ProgramImageQuery) All(ctx context.Context) ([]*ProgramImage, error) {
	ctx = setContextOp(ctx, piq.ctx, "All")
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProgramImage, *ProgramImageQuery]()
	return withInterceptors[[]*ProgramImage](ctx, piq, qr, piq.inters)
}

// AllX is like All, but panics if an error occurs.
func (piq *ProgramImageQuery) AllX(ctx context.Context) []*ProgramImage {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProgramImage IDs.
func (piq *ProgramImageQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if piq.ctx.Unique == nil && piq.path != nil {
		piq.Unique(true)
	}
	ctx = setContextOp(ctx, piq.ctx, "IDs")
	if err = piq.Select(programimage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *ProgramImageQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *ProgramImageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, piq.ctx, "Count")
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, piq, querierCount[*ProgramImageQuery](), piq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (piq *ProgramImageQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *ProgramImageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, piq.ctx, "Exist")
	switch _, err := piq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *ProgramImageQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProgramImageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *ProgramImageQuery) Clone() *ProgramImageQuery {
	if piq == nil {
		return nil
	}
	return &ProgramImageQuery{
		config:             piq.config,
		ctx:                piq.ctx.Clone(),
		order:              append([]programimage.OrderOption{}, piq.order...),
		inters:             append([]Interceptor{}, piq.inters...),
		predicates:         append([]predicate.ProgramImage{}, piq.predicates...),
		withProgramVersion: piq.withProgramVersion.Clone(),
		withImage:          piq.withImage.Clone(),
		// clone intermediate query.
		sql:  piq.sql.Clone(),
		path: piq.path,
	}
}

// WithProgramVersion tells the query-builder to eager-load the nodes that are connected to
// the "program_version" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProgramImageQuery) WithProgramVersion(opts ...func(*ProgramVersionQuery)) *ProgramImageQuery {
	query := (&ProgramVersionClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withProgramVersion = query
	return piq
}

// WithImage tells the query-builder to eager-load the nodes that are connected to
// the "image" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProgramImageQuery) WithImage(opts ...func(*ImageQuery)) *ProgramImageQuery {
	query := (&ImageClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withImage = query
	return piq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Order uint `json:"order,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProgramImage.Query().
//		GroupBy(programimage.FieldOrder).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (piq *ProgramImageQuery) GroupBy(field string, fields ...string) *ProgramImageGroupBy {
	piq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProgramImageGroupBy{build: piq}
	grbuild.flds = &piq.ctx.Fields
	grbuild.label = programimage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Order uint `json:"order,omitempty"`
//	}
//
//	client.ProgramImage.Query().
//		Select(programimage.FieldOrder).
//		Scan(ctx, &v)
func (piq *ProgramImageQuery) Select(fields ...string) *ProgramImageSelect {
	piq.ctx.Fields = append(piq.ctx.Fields, fields...)
	sbuild := &ProgramImageSelect{ProgramImageQuery: piq}
	sbuild.label = programimage.Label
	sbuild.flds, sbuild.scan = &piq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProgramImageSelect configured with the given aggregations.
func (piq *ProgramImageQuery) Aggregate(fns ...AggregateFunc) *ProgramImageSelect {
	return piq.Select().Aggregate(fns...)
}

func (piq *ProgramImageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range piq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, piq); err != nil {
				return err
			}
		}
	}
	for _, f := range piq.ctx.Fields {
		if !programimage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *ProgramImageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProgramImage, error) {
	var (
		nodes       = []*ProgramImage{}
		_spec       = piq.querySpec()
		loadedTypes = [2]bool{
			piq.withProgramVersion != nil,
			piq.withImage != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProgramImage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProgramImage{config: piq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := piq.withProgramVersion; query != nil {
		if err := piq.loadProgramVersion(ctx, query, nodes, nil,
			func(n *ProgramImage, e *ProgramVersion) { n.Edges.ProgramVersion = e }); err != nil {
			return nil, err
		}
	}
	if query := piq.withImage; query != nil {
		if err := piq.loadImage(ctx, query, nodes, nil,
			func(n *ProgramImage, e *Image) { n.Edges.Image = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (piq *ProgramImageQuery) loadProgramVersion(ctx context.Context, query *ProgramVersionQuery, nodes []*ProgramImage, init func(*ProgramImage), assign func(*ProgramImage, *ProgramVersion)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*ProgramImage)
	for i := range nodes {
		fk := nodes[i].ProgramVersionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(programversion.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "program_version_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (piq *ProgramImageQuery) loadImage(ctx context.Context, query *ImageQuery, nodes []*ProgramImage, init func(*ProgramImage), assign func(*ProgramImage, *Image)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*ProgramImage)
	for i := range nodes {
		fk := nodes[i].ImageID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(image.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "image_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (piq *ProgramImageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	_spec.Node.Columns = piq.ctx.Fields
	if len(piq.ctx.Fields) > 0 {
		_spec.Unique = piq.ctx.Unique != nil && *piq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *ProgramImageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(programimage.Table, programimage.Columns, sqlgraph.NewFieldSpec(programimage.FieldID, field.TypeUint64))
	_spec.From = piq.sql
	if unique := piq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if piq.path != nil {
		_spec.Unique = true
	}
	if fields := piq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, programimage.FieldID)
		for i := range fields {
			if fields[i] != programimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if piq.withProgramVersion != nil {
			_spec.Node.AddColumnOnce(programimage.FieldProgramVersionID)
		}
		if piq.withImage != nil {
			_spec.Node.AddColumnOnce(programimage.FieldImageID)
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *ProgramImageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(programimage.Table)
	columns := piq.ctx.Fields
	if len(columns) == 0 {
		columns = programimage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.ctx.Unique != nil && *piq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProgramImageGroupBy is the group-by builder for ProgramImage entities.
type ProgramImageGroupBy struct {
	selector
	build *ProgramImageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *ProgramImageGroupBy) Aggregate(fns ...AggregateFunc) *ProgramImageGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the selector query and scans the result into the given value.
func (pigb *ProgramImageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pigb.build.ctx, "GroupBy")
	if err := pigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProgramImageQuery, *ProgramImageGroupBy](ctx, pigb.build, pigb, pigb.build.inters, v)
}

func (pigb *ProgramImageGroupBy) sqlScan(ctx context.Context, root *ProgramImageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pigb.flds)+len(pigb.fns))
		for _, f := range *pigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProgramImageSelect is the builder for selecting fields of ProgramImage entities.
type ProgramImageSelect struct {
	*ProgramImageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pis *ProgramImageSelect) Aggregate(fns ...AggregateFunc) *ProgramImageSelect {
	pis.fns = append(pis.fns, fns...)
	return pis
}

// Scan applies the selector query and scans the result into the given value.
func (pis *ProgramImageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pis.ctx, "Select")
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProgramImageQuery, *ProgramImageSelect](ctx, pis.ProgramImageQuery, pis, pis.inters, v)
}

func (pis *ProgramImageSelect) sqlScan(ctx context.Context, root *ProgramImageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pis.fns))
	for _, fn := range pis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}