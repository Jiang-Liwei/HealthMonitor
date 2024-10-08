// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"healthmonitor/ent/food"
	"healthmonitor/ent/usermeal"
	"healthmonitor/ent/usermealfood"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserMealFoodCreate is the builder for creating a UserMealFood entity.
type UserMealFoodCreate struct {
	config
	mutation *UserMealFoodMutation
	hooks    []Hook
}

// SetUserMealID sets the "user_meal" edge to the UserMeal entity by ID.
func (umfc *UserMealFoodCreate) SetUserMealID(id int) *UserMealFoodCreate {
	umfc.mutation.SetUserMealID(id)
	return umfc
}

// SetNillableUserMealID sets the "user_meal" edge to the UserMeal entity by ID if the given value is not nil.
func (umfc *UserMealFoodCreate) SetNillableUserMealID(id *int) *UserMealFoodCreate {
	if id != nil {
		umfc = umfc.SetUserMealID(*id)
	}
	return umfc
}

// SetUserMeal sets the "user_meal" edge to the UserMeal entity.
func (umfc *UserMealFoodCreate) SetUserMeal(u *UserMeal) *UserMealFoodCreate {
	return umfc.SetUserMealID(u.ID)
}

// SetFoodID sets the "food" edge to the Food entity by ID.
func (umfc *UserMealFoodCreate) SetFoodID(id uuid.UUID) *UserMealFoodCreate {
	umfc.mutation.SetFoodID(id)
	return umfc
}

// SetNillableFoodID sets the "food" edge to the Food entity by ID if the given value is not nil.
func (umfc *UserMealFoodCreate) SetNillableFoodID(id *uuid.UUID) *UserMealFoodCreate {
	if id != nil {
		umfc = umfc.SetFoodID(*id)
	}
	return umfc
}

// SetFood sets the "food" edge to the Food entity.
func (umfc *UserMealFoodCreate) SetFood(f *Food) *UserMealFoodCreate {
	return umfc.SetFoodID(f.ID)
}

// Mutation returns the UserMealFoodMutation object of the builder.
func (umfc *UserMealFoodCreate) Mutation() *UserMealFoodMutation {
	return umfc.mutation
}

// Save creates the UserMealFood in the database.
func (umfc *UserMealFoodCreate) Save(ctx context.Context) (*UserMealFood, error) {
	return withHooks(ctx, umfc.sqlSave, umfc.mutation, umfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (umfc *UserMealFoodCreate) SaveX(ctx context.Context) *UserMealFood {
	v, err := umfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (umfc *UserMealFoodCreate) Exec(ctx context.Context) error {
	_, err := umfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umfc *UserMealFoodCreate) ExecX(ctx context.Context) {
	if err := umfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (umfc *UserMealFoodCreate) check() error {
	return nil
}

func (umfc *UserMealFoodCreate) sqlSave(ctx context.Context) (*UserMealFood, error) {
	if err := umfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := umfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, umfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	umfc.mutation.id = &_node.ID
	umfc.mutation.done = true
	return _node, nil
}

func (umfc *UserMealFoodCreate) createSpec() (*UserMealFood, *sqlgraph.CreateSpec) {
	var (
		_node = &UserMealFood{config: umfc.config}
		_spec = sqlgraph.NewCreateSpec(usermealfood.Table, sqlgraph.NewFieldSpec(usermealfood.FieldID, field.TypeInt))
	)
	if nodes := umfc.mutation.UserMealIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermealfood.UserMealTable,
			Columns: []string{usermealfood.UserMealColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usermeal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_meal_food = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := umfc.mutation.FoodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermealfood.FoodTable,
			Columns: []string{usermealfood.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.food_user_meal = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserMealFoodCreateBulk is the builder for creating many UserMealFood entities in bulk.
type UserMealFoodCreateBulk struct {
	config
	err      error
	builders []*UserMealFoodCreate
}

// Save creates the UserMealFood entities in the database.
func (umfcb *UserMealFoodCreateBulk) Save(ctx context.Context) ([]*UserMealFood, error) {
	if umfcb.err != nil {
		return nil, umfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(umfcb.builders))
	nodes := make([]*UserMealFood, len(umfcb.builders))
	mutators := make([]Mutator, len(umfcb.builders))
	for i := range umfcb.builders {
		func(i int, root context.Context) {
			builder := umfcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMealFoodMutation)
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
					_, err = mutators[i+1].Mutate(root, umfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, umfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, umfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (umfcb *UserMealFoodCreateBulk) SaveX(ctx context.Context) []*UserMealFood {
	v, err := umfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (umfcb *UserMealFoodCreateBulk) Exec(ctx context.Context) error {
	_, err := umfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umfcb *UserMealFoodCreateBulk) ExecX(ctx context.Context) {
	if err := umfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
