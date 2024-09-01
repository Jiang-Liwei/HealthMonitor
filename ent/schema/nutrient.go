package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Nutrient holds the schema definition for the Nutrient entity.
type Nutrient struct {
	ent.Schema
}

// Annotations of the BloodStatusRecord.
func (Nutrient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Options: "COMMENT='营养表'"},
		entsql.WithComments(true),
	}
}

// Fields of the Nutrient.
func (Nutrient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).
			Comment("营养ID"),
		field.String("name").
			Comment("营养名称"),
		field.Enum("effect").
			Values("beneficial", "neutral", "harmful").
			Comment("对血压等的影响，益处、中性、害处"),
	}
}

// Edges of the Nutrient.
func (Nutrient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("food", FoodNutrientsRelationships.Type),
	}
}
