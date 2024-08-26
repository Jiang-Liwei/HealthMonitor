package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// AdminRoles holds the schema definition for the AdminRoles entity.
type AdminRoles struct {
	ent.Schema
}

// Fields of the AdminRoles.
func (AdminRoles) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New),
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"mysql": "TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP",
		}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			"mysql": "TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
		}),
	}
}

// Edges of the AdminRoles.
func (AdminRoles) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", AdminRolePermission.Type),
		edge.To("users", AdminUserRole.Type),
	}
}
