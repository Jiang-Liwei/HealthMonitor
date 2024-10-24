// Code generated by ent, DO NOT EDIT.

package adminjwtexpiredtokens

import (
	"healthmonitor/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldLTE(FieldID, id))
}

// Jti applies equality check predicate on the "jti" field. It's identical to JtiEQ.
func Jti(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEQ(FieldJti, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEQ(FieldExpiresAt, v))
}

// RevokedAt applies equality check predicate on the "revoked_at" field. It's identical to RevokedAtEQ.
func RevokedAt(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEQ(FieldRevokedAt, v))
}

// JtiEQ applies the EQ predicate on the "jti" field.
func JtiEQ(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEQ(FieldJti, v))
}

// JtiNEQ applies the NEQ predicate on the "jti" field.
func JtiNEQ(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNEQ(FieldJti, v))
}

// JtiIn applies the In predicate on the "jti" field.
func JtiIn(vs ...string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldIn(FieldJti, vs...))
}

// JtiNotIn applies the NotIn predicate on the "jti" field.
func JtiNotIn(vs ...string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNotIn(FieldJti, vs...))
}

// JtiGT applies the GT predicate on the "jti" field.
func JtiGT(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldGT(FieldJti, v))
}

// JtiGTE applies the GTE predicate on the "jti" field.
func JtiGTE(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldGTE(FieldJti, v))
}

// JtiLT applies the LT predicate on the "jti" field.
func JtiLT(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldLT(FieldJti, v))
}

// JtiLTE applies the LTE predicate on the "jti" field.
func JtiLTE(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldLTE(FieldJti, v))
}

// JtiContains applies the Contains predicate on the "jti" field.
func JtiContains(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldContains(FieldJti, v))
}

// JtiHasPrefix applies the HasPrefix predicate on the "jti" field.
func JtiHasPrefix(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldHasPrefix(FieldJti, v))
}

// JtiHasSuffix applies the HasSuffix predicate on the "jti" field.
func JtiHasSuffix(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldHasSuffix(FieldJti, v))
}

// JtiEqualFold applies the EqualFold predicate on the "jti" field.
func JtiEqualFold(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEqualFold(FieldJti, v))
}

// JtiContainsFold applies the ContainsFold predicate on the "jti" field.
func JtiContainsFold(v string) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldContainsFold(FieldJti, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldLTE(FieldExpiresAt, v))
}

// RevokedAtEQ applies the EQ predicate on the "revoked_at" field.
func RevokedAtEQ(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldEQ(FieldRevokedAt, v))
}

// RevokedAtNEQ applies the NEQ predicate on the "revoked_at" field.
func RevokedAtNEQ(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNEQ(FieldRevokedAt, v))
}

// RevokedAtIn applies the In predicate on the "revoked_at" field.
func RevokedAtIn(vs ...int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldIn(FieldRevokedAt, vs...))
}

// RevokedAtNotIn applies the NotIn predicate on the "revoked_at" field.
func RevokedAtNotIn(vs ...int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNotIn(FieldRevokedAt, vs...))
}

// RevokedAtGT applies the GT predicate on the "revoked_at" field.
func RevokedAtGT(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldGT(FieldRevokedAt, v))
}

// RevokedAtGTE applies the GTE predicate on the "revoked_at" field.
func RevokedAtGTE(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldGTE(FieldRevokedAt, v))
}

// RevokedAtLT applies the LT predicate on the "revoked_at" field.
func RevokedAtLT(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldLT(FieldRevokedAt, v))
}

// RevokedAtLTE applies the LTE predicate on the "revoked_at" field.
func RevokedAtLTE(v int) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldLTE(FieldRevokedAt, v))
}

// RevokedAtIsNil applies the IsNil predicate on the "revoked_at" field.
func RevokedAtIsNil() predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldIsNull(FieldRevokedAt))
}

// RevokedAtNotNil applies the NotNil predicate on the "revoked_at" field.
func RevokedAtNotNil() predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.FieldNotNull(FieldRevokedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AdminJWTExpiredTokens) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AdminJWTExpiredTokens) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AdminJWTExpiredTokens) predicate.AdminJWTExpiredTokens {
	return predicate.AdminJWTExpiredTokens(sql.NotPredicates(p))
}
