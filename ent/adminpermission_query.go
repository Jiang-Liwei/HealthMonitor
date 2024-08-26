// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/adminpermission"
	"HealthMonitor/ent/adminrolepermission"
	"HealthMonitor/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AdminPermissionQuery is the builder for querying AdminPermission entities.
type AdminPermissionQuery struct {
	config
	ctx        *QueryContext
	order      []adminpermission.OrderOption
	inters     []Interceptor
	predicates []predicate.AdminPermission
	withRoles  *AdminRolePermissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminPermissionQuery builder.
func (apq *AdminPermissionQuery) Where(ps ...predicate.AdminPermission) *AdminPermissionQuery {
	apq.predicates = append(apq.predicates, ps...)
	return apq
}

// Limit the number of records to be returned by this query.
func (apq *AdminPermissionQuery) Limit(limit int) *AdminPermissionQuery {
	apq.ctx.Limit = &limit
	return apq
}

// Offset to start from.
func (apq *AdminPermissionQuery) Offset(offset int) *AdminPermissionQuery {
	apq.ctx.Offset = &offset
	return apq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (apq *AdminPermissionQuery) Unique(unique bool) *AdminPermissionQuery {
	apq.ctx.Unique = &unique
	return apq
}

// Order specifies how the records should be ordered.
func (apq *AdminPermissionQuery) Order(o ...adminpermission.OrderOption) *AdminPermissionQuery {
	apq.order = append(apq.order, o...)
	return apq
}

// QueryRoles chains the current query on the "roles" edge.
func (apq *AdminPermissionQuery) QueryRoles() *AdminRolePermissionQuery {
	query := (&AdminRolePermissionClient{config: apq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adminpermission.Table, adminpermission.FieldID, selector),
			sqlgraph.To(adminrolepermission.Table, adminrolepermission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adminpermission.RolesTable, adminpermission.RolesColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AdminPermission entity from the query.
// Returns a *NotFoundError when no AdminPermission was found.
func (apq *AdminPermissionQuery) First(ctx context.Context) (*AdminPermission, error) {
	nodes, err := apq.Limit(1).All(setContextOp(ctx, apq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adminpermission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (apq *AdminPermissionQuery) FirstX(ctx context.Context) *AdminPermission {
	node, err := apq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdminPermission ID from the query.
// Returns a *NotFoundError when no AdminPermission ID was found.
func (apq *AdminPermissionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = apq.Limit(1).IDs(setContextOp(ctx, apq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adminpermission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (apq *AdminPermissionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := apq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdminPermission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdminPermission entity is found.
// Returns a *NotFoundError when no AdminPermission entities are found.
func (apq *AdminPermissionQuery) Only(ctx context.Context) (*AdminPermission, error) {
	nodes, err := apq.Limit(2).All(setContextOp(ctx, apq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adminpermission.Label}
	default:
		return nil, &NotSingularError{adminpermission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (apq *AdminPermissionQuery) OnlyX(ctx context.Context) *AdminPermission {
	node, err := apq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdminPermission ID in the query.
// Returns a *NotSingularError when more than one AdminPermission ID is found.
// Returns a *NotFoundError when no entities are found.
func (apq *AdminPermissionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = apq.Limit(2).IDs(setContextOp(ctx, apq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adminpermission.Label}
	default:
		err = &NotSingularError{adminpermission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (apq *AdminPermissionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := apq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdminPermissions.
func (apq *AdminPermissionQuery) All(ctx context.Context) ([]*AdminPermission, error) {
	ctx = setContextOp(ctx, apq.ctx, ent.OpQueryAll)
	if err := apq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AdminPermission, *AdminPermissionQuery]()
	return withInterceptors[[]*AdminPermission](ctx, apq, qr, apq.inters)
}

// AllX is like All, but panics if an error occurs.
func (apq *AdminPermissionQuery) AllX(ctx context.Context) []*AdminPermission {
	nodes, err := apq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdminPermission IDs.
func (apq *AdminPermissionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if apq.ctx.Unique == nil && apq.path != nil {
		apq.Unique(true)
	}
	ctx = setContextOp(ctx, apq.ctx, ent.OpQueryIDs)
	if err = apq.Select(adminpermission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (apq *AdminPermissionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := apq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (apq *AdminPermissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, apq.ctx, ent.OpQueryCount)
	if err := apq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, apq, querierCount[*AdminPermissionQuery](), apq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (apq *AdminPermissionQuery) CountX(ctx context.Context) int {
	count, err := apq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (apq *AdminPermissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, apq.ctx, ent.OpQueryExist)
	switch _, err := apq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (apq *AdminPermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := apq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminPermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (apq *AdminPermissionQuery) Clone() *AdminPermissionQuery {
	if apq == nil {
		return nil
	}
	return &AdminPermissionQuery{
		config:     apq.config,
		ctx:        apq.ctx.Clone(),
		order:      append([]adminpermission.OrderOption{}, apq.order...),
		inters:     append([]Interceptor{}, apq.inters...),
		predicates: append([]predicate.AdminPermission{}, apq.predicates...),
		withRoles:  apq.withRoles.Clone(),
		// clone intermediate query.
		sql:  apq.sql.Clone(),
		path: apq.path,
	}
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *AdminPermissionQuery) WithRoles(opts ...func(*AdminRolePermissionQuery)) *AdminPermissionQuery {
	query := (&AdminRolePermissionClient{config: apq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	apq.withRoles = query
	return apq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdminPermission.Query().
//		GroupBy(adminpermission.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (apq *AdminPermissionQuery) GroupBy(field string, fields ...string) *AdminPermissionGroupBy {
	apq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AdminPermissionGroupBy{build: apq}
	grbuild.flds = &apq.ctx.Fields
	grbuild.label = adminpermission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.AdminPermission.Query().
//		Select(adminpermission.FieldName).
//		Scan(ctx, &v)
func (apq *AdminPermissionQuery) Select(fields ...string) *AdminPermissionSelect {
	apq.ctx.Fields = append(apq.ctx.Fields, fields...)
	sbuild := &AdminPermissionSelect{AdminPermissionQuery: apq}
	sbuild.label = adminpermission.Label
	sbuild.flds, sbuild.scan = &apq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AdminPermissionSelect configured with the given aggregations.
func (apq *AdminPermissionQuery) Aggregate(fns ...AggregateFunc) *AdminPermissionSelect {
	return apq.Select().Aggregate(fns...)
}

func (apq *AdminPermissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range apq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, apq); err != nil {
				return err
			}
		}
	}
	for _, f := range apq.ctx.Fields {
		if !adminpermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if apq.path != nil {
		prev, err := apq.path(ctx)
		if err != nil {
			return err
		}
		apq.sql = prev
	}
	return nil
}

func (apq *AdminPermissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdminPermission, error) {
	var (
		nodes       = []*AdminPermission{}
		_spec       = apq.querySpec()
		loadedTypes = [1]bool{
			apq.withRoles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AdminPermission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AdminPermission{config: apq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, apq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := apq.withRoles; query != nil {
		if err := apq.loadRoles(ctx, query, nodes,
			func(n *AdminPermission) { n.Edges.Roles = []*AdminRolePermission{} },
			func(n *AdminPermission, e *AdminRolePermission) { n.Edges.Roles = append(n.Edges.Roles, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (apq *AdminPermissionQuery) loadRoles(ctx context.Context, query *AdminRolePermissionQuery, nodes []*AdminPermission, init func(*AdminPermission), assign func(*AdminPermission, *AdminRolePermission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*AdminPermission)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AdminRolePermission(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(adminpermission.RolesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.admin_permission_roles
		if fk == nil {
			return fmt.Errorf(`foreign-key "admin_permission_roles" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "admin_permission_roles" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (apq *AdminPermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := apq.querySpec()
	_spec.Node.Columns = apq.ctx.Fields
	if len(apq.ctx.Fields) > 0 {
		_spec.Unique = apq.ctx.Unique != nil && *apq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, apq.driver, _spec)
}

func (apq *AdminPermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(adminpermission.Table, adminpermission.Columns, sqlgraph.NewFieldSpec(adminpermission.FieldID, field.TypeUUID))
	_spec.From = apq.sql
	if unique := apq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if apq.path != nil {
		_spec.Unique = true
	}
	if fields := apq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminpermission.FieldID)
		for i := range fields {
			if fields[i] != adminpermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := apq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := apq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := apq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := apq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (apq *AdminPermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(apq.driver.Dialect())
	t1 := builder.Table(adminpermission.Table)
	columns := apq.ctx.Fields
	if len(columns) == 0 {
		columns = adminpermission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if apq.sql != nil {
		selector = apq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if apq.ctx.Unique != nil && *apq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range apq.predicates {
		p(selector)
	}
	for _, p := range apq.order {
		p(selector)
	}
	if offset := apq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := apq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdminPermissionGroupBy is the group-by builder for AdminPermission entities.
type AdminPermissionGroupBy struct {
	selector
	build *AdminPermissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (apgb *AdminPermissionGroupBy) Aggregate(fns ...AggregateFunc) *AdminPermissionGroupBy {
	apgb.fns = append(apgb.fns, fns...)
	return apgb
}

// Scan applies the selector query and scans the result into the given value.
func (apgb *AdminPermissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, apgb.build.ctx, ent.OpQueryGroupBy)
	if err := apgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminPermissionQuery, *AdminPermissionGroupBy](ctx, apgb.build, apgb, apgb.build.inters, v)
}

func (apgb *AdminPermissionGroupBy) sqlScan(ctx context.Context, root *AdminPermissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(apgb.fns))
	for _, fn := range apgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*apgb.flds)+len(apgb.fns))
		for _, f := range *apgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*apgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := apgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AdminPermissionSelect is the builder for selecting fields of AdminPermission entities.
type AdminPermissionSelect struct {
	*AdminPermissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aps *AdminPermissionSelect) Aggregate(fns ...AggregateFunc) *AdminPermissionSelect {
	aps.fns = append(aps.fns, fns...)
	return aps
}

// Scan applies the selector query and scans the result into the given value.
func (aps *AdminPermissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aps.ctx, ent.OpQuerySelect)
	if err := aps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminPermissionQuery, *AdminPermissionSelect](ctx, aps.AdminPermissionQuery, aps, aps.inters, v)
}

func (aps *AdminPermissionSelect) sqlScan(ctx context.Context, root *AdminPermissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aps.fns))
	for _, fn := range aps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}