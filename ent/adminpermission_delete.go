// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/adminpermission"
	"HealthMonitor/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminPermissionDelete is the builder for deleting a AdminPermission entity.
type AdminPermissionDelete struct {
	config
	hooks    []Hook
	mutation *AdminPermissionMutation
}

// Where appends a list predicates to the AdminPermissionDelete builder.
func (apd *AdminPermissionDelete) Where(ps ...predicate.AdminPermission) *AdminPermissionDelete {
	apd.mutation.Where(ps...)
	return apd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (apd *AdminPermissionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, apd.sqlExec, apd.mutation, apd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (apd *AdminPermissionDelete) ExecX(ctx context.Context) int {
	n, err := apd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (apd *AdminPermissionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(adminpermission.Table, sqlgraph.NewFieldSpec(adminpermission.FieldID, field.TypeUUID))
	if ps := apd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, apd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	apd.mutation.done = true
	return affected, err
}

// AdminPermissionDeleteOne is the builder for deleting a single AdminPermission entity.
type AdminPermissionDeleteOne struct {
	apd *AdminPermissionDelete
}

// Where appends a list predicates to the AdminPermissionDelete builder.
func (apdo *AdminPermissionDeleteOne) Where(ps ...predicate.AdminPermission) *AdminPermissionDeleteOne {
	apdo.apd.mutation.Where(ps...)
	return apdo
}

// Exec executes the deletion query.
func (apdo *AdminPermissionDeleteOne) Exec(ctx context.Context) error {
	n, err := apdo.apd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adminpermission.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (apdo *AdminPermissionDeleteOne) ExecX(ctx context.Context) {
	if err := apdo.Exec(ctx); err != nil {
		panic(err)
	}
}