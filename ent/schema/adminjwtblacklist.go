package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// AdminJWTBlacklist holds the schema definition for the AdminJWTBlacklist entity.
type AdminJWTBlacklist struct {
	ent.Schema
}

// Fields of the AdminJWTBlacklist.
func (AdminJWTBlacklist) Fields() []ent.Field {
	return []ent.Field{
		field.String("jti").Unique(), // JWT ID (Token的唯一标识)
		field.Int("expires_at"),      // JWT的过期时间
		field.Int("revoked_at").
			Default(int(time.Now().Unix())), // 撤销时间
	}
}

// Edges of the AdminJWTBlacklist.
func (AdminJWTBlacklist) Edges() []ent.Edge {
	return nil
}
