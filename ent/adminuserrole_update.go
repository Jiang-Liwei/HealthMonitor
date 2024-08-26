// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/adminroles"
	"HealthMonitor/ent/adminuser"
	"HealthMonitor/ent/adminuserrole"
	"HealthMonitor/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AdminUserRoleUpdate is the builder for updating AdminUserRole entities.
type AdminUserRoleUpdate struct {
	config
	hooks    []Hook
	mutation *AdminUserRoleMutation
}

// Where appends a list predicates to the AdminUserRoleUpdate builder.
func (auru *AdminUserRoleUpdate) Where(ps ...predicate.AdminUserRole) *AdminUserRoleUpdate {
	auru.mutation.Where(ps...)
	return auru
}

// SetUserID sets the "user" edge to the AdminUser entity by ID.
func (auru *AdminUserRoleUpdate) SetUserID(id uuid.UUID) *AdminUserRoleUpdate {
	auru.mutation.SetUserID(id)
	return auru
}

// SetNillableUserID sets the "user" edge to the AdminUser entity by ID if the given value is not nil.
func (auru *AdminUserRoleUpdate) SetNillableUserID(id *uuid.UUID) *AdminUserRoleUpdate {
	if id != nil {
		auru = auru.SetUserID(*id)
	}
	return auru
}

// SetUser sets the "user" edge to the AdminUser entity.
func (auru *AdminUserRoleUpdate) SetUser(a *AdminUser) *AdminUserRoleUpdate {
	return auru.SetUserID(a.ID)
}

// SetRoleID sets the "role" edge to the AdminRoles entity by ID.
func (auru *AdminUserRoleUpdate) SetRoleID(id uuid.UUID) *AdminUserRoleUpdate {
	auru.mutation.SetRoleID(id)
	return auru
}

// SetNillableRoleID sets the "role" edge to the AdminRoles entity by ID if the given value is not nil.
func (auru *AdminUserRoleUpdate) SetNillableRoleID(id *uuid.UUID) *AdminUserRoleUpdate {
	if id != nil {
		auru = auru.SetRoleID(*id)
	}
	return auru
}

// SetRole sets the "role" edge to the AdminRoles entity.
func (auru *AdminUserRoleUpdate) SetRole(a *AdminRoles) *AdminUserRoleUpdate {
	return auru.SetRoleID(a.ID)
}

// Mutation returns the AdminUserRoleMutation object of the builder.
func (auru *AdminUserRoleUpdate) Mutation() *AdminUserRoleMutation {
	return auru.mutation
}

// ClearUser clears the "user" edge to the AdminUser entity.
func (auru *AdminUserRoleUpdate) ClearUser() *AdminUserRoleUpdate {
	auru.mutation.ClearUser()
	return auru
}

