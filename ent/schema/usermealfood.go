package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// UserMealFood holds the schema definition for the UserMealFood entity.
type UserMealFood struct {
	ent.Schema
}

// Annotations of the UserMealFood.
func (UserMealFood) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Options: "COMMENT=餐食与食物的关联表"},
	}
}

// Fields of the UserMealFood.
func (UserMealFood) Fields() []ent.Field {
	return nil
}

// Edges of the UserMealFood.
func (UserMealFood) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_meal", UserMeal.Type).
			Ref("food").
			Unique().
			Comment("关联的餐食记录"),
		edge.From("food", Food.Type).
			Ref("user_meal").
			Unique().
			Comment("关联的食物"),
	}
}
