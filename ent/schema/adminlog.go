package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// AdminLog holds the schema definition for the AdminLog entity.
type AdminLog struct {
	ent.Schema
}

// Fields of the AdminLog.
func (AdminLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New),
		field.String("action"),
		field.String("ip_address"),
		field.Int("created_at").Default(int(time.Now().Unix())),
	}
}

// Edges of the AdminLog.
func (AdminLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", AdminUser.Type).Ref("logs").Unique(),
	}
}
