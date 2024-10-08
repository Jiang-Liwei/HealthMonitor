// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"healthmonitor/ent/adminjwtblacklist"
	"healthmonitor/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminJWTBlacklistQuery is the builder for querying AdminJWTBlacklist entities.
type AdminJWTBlacklistQuery struct {
	config
	ctx        *QueryContext
	order      []adminjwtblacklist.OrderOption
	inters     []Interceptor
	predicates []predicate.AdminJWTBlacklist
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminJWTBlacklistQuery builder.
func (ajbq *AdminJWTBlacklistQuery) Where(ps ...predicate.AdminJWTBlacklist) *AdminJWTBlacklistQuery {
	ajbq.predicates = append(ajbq.predicates, ps...)
	return ajbq
}

// Limit the number of records to be returned by this query.
func (ajbq *AdminJWTBlacklistQuery) Limit(limit int) *AdminJWTBlacklistQuery {
	ajbq.ctx.Limit = &limit
	return ajbq
}

// Offset to start from.
func (ajbq *AdminJWTBlacklistQuery) Offset(offset int) *AdminJWTBlacklistQuery {
	ajbq.ctx.Offset = &offset
	return ajbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ajbq *AdminJWTBlacklistQuery) Unique(unique bool) *AdminJWTBlacklistQuery {
	ajbq.ctx.Unique = &unique
	return ajbq
}

// Order specifies how the records should be ordered.
func (ajbq *AdminJWTBlacklistQuery) Order(o ...adminjwtblacklist.OrderOption) *AdminJWTBlacklistQuery {
	ajbq.order = append(ajbq.order, o...)
	return ajbq
}

// First returns the first AdminJWTBlacklist entity from the query.
// Returns a *NotFoundError when no AdminJWTBlacklist was found.
func (ajbq *AdminJWTBlacklistQuery) First(ctx context.Context) (*AdminJWTBlacklist, error) {
	nodes, err := ajbq.Limit(1).All(setContextOp(ctx, ajbq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adminjwtblacklist.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ajbq *AdminJWTBlacklistQuery) FirstX(ctx context.Context) *AdminJWTBlacklist {
	node, err := ajbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdminJWTBlacklist ID from the query.
// Returns a *NotFoundError when no AdminJWTBlacklist ID was found.
func (ajbq *AdminJWTBlacklistQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ajbq.Limit(1).IDs(setContextOp(ctx, ajbq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adminjwtblacklist.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ajbq *AdminJWTBlacklistQuery) FirstIDX(ctx context.Context) int {
	id, err := ajbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdminJWTBlacklist entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdminJWTBlacklist entity is found.
// Returns a *NotFoundError when no AdminJWTBlacklist entities are found.
func (ajbq *AdminJWTBlacklistQuery) Only(ctx context.Context) (*AdminJWTBlacklist, error) {
	nodes, err := ajbq.Limit(2).All(setContextOp(ctx, ajbq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adminjwtblacklist.Label}
	default:
		return nil, &NotSingularError{adminjwtblacklist.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ajbq *AdminJWTBlacklistQuery) OnlyX(ctx context.Context) *AdminJWTBlacklist {
	node, err := ajbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdminJWTBlacklist ID in the query.
// Returns a *NotSingularError when more than one AdminJWTBlacklist ID is found.
// Returns a *NotFoundError when no entities are found.
func (ajbq *AdminJWTBlacklistQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ajbq.Limit(2).IDs(setContextOp(ctx, ajbq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adminjwtblacklist.Label}
	default:
		err = &NotSingularError{adminjwtblacklist.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ajbq *AdminJWTBlacklistQuery) OnlyIDX(ctx context.Context) int {
	id, err := ajbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdminJWTBlacklists.
func (ajbq *AdminJWTBlacklistQuery) All(ctx context.Context) ([]*AdminJWTBlacklist, error) {
	ctx = setContextOp(ctx, ajbq.ctx, ent.OpQueryAll)
	if err := ajbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AdminJWTBlacklist, *AdminJWTBlacklistQuery]()
	return withInterceptors[[]*AdminJWTBlacklist](ctx, ajbq, qr, ajbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ajbq *AdminJWTBlacklistQuery) AllX(ctx context.Context) []*AdminJWTBlacklist {
	nodes, err := ajbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdminJWTBlacklist IDs.
func (ajbq *AdminJWTBlacklistQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ajbq.ctx.Unique == nil && ajbq.path != nil {
		ajbq.Unique(true)
	}
	ctx = setContextOp(ctx, ajbq.ctx, ent.OpQueryIDs)
	if err = ajbq.Select(adminjwtblacklist.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ajbq *AdminJWTBlacklistQuery) IDsX(ctx context.Context) []int {
	ids, err := ajbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ajbq *AdminJWTBlacklistQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ajbq.ctx, ent.OpQueryCount)
	if err := ajbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ajbq, querierCount[*AdminJWTBlacklistQuery](), ajbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ajbq *AdminJWTBlacklistQuery) CountX(ctx context.Context) int {
	count, err := ajbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ajbq *AdminJWTBlacklistQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ajbq.ctx, ent.OpQueryExist)
	switch _, err := ajbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ajbq *AdminJWTBlacklistQuery) ExistX(ctx context.Context) bool {
	exist, err := ajbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminJWTBlacklistQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ajbq *AdminJWTBlacklistQuery) Clone() *AdminJWTBlacklistQuery {
	if ajbq == nil {
		return nil
	}
	return &AdminJWTBlacklistQuery{
		config:     ajbq.config,
		ctx:        ajbq.ctx.Clone(),
		order:      append([]adminjwtblacklist.OrderOption{}, ajbq.order...),
		inters:     append([]Interceptor{}, ajbq.inters...),
		predicates: append([]predicate.AdminJWTBlacklist{}, ajbq.predicates...),
		// clone intermediate query.
		sql:  ajbq.sql.Clone(),
		path: ajbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Jti string `json:"jti,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdminJWTBlacklist.Query().
//		GroupBy(adminjwtblacklist.FieldJti).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ajbq *AdminJWTBlacklistQuery) GroupBy(field string, fields ...string) *AdminJWTBlacklistGroupBy {
	ajbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AdminJWTBlacklistGroupBy{build: ajbq}
	grbuild.flds = &ajbq.ctx.Fields
	grbuild.label = adminjwtblacklist.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Jti string `json:"jti,omitempty"`
//	}
//
//	client.AdminJWTBlacklist.Query().
//		Select(adminjwtblacklist.FieldJti).
//		Scan(ctx, &v)
func (ajbq *AdminJWTBlacklistQuery) Select(fields ...string) *AdminJWTBlacklistSelect {
	ajbq.ctx.Fields = append(ajbq.ctx.Fields, fields...)
	sbuild := &AdminJWTBlacklistSelect{AdminJWTBlacklistQuery: ajbq}
	sbuild.label = adminjwtblacklist.Label
	sbuild.flds, sbuild.scan = &ajbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AdminJWTBlacklistSelect configured with the given aggregations.
func (ajbq *AdminJWTBlacklistQuery) Aggregate(fns ...AggregateFunc) *AdminJWTBlacklistSelect {
	return ajbq.Select().Aggregate(fns...)
}

func (ajbq *AdminJWTBlacklistQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ajbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ajbq); err != nil {
				return err
			}
		}
	}
	for _, f := range ajbq.ctx.Fields {
		if !adminjwtblacklist.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ajbq.path != nil {
		prev, err := ajbq.path(ctx)
		if err != nil {
			return err
		}
		ajbq.sql = prev
	}
	return nil
}

func (ajbq *AdminJWTBlacklistQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdminJWTBlacklist, error) {
	var (
		nodes = []*AdminJWTBlacklist{}
		_spec = ajbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AdminJWTBlacklist).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AdminJWTBlacklist{config: ajbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ajbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ajbq *AdminJWTBlacklistQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ajbq.querySpec()
	_spec.Node.Columns = ajbq.ctx.Fields
	if len(ajbq.ctx.Fields) > 0 {
		_spec.Unique = ajbq.ctx.Unique != nil && *ajbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ajbq.driver, _spec)
}

func (ajbq *AdminJWTBlacklistQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(adminjwtblacklist.Table, adminjwtblacklist.Columns, sqlgraph.NewFieldSpec(adminjwtblacklist.FieldID, field.TypeInt))
	_spec.From = ajbq.sql
	if unique := ajbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ajbq.path != nil {
		_spec.Unique = true
	}
	if fields := ajbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminjwtblacklist.FieldID)
		for i := range fields {
			if fields[i] != adminjwtblacklist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ajbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ajbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ajbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ajbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ajbq *AdminJWTBlacklistQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ajbq.driver.Dialect())
	t1 := builder.Table(adminjwtblacklist.Table)
	columns := ajbq.ctx.Fields
	if len(columns) == 0 {
		columns = adminjwtblacklist.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ajbq.sql != nil {
		selector = ajbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ajbq.ctx.Unique != nil && *ajbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ajbq.predicates {
		p(selector)
	}
	for _, p := range ajbq.order {
		p(selector)
	}
	if offset := ajbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ajbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdminJWTBlacklistGroupBy is the group-by builder for AdminJWTBlacklist entities.
type AdminJWTBlacklistGroupBy struct {
	selector
	build *AdminJWTBlacklistQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ajbgb *AdminJWTBlacklistGroupBy) Aggregate(fns ...AggregateFunc) *AdminJWTBlacklistGroupBy {
	ajbgb.fns = append(ajbgb.fns, fns...)
	return ajbgb
}

// Scan applies the selector query and scans the result into the given value.
func (ajbgb *AdminJWTBlacklistGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ajbgb.build.ctx, ent.OpQueryGroupBy)
	if err := ajbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminJWTBlacklistQuery, *AdminJWTBlacklistGroupBy](ctx, ajbgb.build, ajbgb, ajbgb.build.inters, v)
}

func (ajbgb *AdminJWTBlacklistGroupBy) sqlScan(ctx context.Context, root *AdminJWTBlacklistQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ajbgb.fns))
	for _, fn := range ajbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ajbgb.flds)+len(ajbgb.fns))
		for _, f := range *ajbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ajbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ajbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AdminJWTBlacklistSelect is the builder for selecting fields of AdminJWTBlacklist entities.
type AdminJWTBlacklistSelect struct {
	*AdminJWTBlacklistQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ajbs *AdminJWTBlacklistSelect) Aggregate(fns ...AggregateFunc) *AdminJWTBlacklistSelect {
	ajbs.fns = append(ajbs.fns, fns...)
	return ajbs
}

// Scan applies the selector query and scans the result into the given value.
func (ajbs *AdminJWTBlacklistSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ajbs.ctx, ent.OpQuerySelect)
	if err := ajbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminJWTBlacklistQuery, *AdminJWTBlacklistSelect](ctx, ajbs.AdminJWTBlacklistQuery, ajbs, ajbs.inters, v)
}

func (ajbs *AdminJWTBlacklistSelect) sqlScan(ctx context.Context, root *AdminJWTBlacklistQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ajbs.fns))
	for _, fn := range ajbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ajbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ajbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
