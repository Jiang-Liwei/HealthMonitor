// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/food"
	"HealthMonitor/ent/foodingredients"
	"HealthMonitor/ent/ingredients"
	"HealthMonitor/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FoodIngredientsQuery is the builder for querying FoodIngredients entities.
type FoodIngredientsQuery struct {
	config
	ctx            *QueryContext
	order          []foodingredients.OrderOption
	inters         []Interceptor
	predicates     []predicate.FoodIngredients
	withFood       *FoodQuery
	withIngredient *IngredientsQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FoodIngredientsQuery builder.
func (fiq *FoodIngredientsQuery) Where(ps ...predicate.FoodIngredients) *FoodIngredientsQuery {
	fiq.predicates = append(fiq.predicates, ps...)
	return fiq
}

// Limit the number of records to be returned by this query.
func (fiq *FoodIngredientsQuery) Limit(limit int) *FoodIngredientsQuery {
	fiq.ctx.Limit = &limit
	return fiq
}

// Offset to start from.
func (fiq *FoodIngredientsQuery) Offset(offset int) *FoodIngredientsQuery {
	fiq.ctx.Offset = &offset
	return fiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fiq *FoodIngredientsQuery) Unique(unique bool) *FoodIngredientsQuery {
	fiq.ctx.Unique = &unique
	return fiq
}

// Order specifies how the records should be ordered.
func (fiq *FoodIngredientsQuery) Order(o ...foodingredients.OrderOption) *FoodIngredientsQuery {
	fiq.order = append(fiq.order, o...)
	return fiq
}

// QueryFood chains the current query on the "food" edge.
func (fiq *FoodIngredientsQuery) QueryFood() *FoodQuery {
	query := (&FoodClient{config: fiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(foodingredients.Table, foodingredients.FieldID, selector),
			sqlgraph.To(food.Table, food.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, foodingredients.FoodTable, foodingredients.FoodColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIngredient chains the current query on the "ingredient" edge.
func (fiq *FoodIngredientsQuery) QueryIngredient() *IngredientsQuery {
	query := (&IngredientsClient{config: fiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(foodingredients.Table, foodingredients.FieldID, selector),
			sqlgraph.To(ingredients.Table, ingredients.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, foodingredients.IngredientTable, foodingredients.IngredientColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FoodIngredients entity from the query.
// Returns a *NotFoundError when no FoodIngredients was found.
func (fiq *FoodIngredientsQuery) First(ctx context.Context) (*FoodIngredients, error) {
	nodes, err := fiq.Limit(1).All(setContextOp(ctx, fiq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{foodingredients.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fiq *FoodIngredientsQuery) FirstX(ctx context.Context) *FoodIngredients {
	node, err := fiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FoodIngredients ID from the query.
// Returns a *NotFoundError when no FoodIngredients ID was found.
func (fiq *FoodIngredientsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fiq.Limit(1).IDs(setContextOp(ctx, fiq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{foodingredients.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fiq *FoodIngredientsQuery) FirstIDX(ctx context.Context) int {
	id, err := fiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FoodIngredients entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FoodIngredients entity is found.
// Returns a *NotFoundError when no FoodIngredients entities are found.
func (fiq *FoodIngredientsQuery) Only(ctx context.Context) (*FoodIngredients, error) {
	nodes, err := fiq.Limit(2).All(setContextOp(ctx, fiq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{foodingredients.Label}
	default:
		return nil, &NotSingularError{foodingredients.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fiq *FoodIngredientsQuery) OnlyX(ctx context.Context) *FoodIngredients {
	node, err := fiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FoodIngredients ID in the query.
// Returns a *NotSingularError when more than one FoodIngredients ID is found.
// Returns a *NotFoundError when no entities are found.
func (fiq *FoodIngredientsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fiq.Limit(2).IDs(setContextOp(ctx, fiq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{foodingredients.Label}
	default:
		err = &NotSingularError{foodingredients.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fiq *FoodIngredientsQuery) OnlyIDX(ctx context.Context) int {
	id, err := fiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FoodIngredientsSlice.
func (fiq *FoodIngredientsQuery) All(ctx context.Context) ([]*FoodIngredients, error) {
	ctx = setContextOp(ctx, fiq.ctx, ent.OpQueryAll)
	if err := fiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FoodIngredients, *FoodIngredientsQuery]()
	return withInterceptors[[]*FoodIngredients](ctx, fiq, qr, fiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fiq *FoodIngredientsQuery) AllX(ctx context.Context) []*FoodIngredients {
	nodes, err := fiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FoodIngredients IDs.
func (fiq *FoodIngredientsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fiq.ctx.Unique == nil && fiq.path != nil {
		fiq.Unique(true)
	}
	ctx = setContextOp(ctx, fiq.ctx, ent.OpQueryIDs)
	if err = fiq.Select(foodingredients.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fiq *FoodIngredientsQuery) IDsX(ctx context.Context) []int {
	ids, err := fiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fiq *FoodIngredientsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fiq.ctx, ent.OpQueryCount)
	if err := fiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fiq, querierCount[*FoodIngredientsQuery](), fiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fiq *FoodIngredientsQuery) CountX(ctx context.Context) int {
	count, err := fiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fiq *FoodIngredientsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fiq.ctx, ent.OpQueryExist)
	switch _, err := fiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fiq *FoodIngredientsQuery) ExistX(ctx context.Context) bool {
	exist, err := fiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FoodIngredientsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fiq *FoodIngredientsQuery) Clone() *FoodIngredientsQuery {
	if fiq == nil {
		return nil
	}
	return &FoodIngredientsQuery{
		config:         fiq.config,
		ctx:            fiq.ctx.Clone(),
		order:          append([]foodingredients.OrderOption{}, fiq.order...),
		inters:         append([]Interceptor{}, fiq.inters...),
		predicates:     append([]predicate.FoodIngredients{}, fiq.predicates...),
		withFood:       fiq.withFood.Clone(),
		withIngredient: fiq.withIngredient.Clone(),
		// clone intermediate query.
		sql:  fiq.sql.Clone(),
		path: fiq.path,
	}
}

// WithFood tells the query-builder to eager-load the nodes that are connected to
// the "food" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FoodIngredientsQuery) WithFood(opts ...func(*FoodQuery)) *FoodIngredientsQuery {
	query := (&FoodClient{config: fiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFood = query
	return fiq
}

// WithIngredient tells the query-builder to eager-load the nodes that are connected to
// the "ingredient" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FoodIngredientsQuery) WithIngredient(opts ...func(*IngredientsQuery)) *FoodIngredientsQuery {
	query := (&IngredientsClient{config: fiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fiq.withIngredient = query
	return fiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (fiq *FoodIngredientsQuery) GroupBy(field string, fields ...string) *FoodIngredientsGroupBy {
	fiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FoodIngredientsGroupBy{build: fiq}
	grbuild.flds = &fiq.ctx.Fields
	grbuild.label = foodingredients.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (fiq *FoodIngredientsQuery) Select(fields ...string) *FoodIngredientsSelect {
	fiq.ctx.Fields = append(fiq.ctx.Fields, fields...)
	sbuild := &FoodIngredientsSelect{FoodIngredientsQuery: fiq}
	sbuild.label = foodingredients.Label
	sbuild.flds, sbuild.scan = &fiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FoodIngredientsSelect configured with the given aggregations.
func (fiq *FoodIngredientsQuery) Aggregate(fns ...AggregateFunc) *FoodIngredientsSelect {
	return fiq.Select().Aggregate(fns...)
}

func (fiq *FoodIngredientsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fiq); err != nil {
				return err
			}
		}
	}
	for _, f := range fiq.ctx.Fields {
		if !foodingredients.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fiq.path != nil {
		prev, err := fiq.path(ctx)
		if err != nil {
			return err
		}
		fiq.sql = prev
	}
	return nil
}

func (fiq *FoodIngredientsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FoodIngredients, error) {
	var (
		nodes       = []*FoodIngredients{}
		withFKs     = fiq.withFKs
		_spec       = fiq.querySpec()
		loadedTypes = [2]bool{
			fiq.withFood != nil,
			fiq.withIngredient != nil,
		}
	)
	if fiq.withFood != nil || fiq.withIngredient != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, foodingredients.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FoodIngredients).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FoodIngredients{config: fiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fiq.withFood; query != nil {
		if err := fiq.loadFood(ctx, query, nodes, nil,
			func(n *FoodIngredients, e *Food) { n.Edges.Food = e }); err != nil {
			return nil, err
		}
	}
	if query := fiq.withIngredient; query != nil {
		if err := fiq.loadIngredient(ctx, query, nodes, nil,
			func(n *FoodIngredients, e *Ingredients) { n.Edges.Ingredient = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fiq *FoodIngredientsQuery) loadFood(ctx context.Context, query *FoodQuery, nodes []*FoodIngredients, init func(*FoodIngredients), assign func(*FoodIngredients, *Food)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*FoodIngredients)
	for i := range nodes {
		if nodes[i].food_ingredients == nil {
			continue
		}
		fk := *nodes[i].food_ingredients
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(food.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "food_ingredients" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fiq *FoodIngredientsQuery) loadIngredient(ctx context.Context, query *IngredientsQuery, nodes []*FoodIngredients, init func(*FoodIngredients), assign func(*FoodIngredients, *Ingredients)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*FoodIngredients)
	for i := range nodes {
		if nodes[i].ingredients_food == nil {
			continue
		}
		fk := *nodes[i].ingredients_food
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(ingredients.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ingredients_food" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fiq *FoodIngredientsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fiq.querySpec()
	_spec.Node.Columns = fiq.ctx.Fields
	if len(fiq.ctx.Fields) > 0 {
		_spec.Unique = fiq.ctx.Unique != nil && *fiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fiq.driver, _spec)
}

func (fiq *FoodIngredientsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(foodingredients.Table, foodingredients.Columns, sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt))
	_spec.From = fiq.sql
	if unique := fiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fiq.path != nil {
		_spec.Unique = true
	}
	if fields := fiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, foodingredients.FieldID)
		for i := range fields {
			if fields[i] != foodingredients.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fiq *FoodIngredientsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fiq.driver.Dialect())
	t1 := builder.Table(foodingredients.Table)
	columns := fiq.ctx.Fields
	if len(columns) == 0 {
		columns = foodingredients.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fiq.sql != nil {
		selector = fiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fiq.ctx.Unique != nil && *fiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fiq.predicates {
		p(selector)
	}
	for _, p := range fiq.order {
		p(selector)
	}
	if offset := fiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FoodIngredientsGroupBy is the group-by builder for FoodIngredients entities.
type FoodIngredientsGroupBy struct {
	selector
	build *FoodIngredientsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (figb *FoodIngredientsGroupBy) Aggregate(fns ...AggregateFunc) *FoodIngredientsGroupBy {
	figb.fns = append(figb.fns, fns...)
	return figb
}

// Scan applies the selector query and scans the result into the given value.
func (figb *FoodIngredientsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, figb.build.ctx, ent.OpQueryGroupBy)
	if err := figb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FoodIngredientsQuery, *FoodIngredientsGroupBy](ctx, figb.build, figb, figb.build.inters, v)
}

func (figb *FoodIngredientsGroupBy) sqlScan(ctx context.Context, root *FoodIngredientsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(figb.fns))
	for _, fn := range figb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*figb.flds)+len(figb.fns))
		for _, f := range *figb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*figb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := figb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FoodIngredientsSelect is the builder for selecting fields of FoodIngredients entities.
type FoodIngredientsSelect struct {
	*FoodIngredientsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fis *FoodIngredientsSelect) Aggregate(fns ...AggregateFunc) *FoodIngredientsSelect {
	fis.fns = append(fis.fns, fns...)
	return fis
}

// Scan applies the selector query and scans the result into the given value.
func (fis *FoodIngredientsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fis.ctx, ent.OpQuerySelect)
	if err := fis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FoodIngredientsQuery, *FoodIngredientsSelect](ctx, fis.FoodIngredientsQuery, fis, fis.inters, v)
}

func (fis *FoodIngredientsSelect) sqlScan(ctx context.Context, root *FoodIngredientsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fis.fns))
	for _, fn := range fis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
