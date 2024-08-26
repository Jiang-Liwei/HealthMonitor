// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/adminrolepermission"
	"HealthMonitor/ent/adminroles"
	"HealthMonitor/ent/adminuserrole"
	"HealthMonitor/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminRolesUpdate is the builder for updating AdminRoles entities.
type AdminRolesUpdate struct {
	config
	hooks    []Hook
	mutation *AdminRolesMutation
}

// Where appends a list predicates to the AdminRolesUpdate builder.
func (aru *AdminRolesUpdate) Where(ps ...predicate.AdminRoles) *AdminRolesUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetName sets the "name" field.
func (aru *AdminRolesUpdate) SetName(s string) *AdminRolesUpdate {
	aru.mutation.SetName(s)
	return aru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (aru *AdminRolesUpdate) SetNillableName(s *string) *AdminRolesUpdate {
	if s != nil {
		aru.SetName(*s)
	}
	return aru
}

// SetDescription sets the "description" field.
func (aru *AdminRolesUpdate) SetDescription(s string) *AdminRolesUpdate {
	aru.mutation.SetDescription(s)
	return aru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (aru *AdminRolesUpdate) SetNillableDescription(s *string) *AdminRolesUpdate {
	if s != nil {
		aru.SetDescription(*s)
	}
	return aru
}

// ClearDescription clears the value of the "description" field.
func (aru *AdminRolesUpdate) ClearDescription() *AdminRolesUpdate {
	aru.mutation.ClearDescription()
	return aru
}

// SetCreatedAt sets the "created_at" field.
func (aru *AdminRolesUpdate) SetCreatedAt(t time.Time) *AdminRolesUpdate {
	aru.mutation.SetCreatedAt(t)
	return aru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aru *AdminRolesUpdate) SetNillableCreatedAt(t *time.Time) *AdminRolesUpdate {
	if t != nil {
		aru.SetCreatedAt(*t)
	}
	return aru
}

// SetUpdatedAt sets the "updated_at" field.
func (aru *AdminRolesUpdate) SetUpdatedAt(t time.Time) *AdminRolesUpdate {
	aru.mutation.SetUpdatedAt(t)
	return aru
}

// AddPermissionIDs adds the "permissions" edge to the AdminRolePermission entity by IDs.
func (aru *AdminRolesUpdate) AddPermissionIDs(ids ...int) *AdminRolesUpdate {
	aru.mutation.AddPermissionIDs(ids...)
	return aru
}

// AddPermissions adds the "permissions" edges to the AdminRolePermission entity.
func (aru *AdminRolesUpdate) AddPermissions(a ...*AdminRolePermission) *AdminRolesUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.AddPermissionIDs(ids...)
}

// AddUserIDs adds the "users" edge to the AdminUserRole entity by IDs.
func (aru *AdminRolesUpdate) AddUserIDs(ids ...int) *AdminRolesUpdate {
	aru.mutation.AddUserIDs(ids...)
	return aru
}

// AddUsers adds the "users" edges to the AdminUserRole entity.
func (aru *AdminRolesUpdate) AddUsers(a ...*AdminUserRole) *AdminRolesUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.AddUserIDs(ids...)
}

// Mutation returns the AdminRolesMutation object of the builder.
func (aru *AdminRolesUpdate) Mutation() *AdminRolesMutation {
	return aru.mutation
}

// ClearPermissions clears all "permissions" edges to the AdminRolePermission entity.
func (aru *AdminRolesUpdate) ClearPermissions() *AdminRolesUpdate {
	aru.mutation.ClearPermissions()
	return aru
}

// RemovePermissionIDs removes the "permissions" edge to AdminRolePermission entities by IDs.
func (aru *AdminRolesUpdate) RemovePermissionIDs(ids ...int) *AdminRolesUpdate {
	aru.mutation.RemovePermissionIDs(ids...)
	return aru
}

// RemovePermissions removes "permissions" edges to AdminRolePermission entities.
func (aru *AdminRolesUpdate) RemovePermissions(a ...*AdminRolePermission) *AdminRolesUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.RemovePermissionIDs(ids...)
}

// ClearUsers clears all "users" edges to the AdminUserRole entity.
func (aru *AdminRolesUpdate) ClearUsers() *AdminRolesUpdate {
	aru.mutation.ClearUsers()
	return aru
}

// RemoveUserIDs removes the "users" edge to AdminUserRole entities by IDs.
func (aru *AdminRolesUpdate) RemoveUserIDs(ids ...int) *AdminRolesUpdate {
	aru.mutation.RemoveUserIDs(ids...)
	return aru
}

