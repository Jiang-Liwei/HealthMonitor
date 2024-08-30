package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// AdminUser holds the schema definition for the AdminUser entity.
type AdminUser struct {
	ent.Schema
}

// Fields of the AdminUser.
func (AdminUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New),
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("password_hash"),
		field.Bool("is_active").Default(true),
		field.Int("last_login_at").Optional(), // 上次登录时间
		field.Int("jwt_issued_at").Optional(), // JWT最近签发时间
		field.Int("created_at").
			Default(int(time.Now().Unix())), // 创建时间，默认当前时间
		field.Int("updated_at").
			Default(int(time.Now().Unix())).
			UpdateDefault(func() int { return int(time.Now().Unix()) }), // 更新时间，每次更新时自动更新为当前时间
		field.Int("deleted_at").Optional(),
	}
}

// Edges of the AdminUser.
func (AdminUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", AdminUserRole.Type),
		edge.To("logs", AdminLog.Type),
	}
}
