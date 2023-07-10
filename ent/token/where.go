// Code generated by ent, DO NOT EDIT.

package token

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldUpdatedAt, v))
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldSymbol, v))
}

// ContractAddress applies equality check predicate on the "contract_address" field. It's identical to ContractAddressEQ.
func ContractAddress(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldContractAddress, v))
}

// Decimals applies equality check predicate on the "decimals" field. It's identical to DecimalsEQ.
func Decimals(v int8) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldDecimals, v))
}

// IsEnabled applies equality check predicate on the "is_enabled" field. It's identical to IsEnabledEQ.
func IsEnabled(v bool) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldIsEnabled, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldUpdatedAt, v))
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldSymbol, v))
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldSymbol, v))
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldSymbol, vs...))
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldSymbol, vs...))
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldSymbol, v))
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldSymbol, v))
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldSymbol, v))
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldSymbol, v))
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.Token {
	return predicate.Token(sql.FieldContains(FieldSymbol, v))
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasPrefix(FieldSymbol, v))
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasSuffix(FieldSymbol, v))
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldSymbol, v))
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldSymbol, v))
}

// ContractAddressEQ applies the EQ predicate on the "contract_address" field.
func ContractAddressEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldContractAddress, v))
}

// ContractAddressNEQ applies the NEQ predicate on the "contract_address" field.
func ContractAddressNEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldContractAddress, v))
}

// ContractAddressIn applies the In predicate on the "contract_address" field.
func ContractAddressIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldContractAddress, vs...))
}

// ContractAddressNotIn applies the NotIn predicate on the "contract_address" field.
func ContractAddressNotIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldContractAddress, vs...))
}

// ContractAddressGT applies the GT predicate on the "contract_address" field.
func ContractAddressGT(v string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldContractAddress, v))
}

// ContractAddressGTE applies the GTE predicate on the "contract_address" field.
func ContractAddressGTE(v string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldContractAddress, v))
}

// ContractAddressLT applies the LT predicate on the "contract_address" field.
func ContractAddressLT(v string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldContractAddress, v))
}

// ContractAddressLTE applies the LTE predicate on the "contract_address" field.
func ContractAddressLTE(v string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldContractAddress, v))
}

// ContractAddressContains applies the Contains predicate on the "contract_address" field.
func ContractAddressContains(v string) predicate.Token {
	return predicate.Token(sql.FieldContains(FieldContractAddress, v))
}

// ContractAddressHasPrefix applies the HasPrefix predicate on the "contract_address" field.
func ContractAddressHasPrefix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasPrefix(FieldContractAddress, v))
}

// ContractAddressHasSuffix applies the HasSuffix predicate on the "contract_address" field.
func ContractAddressHasSuffix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasSuffix(FieldContractAddress, v))
}

// ContractAddressEqualFold applies the EqualFold predicate on the "contract_address" field.
func ContractAddressEqualFold(v string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldContractAddress, v))
}

// ContractAddressContainsFold applies the ContainsFold predicate on the "contract_address" field.
func ContractAddressContainsFold(v string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldContractAddress, v))
}

// DecimalsEQ applies the EQ predicate on the "decimals" field.
func DecimalsEQ(v int8) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldDecimals, v))
}

// DecimalsNEQ applies the NEQ predicate on the "decimals" field.
func DecimalsNEQ(v int8) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldDecimals, v))
}

// DecimalsIn applies the In predicate on the "decimals" field.
func DecimalsIn(vs ...int8) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldDecimals, vs...))
}

// DecimalsNotIn applies the NotIn predicate on the "decimals" field.
func DecimalsNotIn(vs ...int8) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldDecimals, vs...))
}

// DecimalsGT applies the GT predicate on the "decimals" field.
func DecimalsGT(v int8) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldDecimals, v))
}

// DecimalsGTE applies the GTE predicate on the "decimals" field.
func DecimalsGTE(v int8) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldDecimals, v))
}

// DecimalsLT applies the LT predicate on the "decimals" field.
func DecimalsLT(v int8) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldDecimals, v))
}

// DecimalsLTE applies the LTE predicate on the "decimals" field.
func DecimalsLTE(v int8) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldDecimals, v))
}

// IsEnabledEQ applies the EQ predicate on the "is_enabled" field.
func IsEnabledEQ(v bool) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldIsEnabled, v))
}

// IsEnabledNEQ applies the NEQ predicate on the "is_enabled" field.
func IsEnabledNEQ(v bool) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldIsEnabled, v))
}

// HasNetwork applies the HasEdge predicate on the "network" edge.
func HasNetwork() predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NetworkTable, NetworkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNetworkWith applies the HasEdge predicate on the "network" edge with a given conditions (other predicates).
func HasNetworkWith(preds ...predicate.Network) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := newNetworkStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPaymentOrders applies the HasEdge predicate on the "payment_orders" edge.
func HasPaymentOrders() predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentOrdersTable, PaymentOrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentOrdersWith applies the HasEdge predicate on the "payment_orders" edge with a given conditions (other predicates).
func HasPaymentOrdersWith(preds ...predicate.PaymentOrder) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := newPaymentOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
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
func Not(p predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		p(s.Not())
	})
}
