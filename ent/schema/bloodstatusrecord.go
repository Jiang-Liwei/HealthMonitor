package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BloodStatusRecord holds the schema definition for the BloodStatusRecord entity.
type BloodStatusRecord struct {
	ent.Schema
}

// Annotations of the BloodStatusRecord.
func (BloodStatusRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Options: "COMMENT='血压状态记录表'"},
	}
}

// Fields of the BloodStatusRecord.
func (BloodStatusRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).
			Comment("记录的唯一标识符"),
		field.UUID("user_id", uuid.New()).
			Comment("用户id"),
		field.Time("record_date").
			Comment("记录日期"),
		field.Enum("time_of_day").
			Values("morning", "noon", "evening").
			Comment("记录时间段，早晨、中午、晚上"),
		field.Enum("before_after_meals").
			Values("before", "after").
			Comment("餐前餐后，前、后"),
		field.Float("systolic_pressure").
			Comment("收缩压"),
		field.Float("diastolic_pressure").
			Comment("舒张压"),
		field.Float("pulse").
			Comment("脉搏"),
	}
}

// Edges of the BloodStatusRecord.
func (BloodStatusRecord) Edges() []ent.Edge {
	return nil
}
