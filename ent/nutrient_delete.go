// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthmonitor/ent/nutrient"
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NutrientDelete is the builder for deleting a Nutrient entity.
type NutrientDelete struct {
	config
	hooks    []Hook
	mutation *NutrientMutation
}

// Where appends a list predicates to the NutrientDelete builder.
func (nd *NutrientDelete) Where(ps ...predicate.Nutrient) *NutrientDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NutrientDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nd.sqlExec, nd.mutation, nd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NutrientDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NutrientDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(nutrient.Table, sqlgraph.NewFieldSpec(nutrient.FieldID, field.TypeUUID))
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nd.mutation.done = true
	return affected, err
}

// NutrientDeleteOne is the builder for deleting a single Nutrient entity.
type NutrientDeleteOne struct {
	nd *NutrientDelete
}

// Where appends a list predicates to the NutrientDelete builder.
func (ndo *NutrientDeleteOne) Where(ps ...predicate.Nutrient) *NutrientDeleteOne {
	ndo.nd.mutation.Where(ps...)
	return ndo
}

// Exec executes the deletion query.
func (ndo *NutrientDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nutrient.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NutrientDeleteOne) ExecX(ctx context.Context) {
	if err := ndo.Exec(ctx); err != nil {
		panic(err)
	}
}
