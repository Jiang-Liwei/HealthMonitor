// Code generated by ent, DO NOT EDIT.

package usermealfood

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the usermealfood type in the database.
	Label = "user_meal_food"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeUserMeal holds the string denoting the user_meal edge name in mutations.
	EdgeUserMeal = "user_meal"
	// EdgeFood holds the string denoting the food edge name in mutations.
	EdgeFood = "food"
	// Table holds the table name of the usermealfood in the database.
	Table = "user_meal_foods"
	// UserMealTable is the table that holds the user_meal relation/edge.
	UserMealTable = "user_meal_foods"
	// UserMealInverseTable is the table name for the UserMeal entity.
	// It exists in this package in order to avoid circular dependency with the "usermeal" package.
	UserMealInverseTable = "user_meals"
	// UserMealColumn is the table column denoting the user_meal relation/edge.
	UserMealColumn = "user_meal_food"
	// FoodTable is the table that holds the food relation/edge.
	FoodTable = "user_meal_foods"
	// FoodInverseTable is the table name for the Food entity.
	// It exists in this package in order to avoid circular dependency with the "food" package.
	FoodInverseTable = "foods"
	// FoodColumn is the table column denoting the food relation/edge.
	FoodColumn = "food_user_meal"
)

// Columns holds all SQL columns for usermealfood fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_meal_foods"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"food_user_meal",
	"user_meal_food",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the UserMealFood queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserMealField orders the results by user_meal field.
func ByUserMealField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserMealStep(), sql.OrderByField(field, opts...))
	}
}

// ByFoodField orders the results by food field.
func ByFoodField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFoodStep(), sql.OrderByField(field, opts...))
	}
}
func newUserMealStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserMealInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserMealTable, UserMealColumn),
	)
}
func newFoodStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FoodInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FoodTable, FoodColumn),
	)
}