// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthmonitor/ent/foodingredients"
	"healthmonitor/ent/ingredients"
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IngredientsUpdate is the builder for updating Ingredients entities.
type IngredientsUpdate struct {
	config
	hooks    []Hook
	mutation *IngredientsMutation
}

// Where appends a list predicates to the IngredientsUpdate builder.
func (iu *IngredientsUpdate) Where(ps ...predicate.Ingredients) *IngredientsUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *IngredientsUpdate) SetName(s string) *IngredientsUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iu *IngredientsUpdate) SetNillableName(s *string) *IngredientsUpdate {
	if s != nil {
		iu.SetName(*s)
	}
	return iu
}

// SetEffect sets the "effect" field.
func (iu *IngredientsUpdate) SetEffect(i ingredients.Effect) *IngredientsUpdate {
	iu.mutation.SetEffect(i)
	return iu
}

// SetNillableEffect sets the "effect" field if the given value is not nil.
func (iu *IngredientsUpdate) SetNillableEffect(i *ingredients.Effect) *IngredientsUpdate {
	if i != nil {
		iu.SetEffect(*i)
	}
	return iu
}

// AddFoodIDs adds the "food" edge to the FoodIngredients entity by IDs.
func (iu *IngredientsUpdate) AddFoodIDs(ids ...int) *IngredientsUpdate {
	iu.mutation.AddFoodIDs(ids...)
	return iu
}

// AddFood adds the "food" edges to the FoodIngredients entity.
func (iu *IngredientsUpdate) AddFood(f ...*FoodIngredients) *IngredientsUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return iu.AddFoodIDs(ids...)
}

// Mutation returns the IngredientsMutation object of the builder.
func (iu *IngredientsUpdate) Mutation() *IngredientsMutation {
	return iu.mutation
}

// ClearFood clears all "food" edges to the FoodIngredients entity.
func (iu *IngredientsUpdate) ClearFood() *IngredientsUpdate {
	iu.mutation.ClearFood()
	return iu
}

// RemoveFoodIDs removes the "food" edge to FoodIngredients entities by IDs.
func (iu *IngredientsUpdate) RemoveFoodIDs(ids ...int) *IngredientsUpdate {
	iu.mutation.RemoveFoodIDs(ids...)
	return iu
}

// RemoveFood removes "food" edges to FoodIngredients entities.
func (iu *IngredientsUpdate) RemoveFood(f ...*FoodIngredients) *IngredientsUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return iu.RemoveFoodIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IngredientsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IngredientsUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IngredientsUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IngredientsUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IngredientsUpdate) check() error {
	if v, ok := iu.mutation.Effect(); ok {
		if err := ingredients.EffectValidator(v); err != nil {
			return &ValidationError{Name: "effect", err: fmt.Errorf(`ent: validator failed for field "Ingredients.effect": %w`, err)}
		}
	}
	return nil
}

func (iu *IngredientsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(ingredients.Table, ingredients.Columns, sqlgraph.NewFieldSpec(ingredients.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(ingredients.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Effect(); ok {
		_spec.SetField(ingredients.FieldEffect, field.TypeEnum, value)
	}
	if iu.mutation.FoodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingredients.FoodTable,
			Columns: []string{ingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedFoodIDs(); len(nodes) > 0 && !iu.mutation.FoodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingredients.FoodTable,
			Columns: []string{ingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.FoodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingredients.FoodTable,
			Columns: []string{ingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ingredients.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IngredientsUpdateOne is the builder for updating a single Ingredients entity.
type IngredientsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IngredientsMutation
}

// SetName sets the "name" field.
func (iuo *IngredientsUpdateOne) SetName(s string) *IngredientsUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iuo *IngredientsUpdateOne) SetNillableName(s *string) *IngredientsUpdateOne {
	if s != nil {
		iuo.SetName(*s)
	}
	return iuo
}

// SetEffect sets the "effect" field.
func (iuo *IngredientsUpdateOne) SetEffect(i ingredients.Effect) *IngredientsUpdateOne {
	iuo.mutation.SetEffect(i)
	return iuo
}

// SetNillableEffect sets the "effect" field if the given value is not nil.
func (iuo *IngredientsUpdateOne) SetNillableEffect(i *ingredients.Effect) *IngredientsUpdateOne {
	if i != nil {
		iuo.SetEffect(*i)
	}
	return iuo
}

// AddFoodIDs adds the "food" edge to the FoodIngredients entity by IDs.
func (iuo *IngredientsUpdateOne) AddFoodIDs(ids ...int) *IngredientsUpdateOne {
	iuo.mutation.AddFoodIDs(ids...)
	return iuo
}

// AddFood adds the "food" edges to the FoodIngredients entity.
func (iuo *IngredientsUpdateOne) AddFood(f ...*FoodIngredients) *IngredientsUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return iuo.AddFoodIDs(ids...)
}

// Mutation returns the IngredientsMutation object of the builder.
func (iuo *IngredientsUpdateOne) Mutation() *IngredientsMutation {
	return iuo.mutation
}

// ClearFood clears all "food" edges to the FoodIngredients entity.
func (iuo *IngredientsUpdateOne) ClearFood() *IngredientsUpdateOne {
	iuo.mutation.ClearFood()
	return iuo
}

// RemoveFoodIDs removes the "food" edge to FoodIngredients entities by IDs.
func (iuo *IngredientsUpdateOne) RemoveFoodIDs(ids ...int) *IngredientsUpdateOne {
	iuo.mutation.RemoveFoodIDs(ids...)
	return iuo
}

// RemoveFood removes "food" edges to FoodIngredients entities.
func (iuo *IngredientsUpdateOne) RemoveFood(f ...*FoodIngredients) *IngredientsUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return iuo.RemoveFoodIDs(ids...)
}

// Where appends a list predicates to the IngredientsUpdate builder.
func (iuo *IngredientsUpdateOne) Where(ps ...predicate.Ingredients) *IngredientsUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IngredientsUpdateOne) Select(field string, fields ...string) *IngredientsUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Ingredients entity.
func (iuo *IngredientsUpdateOne) Save(ctx context.Context) (*Ingredients, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IngredientsUpdateOne) SaveX(ctx context.Context) *Ingredients {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IngredientsUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IngredientsUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IngredientsUpdateOne) check() error {
	if v, ok := iuo.mutation.Effect(); ok {
		if err := ingredients.EffectValidator(v); err != nil {
			return &ValidationError{Name: "effect", err: fmt.Errorf(`ent: validator failed for field "Ingredients.effect": %w`, err)}
		}
	}
	return nil
}

func (iuo *IngredientsUpdateOne) sqlSave(ctx context.Context) (_node *Ingredients, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(ingredients.Table, ingredients.Columns, sqlgraph.NewFieldSpec(ingredients.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ingredients.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ingredients.FieldID)
		for _, f := range fields {
			if !ingredients.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ingredients.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(ingredients.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Effect(); ok {
		_spec.SetField(ingredients.FieldEffect, field.TypeEnum, value)
	}
	if iuo.mutation.FoodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingredients.FoodTable,
			Columns: []string{ingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedFoodIDs(); len(nodes) > 0 && !iuo.mutation.FoodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingredients.FoodTable,
			Columns: []string{ingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.FoodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ingredients.FoodTable,
			Columns: []string{ingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ingredients{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ingredients.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
