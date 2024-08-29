package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// FoodNutrientsRelationships holds the schema definition for the FoodNutrientsRelationships entity.
type FoodNutrientsRelationships struct {
	ent.Schema
}

// Annotations of the FoodNutrients.
func (FoodNutrientsRelationships) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Options: "COMMENT='食物与营养的关联表'"},
	}
}

// Fields of the FoodNutrients.
func (FoodNutrientsRelationships) Fields() []ent.Field {
	return nil
}

// Edges of the FoodNutrients.
func (FoodNutrientsRelationships) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("food", Food.Type).
			Ref("nutrient").
			Unique().
			Comment("关联食物"),
		edge.From("nutrient", Nutrient.Type).
			Ref("food").
			Unique().
			Comment("关联营养"),
	}
}
