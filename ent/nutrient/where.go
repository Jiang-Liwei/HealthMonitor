// Code generated by ent, DO NOT EDIT.

package nutrient

import (
	"HealthMonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldContainsFold(FieldName, v))
}

// EffectEQ applies the EQ predicate on the "effect" field.
func EffectEQ(v Effect) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldEQ(FieldEffect, v))
}

// EffectNEQ applies the NEQ predicate on the "effect" field.
func EffectNEQ(v Effect) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldNEQ(FieldEffect, v))
}

// EffectIn applies the In predicate on the "effect" field.
func EffectIn(vs ...Effect) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldIn(FieldEffect, vs...))
}

// EffectNotIn applies the NotIn predicate on the "effect" field.
func EffectNotIn(vs ...Effect) predicate.Nutrient {
	return predicate.Nutrient(sql.FieldNotIn(FieldEffect, vs...))
}

// HasFood applies the HasEdge predicate on the "food" edge.
func HasFood() predicate.Nutrient {
	return predicate.Nutrient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FoodTable, FoodColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFoodWith applies the HasEdge predicate on the "food" edge with a given conditions (other predicates).
func HasFoodWith(preds ...predicate.FoodNutrients) predicate.Nutrient {
	return predicate.Nutrient(func(s *sql.Selector) {
		step := newFoodStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Nutrient) predicate.Nutrient {
	return predicate.Nutrient(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Nutrient) predicate.Nutrient {
	return predicate.Nutrient(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Nutrient) predicate.Nutrient {
	return predicate.Nutrient(sql.NotPredicates(p))
}
