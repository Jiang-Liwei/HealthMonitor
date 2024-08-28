// Code generated by ent, DO NOT EDIT.

package bloodstatusrecord

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the bloodstatusrecord type in the database.
	Label = "blood_status_record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRecordDate holds the string denoting the record_date field in the database.
	FieldRecordDate = "record_date"
	// FieldTimeOfDay holds the string denoting the time_of_day field in the database.
	FieldTimeOfDay = "time_of_day"
	// FieldBeforeAfterMeals holds the string denoting the before_after_meals field in the database.
	FieldBeforeAfterMeals = "before_after_meals"
	// FieldSystolicPressure holds the string denoting the systolic_pressure field in the database.
	FieldSystolicPressure = "systolic_pressure"
	// FieldDiastolicPressure holds the string denoting the diastolic_pressure field in the database.
	FieldDiastolicPressure = "diastolic_pressure"
	// FieldPulse holds the string denoting the pulse field in the database.
	FieldPulse = "pulse"
	// Table holds the table name of the bloodstatusrecord in the database.
	Table = "blood_status_records"
)

// Columns holds all SQL columns for bloodstatusrecord fields.
var Columns = []string{
	FieldID,
	FieldRecordDate,
	FieldTimeOfDay,
	FieldBeforeAfterMeals,
	FieldSystolicPressure,
	FieldDiastolicPressure,
	FieldPulse,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// TimeOfDay defines the type for the "time_of_day" enum field.
type TimeOfDay string

// TimeOfDay values.
const (
	TimeOfDayMorning TimeOfDay = "morning"
	TimeOfDayNoon    TimeOfDay = "noon"
	TimeOfDayEvening TimeOfDay = "evening"
)

func (tod TimeOfDay) String() string {
	return string(tod)
}

// TimeOfDayValidator is a validator for the "time_of_day" field enum values. It is called by the builders before save.
func TimeOfDayValidator(tod TimeOfDay) error {
	switch tod {
	case TimeOfDayMorning, TimeOfDayNoon, TimeOfDayEvening:
		return nil
	default:
		return fmt.Errorf("bloodstatusrecord: invalid enum value for time_of_day field: %q", tod)
	}
}

// BeforeAfterMeals defines the type for the "before_after_meals" enum field.
type BeforeAfterMeals string

// BeforeAfterMeals values.
const (
	BeforeAfterMealsBefore BeforeAfterMeals = "before"
	BeforeAfterMealsAfter  BeforeAfterMeals = "after"
)

func (bam BeforeAfterMeals) String() string {
	return string(bam)
}

// BeforeAfterMealsValidator is a validator for the "before_after_meals" field enum values. It is called by the builders before save.
func BeforeAfterMealsValidator(bam BeforeAfterMeals) error {
	switch bam {
	case BeforeAfterMealsBefore, BeforeAfterMealsAfter:
		return nil
	default:
		return fmt.Errorf("bloodstatusrecord: invalid enum value for before_after_meals field: %q", bam)
	}
}

// OrderOption defines the ordering options for the BloodStatusRecord queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRecordDate orders the results by the record_date field.
func ByRecordDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecordDate, opts...).ToFunc()
}

// ByTimeOfDay orders the results by the time_of_day field.
func ByTimeOfDay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeOfDay, opts...).ToFunc()
}

// ByBeforeAfterMeals orders the results by the before_after_meals field.
func ByBeforeAfterMeals(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBeforeAfterMeals, opts...).ToFunc()
}

// BySystolicPressure orders the results by the systolic_pressure field.
func BySystolicPressure(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSystolicPressure, opts...).ToFunc()
}

// ByDiastolicPressure orders the results by the diastolic_pressure field.
func ByDiastolicPressure(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDiastolicPressure, opts...).ToFunc()
}

// ByPulse orders the results by the pulse field.
func ByPulse(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPulse, opts...).ToFunc()
}
