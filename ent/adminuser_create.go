// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/adminlog"
	"HealthMonitor/ent/adminuser"
	"HealthMonitor/ent/adminuserrole"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AdminUserCreate is the builder for creating a AdminUser entity.
type AdminUserCreate struct {
	config
	mutation *AdminUserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (auc *AdminUserCreate) SetUsername(s string) *AdminUserCreate {
	auc.mutation.SetUsername(s)
	return auc
}

// SetEmail sets the "email" field.
func (auc *AdminUserCreate) SetEmail(s string) *AdminUserCreate {
	auc.mutation.SetEmail(s)
	return auc
}

// SetPasswordHash sets the "password_hash" field.
func (auc *AdminUserCreate) SetPasswordHash(s string) *AdminUserCreate {
	auc.mutation.SetPasswordHash(s)
	return auc
}

// SetIsActive sets the "is_active" field.
func (auc *AdminUserCreate) SetIsActive(b bool) *AdminUserCreate {
	auc.mutation.SetIsActive(b)
	return auc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (auc *AdminUserCreate) SetNillableIsActive(b *bool) *AdminUserCreate {
	if b != nil {
		auc.SetIsActive(*b)
	}
	return auc
}

// SetLastLoginAt sets the "last_login_at" field.
func (auc *AdminUserCreate) SetLastLoginAt(t time.Time) *AdminUserCreate {
	auc.mutation.SetLastLoginAt(t)
	return auc
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (auc *AdminUserCreate) SetNillableLastLoginAt(t *time.Time) *AdminUserCreate {
	if t != nil {
		auc.SetLastLoginAt(*t)
	}
	return auc
}

// SetJwtIssuedAt sets the "jwt_issued_at" field.
func (auc *AdminUserCreate) SetJwtIssuedAt(t time.Time) *AdminUserCreate {
	auc.mutation.SetJwtIssuedAt(t)
	return auc
}

// SetNillableJwtIssuedAt sets the "jwt_issued_at" field if the given value is not nil.
func (auc *AdminUserCreate) SetNillableJwtIssuedAt(t *time.Time) *AdminUserCreate {
	if t != nil {
		auc.SetJwtIssuedAt(*t)
	}
	return auc
}

// SetCreatedAt sets the "created_at" field.
func (auc *AdminUserCreate) SetCreatedAt(t time.Time) *AdminUserCreate {
	auc.mutation.SetCreatedAt(t)
	return auc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auc *AdminUserCreate) SetNillableCreatedAt(t *time.Time) *AdminUserCreate {
	if t != nil {
		auc.SetCreatedAt(*t)
	}
	return auc
}

// SetUpdatedAt sets the "updated_at" field.
func (auc *AdminUserCreate) SetUpdatedAt(t time.Time) *AdminUserCreate {
	auc.mutation.SetUpdatedAt(t)
	return auc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auc *AdminUserCreate) SetNillableUpdatedAt(t *time.Time) *AdminUserCreate {
	if t != nil {
		auc.SetUpdatedAt(*t)
	}
	return auc
}

// SetID sets the "id" field.
func (auc *AdminUserCreate) SetID(u uuid.UUID) *AdminUserCreate {
	auc.mutation.SetID(u)
	return auc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (auc *AdminUserCreate) SetNillableID(u *uuid.UUID) *AdminUserCreate {
	if u != nil {
		auc.SetID(*u)
	}
	return auc
}

// AddRoleIDs adds the "roles" edge to the AdminUserRole entity by IDs.
func (auc *AdminUserCreate) AddRoleIDs(ids ...int) *AdminUserCreate {
	auc.mutation.AddRoleIDs(ids...)
	return auc
}

// AddRoles adds the "roles" edges to the AdminUserRole entity.
func (auc *AdminUserCreate) AddRoles(a ...*AdminUserRole) *AdminUserCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auc.AddRoleIDs(ids...)
}

// AddLogIDs adds the "logs" edge to the AdminLog entity by IDs.
func (auc *AdminUserCreate) AddLogIDs(ids ...uuid.UUID) *AdminUserCreate {
	auc.mutation.AddLogIDs(ids...)
	return auc
}

// AddLogs adds the "logs" edges to the AdminLog entity.
func (auc *AdminUserCreate) AddLogs(a ...*AdminLog) *AdminUserCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auc.AddLogIDs(ids...)
}

// Mutation returns the AdminUserMutation object of the builder.
func (auc *AdminUserCreate) Mutation() *AdminUserMutation {
	return auc.mutation
}

// Save creates the AdminUser in the database.
func (auc *AdminUserCreate) Save(ctx context.Context) (*AdminUser, error) {
	auc.defaults()
	return withHooks(ctx, auc.sqlSave, auc.mutation, auc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (auc *AdminUserCreate) SaveX(ctx context.Context) *AdminUser {
	v, err := auc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auc *AdminUserCreate) Exec(ctx context.Context) error {
	_, err := auc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auc *AdminUserCreate) ExecX(ctx context.Context) {
	if err := auc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auc *AdminUserCreate) defaults() {
	if _, ok := auc.mutation.IsActive(); !ok {
		v := adminuser.DefaultIsActive
		auc.mutation.SetIsActive(v)
	}
	if _, ok := auc.mutation.CreatedAt(); !ok {
		v := adminuser.DefaultCreatedAt()
		auc.mutation.SetCreatedAt(v)
	}
	if _, ok := auc.mutation.UpdatedAt(); !ok {
		v := adminuser.DefaultUpdatedAt()
		auc.mutation.SetUpdatedAt(v)
	}
	if _, ok := auc.mutation.ID(); !ok {
		v := adminuser.DefaultID()
		auc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auc *AdminUserCreate) check() error {
	if _, ok := auc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "AdminUser.username"`)}
	}
	if _, ok := auc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "AdminUser.email"`)}
	}
	if _, ok := auc.mutation.PasswordHash(); !ok {
		return &ValidationError{Name: "password_hash", err: errors.New(`ent: missing required field "AdminUser.password_hash"`)}
	}
	if _, ok := auc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "AdminUser.is_active"`)}
	}
	if _, ok := auc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AdminUser.created_at"`)}
	}
	if _, ok := auc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AdminUser.updated_at"`)}
	}
	return nil
}

