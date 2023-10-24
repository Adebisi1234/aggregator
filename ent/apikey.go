// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/paycrest/paycrest-protocol/ent/apikey"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
	"github.com/paycrest/paycrest-protocol/ent/senderprofile"
	"github.com/paycrest/paycrest-protocol/ent/validatorprofile"
)

// APIKey is the model entity for the APIKey schema.
type APIKey struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Secret holds the value of the "secret" field.
	Secret string `json:"secret,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the APIKeyQuery when eager-loading is set.
	Edges                     APIKeyEdges `json:"edges"`
	provider_profile_api_key  *string
	sender_profile_api_key    *uuid.UUID
	validator_profile_api_key *uuid.UUID
	selectValues              sql.SelectValues
}

// APIKeyEdges holds the relations/edges for other nodes in the graph.
type APIKeyEdges struct {
	// SenderProfile holds the value of the sender_profile edge.
	SenderProfile *SenderProfile `json:"sender_profile,omitempty"`
	// ProviderProfile holds the value of the provider_profile edge.
	ProviderProfile *ProviderProfile `json:"provider_profile,omitempty"`
	// ValidatorProfile holds the value of the validator_profile edge.
	ValidatorProfile *ValidatorProfile `json:"validator_profile,omitempty"`
	// PaymentOrders holds the value of the payment_orders edge.
	PaymentOrders []*PaymentOrder `json:"payment_orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// SenderProfileOrErr returns the SenderProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e APIKeyEdges) SenderProfileOrErr() (*SenderProfile, error) {
	if e.loadedTypes[0] {
		if e.SenderProfile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: senderprofile.Label}
		}
		return e.SenderProfile, nil
	}
	return nil, &NotLoadedError{edge: "sender_profile"}
}

// ProviderProfileOrErr returns the ProviderProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e APIKeyEdges) ProviderProfileOrErr() (*ProviderProfile, error) {
	if e.loadedTypes[1] {
		if e.ProviderProfile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: providerprofile.Label}
		}
		return e.ProviderProfile, nil
	}
	return nil, &NotLoadedError{edge: "provider_profile"}
}

// ValidatorProfileOrErr returns the ValidatorProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e APIKeyEdges) ValidatorProfileOrErr() (*ValidatorProfile, error) {
	if e.loadedTypes[2] {
		if e.ValidatorProfile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: validatorprofile.Label}
		}
		return e.ValidatorProfile, nil
	}
	return nil, &NotLoadedError{edge: "validator_profile"}
}

// PaymentOrdersOrErr returns the PaymentOrders value or an error if the edge
// was not loaded in eager-loading.
func (e APIKeyEdges) PaymentOrdersOrErr() ([]*PaymentOrder, error) {
	if e.loadedTypes[3] {
		return e.PaymentOrders, nil
	}
	return nil, &NotLoadedError{edge: "payment_orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*APIKey) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apikey.FieldSecret:
			values[i] = new(sql.NullString)
		case apikey.FieldID:
			values[i] = new(uuid.UUID)
		case apikey.ForeignKeys[0]: // provider_profile_api_key
			values[i] = new(sql.NullString)
		case apikey.ForeignKeys[1]: // sender_profile_api_key
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case apikey.ForeignKeys[2]: // validator_profile_api_key
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the APIKey fields.
func (ak *APIKey) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apikey.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ak.ID = *value
			}
		case apikey.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				ak.Secret = value.String
			}
		case apikey.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_profile_api_key", values[i])
			} else if value.Valid {
				ak.provider_profile_api_key = new(string)
				*ak.provider_profile_api_key = value.String
			}
		case apikey.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field sender_profile_api_key", values[i])
			} else if value.Valid {
				ak.sender_profile_api_key = new(uuid.UUID)
				*ak.sender_profile_api_key = *value.S.(*uuid.UUID)
			}
		case apikey.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field validator_profile_api_key", values[i])
			} else if value.Valid {
				ak.validator_profile_api_key = new(uuid.UUID)
				*ak.validator_profile_api_key = *value.S.(*uuid.UUID)
			}
		default:
			ak.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the APIKey.
// This includes values selected through modifiers, order, etc.
func (ak *APIKey) Value(name string) (ent.Value, error) {
	return ak.selectValues.Get(name)
}

// QuerySenderProfile queries the "sender_profile" edge of the APIKey entity.
func (ak *APIKey) QuerySenderProfile() *SenderProfileQuery {
	return NewAPIKeyClient(ak.config).QuerySenderProfile(ak)
}

// QueryProviderProfile queries the "provider_profile" edge of the APIKey entity.
func (ak *APIKey) QueryProviderProfile() *ProviderProfileQuery {
	return NewAPIKeyClient(ak.config).QueryProviderProfile(ak)
}

// QueryValidatorProfile queries the "validator_profile" edge of the APIKey entity.
func (ak *APIKey) QueryValidatorProfile() *ValidatorProfileQuery {
	return NewAPIKeyClient(ak.config).QueryValidatorProfile(ak)
}

// QueryPaymentOrders queries the "payment_orders" edge of the APIKey entity.
func (ak *APIKey) QueryPaymentOrders() *PaymentOrderQuery {
	return NewAPIKeyClient(ak.config).QueryPaymentOrders(ak)
}

// Update returns a builder for updating this APIKey.
// Note that you need to call APIKey.Unwrap() before calling this method if this APIKey
// was returned from a transaction, and the transaction was committed or rolled back.
func (ak *APIKey) Update() *APIKeyUpdateOne {
	return NewAPIKeyClient(ak.config).UpdateOne(ak)
}

// Unwrap unwraps the APIKey entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ak *APIKey) Unwrap() *APIKey {
	_tx, ok := ak.config.driver.(*txDriver)
	if !ok {
		panic("ent: APIKey is not a transactional entity")
	}
	ak.config.driver = _tx.drv
	return ak
}

// String implements the fmt.Stringer.
func (ak *APIKey) String() string {
	var builder strings.Builder
	builder.WriteString("APIKey(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ak.ID))
	builder.WriteString("secret=")
	builder.WriteString(ak.Secret)
	builder.WriteByte(')')
	return builder.String()
}

// APIKeys is a parsable slice of APIKey.
type APIKeys []*APIKey
