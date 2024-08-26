package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AdminJWTBlacklist holds the schema definition for the AdminJWTBlacklist entity.
type AdminJWTBlacklist struct {
	ent.Schema
}

// Fields of the AdminJWTBlacklist.
func (AdminJWTBlacklist) Fields() []ent.Field {
	return []ent.Field{
		field.String("jti").Unique(), // JWT ID (Token的唯一标识)
		field.Time("expires_at").SchemaType(map[string]string{
			"mysql": "TIMESTAMP(0)",
		}), // JWT的过期时间
		field.Time("revoked_at").Optional().SchemaType(map[string]string{
			"mysql": "TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP",
		}), // 撤销时间
	}
}

// Edges of the AdminJWTBlacklist.
func (AdminJWTBlacklist) Edges() []ent.Edge {
	return nil
}
