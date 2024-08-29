package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Ingredients holds the schema definition for the Ingredients entity.
type Ingredients struct {
	ent.Schema
}

// Annotations of the Ingredients.
func (Ingredients) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Options: "COMMENT='食材表'"},
	}
}

// Fields of the Ingredients.
func (Ingredients) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).
			Comment("食材ID"),
		field.String("name").
			Comment("食材名称"),
		field.Enum("effect").
			Values("beneficial", "neutral", "harmful").
			Comment("对血压等的影响，益处、中性、害处"),
	}
}

// Edges of the Ingredients.
func (Ingredients) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("food", FoodIngredients.Type),
	}
}
