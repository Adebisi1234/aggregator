// Code generated by ent, DO NOT EDIT.

package paymentorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the paymentorder type in the database.
	Label = "payment_order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldAmountPaid holds the string denoting the amount_paid field in the database.
	FieldAmountPaid = "amount_paid"
	// FieldTxHash holds the string denoting the tx_hash field in the database.
	FieldTxHash = "tx_hash"
	// FieldReceiveAddressText holds the string denoting the receive_address_text field in the database.
	FieldReceiveAddressText = "receive_address_text"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLastUsed holds the string denoting the last_used field in the database.
	FieldLastUsed = "last_used"
	// EdgeSenderProfile holds the string denoting the sender_profile edge name in mutations.
	EdgeSenderProfile = "sender_profile"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// EdgeReceiveAddress holds the string denoting the receive_address edge name in mutations.
	EdgeReceiveAddress = "receive_address"
	// EdgeRecipient holds the string denoting the recipient edge name in mutations.
	EdgeRecipient = "recipient"
	// Table holds the table name of the paymentorder in the database.
	Table = "payment_orders"
	// SenderProfileTable is the table that holds the sender_profile relation/edge.
	SenderProfileTable = "payment_orders"
	// SenderProfileInverseTable is the table name for the SenderProfile entity.
	// It exists in this package in order to avoid circular dependency with the "senderprofile" package.
	SenderProfileInverseTable = "sender_profiles"
	// SenderProfileColumn is the table column denoting the sender_profile relation/edge.
	SenderProfileColumn = "sender_profile_payment_orders"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "payment_orders"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "token_payment_orders"
	// ReceiveAddressTable is the table that holds the receive_address relation/edge.
	ReceiveAddressTable = "receive_addresses"
	// ReceiveAddressInverseTable is the table name for the ReceiveAddress entity.
	// It exists in this package in order to avoid circular dependency with the "receiveaddress" package.
	ReceiveAddressInverseTable = "receive_addresses"
	// ReceiveAddressColumn is the table column denoting the receive_address relation/edge.
	ReceiveAddressColumn = "payment_order_receive_address"
	// RecipientTable is the table that holds the recipient relation/edge.
	RecipientTable = "payment_order_recipients"
	// RecipientInverseTable is the table name for the PaymentOrderRecipient entity.
	// It exists in this package in order to avoid circular dependency with the "paymentorderrecipient" package.
	RecipientInverseTable = "payment_order_recipients"
	// RecipientColumn is the table column denoting the recipient relation/edge.
	RecipientColumn = "payment_order_recipient"
)

// Columns holds all SQL columns for paymentorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAmount,
	FieldAmountPaid,
	FieldTxHash,
	FieldReceiveAddressText,
	FieldStatus,
	FieldLastUsed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "payment_orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"api_key_payment_orders",
	"sender_profile_payment_orders",
	"token_payment_orders",
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
	// TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	TxHashValidator func(string) error
	// ReceiveAddressTextValidator is a validator for the "receive_address_text" field. It is called by the builders before save.
	ReceiveAddressTextValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusInitiated is the default value of the Status enum.
const DefaultStatus = StatusInitiated

// Status values.
const (
	StatusInitiated Status = "initiated"
	StatusPending   Status = "pending"
	StatusSettled   Status = "settled"
	StatusCancelled Status = "cancelled"
	StatusFailed    Status = "failed"
	StatusRefunded  Status = "refunded"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInitiated, StatusPending, StatusSettled, StatusCancelled, StatusFailed, StatusRefunded:
		return nil
	default:
		return fmt.Errorf("paymentorder: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the PaymentOrder queries.
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

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByAmountPaid orders the results by the amount_paid field.
func ByAmountPaid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountPaid, opts...).ToFunc()
}

// ByTxHash orders the results by the tx_hash field.
func ByTxHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxHash, opts...).ToFunc()
}

// ByReceiveAddressText orders the results by the receive_address_text field.
func ByReceiveAddressText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceiveAddressText, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLastUsed orders the results by the last_used field.
func ByLastUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUsed, opts...).ToFunc()
}

// BySenderProfileField orders the results by sender_profile field.
func BySenderProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenderProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByTokenField orders the results by token field.
func ByTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokenStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceiveAddressField orders the results by receive_address field.
func ByReceiveAddressField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiveAddressStep(), sql.OrderByField(field, opts...))
	}
}

// ByRecipientField orders the results by recipient field.
func ByRecipientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecipientStep(), sql.OrderByField(field, opts...))
	}
}
func newSenderProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenderProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SenderProfileTable, SenderProfileColumn),
	)
}
func newTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TokenTable, TokenColumn),
	)
}
func newReceiveAddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiveAddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ReceiveAddressTable, ReceiveAddressColumn),
	)
}
func newRecipientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecipientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, RecipientTable, RecipientColumn),
	)
}
