// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthmonitor/ent/food"
	"healthmonitor/ent/foodingredients"
	"healthmonitor/ent/ingredients"
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FoodIngredientsUpdate is the builder for updating FoodIngredients entities.
type FoodIngredientsUpdate struct {
	config
	hooks    []Hook
	mutation *FoodIngredientsMutation
}

// Where appends a list predicates to the FoodIngredientsUpdate builder.
func (fiu *FoodIngredientsUpdate) Where(ps ...predicate.FoodIngredients) *FoodIngredientsUpdate {
	fiu.mutation.Where(ps...)
	return fiu
}

// SetFoodID sets the "food" edge to the Food entity by ID.
func (fiu *FoodIngredientsUpdate) SetFoodID(id uuid.UUID) *FoodIngredientsUpdate {
	fiu.mutation.SetFoodID(id)
	return fiu
}

// SetNillableFoodID sets the "food" edge to the Food entity by ID if the given value is not nil.
func (fiu *FoodIngredientsUpdate) SetNillableFoodID(id *uuid.UUID) *FoodIngredientsUpdate {
	if id != nil {
		fiu = fiu.SetFoodID(*id)
	}
	return fiu
}

// SetFood sets the "food" edge to the Food entity.
func (fiu *FoodIngredientsUpdate) SetFood(f *Food) *FoodIngredientsUpdate {
	return fiu.SetFoodID(f.ID)
}

// SetIngredientID sets the "ingredient" edge to the Ingredients entity by ID.
func (fiu *FoodIngredientsUpdate) SetIngredientID(id uuid.UUID) *FoodIngredientsUpdate {
	fiu.mutation.SetIngredientID(id)
	return fiu
}

// SetNillableIngredientID sets the "ingredient" edge to the Ingredients entity by ID if the given value is not nil.
func (fiu *FoodIngredientsUpdate) SetNillableIngredientID(id *uuid.UUID) *FoodIngredientsUpdate {
	if id != nil {
		fiu = fiu.SetIngredientID(*id)
	}
	return fiu
}

// SetIngredient sets the "ingredient" edge to the Ingredients entity.
func (fiu *FoodIngredientsUpdate) SetIngredient(i *Ingredients) *FoodIngredientsUpdate {
	return fiu.SetIngredientID(i.ID)
}

// Mutation returns the FoodIngredientsMutation object of the builder.
func (fiu *FoodIngredientsUpdate) Mutation() *FoodIngredientsMutation {
	return fiu.mutation
}

// ClearFood clears the "food" edge to the Food entity.
func (fiu *FoodIngredientsUpdate) ClearFood() *FoodIngredientsUpdate {
	fiu.mutation.ClearFood()
	return fiu
}

