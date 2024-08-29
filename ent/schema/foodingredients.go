package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// FoodIngredients holds the schema definition for the FoodIngredients entity.
type FoodIngredients struct {
	ent.Schema
}

// Annotations of the FoodIngredients.
func (FoodIngredients) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Options: "COMMENT='食物与食材的关联表'"},
	}
}

// Fields of the FoodIngredients.
func (FoodIngredients) Fields() []ent.Field {
	return nil
}

// Edges of the FoodIngredients.
func (FoodIngredients) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("food", Food.Type).
			Ref("ingredients").
			Unique().
			Comment("关联的食物"),
		edge.From("ingredient", Ingredients.Type).
			Ref("food").
			Unique().
			Comment("关联的食材"),
	}
}
