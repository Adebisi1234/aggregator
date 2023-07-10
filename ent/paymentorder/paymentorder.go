// Code generated by ent, DO NOT EDIT.

package paymentorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldNetwork holds the string denoting the network field in the database.
	FieldNetwork = "network"
	// FieldTxHash holds the string denoting the tx_hash field in the database.
	FieldTxHash = "tx_hash"
	// FieldReceiveAddress holds the string denoting the receive_address field in the database.
	FieldReceiveAddress = "receive_address"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLastUsed holds the string denoting the last_used field in the database.
	FieldLastUsed = "last_used"
	// EdgeAPIKey holds the string denoting the api_key edge name in mutations.
	EdgeAPIKey = "api_key"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// EdgeReceiveAddressFk holds the string denoting the receive_address_fk edge name in mutations.
	EdgeReceiveAddressFk = "receive_address_fk"
	// EdgeRecipient holds the string denoting the recipient edge name in mutations.
	EdgeRecipient = "recipient"
	// Table holds the table name of the paymentorder in the database.
	Table = "payment_orders"
	// APIKeyTable is the table that holds the api_key relation/edge.
	APIKeyTable = "payment_orders"
	// APIKeyInverseTable is the table name for the APIKey entity.
	// It exists in this package in order to avoid circular dependency with the "apikey" package.
	APIKeyInverseTable = "api_keys"
	// APIKeyColumn is the table column denoting the api_key relation/edge.
	APIKeyColumn = "api_key_payment_orders"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "payment_orders"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "token_payment_orders"
	// ReceiveAddressFkTable is the table that holds the receive_address_fk relation/edge.
	ReceiveAddressFkTable = "receive_addresses"
	// ReceiveAddressFkInverseTable is the table name for the ReceiveAddress entity.
	// It exists in this package in order to avoid circular dependency with the "receiveaddress" package.
	ReceiveAddressFkInverseTable = "receive_addresses"
	// ReceiveAddressFkColumn is the table column denoting the receive_address_fk relation/edge.
	ReceiveAddressFkColumn = "payment_order_receive_address_fk"
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
	FieldNetwork,
	FieldTxHash,
	FieldReceiveAddress,
	FieldStatus,
	FieldLastUsed,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "payment_orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"api_key_payment_orders",
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
	// ReceiveAddressValidator is a validator for the "receive_address" field. It is called by the builders before save.
	ReceiveAddressValidator func(string) error
)

// Network defines the type for the "network" enum field.
type Network string

// Network values.
const (
	NetworkBnbSmartChain Network = "bnb-smart-chain"
	NetworkPolygon       Network = "polygon"
	NetworkTron          Network = "tron"
	NetworkPolygonMumbai Network = "polygon-mumbai"
	NetworkTronShasta    Network = "tron-shasta"
)

func (n Network) String() string {
	return string(n)
}

// NetworkValidator is a validator for the "network" field enum values. It is called by the builders before save.
func NetworkValidator(n Network) error {
	switch n {
	case NetworkBnbSmartChain, NetworkPolygon, NetworkTron, NetworkPolygonMumbai, NetworkTronShasta:
		return nil
	default:
		return fmt.Errorf("paymentorder: invalid enum value for network field: %q", n)
	}
}

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

// ByNetwork orders the results by the network field.
func ByNetwork(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNetwork, opts...).ToFunc()
}

// ByTxHash orders the results by the tx_hash field.
func ByTxHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxHash, opts...).ToFunc()
}

// ByReceiveAddress orders the results by the receive_address field.
func ByReceiveAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceiveAddress, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLastUsed orders the results by the last_used field.
func ByLastUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUsed, opts...).ToFunc()
}

// ByAPIKeyField orders the results by api_key field.
func ByAPIKeyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAPIKeyStep(), sql.OrderByField(field, opts...))
	}
}

// ByTokenField orders the results by token field.
func ByTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokenStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceiveAddressFkField orders the results by receive_address_fk field.
func ByReceiveAddressFkField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiveAddressFkStep(), sql.OrderByField(field, opts...))
	}
}

// ByRecipientField orders the results by recipient field.
func ByRecipientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecipientStep(), sql.OrderByField(field, opts...))
	}
}
func newAPIKeyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(APIKeyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, APIKeyTable, APIKeyColumn),
	)
}
func newTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TokenTable, TokenColumn),
	)
}
func newReceiveAddressFkStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiveAddressFkInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ReceiveAddressFkTable, ReceiveAddressFkColumn),
	)
}
func newRecipientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecipientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, RecipientTable, RecipientColumn),
	)
}