// ClearIngredient clears the "ingredient" edge to the Ingredients entity.
func (fiu *FoodIngredientsUpdate) ClearIngredient() *FoodIngredientsUpdate {
	fiu.mutation.ClearIngredient()
	return fiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fiu *FoodIngredientsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fiu.sqlSave, fiu.mutation, fiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiu *FoodIngredientsUpdate) SaveX(ctx context.Context) int {
	affected, err := fiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fiu *FoodIngredientsUpdate) Exec(ctx context.Context) error {
	_, err := fiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiu *FoodIngredientsUpdate) ExecX(ctx context.Context) {
	if err := fiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fiu *FoodIngredientsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(foodingredients.Table, foodingredients.Columns, sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt))
	if ps := fiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fiu.mutation.FoodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodingredients.FoodTable,
			Columns: []string{foodingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiu.mutation.FoodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodingredients.FoodTable,
			Columns: []string{foodingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fiu.mutation.IngredientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodingredients.IngredientTable,
			Columns: []string{foodingredients.IngredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ingredients.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiu.mutation.IngredientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodingredients.IngredientTable,
			Columns: []string{foodingredients.IngredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ingredients.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{foodingredients.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fiu.mutation.done = true
	return n, nil
}

// FoodIngredientsUpdateOne is the builder for updating a single FoodIngredients entity.
type FoodIngredientsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FoodIngredientsMutation
}

// SetFoodID sets the "food" edge to the Food entity by ID.
func (fiuo *FoodIngredientsUpdateOne) SetFoodID(id uuid.UUID) *FoodIngredientsUpdateOne {
	fiuo.mutation.SetFoodID(id)
	return fiuo
}

// SetNillableFoodID sets the "food" edge to the Food entity by ID if the given value is not nil.
func (fiuo *FoodIngredientsUpdateOne) SetNillableFoodID(id *uuid.UUID) *FoodIngredientsUpdateOne {
	if id != nil {
		fiuo = fiuo.SetFoodID(*id)
	}
	return fiuo
}

// SetFood sets the "food" edge to the Food entity.
func (fiuo *FoodIngredientsUpdateOne) SetFood(f *Food) *FoodIngredientsUpdateOne {
	return fiuo.SetFoodID(f.ID)
}

// SetIngredientID sets the "ingredient" edge to the Ingredients entity by ID.
func (fiuo *FoodIngredientsUpdateOne) SetIngredientID(id uuid.UUID) *FoodIngredientsUpdateOne {
	fiuo.mutation.SetIngredientID(id)
	return fiuo
}

// SetNillableIngredientID sets the "ingredient" edge to the Ingredients entity by ID if the given value is not nil.
func (fiuo *FoodIngredientsUpdateOne) SetNillableIngredientID(id *uuid.UUID) *FoodIngredientsUpdateOne {
	if id != nil {
		fiuo = fiuo.SetIngredientID(*id)
	}
	return fiuo
}

// SetIngredient sets the "ingredient" edge to the Ingredients entity.
func (fiuo *FoodIngredientsUpdateOne) SetIngredient(i *Ingredients) *FoodIngredientsUpdateOne {
	return fiuo.SetIngredientID(i.ID)
}

// Mutation returns the FoodIngredientsMutation object of the builder.
func (fiuo *FoodIngredientsUpdateOne) Mutation() *FoodIngredientsMutation {
	return fiuo.mutation
}

// ClearFood clears the "food" edge to the Food entity.
func (fiuo *FoodIngredientsUpdateOne) ClearFood() *FoodIngredientsUpdateOne {
	fiuo.mutation.ClearFood()
	return fiuo
}

// ClearIngredient clears the "ingredient" edge to the Ingredients entity.
func (fiuo *FoodIngredientsUpdateOne) ClearIngredient() *FoodIngredientsUpdateOne {
	fiuo.mutation.ClearIngredient()
	return fiuo
}

// Where appends a list predicates to the FoodIngredientsUpdate builder.
func (fiuo *FoodIngredientsUpdateOne) Where(ps ...predicate.FoodIngredients) *FoodIngredientsUpdateOne {
	fiuo.mutation.Where(ps...)
	return fiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fiuo *FoodIngredientsUpdateOne) Select(field string, fields ...string) *FoodIngredientsUpdateOne {
	fiuo.fields = append([]string{field}, fields...)
	return fiuo
}

// Save executes the query and returns the updated FoodIngredients entity.
func (fiuo *FoodIngredientsUpdateOne) Save(ctx context.Context) (*FoodIngredients, error) {
	return withHooks(ctx, fiuo.sqlSave, fiuo.mutation, fiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiuo *FoodIngredientsUpdateOne) SaveX(ctx context.Context) *FoodIngredients {
	node, err := fiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fiuo *FoodIngredientsUpdateOne) Exec(ctx context.Context) error {
	_, err := fiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiuo *FoodIngredientsUpdateOne) ExecX(ctx context.Context) {
	if err := fiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fiuo *FoodIngredientsUpdateOne) sqlSave(ctx context.Context) (_node *FoodIngredients, err error) {
	_spec := sqlgraph.NewUpdateSpec(foodingredients.Table, foodingredients.Columns, sqlgraph.NewFieldSpec(foodingredients.FieldID, field.TypeInt))
	id, ok := fiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FoodIngredients.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, foodingredients.FieldID)
		for _, f := range fields {
			if !foodingredients.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != foodingredients.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fiuo.mutation.FoodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodingredients.FoodTable,
			Columns: []string{foodingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiuo.mutation.FoodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodingredients.FoodTable,
			Columns: []string{foodingredients.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fiuo.mutation.IngredientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodingredients.IngredientTable,
			Columns: []string{foodingredients.IngredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ingredients.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiuo.mutation.IngredientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodingredients.IngredientTable,
			Columns: []string{foodingredients.IngredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ingredients.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FoodIngredients{config: fiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{foodingredients.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fiuo.mutation.done = true
	return _node, nil
}
