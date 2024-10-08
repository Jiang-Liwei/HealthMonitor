// Code generated by ent, DO NOT EDIT.

package food

import (
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Food {
	return predicate.Food(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Food {
	return predicate.Food(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Food {
	return predicate.Food(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Food {
	return predicate.Food(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Food {
	return predicate.Food(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Food {
	return predicate.Food(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Food {
	return predicate.Food(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Food {
	return predicate.Food(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Food {
	return predicate.Food(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Food {
	return predicate.Food(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Food {
	return predicate.Food(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Food {
	return predicate.Food(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Food {
	return predicate.Food(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Food {
	return predicate.Food(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Food {
	return predicate.Food(sql.FieldContainsFold(FieldName, v))
}

// EffectEQ applies the EQ predicate on the "effect" field.
func EffectEQ(v Effect) predicate.Food {
	return predicate.Food(sql.FieldEQ(FieldEffect, v))
}

// EffectNEQ applies the NEQ predicate on the "effect" field.
func EffectNEQ(v Effect) predicate.Food {
	return predicate.Food(sql.FieldNEQ(FieldEffect, v))
}

// EffectIn applies the In predicate on the "effect" field.
func EffectIn(vs ...Effect) predicate.Food {
	return predicate.Food(sql.FieldIn(FieldEffect, vs...))
}

// EffectNotIn applies the NotIn predicate on the "effect" field.
func EffectNotIn(vs ...Effect) predicate.Food {
	return predicate.Food(sql.FieldNotIn(FieldEffect, vs...))
}

// HasIngredients applies the HasEdge predicate on the "ingredients" edge.
func HasIngredients() predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IngredientsTable, IngredientsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIngredientsWith applies the HasEdge predicate on the "ingredients" edge with a given conditions (other predicates).
func HasIngredientsWith(preds ...predicate.FoodIngredients) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := newIngredientsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNutrient applies the HasEdge predicate on the "nutrient" edge.
func HasNutrient() predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NutrientTable, NutrientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNutrientWith applies the HasEdge predicate on the "nutrient" edge with a given conditions (other predicates).
func HasNutrientWith(preds ...predicate.FoodNutrientsRelationships) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := newNutrientStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserMeal applies the HasEdge predicate on the "user_meal" edge.
func HasUserMeal() predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserMealTable, UserMealColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserMealWith applies the HasEdge predicate on the "user_meal" edge with a given conditions (other predicates).
func HasUserMealWith(preds ...predicate.UserMealFood) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := newUserMealStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Food) predicate.Food {
	return predicate.Food(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Food) predicate.Food {
	return predicate.Food(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Food) predicate.Food {
	return predicate.Food(sql.NotPredicates(p))
}
