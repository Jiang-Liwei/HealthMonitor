package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Food holds the schema definition for the Food entity.
type Food struct {
	ent.Schema
}

// Annotations of the Food.
func (Food) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Options: "COMMENT='食物表'"},
		entsql.WithComments(true),
	}
}

// Fields of the Food.
func (Food) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).
			Comment("食物ID"),
		field.String("name").
			Comment("食物名称"),
		field.Enum("effect").
			Values("beneficial", "neutral", "harmful").
			Comment("对血压等的影响，益处、中性、害处"),
	}
}

// Edges of the Food.
func (Food) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ingredients", FoodIngredients.Type).
			Comment("食物与食材的关联"),
		edge.To("nutrient", FoodNutrientsRelationships.Type).
			Comment("食物与营养的关联"),
		edge.To("user_meal", UserMealFood.Type).
			Comment("餐食与食物的关联"),
	}
}
