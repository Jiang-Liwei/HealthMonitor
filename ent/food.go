// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthmonitor/ent/food"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Food is the model entity for the Food schema.
type Food struct {
	config `json:"-"`
	// ID of the ent.
	// 食物ID
	ID uuid.UUID `json:"id,omitempty"`
	// 食物名称
	Name string `json:"name,omitempty"`
	// 对血压等的影响，益处、中性、害处
	Effect food.Effect `json:"effect,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FoodQuery when eager-loading is set.
	Edges        FoodEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FoodEdges holds the relations/edges for other nodes in the graph.
type FoodEdges struct {
	// 食物与食材的关联
	Ingredients []*FoodIngredients `json:"ingredients,omitempty"`
	// 食物与营养的关联
	Nutrient []*FoodNutrientsRelationships `json:"nutrient,omitempty"`
	// 餐食与食物的关联
	UserMeal []*UserMealFood `json:"user_meal,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// IngredientsOrErr returns the Ingredients value or an error if the edge
// was not loaded in eager-loading.
func (e FoodEdges) IngredientsOrErr() ([]*FoodIngredients, error) {
	if e.loadedTypes[0] {
		return e.Ingredients, nil
	}
	return nil, &NotLoadedError{edge: "ingredients"}
}

// NutrientOrErr returns the Nutrient value or an error if the edge
// was not loaded in eager-loading.
func (e FoodEdges) NutrientOrErr() ([]*FoodNutrientsRelationships, error) {
	if e.loadedTypes[1] {
		return e.Nutrient, nil
	}
	return nil, &NotLoadedError{edge: "nutrient"}
}

// UserMealOrErr returns the UserMeal value or an error if the edge
// was not loaded in eager-loading.
func (e FoodEdges) UserMealOrErr() ([]*UserMealFood, error) {
	if e.loadedTypes[2] {
		return e.UserMeal, nil
	}
	return nil, &NotLoadedError{edge: "user_meal"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Food) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case food.FieldName, food.FieldEffect:
			values[i] = new(sql.NullString)
		case food.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Food fields.
func (f *Food) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case food.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case food.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case food.FieldEffect:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field effect", values[i])
			} else if value.Valid {
				f.Effect = food.Effect(value.String)
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Food.
// This includes values selected through modifiers, order, etc.
func (f *Food) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryIngredients queries the "ingredients" edge of the Food entity.
func (f *Food) QueryIngredients() *FoodIngredientsQuery {
	return NewFoodClient(f.config).QueryIngredients(f)
}

// QueryNutrient queries the "nutrient" edge of the Food entity.
func (f *Food) QueryNutrient() *FoodNutrientsRelationshipsQuery {
	return NewFoodClient(f.config).QueryNutrient(f)
}

// QueryUserMeal queries the "user_meal" edge of the Food entity.
func (f *Food) QueryUserMeal() *UserMealFoodQuery {
	return NewFoodClient(f.config).QueryUserMeal(f)
}

// Update returns a builder for updating this Food.
// Note that you need to call Food.Unwrap() before calling this method if this Food
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Food) Update() *FoodUpdateOne {
	return NewFoodClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Food entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Food) Unwrap() *Food {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Food is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Food) String() string {
	var builder strings.Builder
	builder.WriteString("Food(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("effect=")
	builder.WriteString(fmt.Sprintf("%v", f.Effect))
	builder.WriteByte(')')
	return builder.String()
}

// Foods is a parsable slice of Food.
type Foods []*Food
