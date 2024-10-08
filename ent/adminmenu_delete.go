// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthmonitor/ent/adminmenu"
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminMenuDelete is the builder for deleting a AdminMenu entity.
type AdminMenuDelete struct {
	config
	hooks    []Hook
	mutation *AdminMenuMutation
}

// Where appends a list predicates to the AdminMenuDelete builder.
func (amd *AdminMenuDelete) Where(ps ...predicate.AdminMenu) *AdminMenuDelete {
	amd.mutation.Where(ps...)
	return amd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (amd *AdminMenuDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, amd.sqlExec, amd.mutation, amd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (amd *AdminMenuDelete) ExecX(ctx context.Context) int {
	n, err := amd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (amd *AdminMenuDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(adminmenu.Table, sqlgraph.NewFieldSpec(adminmenu.FieldID, field.TypeUUID))
	if ps := amd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, amd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	amd.mutation.done = true
	return affected, err
}

// AdminMenuDeleteOne is the builder for deleting a single AdminMenu entity.
type AdminMenuDeleteOne struct {
	amd *AdminMenuDelete
}

// Where appends a list predicates to the AdminMenuDelete builder.
func (amdo *AdminMenuDeleteOne) Where(ps ...predicate.AdminMenu) *AdminMenuDeleteOne {
	amdo.amd.mutation.Where(ps...)
	return amdo
}

// Exec executes the deletion query.
func (amdo *AdminMenuDeleteOne) Exec(ctx context.Context) error {
	n, err := amdo.amd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adminmenu.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (amdo *AdminMenuDeleteOne) ExecX(ctx context.Context) {
	if err := amdo.Exec(ctx); err != nil {
		panic(err)
	}
}
