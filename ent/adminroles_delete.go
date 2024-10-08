// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthmonitor/ent/adminroles"
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminRolesDelete is the builder for deleting a AdminRoles entity.
type AdminRolesDelete struct {
	config
	hooks    []Hook
	mutation *AdminRolesMutation
}

// Where appends a list predicates to the AdminRolesDelete builder.
func (ard *AdminRolesDelete) Where(ps ...predicate.AdminRoles) *AdminRolesDelete {
	ard.mutation.Where(ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *AdminRolesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ard.sqlExec, ard.mutation, ard.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *AdminRolesDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *AdminRolesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(adminroles.Table, sqlgraph.NewFieldSpec(adminroles.FieldID, field.TypeUUID))
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ard.mutation.done = true
	return affected, err
}

// AdminRolesDeleteOne is the builder for deleting a single AdminRoles entity.
type AdminRolesDeleteOne struct {
	ard *AdminRolesDelete
}

// Where appends a list predicates to the AdminRolesDelete builder.
func (ardo *AdminRolesDeleteOne) Where(ps ...predicate.AdminRoles) *AdminRolesDeleteOne {
	ardo.ard.mutation.Where(ps...)
	return ardo
}

// Exec executes the deletion query.
func (ardo *AdminRolesDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adminroles.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *AdminRolesDeleteOne) ExecX(ctx context.Context) {
	if err := ardo.Exec(ctx); err != nil {
		panic(err)
	}
}
