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
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"mysql": "TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP",
		}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			"mysql": "TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
		}),
	}
}

// Edges of the AdminPermission.
func (AdminPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", AdminRolePermission.Type),
	}
}
