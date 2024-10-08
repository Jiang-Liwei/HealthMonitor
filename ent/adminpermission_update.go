// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthmonitor/ent/adminpermission"
	"healthmonitor/ent/adminrolepermission"
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminPermissionUpdate is the builder for updating AdminPermission entities.
type AdminPermissionUpdate struct {
	config
	hooks    []Hook
	mutation *AdminPermissionMutation
}

// Where appends a list predicates to the AdminPermissionUpdate builder.
func (apu *AdminPermissionUpdate) Where(ps ...predicate.AdminPermission) *AdminPermissionUpdate {
	apu.mutation.Where(ps...)
	return apu
}

// SetName sets the "name" field.
func (apu *AdminPermissionUpdate) SetName(s string) *AdminPermissionUpdate {
	apu.mutation.SetName(s)
	return apu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (apu *AdminPermissionUpdate) SetNillableName(s *string) *AdminPermissionUpdate {
	if s != nil {
		apu.SetName(*s)
	}
	return apu
}

// SetDescription sets the "description" field.
func (apu *AdminPermissionUpdate) SetDescription(s string) *AdminPermissionUpdate {
	apu.mutation.SetDescription(s)
	return apu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (apu *AdminPermissionUpdate) SetNillableDescription(s *string) *AdminPermissionUpdate {
	if s != nil {
		apu.SetDescription(*s)
	}
	return apu
}

// ClearDescription clears the value of the "description" field.
func (apu *AdminPermissionUpdate) ClearDescription() *AdminPermissionUpdate {
	apu.mutation.ClearDescription()
	return apu
}

// SetPath sets the "path" field.
func (apu *AdminPermissionUpdate) SetPath(s string) *AdminPermissionUpdate {
	apu.mutation.SetPath(s)
	return apu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (apu *AdminPermissionUpdate) SetNillablePath(s *string) *AdminPermissionUpdate {
	if s != nil {
		apu.SetPath(*s)
	}
	return apu
}

// SetMethod sets the "method" field.
func (apu *AdminPermissionUpdate) SetMethod(s string) *AdminPermissionUpdate {
	apu.mutation.SetMethod(s)
	return apu
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (apu *AdminPermissionUpdate) SetNillableMethod(s *string) *AdminPermissionUpdate {
	if s != nil {
		apu.SetMethod(*s)
	}
	return apu
}

// SetCreatedAt sets the "created_at" field.
func (apu *AdminPermissionUpdate) SetCreatedAt(i int) *AdminPermissionUpdate {
	apu.mutation.ResetCreatedAt()
	apu.mutation.SetCreatedAt(i)
	return apu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (apu *AdminPermissionUpdate) SetNillableCreatedAt(i *int) *AdminPermissionUpdate {
	if i != nil {
		apu.SetCreatedAt(*i)
	}
	return apu
}

// AddCreatedAt adds i to the "created_at" field.
func (apu *AdminPermissionUpdate) AddCreatedAt(i int) *AdminPermissionUpdate {
	apu.mutation.AddCreatedAt(i)
	return apu
}

// SetUpdatedAt sets the "updated_at" field.
func (apu *AdminPermissionUpdate) SetUpdatedAt(i int) *AdminPermissionUpdate {
	apu.mutation.ResetUpdatedAt()
	apu.mutation.SetUpdatedAt(i)
	return apu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (apu *AdminPermissionUpdate) AddUpdatedAt(i int) *AdminPermissionUpdate {
	apu.mutation.AddUpdatedAt(i)
	return apu
}

// SetDeletedAt sets the "deleted_at" field.
func (apu *AdminPermissionUpdate) SetDeletedAt(i int) *AdminPermissionUpdate {
	apu.mutation.ResetDeletedAt()
	apu.mutation.SetDeletedAt(i)
	return apu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (apu *AdminPermissionUpdate) SetNillableDeletedAt(i *int) *AdminPermissionUpdate {
	if i != nil {
		apu.SetDeletedAt(*i)
	}
	return apu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (apu *AdminPermissionUpdate) AddDeletedAt(i int) *AdminPermissionUpdate {
	apu.mutation.AddDeletedAt(i)
	return apu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (apu *AdminPermissionUpdate) ClearDeletedAt() *AdminPermissionUpdate {
	apu.mutation.ClearDeletedAt()
	return apu
}

// AddRoleIDs adds the "roles" edge to the AdminRolePermission entity by IDs.
func (apu *AdminPermissionUpdate) AddRoleIDs(ids ...int) *AdminPermissionUpdate {
	apu.mutation.AddRoleIDs(ids...)
	return apu
}

// AddRoles adds the "roles" edges to the AdminRolePermission entity.
func (apu *AdminPermissionUpdate) AddRoles(a ...*AdminRolePermission) *AdminPermissionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apu.AddRoleIDs(ids...)
}

// Mutation returns the AdminPermissionMutation object of the builder.
func (apu *AdminPermissionUpdate) Mutation() *AdminPermissionMutation {
	return apu.mutation
}

// ClearRoles clears all "roles" edges to the AdminRolePermission entity.
func (apu *AdminPermissionUpdate) ClearRoles() *AdminPermissionUpdate {
	apu.mutation.ClearRoles()
	return apu
}

// RemoveRoleIDs removes the "roles" edge to AdminRolePermission entities by IDs.
func (apu *AdminPermissionUpdate) RemoveRoleIDs(ids ...int) *AdminPermissionUpdate {
	apu.mutation.RemoveRoleIDs(ids...)
	return apu
}

// RemoveRoles removes "roles" edges to AdminRolePermission entities.
func (apu *AdminPermissionUpdate) RemoveRoles(a ...*AdminRolePermission) *AdminPermissionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apu *AdminPermissionUpdate) Save(ctx context.Context) (int, error) {
	apu.defaults()
	return withHooks(ctx, apu.sqlSave, apu.mutation, apu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apu *AdminPermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := apu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apu *AdminPermissionUpdate) Exec(ctx context.Context) error {
	_, err := apu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apu *AdminPermissionUpdate) ExecX(ctx context.Context) {
	if err := apu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apu *AdminPermissionUpdate) defaults() {
	if _, ok := apu.mutation.UpdatedAt(); !ok {
		v := adminpermission.UpdateDefaultUpdatedAt()
		apu.mutation.SetUpdatedAt(v)
	}
}

func (apu *AdminPermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(adminpermission.Table, adminpermission.Columns, sqlgraph.NewFieldSpec(adminpermission.FieldID, field.TypeUUID))
	if ps := apu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apu.mutation.Name(); ok {
		_spec.SetField(adminpermission.FieldName, field.TypeString, value)
	}
	if value, ok := apu.mutation.Description(); ok {
		_spec.SetField(adminpermission.FieldDescription, field.TypeString, value)
	}
	if apu.mutation.DescriptionCleared() {
		_spec.ClearField(adminpermission.FieldDescription, field.TypeString)
	}
	if value, ok := apu.mutation.Path(); ok {
		_spec.SetField(adminpermission.FieldPath, field.TypeString, value)
	}
	if value, ok := apu.mutation.Method(); ok {
		_spec.SetField(adminpermission.FieldMethod, field.TypeString, value)
	}
	if value, ok := apu.mutation.CreatedAt(); ok {
		_spec.SetField(adminpermission.FieldCreatedAt, field.TypeInt, value)
	}
	if value, ok := apu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(adminpermission.FieldCreatedAt, field.TypeInt, value)
	}
	if value, ok := apu.mutation.UpdatedAt(); ok {
		_spec.SetField(adminpermission.FieldUpdatedAt, field.TypeInt, value)
	}
	if value, ok := apu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(adminpermission.FieldUpdatedAt, field.TypeInt, value)
	}
	if value, ok := apu.mutation.DeletedAt(); ok {
		_spec.SetField(adminpermission.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := apu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(adminpermission.FieldDeletedAt, field.TypeInt, value)
	}
	if apu.mutation.DeletedAtCleared() {
		_spec.ClearField(adminpermission.FieldDeletedAt, field.TypeInt)
	}
	if apu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminpermission.RolesTable,
			Columns: []string{adminpermission.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminrolepermission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !apu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminpermission.RolesTable,
			Columns: []string{adminpermission.RolesColumn},
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
	if nodes := apu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminpermission.RolesTable,
			Columns: []string{adminpermission.RolesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, apu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminpermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	apu.mutation.done = true
	return n, nil
}

// AdminPermissionUpdateOne is the builder for updating a single AdminPermission entity.
type AdminPermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdminPermissionMutation
}

// SetName sets the "name" field.
func (apuo *AdminPermissionUpdateOne) SetName(s string) *AdminPermissionUpdateOne {
	apuo.mutation.SetName(s)
	return apuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (apuo *AdminPermissionUpdateOne) SetNillableName(s *string) *AdminPermissionUpdateOne {
	if s != nil {
		apuo.SetName(*s)
	}
	return apuo
}

// SetDescription sets the "description" field.
func (apuo *AdminPermissionUpdateOne) SetDescription(s string) *AdminPermissionUpdateOne {
	apuo.mutation.SetDescription(s)
	return apuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (apuo *AdminPermissionUpdateOne) SetNillableDescription(s *string) *AdminPermissionUpdateOne {
	if s != nil {
		apuo.SetDescription(*s)
	}
	return apuo
}

// ClearDescription clears the value of the "description" field.
func (apuo *AdminPermissionUpdateOne) ClearDescription() *AdminPermissionUpdateOne {
	apuo.mutation.ClearDescription()
	return apuo
}

// SetPath sets the "path" field.
func (apuo *AdminPermissionUpdateOne) SetPath(s string) *AdminPermissionUpdateOne {
	apuo.mutation.SetPath(s)
	return apuo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (apuo *AdminPermissionUpdateOne) SetNillablePath(s *string) *AdminPermissionUpdateOne {
	if s != nil {
		apuo.SetPath(*s)
	}
	return apuo
}

// SetMethod sets the "method" field.
func (apuo *AdminPermissionUpdateOne) SetMethod(s string) *AdminPermissionUpdateOne {
	apuo.mutation.SetMethod(s)
	return apuo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (apuo *AdminPermissionUpdateOne) SetNillableMethod(s *string) *AdminPermissionUpdateOne {
	if s != nil {
		apuo.SetMethod(*s)
	}
	return apuo
}

// SetCreatedAt sets the "created_at" field.
func (apuo *AdminPermissionUpdateOne) SetCreatedAt(i int) *AdminPermissionUpdateOne {
	apuo.mutation.ResetCreatedAt()
	apuo.mutation.SetCreatedAt(i)
	return apuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (apuo *AdminPermissionUpdateOne) SetNillableCreatedAt(i *int) *AdminPermissionUpdateOne {
	if i != nil {
		apuo.SetCreatedAt(*i)
	}
	return apuo
}

// AddCreatedAt adds i to the "created_at" field.
func (apuo *AdminPermissionUpdateOne) AddCreatedAt(i int) *AdminPermissionUpdateOne {
	apuo.mutation.AddCreatedAt(i)
	return apuo
}

// SetUpdatedAt sets the "updated_at" field.
func (apuo *AdminPermissionUpdateOne) SetUpdatedAt(i int) *AdminPermissionUpdateOne {
	apuo.mutation.ResetUpdatedAt()
	apuo.mutation.SetUpdatedAt(i)
	return apuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (apuo *AdminPermissionUpdateOne) AddUpdatedAt(i int) *AdminPermissionUpdateOne {
	apuo.mutation.AddUpdatedAt(i)
	return apuo
}

// SetDeletedAt sets the "deleted_at" field.
func (apuo *AdminPermissionUpdateOne) SetDeletedAt(i int) *AdminPermissionUpdateOne {
	apuo.mutation.ResetDeletedAt()
	apuo.mutation.SetDeletedAt(i)
	return apuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (apuo *AdminPermissionUpdateOne) SetNillableDeletedAt(i *int) *AdminPermissionUpdateOne {
	if i != nil {
		apuo.SetDeletedAt(*i)
	}
	return apuo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (apuo *AdminPermissionUpdateOne) AddDeletedAt(i int) *AdminPermissionUpdateOne {
	apuo.mutation.AddDeletedAt(i)
	return apuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (apuo *AdminPermissionUpdateOne) ClearDeletedAt() *AdminPermissionUpdateOne {
	apuo.mutation.ClearDeletedAt()
	return apuo
}

// AddRoleIDs adds the "roles" edge to the AdminRolePermission entity by IDs.
func (apuo *AdminPermissionUpdateOne) AddRoleIDs(ids ...int) *AdminPermissionUpdateOne {
	apuo.mutation.AddRoleIDs(ids...)
	return apuo
}

// AddRoles adds the "roles" edges to the AdminRolePermission entity.
func (apuo *AdminPermissionUpdateOne) AddRoles(a ...*AdminRolePermission) *AdminPermissionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apuo.AddRoleIDs(ids...)
}

// Mutation returns the AdminPermissionMutation object of the builder.
func (apuo *AdminPermissionUpdateOne) Mutation() *AdminPermissionMutation {
	return apuo.mutation
}

// ClearRoles clears all "roles" edges to the AdminRolePermission entity.
func (apuo *AdminPermissionUpdateOne) ClearRoles() *AdminPermissionUpdateOne {
	apuo.mutation.ClearRoles()
	return apuo
}

// RemoveRoleIDs removes the "roles" edge to AdminRolePermission entities by IDs.
func (apuo *AdminPermissionUpdateOne) RemoveRoleIDs(ids ...int) *AdminPermissionUpdateOne {
	apuo.mutation.RemoveRoleIDs(ids...)
	return apuo
}

// RemoveRoles removes "roles" edges to AdminRolePermission entities.
func (apuo *AdminPermissionUpdateOne) RemoveRoles(a ...*AdminRolePermission) *AdminPermissionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apuo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the AdminPermissionUpdate builder.
func (apuo *AdminPermissionUpdateOne) Where(ps ...predicate.AdminPermission) *AdminPermissionUpdateOne {
	apuo.mutation.Where(ps...)
	return apuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apuo *AdminPermissionUpdateOne) Select(field string, fields ...string) *AdminPermissionUpdateOne {
	apuo.fields = append([]string{field}, fields...)
	return apuo
}

// Save executes the query and returns the updated AdminPermission entity.
func (apuo *AdminPermissionUpdateOne) Save(ctx context.Context) (*AdminPermission, error) {
	apuo.defaults()
	return withHooks(ctx, apuo.sqlSave, apuo.mutation, apuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apuo *AdminPermissionUpdateOne) SaveX(ctx context.Context) *AdminPermission {
	node, err := apuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apuo *AdminPermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := apuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apuo *AdminPermissionUpdateOne) ExecX(ctx context.Context) {
	if err := apuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apuo *AdminPermissionUpdateOne) defaults() {
	if _, ok := apuo.mutation.UpdatedAt(); !ok {
		v := adminpermission.UpdateDefaultUpdatedAt()
		apuo.mutation.SetUpdatedAt(v)
	}
}

func (apuo *AdminPermissionUpdateOne) sqlSave(ctx context.Context) (_node *AdminPermission, err error) {
	_spec := sqlgraph.NewUpdateSpec(adminpermission.Table, adminpermission.Columns, sqlgraph.NewFieldSpec(adminpermission.FieldID, field.TypeUUID))
	id, ok := apuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AdminPermission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := apuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminpermission.FieldID)
		for _, f := range fields {
			if !adminpermission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adminpermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apuo.mutation.Name(); ok {
		_spec.SetField(adminpermission.FieldName, field.TypeString, value)
	}
	if value, ok := apuo.mutation.Description(); ok {
		_spec.SetField(adminpermission.FieldDescription, field.TypeString, value)
	}
	if apuo.mutation.DescriptionCleared() {
		_spec.ClearField(adminpermission.FieldDescription, field.TypeString)
	}
	if value, ok := apuo.mutation.Path(); ok {
		_spec.SetField(adminpermission.FieldPath, field.TypeString, value)
	}
	if value, ok := apuo.mutation.Method(); ok {
		_spec.SetField(adminpermission.FieldMethod, field.TypeString, value)
	}
	if value, ok := apuo.mutation.CreatedAt(); ok {
		_spec.SetField(adminpermission.FieldCreatedAt, field.TypeInt, value)
	}
	if value, ok := apuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(adminpermission.FieldCreatedAt, field.TypeInt, value)
	}
	if value, ok := apuo.mutation.UpdatedAt(); ok {
		_spec.SetField(adminpermission.FieldUpdatedAt, field.TypeInt, value)
	}
	if value, ok := apuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(adminpermission.FieldUpdatedAt, field.TypeInt, value)
	}
	if value, ok := apuo.mutation.DeletedAt(); ok {
		_spec.SetField(adminpermission.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := apuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(adminpermission.FieldDeletedAt, field.TypeInt, value)
	}
	if apuo.mutation.DeletedAtCleared() {
		_spec.ClearField(adminpermission.FieldDeletedAt, field.TypeInt)
	}
	if apuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminpermission.RolesTable,
			Columns: []string{adminpermission.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminrolepermission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !apuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminpermission.RolesTable,
			Columns: []string{adminpermission.RolesColumn},
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
	if nodes := apuo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminpermission.RolesTable,
			Columns: []string{adminpermission.RolesColumn},
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
	_node = &AdminPermission{config: apuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminpermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	apuo.mutation.done = true
	return _node, nil
}
