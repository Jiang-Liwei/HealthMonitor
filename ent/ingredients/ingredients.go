// Code generated by ent, DO NOT EDIT.

package ingredients

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ingredients type in the database.
	Label = "ingredients"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEffect holds the string denoting the effect field in the database.
	FieldEffect = "effect"
	// EdgeFood holds the string denoting the food edge name in mutations.
	EdgeFood = "food"
	// Table holds the table name of the ingredients in the database.
	Table = "ingredients"
	// FoodTable is the table that holds the food relation/edge.
	FoodTable = "food_ingredients"
	// FoodInverseTable is the table name for the FoodIngredients entity.
	// It exists in this package in order to avoid circular dependency with the "foodingredients" package.
	FoodInverseTable = "food_ingredients"
	// FoodColumn is the table column denoting the food relation/edge.
	FoodColumn = "ingredients_food"
)

// Columns holds all SQL columns for ingredients fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEffect,
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

// Effect defines the type for the "effect" enum field.
type Effect string

// Effect values.
const (
	EffectBeneficial Effect = "beneficial"
	EffectNeutral    Effect = "neutral"
	EffectHarmful    Effect = "harmful"
)

func (e Effect) String() string {
	return string(e)
}

// EffectValidator is a validator for the "effect" field enum values. It is called by the builders before save.
func EffectValidator(e Effect) error {
	switch e {
	case EffectBeneficial, EffectNeutral, EffectHarmful:
		return nil
	default:
		return fmt.Errorf("ingredients: invalid enum value for effect field: %q", e)
	}
}

// OrderOption defines the ordering options for the Ingredients queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEffect orders the results by the effect field.
func ByEffect(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEffect, opts...).ToFunc()
}

// ByFoodCount orders the results by food count.
func ByFoodCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFoodStep(), opts...)
	}
}

// ByFood orders the results by food terms.
func ByFood(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFoodStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFoodStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FoodInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FoodTable, FoodColumn),
	)
}