// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/bloodstatusrecord"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// BloodStatusRecord is the model entity for the BloodStatusRecord schema.
type BloodStatusRecord struct {
	config `json:"-"`
	// ID of the ent.
	// 记录的唯一标识符
	ID uuid.UUID `json:"id,omitempty"`
	// 记录日期
	RecordDate time.Time `json:"record_date,omitempty"`
	// 记录时间段，早晨、中午、晚上
	TimeOfDay bloodstatusrecord.TimeOfDay `json:"time_of_day,omitempty"`
	// 餐前餐后，前、后
	BeforeAfterMeals bloodstatusrecord.BeforeAfterMeals `json:"before_after_meals,omitempty"`
	// 收缩压
	SystolicPressure float64 `json:"systolic_pressure,omitempty"`
	// 舒张压
	DiastolicPressure float64 `json:"diastolic_pressure,omitempty"`
	// 脉搏
	Pulse        float64 `json:"pulse,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BloodStatusRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bloodstatusrecord.FieldSystolicPressure, bloodstatusrecord.FieldDiastolicPressure, bloodstatusrecord.FieldPulse:
			values[i] = new(sql.NullFloat64)
		case bloodstatusrecord.FieldTimeOfDay, bloodstatusrecord.FieldBeforeAfterMeals:
			values[i] = new(sql.NullString)
		case bloodstatusrecord.FieldRecordDate:
			values[i] = new(sql.NullTime)
		case bloodstatusrecord.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BloodStatusRecord fields.
func (bsr *BloodStatusRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bloodstatusrecord.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				bsr.ID = *value
			}
		case bloodstatusrecord.FieldRecordDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field record_date", values[i])
			} else if value.Valid {
				bsr.RecordDate = value.Time
			}
		case bloodstatusrecord.FieldTimeOfDay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field time_of_day", values[i])
			} else if value.Valid {
				bsr.TimeOfDay = bloodstatusrecord.TimeOfDay(value.String)
			}
		case bloodstatusrecord.FieldBeforeAfterMeals:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field before_after_meals", values[i])
			} else if value.Valid {
				bsr.BeforeAfterMeals = bloodstatusrecord.BeforeAfterMeals(value.String)
			}
		case bloodstatusrecord.FieldSystolicPressure:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field systolic_pressure", values[i])
			} else if value.Valid {
				bsr.SystolicPressure = value.Float64
			}
		case bloodstatusrecord.FieldDiastolicPressure:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field diastolic_pressure", values[i])
			} else if value.Valid {
				bsr.DiastolicPressure = value.Float64
			}
		case bloodstatusrecord.FieldPulse:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field pulse", values[i])
			} else if value.Valid {
				bsr.Pulse = value.Float64
			}
		default:
			bsr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BloodStatusRecord.
// This includes values selected through modifiers, order, etc.
func (bsr *BloodStatusRecord) Value(name string) (ent.Value, error) {
	return bsr.selectValues.Get(name)
}

// Update returns a builder for updating this BloodStatusRecord.
// Note that you need to call BloodStatusRecord.Unwrap() before calling this method if this BloodStatusRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (bsr *BloodStatusRecord) Update() *BloodStatusRecordUpdateOne {
	return NewBloodStatusRecordClient(bsr.config).UpdateOne(bsr)
}

// Unwrap unwraps the BloodStatusRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bsr *BloodStatusRecord) Unwrap() *BloodStatusRecord {
	_tx, ok := bsr.config.driver.(*txDriver)
	if !ok {
		panic("ent: BloodStatusRecord is not a transactional entity")
	}
	bsr.config.driver = _tx.drv
	return bsr
}

// String implements the fmt.Stringer.
func (bsr *BloodStatusRecord) String() string {
	var builder strings.Builder
	builder.WriteString("BloodStatusRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bsr.ID))
	builder.WriteString("record_date=")
	builder.WriteString(bsr.RecordDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("time_of_day=")
	builder.WriteString(fmt.Sprintf("%v", bsr.TimeOfDay))
	builder.WriteString(", ")
	builder.WriteString("before_after_meals=")
	builder.WriteString(fmt.Sprintf("%v", bsr.BeforeAfterMeals))
	builder.WriteString(", ")
	builder.WriteString("systolic_pressure=")
	builder.WriteString(fmt.Sprintf("%v", bsr.SystolicPressure))
	builder.WriteString(", ")
	builder.WriteString("diastolic_pressure=")
	builder.WriteString(fmt.Sprintf("%v", bsr.DiastolicPressure))
	builder.WriteString(", ")
	builder.WriteString("pulse=")
	builder.WriteString(fmt.Sprintf("%v", bsr.Pulse))
	builder.WriteByte(')')
	return builder.String()
}

// BloodStatusRecords is a parsable slice of BloodStatusRecord.
type BloodStatusRecords []*BloodStatusRecord