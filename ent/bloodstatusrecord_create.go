// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthmonitor/ent/bloodstatusrecord"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BloodStatusRecordCreate is the builder for creating a BloodStatusRecord entity.
type BloodStatusRecordCreate struct {
	config
	mutation *BloodStatusRecordMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (bsrc *BloodStatusRecordCreate) SetUserID(u uuid.UUID) *BloodStatusRecordCreate {
	bsrc.mutation.SetUserID(u)
	return bsrc
}

// SetRecordDate sets the "record_date" field.
func (bsrc *BloodStatusRecordCreate) SetRecordDate(u uint) *BloodStatusRecordCreate {
	bsrc.mutation.SetRecordDate(u)
	return bsrc
}

// SetTimeOfDay sets the "time_of_day" field.
func (bsrc *BloodStatusRecordCreate) SetTimeOfDay(bod bloodstatusrecord.TimeOfDay) *BloodStatusRecordCreate {
	bsrc.mutation.SetTimeOfDay(bod)
	return bsrc
}

// SetBeforeAfterMeals sets the "before_after_meals" field.
func (bsrc *BloodStatusRecordCreate) SetBeforeAfterMeals(bam bloodstatusrecord.BeforeAfterMeals) *BloodStatusRecordCreate {
	bsrc.mutation.SetBeforeAfterMeals(bam)
	return bsrc
}

// SetSystolicPressure sets the "systolic_pressure" field.
func (bsrc *BloodStatusRecordCreate) SetSystolicPressure(u uint8) *BloodStatusRecordCreate {
	bsrc.mutation.SetSystolicPressure(u)
	return bsrc
}

// SetDiastolicPressure sets the "diastolic_pressure" field.
func (bsrc *BloodStatusRecordCreate) SetDiastolicPressure(u uint8) *BloodStatusRecordCreate {
	bsrc.mutation.SetDiastolicPressure(u)
	return bsrc
}

// SetPulse sets the "pulse" field.
func (bsrc *BloodStatusRecordCreate) SetPulse(u uint8) *BloodStatusRecordCreate {
	bsrc.mutation.SetPulse(u)
	return bsrc
}

// SetMood sets the "mood" field.
func (bsrc *BloodStatusRecordCreate) SetMood(b bloodstatusrecord.Mood) *BloodStatusRecordCreate {
	bsrc.mutation.SetMood(b)
	return bsrc
}

// SetStatusSummary sets the "status_summary" field.
func (bsrc *BloodStatusRecordCreate) SetStatusSummary(bs bloodstatusrecord.StatusSummary) *BloodStatusRecordCreate {
	bsrc.mutation.SetStatusSummary(bs)
	return bsrc
}

// SetCreatedAt sets the "created_at" field.
func (bsrc *BloodStatusRecordCreate) SetCreatedAt(i int64) *BloodStatusRecordCreate {
	bsrc.mutation.SetCreatedAt(i)
	return bsrc
}

// SetUpdatedAt sets the "updated_at" field.
func (bsrc *BloodStatusRecordCreate) SetUpdatedAt(i int64) *BloodStatusRecordCreate {
	bsrc.mutation.SetUpdatedAt(i)
	return bsrc
}

// SetDeletedAt sets the "deleted_at" field.
func (bsrc *BloodStatusRecordCreate) SetDeletedAt(i int64) *BloodStatusRecordCreate {
	bsrc.mutation.SetDeletedAt(i)
	return bsrc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bsrc *BloodStatusRecordCreate) SetNillableDeletedAt(i *int64) *BloodStatusRecordCreate {
	if i != nil {
		bsrc.SetDeletedAt(*i)
	}
	return bsrc
}

// SetID sets the "id" field.
func (bsrc *BloodStatusRecordCreate) SetID(u uuid.UUID) *BloodStatusRecordCreate {
	bsrc.mutation.SetID(u)
	return bsrc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bsrc *BloodStatusRecordCreate) SetNillableID(u *uuid.UUID) *BloodStatusRecordCreate {
	if u != nil {
		bsrc.SetID(*u)
	}
	return bsrc
}

// Mutation returns the BloodStatusRecordMutation object of the builder.
func (bsrc *BloodStatusRecordCreate) Mutation() *BloodStatusRecordMutation {
	return bsrc.mutation
}

// Save creates the BloodStatusRecord in the database.
func (bsrc *BloodStatusRecordCreate) Save(ctx context.Context) (*BloodStatusRecord, error) {
	bsrc.defaults()
	return withHooks(ctx, bsrc.sqlSave, bsrc.mutation, bsrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bsrc *BloodStatusRecordCreate) SaveX(ctx context.Context) *BloodStatusRecord {
	v, err := bsrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bsrc *BloodStatusRecordCreate) Exec(ctx context.Context) error {
	_, err := bsrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsrc *BloodStatusRecordCreate) ExecX(ctx context.Context) {
	if err := bsrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bsrc *BloodStatusRecordCreate) defaults() {
	if _, ok := bsrc.mutation.DeletedAt(); !ok {
		v := bloodstatusrecord.DefaultDeletedAt
		bsrc.mutation.SetDeletedAt(v)
	}
	if _, ok := bsrc.mutation.ID(); !ok {
		v := bloodstatusrecord.DefaultID()
		bsrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bsrc *BloodStatusRecordCreate) check() error {
	if _, ok := bsrc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "BloodStatusRecord.user_id"`)}
	}
	if _, ok := bsrc.mutation.RecordDate(); !ok {
		return &ValidationError{Name: "record_date", err: errors.New(`ent: missing required field "BloodStatusRecord.record_date"`)}
	}
	if _, ok := bsrc.mutation.TimeOfDay(); !ok {
		return &ValidationError{Name: "time_of_day", err: errors.New(`ent: missing required field "BloodStatusRecord.time_of_day"`)}
	}
	if v, ok := bsrc.mutation.TimeOfDay(); ok {
		if err := bloodstatusrecord.TimeOfDayValidator(v); err != nil {
			return &ValidationError{Name: "time_of_day", err: fmt.Errorf(`ent: validator failed for field "BloodStatusRecord.time_of_day": %w`, err)}
		}
	}
	if _, ok := bsrc.mutation.BeforeAfterMeals(); !ok {
		return &ValidationError{Name: "before_after_meals", err: errors.New(`ent: missing required field "BloodStatusRecord.before_after_meals"`)}
	}
	if v, ok := bsrc.mutation.BeforeAfterMeals(); ok {
		if err := bloodstatusrecord.BeforeAfterMealsValidator(v); err != nil {
			return &ValidationError{Name: "before_after_meals", err: fmt.Errorf(`ent: validator failed for field "BloodStatusRecord.before_after_meals": %w`, err)}
		}
	}
	if _, ok := bsrc.mutation.SystolicPressure(); !ok {
		return &ValidationError{Name: "systolic_pressure", err: errors.New(`ent: missing required field "BloodStatusRecord.systolic_pressure"`)}
	}
	if _, ok := bsrc.mutation.DiastolicPressure(); !ok {
		return &ValidationError{Name: "diastolic_pressure", err: errors.New(`ent: missing required field "BloodStatusRecord.diastolic_pressure"`)}
	}
	if _, ok := bsrc.mutation.Pulse(); !ok {
		return &ValidationError{Name: "pulse", err: errors.New(`ent: missing required field "BloodStatusRecord.pulse"`)}
	}
	if _, ok := bsrc.mutation.Mood(); !ok {
		return &ValidationError{Name: "mood", err: errors.New(`ent: missing required field "BloodStatusRecord.mood"`)}
	}
	if v, ok := bsrc.mutation.Mood(); ok {
		if err := bloodstatusrecord.MoodValidator(v); err != nil {
			return &ValidationError{Name: "mood", err: fmt.Errorf(`ent: validator failed for field "BloodStatusRecord.mood": %w`, err)}
		}
	}
	if _, ok := bsrc.mutation.StatusSummary(); !ok {
		return &ValidationError{Name: "status_summary", err: errors.New(`ent: missing required field "BloodStatusRecord.status_summary"`)}
	}
	if v, ok := bsrc.mutation.StatusSummary(); ok {
		if err := bloodstatusrecord.StatusSummaryValidator(v); err != nil {
			return &ValidationError{Name: "status_summary", err: fmt.Errorf(`ent: validator failed for field "BloodStatusRecord.status_summary": %w`, err)}
		}
	}
	if _, ok := bsrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BloodStatusRecord.created_at"`)}
	}
	if _, ok := bsrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BloodStatusRecord.updated_at"`)}
	}
	if _, ok := bsrc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "BloodStatusRecord.deleted_at"`)}
	}
	return nil
}

func (bsrc *BloodStatusRecordCreate) sqlSave(ctx context.Context) (*BloodStatusRecord, error) {
	if err := bsrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bsrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bsrc.driver, _spec); err != nil {
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
	bsrc.mutation.id = &_node.ID
	bsrc.mutation.done = true
	return _node, nil
}

func (bsrc *BloodStatusRecordCreate) createSpec() (*BloodStatusRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &BloodStatusRecord{config: bsrc.config}
		_spec = sqlgraph.NewCreateSpec(bloodstatusrecord.Table, sqlgraph.NewFieldSpec(bloodstatusrecord.FieldID, field.TypeUUID))
	)
	if id, ok := bsrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bsrc.mutation.UserID(); ok {
		_spec.SetField(bloodstatusrecord.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := bsrc.mutation.RecordDate(); ok {
		_spec.SetField(bloodstatusrecord.FieldRecordDate, field.TypeUint, value)
		_node.RecordDate = value
	}
	if value, ok := bsrc.mutation.TimeOfDay(); ok {
		_spec.SetField(bloodstatusrecord.FieldTimeOfDay, field.TypeEnum, value)
		_node.TimeOfDay = value
	}
	if value, ok := bsrc.mutation.BeforeAfterMeals(); ok {
		_spec.SetField(bloodstatusrecord.FieldBeforeAfterMeals, field.TypeEnum, value)
		_node.BeforeAfterMeals = value
	}
	if value, ok := bsrc.mutation.SystolicPressure(); ok {
		_spec.SetField(bloodstatusrecord.FieldSystolicPressure, field.TypeUint8, value)
		_node.SystolicPressure = value
	}
	if value, ok := bsrc.mutation.DiastolicPressure(); ok {
		_spec.SetField(bloodstatusrecord.FieldDiastolicPressure, field.TypeUint8, value)
		_node.DiastolicPressure = value
	}
	if value, ok := bsrc.mutation.Pulse(); ok {
		_spec.SetField(bloodstatusrecord.FieldPulse, field.TypeUint8, value)
		_node.Pulse = value
	}
	if value, ok := bsrc.mutation.Mood(); ok {
		_spec.SetField(bloodstatusrecord.FieldMood, field.TypeEnum, value)
		_node.Mood = value
	}
	if value, ok := bsrc.mutation.StatusSummary(); ok {
		_spec.SetField(bloodstatusrecord.FieldStatusSummary, field.TypeEnum, value)
		_node.StatusSummary = value
	}
	if value, ok := bsrc.mutation.CreatedAt(); ok {
		_spec.SetField(bloodstatusrecord.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := bsrc.mutation.UpdatedAt(); ok {
		_spec.SetField(bloodstatusrecord.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := bsrc.mutation.DeletedAt(); ok {
		_spec.SetField(bloodstatusrecord.FieldDeletedAt, field.TypeInt64, value)
		_node.DeletedAt = value
	}
	return _node, _spec
}

// BloodStatusRecordCreateBulk is the builder for creating many BloodStatusRecord entities in bulk.
type BloodStatusRecordCreateBulk struct {
	config
	err      error
	builders []*BloodStatusRecordCreate
}

// Save creates the BloodStatusRecord entities in the database.
func (bsrcb *BloodStatusRecordCreateBulk) Save(ctx context.Context) ([]*BloodStatusRecord, error) {
	if bsrcb.err != nil {
		return nil, bsrcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bsrcb.builders))
	nodes := make([]*BloodStatusRecord, len(bsrcb.builders))
	mutators := make([]Mutator, len(bsrcb.builders))
	for i := range bsrcb.builders {
		func(i int, root context.Context) {
			builder := bsrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BloodStatusRecordMutation)
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
					_, err = mutators[i+1].Mutate(root, bsrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bsrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bsrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bsrcb *BloodStatusRecordCreateBulk) SaveX(ctx context.Context) []*BloodStatusRecord {
	v, err := bsrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bsrcb *BloodStatusRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := bsrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsrcb *BloodStatusRecordCreateBulk) ExecX(ctx context.Context) {
	if err := bsrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
