// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/foodnutrients"
	"HealthMonitor/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FoodNutrientsDelete is the builder for deleting a FoodNutrients entity.
type FoodNutrientsDelete struct {
	config
	hooks    []Hook
	mutation *FoodNutrientsMutation
}

// Where appends a list predicates to the FoodNutrientsDelete builder.
func (fnd *FoodNutrientsDelete) Where(ps ...predicate.FoodNutrients) *FoodNutrientsDelete {
	fnd.mutation.Where(ps...)
	return fnd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fnd *FoodNutrientsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fnd.sqlExec, fnd.mutation, fnd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fnd *FoodNutrientsDelete) ExecX(ctx context.Context) int {
	n, err := fnd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fnd *FoodNutrientsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(foodnutrients.Table, sqlgraph.NewFieldSpec(foodnutrients.FieldID, field.TypeInt))
	if ps := fnd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fnd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fnd.mutation.done = true
	return affected, err
}

// FoodNutrientsDeleteOne is the builder for deleting a single FoodNutrients entity.
type FoodNutrientsDeleteOne struct {
	fnd *FoodNutrientsDelete
}

// Where appends a list predicates to the FoodNutrientsDelete builder.
func (fndo *FoodNutrientsDeleteOne) Where(ps ...predicate.FoodNutrients) *FoodNutrientsDeleteOne {
	fndo.fnd.mutation.Where(ps...)
	return fndo
}

// Exec executes the deletion query.
func (fndo *FoodNutrientsDeleteOne) Exec(ctx context.Context) error {
	n, err := fndo.fnd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{foodnutrients.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fndo *FoodNutrientsDeleteOne) ExecX(ctx context.Context) {
	if err := fndo.Exec(ctx); err != nil {
		panic(err)
	}
}