package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// AdminRolePermission holds the schema definition for the AdminRolePermission entity.
type AdminRolePermission struct {
	ent.Schema
}

// Fields of the AdminRolePermission.
func (AdminRolePermission) Fields() []ent.Field {
	return nil
}

// Edges of the AdminRolePermission.
func (AdminRolePermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", AdminRoles.Type).
			Ref("permissions").
			Unique(),
		edge.From("permission", AdminPermission.Type).
			Ref("roles").
			Unique(),
	}
}
