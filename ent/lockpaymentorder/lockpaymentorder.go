// Code generated by ent, DO NOT EDIT.

package lockpaymentorder

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the lockpaymentorder type in the database.
	Label = "lock_payment_order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldGatewayID holds the string denoting the gateway_id field in the database.
	FieldGatewayID = "gateway_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldRate holds the string denoting the rate field in the database.
	FieldRate = "rate"
	// FieldOrderPercent holds the string denoting the order_percent field in the database.
	FieldOrderPercent = "order_percent"
	// FieldTxHash holds the string denoting the tx_hash field in the database.
	FieldTxHash = "tx_hash"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldBlockNumber holds the string denoting the block_number field in the database.
	FieldBlockNumber = "block_number"
	// FieldInstitution holds the string denoting the institution field in the database.
	FieldInstitution = "institution"
	// FieldAccountIdentifier holds the string denoting the account_identifier field in the database.
	FieldAccountIdentifier = "account_identifier"
	// FieldAccountName holds the string denoting the account_name field in the database.
	FieldAccountName = "account_name"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldCancellationCount holds the string denoting the cancellation_count field in the database.
	FieldCancellationCount = "cancellation_count"
	// FieldCancellationReasons holds the string denoting the cancellation_reasons field in the database.
	FieldCancellationReasons = "cancellation_reasons"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// EdgeProvisionBucket holds the string denoting the provision_bucket edge name in mutations.
	EdgeProvisionBucket = "provision_bucket"
	// EdgeProvider holds the string denoting the provider edge name in mutations.
	EdgeProvider = "provider"
	// EdgeFulfillment holds the string denoting the fulfillment edge name in mutations.
	EdgeFulfillment = "fulfillment"
	// EdgeTransactions holds the string denoting the transactions edge name in mutations.
	EdgeTransactions = "transactions"
	// Table holds the table name of the lockpaymentorder in the database.
	Table = "lock_payment_orders"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "lock_payment_orders"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "token_lock_payment_orders"
	// ProvisionBucketTable is the table that holds the provision_bucket relation/edge.
	ProvisionBucketTable = "lock_payment_orders"
	// ProvisionBucketInverseTable is the table name for the ProvisionBucket entity.
	// It exists in this package in order to avoid circular dependency with the "provisionbucket" package.
	ProvisionBucketInverseTable = "provision_buckets"
	// ProvisionBucketColumn is the table column denoting the provision_bucket relation/edge.
	ProvisionBucketColumn = "provision_bucket_lock_payment_orders"
	// ProviderTable is the table that holds the provider relation/edge.
	ProviderTable = "lock_payment_orders"
	// ProviderInverseTable is the table name for the ProviderProfile entity.
	// It exists in this package in order to avoid circular dependency with the "providerprofile" package.
	ProviderInverseTable = "provider_profiles"
	// ProviderColumn is the table column denoting the provider relation/edge.
	ProviderColumn = "provider_profile_assigned_orders"
	// FulfillmentTable is the table that holds the fulfillment relation/edge.
	FulfillmentTable = "lock_order_fulfillments"
	// FulfillmentInverseTable is the table name for the LockOrderFulfillment entity.
	// It exists in this package in order to avoid circular dependency with the "lockorderfulfillment" package.
	FulfillmentInverseTable = "lock_order_fulfillments"
	// FulfillmentColumn is the table column denoting the fulfillment relation/edge.
	FulfillmentColumn = "lock_payment_order_fulfillment"
	// TransactionsTable is the table that holds the transactions relation/edge.
	TransactionsTable = "transaction_logs"
	// TransactionsInverseTable is the table name for the TransactionLog entity.
	// It exists in this package in order to avoid circular dependency with the "transactionlog" package.
	TransactionsInverseTable = "transaction_logs"
	// TransactionsColumn is the table column denoting the transactions relation/edge.
	TransactionsColumn = "lock_payment_order_transactions"
)

// Columns holds all SQL columns for lockpaymentorder fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldGatewayID,
	FieldAmount,
	FieldRate,
	FieldOrderPercent,
	FieldTxHash,
	FieldStatus,
	FieldBlockNumber,
	FieldInstitution,
	FieldAccountIdentifier,
	FieldAccountName,
	FieldMemo,
	FieldCancellationCount,
	FieldCancellationReasons,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "lock_payment_orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"provider_profile_assigned_orders",
	"provision_bucket_lock_payment_orders",
	"token_lock_payment_orders",
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
	// DefaultGatewayID holds the default value on creation for the "gateway_id" field.
	DefaultGatewayID string
	// TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	TxHashValidator func(string) error
	// DefaultCancellationCount holds the default value on creation for the "cancellation_count" field.
	DefaultCancellationCount int
	// DefaultCancellationReasons holds the default value on creation for the "cancellation_reasons" field.
	DefaultCancellationReasons []string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusCancelled  Status = "cancelled"
	StatusFulfilled  Status = "fulfilled"
	StatusValidated  Status = "validated"
	StatusSettled    Status = "settled"
	StatusRefunded   Status = "refunded"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusProcessing, StatusCancelled, StatusFulfilled, StatusValidated, StatusSettled, StatusRefunded:
		return nil
	default:
		return fmt.Errorf("lockpaymentorder: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the LockPaymentOrder queries.
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

// ByGatewayID orders the results by the gateway_id field.
func ByGatewayID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGatewayID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByRate orders the results by the rate field.
func ByRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRate, opts...).ToFunc()
}

// ByOrderPercent orders the results by the order_percent field.
func ByOrderPercent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderPercent, opts...).ToFunc()
}

// ByTxHash orders the results by the tx_hash field.
func ByTxHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxHash, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByBlockNumber orders the results by the block_number field.
func ByBlockNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlockNumber, opts...).ToFunc()
}

// ByInstitution orders the results by the institution field.
func ByInstitution(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstitution, opts...).ToFunc()
}

// ByAccountIdentifier orders the results by the account_identifier field.
func ByAccountIdentifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountIdentifier, opts...).ToFunc()
}

// ByAccountName orders the results by the account_name field.
func ByAccountName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountName, opts...).ToFunc()
}

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// ByCancellationCount orders the results by the cancellation_count field.
func ByCancellationCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCancellationCount, opts...).ToFunc()
}

// ByTokenField orders the results by token field.
func ByTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokenStep(), sql.OrderByField(field, opts...))
	}
}

// ByProvisionBucketField orders the results by provision_bucket field.
func ByProvisionBucketField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProvisionBucketStep(), sql.OrderByField(field, opts...))
	}
}

// ByProviderField orders the results by provider field.
func ByProviderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProviderStep(), sql.OrderByField(field, opts...))
	}
}

// ByFulfillmentField orders the results by fulfillment field.
func ByFulfillmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFulfillmentStep(), sql.OrderByField(field, opts...))
	}
}

// ByTransactionsCount orders the results by transactions count.
func ByTransactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionsStep(), opts...)
	}
}

// ByTransactions orders the results by transactions terms.
func ByTransactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TokenTable, TokenColumn),
	)
}
func newProvisionBucketStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProvisionBucketInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProvisionBucketTable, ProvisionBucketColumn),
	)
}
func newProviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
	)
}
func newFulfillmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FulfillmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, FulfillmentTable, FulfillmentColumn),
	)
}
func newTransactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
	)
}
