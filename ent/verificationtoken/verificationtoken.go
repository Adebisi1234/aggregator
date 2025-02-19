// Code generated by ent, DO NOT EDIT.

package verificationtoken

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the verificationtoken type in the database.
	Label = "verification_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldScope holds the string denoting the scope field in the database.
	FieldScope = "scope"
	// FieldExpiryAt holds the string denoting the expiry_at field in the database.
	FieldExpiryAt = "expiry_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the verificationtoken in the database.
	Table = "verification_tokens"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "verification_tokens"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_verification_token"
)

// Columns holds all SQL columns for verificationtoken fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldToken,
	FieldScope,
	FieldExpiryAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "verification_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_verification_token",
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/paycrest/aggregator/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultExpiryAt holds the default value on creation for the "expiry_at" field.
	DefaultExpiryAt time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Scope defines the type for the "scope" enum field.
type Scope string

// Scope values.
const (
	ScopeEmailVerification Scope = "emailVerification"
	ScopeResetPassword     Scope = "resetPassword"
)

func (s Scope) String() string {
	return string(s)
}

// ScopeValidator is a validator for the "scope" field enum values. It is called by the builders before save.
func ScopeValidator(s Scope) error {
	switch s {
	case ScopeEmailVerification, ScopeResetPassword:
		return nil
	default:
		return fmt.Errorf("verificationtoken: invalid enum value for scope field: %q", s)
	}
}

// OrderOption defines the ordering options for the VerificationToken queries.
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

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByScope orders the results by the scope field.
func ByScope(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScope, opts...).ToFunc()
}

// ByExpiryAt orders the results by the expiry_at field.
func ByExpiryAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiryAt, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