// RemoveUsers removes "users" edges to AdminUserRole entities.
func (aru *AdminRolesUpdate) RemoveUsers(a ...*AdminUserRole) *AdminRolesUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AdminRolesUpdate) Save(ctx context.Context) (int, error) {
	aru.defaults()
	return withHooks(ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AdminRolesUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AdminRolesUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AdminRolesUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aru *AdminRolesUpdate) defaults() {
	if _, ok := aru.mutation.UpdatedAt(); !ok {
		v := adminroles.UpdateDefaultUpdatedAt()
		aru.mutation.SetUpdatedAt(v)
	}
}

func (aru *AdminRolesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(adminroles.Table, adminroles.Columns, sqlgraph.NewFieldSpec(adminroles.FieldID, field.TypeUUID))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.Name(); ok {
		_spec.SetField(adminroles.FieldName, field.TypeString, value)
	}
	if value, ok := aru.mutation.Description(); ok {
		_spec.SetField(adminroles.FieldDescription, field.TypeString, value)
	}
	if aru.mutation.DescriptionCleared() {
		_spec.ClearField(adminroles.FieldDescription, field.TypeString)
	}
	if value, ok := aru.mutation.CreatedAt(); ok {
		_spec.SetField(adminroles.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := aru.mutation.UpdatedAt(); ok {
		_spec.SetField(adminroles.FieldUpdatedAt, field.TypeTime, value)
	}
	if aru.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.PermissionsTable,
			Columns: []string{adminroles.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminrolepermission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !aru.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.PermissionsTable,
			Columns: []string{adminroles.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminrolepermission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.PermissionsTable,
			Columns: []string{adminroles.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminrolepermission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.UsersTable,
			Columns: []string{adminroles.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !aru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.UsersTable,
			Columns: []string{adminroles.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.UsersTable,
			Columns: []string{adminroles.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminroles.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// AdminRolesUpdateOne is the builder for updating a single AdminRoles entity.
type AdminRolesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdminRolesMutation
}

// SetName sets the "name" field.
func (aruo *AdminRolesUpdateOne) SetName(s string) *AdminRolesUpdateOne {
	aruo.mutation.SetName(s)
	return aruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (aruo *AdminRolesUpdateOne) SetNillableName(s *string) *AdminRolesUpdateOne {
	if s != nil {
		aruo.SetName(*s)
	}
	return aruo
}

// SetDescription sets the "description" field.
func (aruo *AdminRolesUpdateOne) SetDescription(s string) *AdminRolesUpdateOne {
	aruo.mutation.SetDescription(s)
	return aruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (aruo *AdminRolesUpdateOne) SetNillableDescription(s *string) *AdminRolesUpdateOne {
	if s != nil {
		aruo.SetDescription(*s)
	}
	return aruo
}

// ClearDescription clears the value of the "description" field.
func (aruo *AdminRolesUpdateOne) ClearDescription() *AdminRolesUpdateOne {
	aruo.mutation.ClearDescription()
	return aruo
}

// SetCreatedAt sets the "created_at" field.
func (aruo *AdminRolesUpdateOne) SetCreatedAt(t time.Time) *AdminRolesUpdateOne {
	aruo.mutation.SetCreatedAt(t)
	return aruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aruo *AdminRolesUpdateOne) SetNillableCreatedAt(t *time.Time) *AdminRolesUpdateOne {
	if t != nil {
		aruo.SetCreatedAt(*t)
	}
	return aruo
}

// SetUpdatedAt sets the "updated_at" field.
func (aruo *AdminRolesUpdateOne) SetUpdatedAt(t time.Time) *AdminRolesUpdateOne {
	aruo.mutation.SetUpdatedAt(t)
	return aruo
}

// AddPermissionIDs adds the "permissions" edge to the AdminRolePermission entity by IDs.
func (aruo *AdminRolesUpdateOne) AddPermissionIDs(ids ...int) *AdminRolesUpdateOne {
	aruo.mutation.AddPermissionIDs(ids...)
	return aruo
}

// AddPermissions adds the "permissions" edges to the AdminRolePermission entity.
func (aruo *AdminRolesUpdateOne) AddPermissions(a ...*AdminRolePermission) *AdminRolesUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.AddPermissionIDs(ids...)
}

// AddUserIDs adds the "users" edge to the AdminUserRole entity by IDs.
func (aruo *AdminRolesUpdateOne) AddUserIDs(ids ...int) *AdminRolesUpdateOne {
	aruo.mutation.AddUserIDs(ids...)
	return aruo
}

// AddUsers adds the "users" edges to the AdminUserRole entity.
func (aruo *AdminRolesUpdateOne) AddUsers(a ...*AdminUserRole) *AdminRolesUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.AddUserIDs(ids...)
}

// Mutation returns the AdminRolesMutation object of the builder.
func (aruo *AdminRolesUpdateOne) Mutation() *AdminRolesMutation {
	return aruo.mutation
}

// ClearPermissions clears all "permissions" edges to the AdminRolePermission entity.
func (aruo *AdminRolesUpdateOne) ClearPermissions() *AdminRolesUpdateOne {
	aruo.mutation.ClearPermissions()
	return aruo
}

// RemovePermissionIDs removes the "permissions" edge to AdminRolePermission entities by IDs.
func (aruo *AdminRolesUpdateOne) RemovePermissionIDs(ids ...int) *AdminRolesUpdateOne {
	aruo.mutation.RemovePermissionIDs(ids...)
	return aruo
}

// RemovePermissions removes "permissions" edges to AdminRolePermission entities.
func (aruo *AdminRolesUpdateOne) RemovePermissions(a ...*AdminRolePermission) *AdminRolesUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.RemovePermissionIDs(ids...)
}

// ClearUsers clears all "users" edges to the AdminUserRole entity.
func (aruo *AdminRolesUpdateOne) ClearUsers() *AdminRolesUpdateOne {
	aruo.mutation.ClearUsers()
	return aruo
}

// RemoveUserIDs removes the "users" edge to AdminUserRole entities by IDs.
func (aruo *AdminRolesUpdateOne) RemoveUserIDs(ids ...int) *AdminRolesUpdateOne {
	aruo.mutation.RemoveUserIDs(ids...)
	return aruo
}

// RemoveUsers removes "users" edges to AdminUserRole entities.
func (aruo *AdminRolesUpdateOne) RemoveUsers(a ...*AdminUserRole) *AdminRolesUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the AdminRolesUpdate builder.
func (aruo *AdminRolesUpdateOne) Where(ps ...predicate.AdminRoles) *AdminRolesUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AdminRolesUpdateOne) Select(field string, fields ...string) *AdminRolesUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AdminRoles entity.
func (aruo *AdminRolesUpdateOne) Save(ctx context.Context) (*AdminRoles, error) {
	aruo.defaults()
	return withHooks(ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AdminRolesUpdateOne) SaveX(ctx context.Context) *AdminRoles {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AdminRolesUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AdminRolesUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aruo *AdminRolesUpdateOne) defaults() {
	if _, ok := aruo.mutation.UpdatedAt(); !ok {
		v := adminroles.UpdateDefaultUpdatedAt()
		aruo.mutation.SetUpdatedAt(v)
	}
}

func (aruo *AdminRolesUpdateOne) sqlSave(ctx context.Context) (_node *AdminRoles, err error) {
	_spec := sqlgraph.NewUpdateSpec(adminroles.Table, adminroles.Columns, sqlgraph.NewFieldSpec(adminroles.FieldID, field.TypeUUID))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AdminRoles.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminroles.FieldID)
		for _, f := range fields {
			if !adminroles.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adminroles.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.Name(); ok {
		_spec.SetField(adminroles.FieldName, field.TypeString, value)
	}
	if value, ok := aruo.mutation.Description(); ok {
		_spec.SetField(adminroles.FieldDescription, field.TypeString, value)
	}
	if aruo.mutation.DescriptionCleared() {
		_spec.ClearField(adminroles.FieldDescription, field.TypeString)
	}
	if value, ok := aruo.mutation.CreatedAt(); ok {
		_spec.SetField(adminroles.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := aruo.mutation.UpdatedAt(); ok {
		_spec.SetField(adminroles.FieldUpdatedAt, field.TypeTime, value)
	}
	if aruo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.PermissionsTable,
			Columns: []string{adminroles.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminrolepermission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !aruo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.PermissionsTable,
			Columns: []string{adminroles.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminrolepermission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.PermissionsTable,
			Columns: []string{adminroles.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminrolepermission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.UsersTable,
			Columns: []string{adminroles.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !aruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.UsersTable,
			Columns: []string{adminroles.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminroles.UsersTable,
			Columns: []string{adminroles.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AdminRoles{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminroles.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}