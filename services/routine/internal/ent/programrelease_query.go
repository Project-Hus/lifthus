// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"routine/internal/ent/predicate"
	"routine/internal/ent/program"
	"routine/internal/ent/programrelease"
	"routine/internal/ent/routine"
	"routine/internal/ent/s3programimage"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProgramReleaseQuery is the builder for querying ProgramRelease entities.
type ProgramReleaseQuery struct {
	config
	ctx                 *QueryContext
	order               []programrelease.OrderOption
	inters              []Interceptor
	predicates          []predicate.ProgramRelease
	withProgram         *ProgramQuery
	withS3ProgramImages *S3ProgramImageQuery
	withRoutines        *RoutineQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProgramReleaseQuery builder.
func (prq *ProgramReleaseQuery) Where(ps ...predicate.ProgramRelease) *ProgramReleaseQuery {
	prq.predicates = append(prq.predicates, ps...)
	return prq
}

// Limit the number of records to be returned by this query.
func (prq *ProgramReleaseQuery) Limit(limit int) *ProgramReleaseQuery {
	prq.ctx.Limit = &limit
	return prq
}

// Offset to start from.
func (prq *ProgramReleaseQuery) Offset(offset int) *ProgramReleaseQuery {
	prq.ctx.Offset = &offset
	return prq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prq *ProgramReleaseQuery) Unique(unique bool) *ProgramReleaseQuery {
	prq.ctx.Unique = &unique
	return prq
}

// Order specifies how the records should be ordered.
func (prq *ProgramReleaseQuery) Order(o ...programrelease.OrderOption) *ProgramReleaseQuery {
	prq.order = append(prq.order, o...)
	return prq
}

// QueryProgram chains the current query on the "program" edge.
func (prq *ProgramReleaseQuery) QueryProgram() *ProgramQuery {
	query := (&ProgramClient{config: prq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(programrelease.Table, programrelease.FieldID, selector),
			sqlgraph.To(program.Table, program.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, programrelease.ProgramTable, programrelease.ProgramColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryS3ProgramImages chains the current query on the "s3_program_images" edge.
func (prq *ProgramReleaseQuery) QueryS3ProgramImages() *S3ProgramImageQuery {
	query := (&S3ProgramImageClient{config: prq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(programrelease.Table, programrelease.FieldID, selector),
			sqlgraph.To(s3programimage.Table, s3programimage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, programrelease.S3ProgramImagesTable, programrelease.S3ProgramImagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRoutines chains the current query on the "routines" edge.
func (prq *ProgramReleaseQuery) QueryRoutines() *RoutineQuery {
	query := (&RoutineClient{config: prq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(programrelease.Table, programrelease.FieldID, selector),
			sqlgraph.To(routine.Table, routine.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, programrelease.RoutinesTable, programrelease.RoutinesColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProgramRelease entity from the query.
// Returns a *NotFoundError when no ProgramRelease was found.
func (prq *ProgramReleaseQuery) First(ctx context.Context) (*ProgramRelease, error) {
	nodes, err := prq.Limit(1).All(setContextOp(ctx, prq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{programrelease.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prq *ProgramReleaseQuery) FirstX(ctx context.Context) *ProgramRelease {
	node, err := prq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProgramRelease ID from the query.
// Returns a *NotFoundError when no ProgramRelease ID was found.
func (prq *ProgramReleaseQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = prq.Limit(1).IDs(setContextOp(ctx, prq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{programrelease.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prq *ProgramReleaseQuery) FirstIDX(ctx context.Context) int64 {
	id, err := prq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProgramRelease entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProgramRelease entity is found.
// Returns a *NotFoundError when no ProgramRelease entities are found.
func (prq *ProgramReleaseQuery) Only(ctx context.Context) (*ProgramRelease, error) {
	nodes, err := prq.Limit(2).All(setContextOp(ctx, prq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{programrelease.Label}
	default:
		return nil, &NotSingularError{programrelease.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prq *ProgramReleaseQuery) OnlyX(ctx context.Context) *ProgramRelease {
	node, err := prq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProgramRelease ID in the query.
// Returns a *NotSingularError when more than one ProgramRelease ID is found.
// Returns a *NotFoundError when no entities are found.
func (prq *ProgramReleaseQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = prq.Limit(2).IDs(setContextOp(ctx, prq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{programrelease.Label}
	default:
		err = &NotSingularError{programrelease.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prq *ProgramReleaseQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := prq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProgramReleases.
func (prq *ProgramReleaseQuery) All(ctx context.Context) ([]*ProgramRelease, error) {
	ctx = setContextOp(ctx, prq.ctx, "All")
	if err := prq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProgramRelease, *ProgramReleaseQuery]()
	return withInterceptors[[]*ProgramRelease](ctx, prq, qr, prq.inters)
}

// AllX is like All, but panics if an error occurs.
func (prq *ProgramReleaseQuery) AllX(ctx context.Context) []*ProgramRelease {
	nodes, err := prq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProgramRelease IDs.
func (prq *ProgramReleaseQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if prq.ctx.Unique == nil && prq.path != nil {
		prq.Unique(true)
	}
	ctx = setContextOp(ctx, prq.ctx, "IDs")
	if err = prq.Select(programrelease.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prq *ProgramReleaseQuery) IDsX(ctx context.Context) []int64 {
	ids, err := prq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prq *ProgramReleaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, prq.ctx, "Count")
	if err := prq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, prq, querierCount[*ProgramReleaseQuery](), prq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (prq *ProgramReleaseQuery) CountX(ctx context.Context) int {
	count, err := prq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prq *ProgramReleaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, prq.ctx, "Exist")
	switch _, err := prq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (prq *ProgramReleaseQuery) ExistX(ctx context.Context) bool {
	exist, err := prq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProgramReleaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prq *ProgramReleaseQuery) Clone() *ProgramReleaseQuery {
	if prq == nil {
		return nil
	}
	return &ProgramReleaseQuery{
		config:              prq.config,
		ctx:                 prq.ctx.Clone(),
		order:               append([]programrelease.OrderOption{}, prq.order...),
		inters:              append([]Interceptor{}, prq.inters...),
		predicates:          append([]predicate.ProgramRelease{}, prq.predicates...),
		withProgram:         prq.withProgram.Clone(),
		withS3ProgramImages: prq.withS3ProgramImages.Clone(),
		withRoutines:        prq.withRoutines.Clone(),
		// clone intermediate query.
		sql:  prq.sql.Clone(),
		path: prq.path,
	}
}

// WithProgram tells the query-builder to eager-load the nodes that are connected to
// the "program" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *ProgramReleaseQuery) WithProgram(opts ...func(*ProgramQuery)) *ProgramReleaseQuery {
	query := (&ProgramClient{config: prq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	prq.withProgram = query
	return prq
}

// WithS3ProgramImages tells the query-builder to eager-load the nodes that are connected to
// the "s3_program_images" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *ProgramReleaseQuery) WithS3ProgramImages(opts ...func(*S3ProgramImageQuery)) *ProgramReleaseQuery {
	query := (&S3ProgramImageClient{config: prq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	prq.withS3ProgramImages = query
	return prq
}

// WithRoutines tells the query-builder to eager-load the nodes that are connected to
// the "routines" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *ProgramReleaseQuery) WithRoutines(opts ...func(*RoutineQuery)) *ProgramReleaseQuery {
	query := (&RoutineClient{config: prq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	prq.withRoutines = query
	return prq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Version int `json:"version,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProgramRelease.Query().
//		GroupBy(programrelease.FieldVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (prq *ProgramReleaseQuery) GroupBy(field string, fields ...string) *ProgramReleaseGroupBy {
	prq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProgramReleaseGroupBy{build: prq}
	grbuild.flds = &prq.ctx.Fields
	grbuild.label = programrelease.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Version int `json:"version,omitempty"`
//	}
//
//	client.ProgramRelease.Query().
//		Select(programrelease.FieldVersion).
//		Scan(ctx, &v)
func (prq *ProgramReleaseQuery) Select(fields ...string) *ProgramReleaseSelect {
	prq.ctx.Fields = append(prq.ctx.Fields, fields...)
	sbuild := &ProgramReleaseSelect{ProgramReleaseQuery: prq}
	sbuild.label = programrelease.Label
	sbuild.flds, sbuild.scan = &prq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProgramReleaseSelect configured with the given aggregations.
func (prq *ProgramReleaseQuery) Aggregate(fns ...AggregateFunc) *ProgramReleaseSelect {
	return prq.Select().Aggregate(fns...)
}

func (prq *ProgramReleaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range prq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, prq); err != nil {
				return err
			}
		}
	}
	for _, f := range prq.ctx.Fields {
		if !programrelease.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prq.path != nil {
		prev, err := prq.path(ctx)
		if err != nil {
			return err
		}
		prq.sql = prev
	}
	return nil
}

func (prq *ProgramReleaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProgramRelease, error) {
	var (
		nodes       = []*ProgramRelease{}
		withFKs     = prq.withFKs
		_spec       = prq.querySpec()
		loadedTypes = [3]bool{
			prq.withProgram != nil,
			prq.withS3ProgramImages != nil,
			prq.withRoutines != nil,
		}
	)
	if prq.withProgram != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, programrelease.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProgramRelease).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProgramRelease{config: prq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, prq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := prq.withProgram; query != nil {
		if err := prq.loadProgram(ctx, query, nodes, nil,
			func(n *ProgramRelease, e *Program) { n.Edges.Program = e }); err != nil {
			return nil, err
		}
	}
	if query := prq.withS3ProgramImages; query != nil {
		if err := prq.loadS3ProgramImages(ctx, query, nodes,
			func(n *ProgramRelease) { n.Edges.S3ProgramImages = []*S3ProgramImage{} },
			func(n *ProgramRelease, e *S3ProgramImage) {
				n.Edges.S3ProgramImages = append(n.Edges.S3ProgramImages, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := prq.withRoutines; query != nil {
		if err := prq.loadRoutines(ctx, query, nodes,
			func(n *ProgramRelease) { n.Edges.Routines = []*Routine{} },
			func(n *ProgramRelease, e *Routine) { n.Edges.Routines = append(n.Edges.Routines, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (prq *ProgramReleaseQuery) loadProgram(ctx context.Context, query *ProgramQuery, nodes []*ProgramRelease, init func(*ProgramRelease), assign func(*ProgramRelease, *Program)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ProgramRelease)
	for i := range nodes {
		if nodes[i].program_program_releases == nil {
			continue
		}
		fk := *nodes[i].program_program_releases
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(program.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "program_program_releases" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (prq *ProgramReleaseQuery) loadS3ProgramImages(ctx context.Context, query *S3ProgramImageQuery, nodes []*ProgramRelease, init func(*ProgramRelease), assign func(*ProgramRelease, *S3ProgramImage)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*ProgramRelease)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(s3programimage.FieldProgramReleaseID)
	}
	query.Where(predicate.S3ProgramImage(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(programrelease.S3ProgramImagesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProgramReleaseID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "program_release_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (prq *ProgramReleaseQuery) loadRoutines(ctx context.Context, query *RoutineQuery, nodes []*ProgramRelease, init func(*ProgramRelease), assign func(*ProgramRelease, *Routine)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*ProgramRelease)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Routine(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(programrelease.RoutinesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.program_release_routines
		if fk == nil {
			return fmt.Errorf(`foreign-key "program_release_routines" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "program_release_routines" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (prq *ProgramReleaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prq.querySpec()
	_spec.Node.Columns = prq.ctx.Fields
	if len(prq.ctx.Fields) > 0 {
		_spec.Unique = prq.ctx.Unique != nil && *prq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, prq.driver, _spec)
}

func (prq *ProgramReleaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(programrelease.Table, programrelease.Columns, sqlgraph.NewFieldSpec(programrelease.FieldID, field.TypeInt64))
	_spec.From = prq.sql
	if unique := prq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if prq.path != nil {
		_spec.Unique = true
	}
	if fields := prq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, programrelease.FieldID)
		for i := range fields {
			if fields[i] != programrelease.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prq *ProgramReleaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prq.driver.Dialect())
	t1 := builder.Table(programrelease.Table)
	columns := prq.ctx.Fields
	if len(columns) == 0 {
		columns = programrelease.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prq.sql != nil {
		selector = prq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if prq.ctx.Unique != nil && *prq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range prq.predicates {
		p(selector)
	}
	for _, p := range prq.order {
		p(selector)
	}
	if offset := prq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProgramReleaseGroupBy is the group-by builder for ProgramRelease entities.
type ProgramReleaseGroupBy struct {
	selector
	build *ProgramReleaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prgb *ProgramReleaseGroupBy) Aggregate(fns ...AggregateFunc) *ProgramReleaseGroupBy {
	prgb.fns = append(prgb.fns, fns...)
	return prgb
}

// Scan applies the selector query and scans the result into the given value.
func (prgb *ProgramReleaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prgb.build.ctx, "GroupBy")
	if err := prgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProgramReleaseQuery, *ProgramReleaseGroupBy](ctx, prgb.build, prgb, prgb.build.inters, v)
}

func (prgb *ProgramReleaseGroupBy) sqlScan(ctx context.Context, root *ProgramReleaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(prgb.fns))
	for _, fn := range prgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*prgb.flds)+len(prgb.fns))
		for _, f := range *prgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*prgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProgramReleaseSelect is the builder for selecting fields of ProgramRelease entities.
type ProgramReleaseSelect struct {
	*ProgramReleaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (prs *ProgramReleaseSelect) Aggregate(fns ...AggregateFunc) *ProgramReleaseSelect {
	prs.fns = append(prs.fns, fns...)
	return prs
}

// Scan applies the selector query and scans the result into the given value.
func (prs *ProgramReleaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prs.ctx, "Select")
	if err := prs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProgramReleaseQuery, *ProgramReleaseSelect](ctx, prs.ProgramReleaseQuery, prs, prs.inters, v)
}

func (prs *ProgramReleaseSelect) sqlScan(ctx context.Context, root *ProgramReleaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(prs.fns))
	for _, fn := range prs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*prs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}