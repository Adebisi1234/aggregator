// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/paycrest/paycrest-protocol/ent/apikey"
	"github.com/paycrest/paycrest-protocol/ent/user"
)

// APIKey is the model entity for the APIKey schema.
type APIKey struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Scope holds the value of the "scope" field.
	Scope apikey.Scope `json:"scope,omitempty"`
	// Secret holds the value of the "secret" field.
	Secret string `json:"secret,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the APIKeyQuery when eager-loading is set.
	Edges         APIKeyEdges `json:"edges"`
	user_api_keys *uuid.UUID
	selectValues  sql.SelectValues
}

// APIKeyEdges holds the relations/edges for other nodes in the graph.
type APIKeyEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e APIKeyEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*APIKey) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apikey.FieldIsActive:
			values[i] = new(sql.NullBool)
		case apikey.FieldName, apikey.FieldScope, apikey.FieldSecret:
			values[i] = new(sql.NullString)
		case apikey.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case apikey.FieldID:
			values[i] = new(uuid.UUID)
		case apikey.ForeignKeys[0]: // user_api_keys
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
		case apikey.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ak.Name = value.String
			}
		case apikey.FieldScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scope", values[i])
			} else if value.Valid {
				ak.Scope = apikey.Scope(value.String)
			}
		case apikey.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				ak.Secret = value.String
			}
		case apikey.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				ak.IsActive = value.Bool
			}
		case apikey.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ak.CreatedAt = value.Time
			}
		case apikey.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_api_keys", values[i])
			} else if value.Valid {
				ak.user_api_keys = new(uuid.UUID)
				*ak.user_api_keys = *value.S.(*uuid.UUID)
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

// QueryUser queries the "user" edge of the APIKey entity.
func (ak *APIKey) QueryUser() *UserQuery {
	return NewAPIKeyClient(ak.config).QueryUser(ak)
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
	builder.WriteString("name=")
	builder.WriteString(ak.Name)
	builder.WriteString(", ")
	builder.WriteString("scope=")
	builder.WriteString(fmt.Sprintf("%v", ak.Scope))
	builder.WriteString(", ")
	builder.WriteString("secret=")
	builder.WriteString(ak.Secret)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", ak.IsActive))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ak.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// APIKeys is a parsable slice of APIKey.
type APIKeys []*APIKey
