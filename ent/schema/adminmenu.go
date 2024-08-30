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
		field.String("icon").Optional(),  // 菜单图标
		field.String("path"),             // 前端路由路径
		field.Uint16("order").Default(0), // 菜单排序
		field.Int("created_at").Default(int(time.Now().Unix())),
		field.Int("updated_at").
			Default(int(time.Now().Unix())).
			UpdateDefault(func() int { return int(time.Now().Unix()) }), // 更新时间
		field.Int("deleted_at").Optional(),
	}
}

// Edges of the AdminMenu.
func (AdminMenu) Edges() []ent.Edge {
	return nil
}