// ClearRole clears the "role" edge to the AdminRoles entity.
func (auru *AdminUserRoleUpdate) ClearRole() *AdminUserRoleUpdate {
	auru.mutation.ClearRole()
	return auru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (auru *AdminUserRoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, auru.sqlSave, auru.mutation, auru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auru *AdminUserRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := auru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auru *AdminUserRoleUpdate) Exec(ctx context.Context) error {
	_, err := auru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auru *AdminUserRoleUpdate) ExecX(ctx context.Context) {
	if err := auru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auru *AdminUserRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(adminuserrole.Table, adminuserrole.Columns, sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt))
	if ps := auru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if auru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminuserrole.UserTable,
			Columns: []string{adminuserrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminuserrole.UserTable,
			Columns: []string{adminuserrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auru.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminuserrole.RoleTable,
			Columns: []string{adminuserrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminroles.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auru.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminuserrole.RoleTable,
			Columns: []string{adminuserrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminroles.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, auru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminuserrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	auru.mutation.done = true
	return n, nil
}

// AdminUserRoleUpdateOne is the builder for updating a single AdminUserRole entity.
type AdminUserRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdminUserRoleMutation
}

// SetUserID sets the "user" edge to the AdminUser entity by ID.
func (auruo *AdminUserRoleUpdateOne) SetUserID(id uuid.UUID) *AdminUserRoleUpdateOne {
	auruo.mutation.SetUserID(id)
	return auruo
}

// SetNillableUserID sets the "user" edge to the AdminUser entity by ID if the given value is not nil.
func (auruo *AdminUserRoleUpdateOne) SetNillableUserID(id *uuid.UUID) *AdminUserRoleUpdateOne {
	if id != nil {
		auruo = auruo.SetUserID(*id)
	}
	return auruo
}

// SetUser sets the "user" edge to the AdminUser entity.
func (auruo *AdminUserRoleUpdateOne) SetUser(a *AdminUser) *AdminUserRoleUpdateOne {
	return auruo.SetUserID(a.ID)
}

// SetRoleID sets the "role" edge to the AdminRoles entity by ID.
func (auruo *AdminUserRoleUpdateOne) SetRoleID(id uuid.UUID) *AdminUserRoleUpdateOne {
	auruo.mutation.SetRoleID(id)
	return auruo
}

// SetNillableRoleID sets the "role" edge to the AdminRoles entity by ID if the given value is not nil.
func (auruo *AdminUserRoleUpdateOne) SetNillableRoleID(id *uuid.UUID) *AdminUserRoleUpdateOne {
	if id != nil {
		auruo = auruo.SetRoleID(*id)
	}
	return auruo
}

// SetRole sets the "role" edge to the AdminRoles entity.
func (auruo *AdminUserRoleUpdateOne) SetRole(a *AdminRoles) *AdminUserRoleUpdateOne {
	return auruo.SetRoleID(a.ID)
}

// Mutation returns the AdminUserRoleMutation object of the builder.
func (auruo *AdminUserRoleUpdateOne) Mutation() *AdminUserRoleMutation {
	return auruo.mutation
}

// ClearUser clears the "user" edge to the AdminUser entity.
func (auruo *AdminUserRoleUpdateOne) ClearUser() *AdminUserRoleUpdateOne {
	auruo.mutation.ClearUser()
	return auruo
}

// ClearRole clears the "role" edge to the AdminRoles entity.
func (auruo *AdminUserRoleUpdateOne) ClearRole() *AdminUserRoleUpdateOne {
	auruo.mutation.ClearRole()
	return auruo
}

// Where appends a list predicates to the AdminUserRoleUpdate builder.
func (auruo *AdminUserRoleUpdateOne) Where(ps ...predicate.AdminUserRole) *AdminUserRoleUpdateOne {
	auruo.mutation.Where(ps...)
	return auruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auruo *AdminUserRoleUpdateOne) Select(field string, fields ...string) *AdminUserRoleUpdateOne {
	auruo.fields = append([]string{field}, fields...)
	return auruo
}

// Save executes the query and returns the updated AdminUserRole entity.
func (auruo *AdminUserRoleUpdateOne) Save(ctx context.Context) (*AdminUserRole, error) {
	return withHooks(ctx, auruo.sqlSave, auruo.mutation, auruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auruo *AdminUserRoleUpdateOne) SaveX(ctx context.Context) *AdminUserRole {
	node, err := auruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auruo *AdminUserRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := auruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auruo *AdminUserRoleUpdateOne) ExecX(ctx context.Context) {
	if err := auruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auruo *AdminUserRoleUpdateOne) sqlSave(ctx context.Context) (_node *AdminUserRole, err error) {
	_spec := sqlgraph.NewUpdateSpec(adminuserrole.Table, adminuserrole.Columns, sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt))
	id, ok := auruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AdminUserRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminuserrole.FieldID)
		for _, f := range fields {
			if !adminuserrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adminuserrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if auruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminuserrole.UserTable,
			Columns: []string{adminuserrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminuserrole.UserTable,
			Columns: []string{adminuserrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auruo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminuserrole.RoleTable,
			Columns: []string{adminuserrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminroles.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auruo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminuserrole.RoleTable,
			Columns: []string{adminuserrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminroles.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AdminUserRole{config: auruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminuserrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auruo.mutation.done = true
	return _node, nil
}