func (auc *AdminUserCreate) sqlSave(ctx context.Context) (*AdminUser, error) {
	if err := auc.check(); err != nil {
		return nil, err
	}
	_node, _spec := auc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	auc.mutation.id = &_node.ID
	auc.mutation.done = true
	return _node, nil
}

func (auc *AdminUserCreate) createSpec() (*AdminUser, *sqlgraph.CreateSpec) {
	var (
		_node = &AdminUser{config: auc.config}
		_spec = sqlgraph.NewCreateSpec(adminuser.Table, sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeUUID))
	)
	if id, ok := auc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := auc.mutation.Username(); ok {
		_spec.SetField(adminuser.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := auc.mutation.Email(); ok {
		_spec.SetField(adminuser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := auc.mutation.PasswordHash(); ok {
		_spec.SetField(adminuser.FieldPasswordHash, field.TypeString, value)
		_node.PasswordHash = value
	}
	if value, ok := auc.mutation.IsActive(); ok {
		_spec.SetField(adminuser.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := auc.mutation.LastLoginAt(); ok {
		_spec.SetField(adminuser.FieldLastLoginAt, field.TypeTime, value)
		_node.LastLoginAt = value
	}
	if value, ok := auc.mutation.JwtIssuedAt(); ok {
		_spec.SetField(adminuser.FieldJwtIssuedAt, field.TypeTime, value)
		_node.JwtIssuedAt = value
	}
	if value, ok := auc.mutation.CreatedAt(); ok {
		_spec.SetField(adminuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := auc.mutation.UpdatedAt(); ok {
		_spec.SetField(adminuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := auc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminuser.RolesTable,
			Columns: []string{adminuser.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminuserrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := auc.mutation.LogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adminuser.LogsTable,
			Columns: []string{adminuser.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(adminlog.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdminUserCreateBulk is the builder for creating many AdminUser entities in bulk.
type AdminUserCreateBulk struct {
	config
	err      error
	builders []*AdminUserCreate
}

// Save creates the AdminUser entities in the database.
func (aucb *AdminUserCreateBulk) Save(ctx context.Context) ([]*AdminUser, error) {
	if aucb.err != nil {
		return nil, aucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(aucb.builders))
	nodes := make([]*AdminUser, len(aucb.builders))
	mutators := make([]Mutator, len(aucb.builders))
	for i := range aucb.builders {
		func(i int, root context.Context) {
			builder := aucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aucb *AdminUserCreateBulk) SaveX(ctx context.Context) []*AdminUser {
	v, err := aucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aucb *AdminUserCreateBulk) Exec(ctx context.Context) error {
	_, err := aucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aucb *AdminUserCreateBulk) ExecX(ctx context.Context) {
	if err := aucb.Exec(ctx); err != nil {
		panic(err)
	}
}
