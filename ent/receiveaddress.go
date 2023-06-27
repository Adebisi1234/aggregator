// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/paycrest/paycrest-protocol/ent/receiveaddress"
)

// ReceiveAddress is the model entity for the ReceiveAddress schema.
type ReceiveAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// AccountIndex holds the value of the "accountIndex" field.
	AccountIndex int `json:"accountIndex,omitempty"`
	// Status holds the value of the "status" field.
	Status       receiveaddress.Status `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ReceiveAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case receiveaddress.FieldID, receiveaddress.FieldAccountIndex:
			values[i] = new(sql.NullInt64)
		case receiveaddress.FieldAddress, receiveaddress.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReceiveAddress fields.
func (ra *ReceiveAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case receiveaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ra.ID = int(value.Int64)
		case receiveaddress.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				ra.Address = value.String
			}
		case receiveaddress.FieldAccountIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field accountIndex", values[i])
			} else if value.Valid {
				ra.AccountIndex = int(value.Int64)
			}
		case receiveaddress.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ra.Status = receiveaddress.Status(value.String)
			}
		default:
			ra.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ReceiveAddress.
// This includes values selected through modifiers, order, etc.
func (ra *ReceiveAddress) Value(name string) (ent.Value, error) {
	return ra.selectValues.Get(name)
}

// Update returns a builder for updating this ReceiveAddress.
// Note that you need to call ReceiveAddress.Unwrap() before calling this method if this ReceiveAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (ra *ReceiveAddress) Update() *ReceiveAddressUpdateOne {
	return NewReceiveAddressClient(ra.config).UpdateOne(ra)
}

// Unwrap unwraps the ReceiveAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ra *ReceiveAddress) Unwrap() *ReceiveAddress {
	_tx, ok := ra.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReceiveAddress is not a transactional entity")
	}
	ra.config.driver = _tx.drv
	return ra
}

// String implements the fmt.Stringer.
func (ra *ReceiveAddress) String() string {
	var builder strings.Builder
	builder.WriteString("ReceiveAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ra.ID))
	builder.WriteString("address=")
	builder.WriteString(ra.Address)
	builder.WriteString(", ")
	builder.WriteString("accountIndex=")
	builder.WriteString(fmt.Sprintf("%v", ra.AccountIndex))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ra.Status))
	builder.WriteByte(')')
	return builder.String()
}

// ReceiveAddresses is a parsable slice of ReceiveAddress.
type ReceiveAddresses []*ReceiveAddress
