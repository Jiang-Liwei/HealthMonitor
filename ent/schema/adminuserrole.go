package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// AdminUserRole holds the schema definition for the AdminUserRole entity.
type AdminUserRole struct {
	ent.Schema
}

// Fields of the AdminUserRole.
func (AdminUserRole) Fields() []ent.Field {
	return nil
}

// Edges of the AdminUserRole.
func (AdminUserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", AdminUser.Type).
			Ref("roles").
			Unique(),
		edge.From("role", AdminRoles.Type).
			Ref("users").
			Unique(),
	}
}
