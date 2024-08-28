// Code generated by ent, DO NOT EDIT.

package ent

import (
	"HealthMonitor/ent/food"
	"HealthMonitor/ent/foodingredients"
	"HealthMonitor/ent/ingredients"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// FoodIngredients is the model entity for the FoodIngredients schema.
type FoodIngredients struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FoodIngredientsQuery when eager-loading is set.
	Edges            FoodIngredientsEdges `json:"edges"`
	food_ingredients *uuid.UUID
	ingredients_food *uuid.UUID
	selectValues     sql.SelectValues
}

// FoodIngredientsEdges holds the relations/edges for other nodes in the graph.
type FoodIngredientsEdges struct {
	// 关联的食物
	Food *Food `json:"food,omitempty"`
	// 关联的食材
	Ingredient *Ingredients `json:"ingredient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FoodOrErr returns the Food value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FoodIngredientsEdges) FoodOrErr() (*Food, error) {
	if e.Food != nil {
		return e.Food, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: food.Label}
	}
	return nil, &NotLoadedError{edge: "food"}
}

// IngredientOrErr returns the Ingredient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FoodIngredientsEdges) IngredientOrErr() (*Ingredients, error) {
	if e.Ingredient != nil {
		return e.Ingredient, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: ingredients.Label}
	}
	return nil, &NotLoadedError{edge: "ingredient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FoodIngredients) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case foodingredients.FieldID:
			values[i] = new(sql.NullInt64)
		case foodingredients.ForeignKeys[0]: // food_ingredients
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case foodingredients.ForeignKeys[1]: // ingredients_food
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FoodIngredients fields.
func (fi *FoodIngredients) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case foodingredients.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fi.ID = int(value.Int64)
		case foodingredients.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field food_ingredients", values[i])
			} else if value.Valid {
				fi.food_ingredients = new(uuid.UUID)
				*fi.food_ingredients = *value.S.(*uuid.UUID)
			}
		case foodingredients.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field ingredients_food", values[i])
			} else if value.Valid {
				fi.ingredients_food = new(uuid.UUID)
				*fi.ingredients_food = *value.S.(*uuid.UUID)
			}
		default:
			fi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FoodIngredients.
// This includes values selected through modifiers, order, etc.
func (fi *FoodIngredients) Value(name string) (ent.Value, error) {
	return fi.selectValues.Get(name)
}

// QueryFood queries the "food" edge of the FoodIngredients entity.
func (fi *FoodIngredients) QueryFood() *FoodQuery {
	return NewFoodIngredientsClient(fi.config).QueryFood(fi)
}

// QueryIngredient queries the "ingredient" edge of the FoodIngredients entity.
func (fi *FoodIngredients) QueryIngredient() *IngredientsQuery {
	return NewFoodIngredientsClient(fi.config).QueryIngredient(fi)
}

// Update returns a builder for updating this FoodIngredients.
// Note that you need to call FoodIngredients.Unwrap() before calling this method if this FoodIngredients
// was returned from a transaction, and the transaction was committed or rolled back.
func (fi *FoodIngredients) Update() *FoodIngredientsUpdateOne {
	return NewFoodIngredientsClient(fi.config).UpdateOne(fi)
}

// Unwrap unwraps the FoodIngredients entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fi *FoodIngredients) Unwrap() *FoodIngredients {
	_tx, ok := fi.config.driver.(*txDriver)
	if !ok {
		panic("ent: FoodIngredients is not a transactional entity")
	}
	fi.config.driver = _tx.drv
	return fi
}

// String implements the fmt.Stringer.
func (fi *FoodIngredients) String() string {
	var builder strings.Builder
	builder.WriteString("FoodIngredients(")
	builder.WriteString(fmt.Sprintf("id=%v", fi.ID))
	builder.WriteByte(')')
	return builder.String()
}

// FoodIngredientsSlice is a parsable slice of FoodIngredients.
type FoodIngredientsSlice []*FoodIngredients
