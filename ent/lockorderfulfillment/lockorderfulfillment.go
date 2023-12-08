// Code generated by ent, DO NOT EDIT.

package lockorderfulfillment

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the lockorderfulfillment type in the database.
	Label = "lock_order_fulfillment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTxID holds the string denoting the tx_id field in the database.
	FieldTxID = "tx_id"
	// FieldValidationStatus holds the string denoting the validation_status field in the database.
	FieldValidationStatus = "validation_status"
	// FieldValidationError holds the string denoting the validation_error field in the database.
	FieldValidationError = "validation_error"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// Table holds the table name of the lockorderfulfillment in the database.
	Table = "lock_order_fulfillments"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "lock_order_fulfillments"
	// OrderInverseTable is the table name for the LockPaymentOrder entity.
	// It exists in this package in order to avoid circular dependency with the "lockpaymentorder" package.
	OrderInverseTable = "lock_payment_orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "lock_payment_order_fulfillment"
)

// Columns holds all SQL columns for lockorderfulfillment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTxID,
	FieldValidationStatus,
	FieldValidationError,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "lock_order_fulfillments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"lock_payment_order_fulfillment",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// ValidationStatus defines the type for the "validation_status" enum field.
type ValidationStatus string

// ValidationStatusPending is the default value of the ValidationStatus enum.
const DefaultValidationStatus = ValidationStatusPending

// ValidationStatus values.
const (
	ValidationStatusPending ValidationStatus = "pending"
	ValidationStatusSuccess ValidationStatus = "success"
	ValidationStatusFailure ValidationStatus = "failure"
)

func (vs ValidationStatus) String() string {
	return string(vs)
}

// ValidationStatusValidator is a validator for the "validation_status" field enum values. It is called by the builders before save.
func ValidationStatusValidator(vs ValidationStatus) error {
	switch vs {
	case ValidationStatusPending, ValidationStatusSuccess, ValidationStatusFailure:
		return nil
	default:
		return fmt.Errorf("lockorderfulfillment: invalid enum value for validation_status field: %q", vs)
	}
}

// OrderOption defines the ordering options for the LockOrderFulfillment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTxID orders the results by the tx_id field.
func ByTxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxID, opts...).ToFunc()
}

// ByValidationStatus orders the results by the validation_status field.
func ByValidationStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidationStatus, opts...).ToFunc()
}

// ByValidationError orders the results by the validation_error field.
func ByValidationError(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidationError, opts...).ToFunc()
}

// ByOrderField orders the results by order field.
func ByOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, OrderTable, OrderColumn),
	)
}
