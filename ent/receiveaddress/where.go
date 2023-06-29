// Code generated by ent, DO NOT EDIT.

package receiveaddress

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldUpdatedAt, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldAddress, v))
}

// AccountIndex applies equality check predicate on the "accountIndex" field. It's identical to AccountIndexEQ.
func AccountIndex(v int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldAccountIndex, v))
}

// LastUsed applies equality check predicate on the "last_used" field. It's identical to LastUsedEQ.
func LastUsed(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldLastUsed, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLTE(FieldUpdatedAt, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldContainsFold(FieldAddress, v))
}

// AccountIndexEQ applies the EQ predicate on the "accountIndex" field.
func AccountIndexEQ(v int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldAccountIndex, v))
}

// AccountIndexNEQ applies the NEQ predicate on the "accountIndex" field.
func AccountIndexNEQ(v int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNEQ(FieldAccountIndex, v))
}

// AccountIndexIn applies the In predicate on the "accountIndex" field.
func AccountIndexIn(vs ...int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldIn(FieldAccountIndex, vs...))
}

// AccountIndexNotIn applies the NotIn predicate on the "accountIndex" field.
func AccountIndexNotIn(vs ...int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNotIn(FieldAccountIndex, vs...))
}

// AccountIndexGT applies the GT predicate on the "accountIndex" field.
func AccountIndexGT(v int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGT(FieldAccountIndex, v))
}

// AccountIndexGTE applies the GTE predicate on the "accountIndex" field.
func AccountIndexGTE(v int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGTE(FieldAccountIndex, v))
}

// AccountIndexLT applies the LT predicate on the "accountIndex" field.
func AccountIndexLT(v int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLT(FieldAccountIndex, v))
}

// AccountIndexLTE applies the LTE predicate on the "accountIndex" field.
func AccountIndexLTE(v int) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLTE(FieldAccountIndex, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNotIn(FieldStatus, vs...))
}

// LastUsedEQ applies the EQ predicate on the "last_used" field.
func LastUsedEQ(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldEQ(FieldLastUsed, v))
}

// LastUsedNEQ applies the NEQ predicate on the "last_used" field.
func LastUsedNEQ(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNEQ(FieldLastUsed, v))
}

// LastUsedIn applies the In predicate on the "last_used" field.
func LastUsedIn(vs ...time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldIn(FieldLastUsed, vs...))
}

// LastUsedNotIn applies the NotIn predicate on the "last_used" field.
func LastUsedNotIn(vs ...time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNotIn(FieldLastUsed, vs...))
}

// LastUsedGT applies the GT predicate on the "last_used" field.
func LastUsedGT(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGT(FieldLastUsed, v))
}

// LastUsedGTE applies the GTE predicate on the "last_used" field.
func LastUsedGTE(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldGTE(FieldLastUsed, v))
}

// LastUsedLT applies the LT predicate on the "last_used" field.
func LastUsedLT(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLT(FieldLastUsed, v))
}

// LastUsedLTE applies the LTE predicate on the "last_used" field.
func LastUsedLTE(v time.Time) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldLTE(FieldLastUsed, v))
}

// LastUsedIsNil applies the IsNil predicate on the "last_used" field.
func LastUsedIsNil() predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldIsNull(FieldLastUsed))
}

// LastUsedNotNil applies the NotNil predicate on the "last_used" field.
func LastUsedNotNil() predicate.ReceiveAddress {
	return predicate.ReceiveAddress(sql.FieldNotNull(FieldLastUsed))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ReceiveAddress) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ReceiveAddress) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ReceiveAddress) predicate.ReceiveAddress {
	return predicate.ReceiveAddress(func(s *sql.Selector) {
		p(s.Not())
	})
}
