package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserMeal holds the schema definition for the UserMeal entity.
type UserMeal struct {
	ent.Schema
}

// Annotations of the BloodStatusRecord.
func (UserMeal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Options: "COMMENT='餐食记录表'"},
		entsql.WithComments(true),
	}
}

// Fields of the UserMeal.
func (UserMeal) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.New()).
			Comment("用户id"),
		field.Int("record_date").
			Comment("记录日期"),
		field.Enum("meal_type").
			Values("breakfast", "lunch", "dinner").
			Comment("餐点类型，早餐、午餐、晚餐"),
		field.String("description").
			Comment("餐点描述"),
	}
}

// Edges of the UserMeal.
func (UserMeal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("food", UserMealFood.Type),
	}
}
