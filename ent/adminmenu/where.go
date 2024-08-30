// Code generated by ent, DO NOT EDIT.

package adminmenu

import (
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldName, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldIcon, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldPath, v))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldOrder, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldDeletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldContainsFold(FieldName, v))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldHasSuffix(FieldIcon, v))
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIsNull(FieldIcon))
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotNull(FieldIcon))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldContainsFold(FieldIcon, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldContainsFold(FieldPath, v))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v uint16) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLTE(FieldOrder, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int) predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.AdminMenu {
	return predicate.AdminMenu(sql.FieldNotNull(FieldDeletedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AdminMenu) predicate.AdminMenu {
	return predicate.AdminMenu(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AdminMenu) predicate.AdminMenu {
	return predicate.AdminMenu(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AdminMenu) predicate.AdminMenu {
	return predicate.AdminMenu(sql.NotPredicates(p))
}
