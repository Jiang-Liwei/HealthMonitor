package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// AdminMenu holds the schema definition for the AdminMenu entity.
type AdminMenu struct {
	ent.Schema
}

// Fields of the AdminMenu.
func (AdminMenu) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New),
		field.String("name"),
		field.String("icon").Optional(), // 菜单图标
		field.String("path"),            // 前端路由路径
		field.Int("order").Default(0),   // 菜单排序
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"mysql": "TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP",
		}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			"mysql": "TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
		}),
	}
}

// Edges of the AdminMenu.
func (AdminMenu) Edges() []ent.Edge {
	return nil
}
