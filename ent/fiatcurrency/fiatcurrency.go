// Code generated by ent, DO NOT EDIT.

package fiatcurrency

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the fiatcurrency type in the database.
	Label = "fiat_currency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldShortName holds the string denoting the short_name field in the database.
	FieldShortName = "short_name"
	// FieldDecimals holds the string denoting the decimals field in the database.
	FieldDecimals = "decimals"
	// FieldSymbol holds the string denoting the symbol field in the database.
	FieldSymbol = "symbol"
	// Table holds the table name of the fiatcurrency in the database.
	Table = "fiat_currencies"
)

// Columns holds all SQL columns for fiatcurrency fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCode,
	FieldShortName,
	FieldDecimals,
	FieldSymbol,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DefaultDecimals holds the default value on creation for the "decimals" field.
	DefaultDecimals int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the FiatCurrency queries.
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

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByShortName orders the results by the short_name field.
func ByShortName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShortName, opts...).ToFunc()
}

// ByDecimals orders the results by the decimals field.
func ByDecimals(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDecimals, opts...).ToFunc()
}

// BySymbol orders the results by the symbol field.
func BySymbol(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbol, opts...).ToFunc()
}
