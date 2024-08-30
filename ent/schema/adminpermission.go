package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// AdminPermission holds the schema definition for the AdminPermission entity.
type AdminPermission struct {
	ent.Schema
}

// Fields of the AdminPermission.
func (AdminPermission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New),
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.String("path"),   // 例如 API 路径
		field.String("method"), // GET, POST 等
		field.Int("created_at").
			Default(int(time.Now().Unix())),
		field.Int("updated_at").
			Default(int(time.Now().Unix())).
			UpdateDefault(func() int { return int(time.Now().Unix()) }),
		field.Int("deleted_at").Optional(),
	}
}

// Edges of the AdminPermission.
func (AdminPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", AdminRolePermission.Type),
	}
}
