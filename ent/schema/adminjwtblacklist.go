package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AdminJWTExpiredTokens holds the schema definition for the AdminJWTExpiredTokens entity.
type AdminJWTExpiredTokens struct {
	ent.Schema
}

// Fields of the AdminJWTExpiredTokens.
func (AdminJWTExpiredTokens) Fields() []ent.Field {
	return []ent.Field{
		field.String("jti").Unique(),       // JWT ID (Token的唯一标识)
		field.Int("expires_at"),            // JWT的过期时间
		field.Int("revoked_at").Optional(), // 撤销时间
	}
}

// Edges of the AdminJWTExpiredTokens.
func (AdminJWTExpiredTokens) Edges() []ent.Edge {
	return nil
}
