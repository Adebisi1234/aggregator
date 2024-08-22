// Code generated by ent, DO NOT EDIT.

package verificationtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldToken, v))
}

// ExpiryAt applies equality check predicate on the "expiry_at" field. It's identical to ExpiryAtEQ.
func ExpiryAt(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldExpiryAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldContainsFold(FieldToken, v))
}

// ScopeEQ applies the EQ predicate on the "scope" field.
func ScopeEQ(v Scope) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldScope, v))
}

// ScopeNEQ applies the NEQ predicate on the "scope" field.
func ScopeNEQ(v Scope) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNEQ(FieldScope, v))
}

// ScopeIn applies the In predicate on the "scope" field.
func ScopeIn(vs ...Scope) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldIn(FieldScope, vs...))
}

// ScopeNotIn applies the NotIn predicate on the "scope" field.
func ScopeNotIn(vs ...Scope) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNotIn(FieldScope, vs...))
}

// ExpiryAtEQ applies the EQ predicate on the "expiry_at" field.
func ExpiryAtEQ(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldEQ(FieldExpiryAt, v))
}

// ExpiryAtNEQ applies the NEQ predicate on the "expiry_at" field.
func ExpiryAtNEQ(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNEQ(FieldExpiryAt, v))
}

// ExpiryAtIn applies the In predicate on the "expiry_at" field.
func ExpiryAtIn(vs ...time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldIn(FieldExpiryAt, vs...))
}

// ExpiryAtNotIn applies the NotIn predicate on the "expiry_at" field.
func ExpiryAtNotIn(vs ...time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldNotIn(FieldExpiryAt, vs...))
}

// ExpiryAtGT applies the GT predicate on the "expiry_at" field.
func ExpiryAtGT(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGT(FieldExpiryAt, v))
}

// ExpiryAtGTE applies the GTE predicate on the "expiry_at" field.
func ExpiryAtGTE(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldGTE(FieldExpiryAt, v))
}

// ExpiryAtLT applies the LT predicate on the "expiry_at" field.
func ExpiryAtLT(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLT(FieldExpiryAt, v))
}

// ExpiryAtLTE applies the LTE predicate on the "expiry_at" field.
func ExpiryAtLTE(v time.Time) predicate.VerificationToken {
	return predicate.VerificationToken(sql.FieldLTE(FieldExpiryAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.VerificationToken {
	return predicate.VerificationToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.VerificationToken {
	return predicate.VerificationToken(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VerificationToken) predicate.VerificationToken {
	return predicate.VerificationToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VerificationToken) predicate.VerificationToken {
	return predicate.VerificationToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VerificationToken) predicate.VerificationToken {
	return predicate.VerificationToken(sql.NotPredicates(p))
}
