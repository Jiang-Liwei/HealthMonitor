// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthmonitor/ent/predicate"
	"healthmonitor/ent/usermealfood"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserMealFoodDelete is the builder for deleting a UserMealFood entity.
type UserMealFoodDelete struct {
	config
	hooks    []Hook
	mutation *UserMealFoodMutation
}

// Where appends a list predicates to the UserMealFoodDelete builder.
func (umfd *UserMealFoodDelete) Where(ps ...predicate.UserMealFood) *UserMealFoodDelete {
	umfd.mutation.Where(ps...)
	return umfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (umfd *UserMealFoodDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, umfd.sqlExec, umfd.mutation, umfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (umfd *UserMealFoodDelete) ExecX(ctx context.Context) int {
	n, err := umfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (umfd *UserMealFoodDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usermealfood.Table, sqlgraph.NewFieldSpec(usermealfood.FieldID, field.TypeInt))
	if ps := umfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, umfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	umfd.mutation.done = true
	return affected, err
}

// UserMealFoodDeleteOne is the builder for deleting a single UserMealFood entity.
type UserMealFoodDeleteOne struct {
	umfd *UserMealFoodDelete
}

// Where appends a list predicates to the UserMealFoodDelete builder.
func (umfdo *UserMealFoodDeleteOne) Where(ps ...predicate.UserMealFood) *UserMealFoodDeleteOne {
	umfdo.umfd.mutation.Where(ps...)
	return umfdo
}

// Exec executes the deletion query.
func (umfdo *UserMealFoodDeleteOne) Exec(ctx context.Context) error {
	n, err := umfdo.umfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usermealfood.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (umfdo *UserMealFoodDeleteOne) ExecX(ctx context.Context) {
	if err := umfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
