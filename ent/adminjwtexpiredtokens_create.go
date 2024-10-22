// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthmonitor/ent/adminjwtexpiredtokens"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminJWTExpiredTokensCreate is the builder for creating a AdminJWTExpiredTokens entity.
type AdminJWTExpiredTokensCreate struct {
	config
	mutation *AdminJWTExpiredTokensMutation
	hooks    []Hook
}

// SetJti sets the "jti" field.
func (ajetc *AdminJWTExpiredTokensCreate) SetJti(s string) *AdminJWTExpiredTokensCreate {
	ajetc.mutation.SetJti(s)
	return ajetc
}

// SetExpiresAt sets the "expires_at" field.
func (ajetc *AdminJWTExpiredTokensCreate) SetExpiresAt(i int) *AdminJWTExpiredTokensCreate {
	ajetc.mutation.SetExpiresAt(i)
	return ajetc
}

// SetRevokedAt sets the "revoked_at" field.
func (ajetc *AdminJWTExpiredTokensCreate) SetRevokedAt(i int) *AdminJWTExpiredTokensCreate {
	ajetc.mutation.SetRevokedAt(i)
	return ajetc
}

// SetNillableRevokedAt sets the "revoked_at" field if the given value is not nil.
func (ajetc *AdminJWTExpiredTokensCreate) SetNillableRevokedAt(i *int) *AdminJWTExpiredTokensCreate {
	if i != nil {
		ajetc.SetRevokedAt(*i)
	}
	return ajetc
}

// Mutation returns the AdminJWTExpiredTokensMutation object of the builder.
func (ajetc *AdminJWTExpiredTokensCreate) Mutation() *AdminJWTExpiredTokensMutation {
	return ajetc.mutation
}

// Save creates the AdminJWTExpiredTokens in the database.
func (ajetc *AdminJWTExpiredTokensCreate) Save(ctx context.Context) (*AdminJWTExpiredTokens, error) {
	return withHooks(ctx, ajetc.sqlSave, ajetc.mutation, ajetc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ajetc *AdminJWTExpiredTokensCreate) SaveX(ctx context.Context) *AdminJWTExpiredTokens {
	v, err := ajetc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ajetc *AdminJWTExpiredTokensCreate) Exec(ctx context.Context) error {
	_, err := ajetc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ajetc *AdminJWTExpiredTokensCreate) ExecX(ctx context.Context) {
	if err := ajetc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ajetc *AdminJWTExpiredTokensCreate) check() error {
	if _, ok := ajetc.mutation.Jti(); !ok {
		return &ValidationError{Name: "jti", err: errors.New(`ent: missing required field "AdminJWTExpiredTokens.jti"`)}
	}
	if _, ok := ajetc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "AdminJWTExpiredTokens.expires_at"`)}
	}
	return nil
}

func (ajetc *AdminJWTExpiredTokensCreate) sqlSave(ctx context.Context) (*AdminJWTExpiredTokens, error) {
	if err := ajetc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ajetc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ajetc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ajetc.mutation.id = &_node.ID
	ajetc.mutation.done = true
	return _node, nil
}

func (ajetc *AdminJWTExpiredTokensCreate) createSpec() (*AdminJWTExpiredTokens, *sqlgraph.CreateSpec) {
	var (
		_node = &AdminJWTExpiredTokens{config: ajetc.config}
		_spec = sqlgraph.NewCreateSpec(adminjwtexpiredtokens.Table, sqlgraph.NewFieldSpec(adminjwtexpiredtokens.FieldID, field.TypeInt))
	)
	if value, ok := ajetc.mutation.Jti(); ok {
		_spec.SetField(adminjwtexpiredtokens.FieldJti, field.TypeString, value)
		_node.Jti = value
	}
	if value, ok := ajetc.mutation.ExpiresAt(); ok {
		_spec.SetField(adminjwtexpiredtokens.FieldExpiresAt, field.TypeInt, value)
		_node.ExpiresAt = value
	}
	if value, ok := ajetc.mutation.RevokedAt(); ok {
		_spec.SetField(adminjwtexpiredtokens.FieldRevokedAt, field.TypeInt, value)
		_node.RevokedAt = value
	}
	return _node, _spec
}

// AdminJWTExpiredTokensCreateBulk is the builder for creating many AdminJWTExpiredTokens entities in bulk.
type AdminJWTExpiredTokensCreateBulk struct {
	config
	err      error
	builders []*AdminJWTExpiredTokensCreate
}

// Save creates the AdminJWTExpiredTokens entities in the database.
func (ajetcb *AdminJWTExpiredTokensCreateBulk) Save(ctx context.Context) ([]*AdminJWTExpiredTokens, error) {
	if ajetcb.err != nil {
		return nil, ajetcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ajetcb.builders))
	nodes := make([]*AdminJWTExpiredTokens, len(ajetcb.builders))
	mutators := make([]Mutator, len(ajetcb.builders))
	for i := range ajetcb.builders {
		func(i int, root context.Context) {
			builder := ajetcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminJWTExpiredTokensMutation)
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
					_, err = mutators[i+1].Mutate(root, ajetcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ajetcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ajetcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ajetcb *AdminJWTExpiredTokensCreateBulk) SaveX(ctx context.Context) []*AdminJWTExpiredTokens {
	v, err := ajetcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ajetcb *AdminJWTExpiredTokensCreateBulk) Exec(ctx context.Context) error {
	_, err := ajetcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ajetcb *AdminJWTExpiredTokensCreateBulk) ExecX(ctx context.Context) {
	if err := ajetcb.Exec(ctx); err != nil {
		panic(err)
	}
}