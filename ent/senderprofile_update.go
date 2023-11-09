// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/apikey"
	"github.com/paycrest/protocol/ent/paymentorder"
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/paycrest/protocol/ent/senderprofile"
	"github.com/shopspring/decimal"
)

// SenderProfileUpdate is the builder for updating SenderProfile entities.
type SenderProfileUpdate struct {
	config
	hooks    []Hook
	mutation *SenderProfileMutation
}

// Where appends a list predicates to the SenderProfileUpdate builder.
func (spu *SenderProfileUpdate) Where(ps ...predicate.SenderProfile) *SenderProfileUpdate {
	spu.mutation.Where(ps...)
	return spu
}

// SetWebhookURL sets the "webhook_url" field.
func (spu *SenderProfileUpdate) SetWebhookURL(s string) *SenderProfileUpdate {
	spu.mutation.SetWebhookURL(s)
	return spu
}

// SetNillableWebhookURL sets the "webhook_url" field if the given value is not nil.
func (spu *SenderProfileUpdate) SetNillableWebhookURL(s *string) *SenderProfileUpdate {
	if s != nil {
		spu.SetWebhookURL(*s)
	}
	return spu
}

// ClearWebhookURL clears the value of the "webhook_url" field.
func (spu *SenderProfileUpdate) ClearWebhookURL() *SenderProfileUpdate {
	spu.mutation.ClearWebhookURL()
	return spu
}

// SetFeePerTokenUnit sets the "fee_per_token_unit" field.
func (spu *SenderProfileUpdate) SetFeePerTokenUnit(d decimal.Decimal) *SenderProfileUpdate {
	spu.mutation.ResetFeePerTokenUnit()
	spu.mutation.SetFeePerTokenUnit(d)
	return spu
}

// AddFeePerTokenUnit adds d to the "fee_per_token_unit" field.
func (spu *SenderProfileUpdate) AddFeePerTokenUnit(d decimal.Decimal) *SenderProfileUpdate {
	spu.mutation.AddFeePerTokenUnit(d)
	return spu
}

// SetFeeAddress sets the "fee_address" field.
func (spu *SenderProfileUpdate) SetFeeAddress(s string) *SenderProfileUpdate {
	spu.mutation.SetFeeAddress(s)
	return spu
}

// SetNillableFeeAddress sets the "fee_address" field if the given value is not nil.
func (spu *SenderProfileUpdate) SetNillableFeeAddress(s *string) *SenderProfileUpdate {
	if s != nil {
		spu.SetFeeAddress(*s)
	}
	return spu
}

// ClearFeeAddress clears the value of the "fee_address" field.
func (spu *SenderProfileUpdate) ClearFeeAddress() *SenderProfileUpdate {
	spu.mutation.ClearFeeAddress()
	return spu
}

// SetRefundAddress sets the "refund_address" field.
func (spu *SenderProfileUpdate) SetRefundAddress(s string) *SenderProfileUpdate {
	spu.mutation.SetRefundAddress(s)
	return spu
}

// SetNillableRefundAddress sets the "refund_address" field if the given value is not nil.
func (spu *SenderProfileUpdate) SetNillableRefundAddress(s *string) *SenderProfileUpdate {
	if s != nil {
		spu.SetRefundAddress(*s)
	}
	return spu
}

// ClearRefundAddress clears the value of the "refund_address" field.
func (spu *SenderProfileUpdate) ClearRefundAddress() *SenderProfileUpdate {
	spu.mutation.ClearRefundAddress()
	return spu
}

// SetDomainWhitelist sets the "domain_whitelist" field.
func (spu *SenderProfileUpdate) SetDomainWhitelist(s []string) *SenderProfileUpdate {
	spu.mutation.SetDomainWhitelist(s)
	return spu
}

// AppendDomainWhitelist appends s to the "domain_whitelist" field.
func (spu *SenderProfileUpdate) AppendDomainWhitelist(s []string) *SenderProfileUpdate {
	spu.mutation.AppendDomainWhitelist(s)
	return spu
}

// SetIsActive sets the "is_active" field.
func (spu *SenderProfileUpdate) SetIsActive(b bool) *SenderProfileUpdate {
	spu.mutation.SetIsActive(b)
	return spu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (spu *SenderProfileUpdate) SetNillableIsActive(b *bool) *SenderProfileUpdate {
	if b != nil {
		spu.SetIsActive(*b)
	}
	return spu
}

// SetUpdatedAt sets the "updated_at" field.
func (spu *SenderProfileUpdate) SetUpdatedAt(t time.Time) *SenderProfileUpdate {
	spu.mutation.SetUpdatedAt(t)
	return spu
}

// SetAPIKeyID sets the "api_key" edge to the APIKey entity by ID.
func (spu *SenderProfileUpdate) SetAPIKeyID(id uuid.UUID) *SenderProfileUpdate {
	spu.mutation.SetAPIKeyID(id)
	return spu
}

// SetNillableAPIKeyID sets the "api_key" edge to the APIKey entity by ID if the given value is not nil.
func (spu *SenderProfileUpdate) SetNillableAPIKeyID(id *uuid.UUID) *SenderProfileUpdate {
	if id != nil {
		spu = spu.SetAPIKeyID(*id)
	}
	return spu
}

// SetAPIKey sets the "api_key" edge to the APIKey entity.
func (spu *SenderProfileUpdate) SetAPIKey(a *APIKey) *SenderProfileUpdate {
	return spu.SetAPIKeyID(a.ID)
}

// AddPaymentOrderIDs adds the "payment_orders" edge to the PaymentOrder entity by IDs.
func (spu *SenderProfileUpdate) AddPaymentOrderIDs(ids ...uuid.UUID) *SenderProfileUpdate {
	spu.mutation.AddPaymentOrderIDs(ids...)
	return spu
}

// AddPaymentOrders adds the "payment_orders" edges to the PaymentOrder entity.
func (spu *SenderProfileUpdate) AddPaymentOrders(p ...*PaymentOrder) *SenderProfileUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return spu.AddPaymentOrderIDs(ids...)
}

// Mutation returns the SenderProfileMutation object of the builder.
func (spu *SenderProfileUpdate) Mutation() *SenderProfileMutation {
	return spu.mutation
}

// ClearAPIKey clears the "api_key" edge to the APIKey entity.
func (spu *SenderProfileUpdate) ClearAPIKey() *SenderProfileUpdate {
	spu.mutation.ClearAPIKey()
	return spu
}

// ClearPaymentOrders clears all "payment_orders" edges to the PaymentOrder entity.
func (spu *SenderProfileUpdate) ClearPaymentOrders() *SenderProfileUpdate {
	spu.mutation.ClearPaymentOrders()
	return spu
}

// RemovePaymentOrderIDs removes the "payment_orders" edge to PaymentOrder entities by IDs.
func (spu *SenderProfileUpdate) RemovePaymentOrderIDs(ids ...uuid.UUID) *SenderProfileUpdate {
	spu.mutation.RemovePaymentOrderIDs(ids...)
	return spu
}

// RemovePaymentOrders removes "payment_orders" edges to PaymentOrder entities.
func (spu *SenderProfileUpdate) RemovePaymentOrders(p ...*PaymentOrder) *SenderProfileUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return spu.RemovePaymentOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spu *SenderProfileUpdate) Save(ctx context.Context) (int, error) {
	spu.defaults()
	return withHooks(ctx, spu.sqlSave, spu.mutation, spu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (spu *SenderProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := spu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spu *SenderProfileUpdate) Exec(ctx context.Context) error {
	_, err := spu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spu *SenderProfileUpdate) ExecX(ctx context.Context) {
	if err := spu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spu *SenderProfileUpdate) defaults() {
	if _, ok := spu.mutation.UpdatedAt(); !ok {
		v := senderprofile.UpdateDefaultUpdatedAt()
		spu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spu *SenderProfileUpdate) check() error {
	if _, ok := spu.mutation.UserID(); spu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SenderProfile.user"`)
	}
	return nil
}

func (spu *SenderProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := spu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(senderprofile.Table, senderprofile.Columns, sqlgraph.NewFieldSpec(senderprofile.FieldID, field.TypeUUID))
	if ps := spu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spu.mutation.WebhookURL(); ok {
		_spec.SetField(senderprofile.FieldWebhookURL, field.TypeString, value)
	}
	if spu.mutation.WebhookURLCleared() {
		_spec.ClearField(senderprofile.FieldWebhookURL, field.TypeString)
	}
	if value, ok := spu.mutation.FeePerTokenUnit(); ok {
		_spec.SetField(senderprofile.FieldFeePerTokenUnit, field.TypeFloat64, value)
	}
	if value, ok := spu.mutation.AddedFeePerTokenUnit(); ok {
		_spec.AddField(senderprofile.FieldFeePerTokenUnit, field.TypeFloat64, value)
	}
	if value, ok := spu.mutation.FeeAddress(); ok {
		_spec.SetField(senderprofile.FieldFeeAddress, field.TypeString, value)
	}
	if spu.mutation.FeeAddressCleared() {
		_spec.ClearField(senderprofile.FieldFeeAddress, field.TypeString)
	}
	if value, ok := spu.mutation.RefundAddress(); ok {
		_spec.SetField(senderprofile.FieldRefundAddress, field.TypeString, value)
	}
	if spu.mutation.RefundAddressCleared() {
		_spec.ClearField(senderprofile.FieldRefundAddress, field.TypeString)
	}
	if value, ok := spu.mutation.DomainWhitelist(); ok {
		_spec.SetField(senderprofile.FieldDomainWhitelist, field.TypeJSON, value)
	}
	if value, ok := spu.mutation.AppendedDomainWhitelist(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, senderprofile.FieldDomainWhitelist, value)
		})
	}
	if value, ok := spu.mutation.IsActive(); ok {
		_spec.SetField(senderprofile.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := spu.mutation.UpdatedAt(); ok {
		_spec.SetField(senderprofile.FieldUpdatedAt, field.TypeTime, value)
	}
	if spu.mutation.APIKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   senderprofile.APIKeyTable,
			Columns: []string{senderprofile.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.APIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   senderprofile.APIKeyTable,
			Columns: []string{senderprofile.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if spu.mutation.PaymentOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   senderprofile.PaymentOrdersTable,
			Columns: []string{senderprofile.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.RemovedPaymentOrdersIDs(); len(nodes) > 0 && !spu.mutation.PaymentOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   senderprofile.PaymentOrdersTable,
			Columns: []string{senderprofile.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.PaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   senderprofile.PaymentOrdersTable,
			Columns: []string{senderprofile.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{senderprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	spu.mutation.done = true
	return n, nil
}

// SenderProfileUpdateOne is the builder for updating a single SenderProfile entity.
type SenderProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SenderProfileMutation
}

// SetWebhookURL sets the "webhook_url" field.
func (spuo *SenderProfileUpdateOne) SetWebhookURL(s string) *SenderProfileUpdateOne {
	spuo.mutation.SetWebhookURL(s)
	return spuo
}

// SetNillableWebhookURL sets the "webhook_url" field if the given value is not nil.
func (spuo *SenderProfileUpdateOne) SetNillableWebhookURL(s *string) *SenderProfileUpdateOne {
	if s != nil {
		spuo.SetWebhookURL(*s)
	}
	return spuo
}

// ClearWebhookURL clears the value of the "webhook_url" field.
func (spuo *SenderProfileUpdateOne) ClearWebhookURL() *SenderProfileUpdateOne {
	spuo.mutation.ClearWebhookURL()
	return spuo
}

// SetFeePerTokenUnit sets the "fee_per_token_unit" field.
func (spuo *SenderProfileUpdateOne) SetFeePerTokenUnit(d decimal.Decimal) *SenderProfileUpdateOne {
	spuo.mutation.ResetFeePerTokenUnit()
	spuo.mutation.SetFeePerTokenUnit(d)
	return spuo
}

// AddFeePerTokenUnit adds d to the "fee_per_token_unit" field.
func (spuo *SenderProfileUpdateOne) AddFeePerTokenUnit(d decimal.Decimal) *SenderProfileUpdateOne {
	spuo.mutation.AddFeePerTokenUnit(d)
	return spuo
}

// SetFeeAddress sets the "fee_address" field.
func (spuo *SenderProfileUpdateOne) SetFeeAddress(s string) *SenderProfileUpdateOne {
	spuo.mutation.SetFeeAddress(s)
	return spuo
}

// SetNillableFeeAddress sets the "fee_address" field if the given value is not nil.
func (spuo *SenderProfileUpdateOne) SetNillableFeeAddress(s *string) *SenderProfileUpdateOne {
	if s != nil {
		spuo.SetFeeAddress(*s)
	}
	return spuo
}

// ClearFeeAddress clears the value of the "fee_address" field.
func (spuo *SenderProfileUpdateOne) ClearFeeAddress() *SenderProfileUpdateOne {
	spuo.mutation.ClearFeeAddress()
	return spuo
}

// SetRefundAddress sets the "refund_address" field.
func (spuo *SenderProfileUpdateOne) SetRefundAddress(s string) *SenderProfileUpdateOne {
	spuo.mutation.SetRefundAddress(s)
	return spuo
}

// SetNillableRefundAddress sets the "refund_address" field if the given value is not nil.
func (spuo *SenderProfileUpdateOne) SetNillableRefundAddress(s *string) *SenderProfileUpdateOne {
	if s != nil {
		spuo.SetRefundAddress(*s)
	}
	return spuo
}

// ClearRefundAddress clears the value of the "refund_address" field.
func (spuo *SenderProfileUpdateOne) ClearRefundAddress() *SenderProfileUpdateOne {
	spuo.mutation.ClearRefundAddress()
	return spuo
}

// SetDomainWhitelist sets the "domain_whitelist" field.
func (spuo *SenderProfileUpdateOne) SetDomainWhitelist(s []string) *SenderProfileUpdateOne {
	spuo.mutation.SetDomainWhitelist(s)
	return spuo
}

// AppendDomainWhitelist appends s to the "domain_whitelist" field.
func (spuo *SenderProfileUpdateOne) AppendDomainWhitelist(s []string) *SenderProfileUpdateOne {
	spuo.mutation.AppendDomainWhitelist(s)
	return spuo
}

// SetIsActive sets the "is_active" field.
func (spuo *SenderProfileUpdateOne) SetIsActive(b bool) *SenderProfileUpdateOne {
	spuo.mutation.SetIsActive(b)
	return spuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (spuo *SenderProfileUpdateOne) SetNillableIsActive(b *bool) *SenderProfileUpdateOne {
	if b != nil {
		spuo.SetIsActive(*b)
	}
	return spuo
}

// SetUpdatedAt sets the "updated_at" field.
func (spuo *SenderProfileUpdateOne) SetUpdatedAt(t time.Time) *SenderProfileUpdateOne {
	spuo.mutation.SetUpdatedAt(t)
	return spuo
}

// SetAPIKeyID sets the "api_key" edge to the APIKey entity by ID.
func (spuo *SenderProfileUpdateOne) SetAPIKeyID(id uuid.UUID) *SenderProfileUpdateOne {
	spuo.mutation.SetAPIKeyID(id)
	return spuo
}

// SetNillableAPIKeyID sets the "api_key" edge to the APIKey entity by ID if the given value is not nil.
func (spuo *SenderProfileUpdateOne) SetNillableAPIKeyID(id *uuid.UUID) *SenderProfileUpdateOne {
	if id != nil {
		spuo = spuo.SetAPIKeyID(*id)
	}
	return spuo
}

// SetAPIKey sets the "api_key" edge to the APIKey entity.
func (spuo *SenderProfileUpdateOne) SetAPIKey(a *APIKey) *SenderProfileUpdateOne {
	return spuo.SetAPIKeyID(a.ID)
}

// AddPaymentOrderIDs adds the "payment_orders" edge to the PaymentOrder entity by IDs.
func (spuo *SenderProfileUpdateOne) AddPaymentOrderIDs(ids ...uuid.UUID) *SenderProfileUpdateOne {
	spuo.mutation.AddPaymentOrderIDs(ids...)
	return spuo
}

// AddPaymentOrders adds the "payment_orders" edges to the PaymentOrder entity.
func (spuo *SenderProfileUpdateOne) AddPaymentOrders(p ...*PaymentOrder) *SenderProfileUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return spuo.AddPaymentOrderIDs(ids...)
}

// Mutation returns the SenderProfileMutation object of the builder.
func (spuo *SenderProfileUpdateOne) Mutation() *SenderProfileMutation {
	return spuo.mutation
}

// ClearAPIKey clears the "api_key" edge to the APIKey entity.
func (spuo *SenderProfileUpdateOne) ClearAPIKey() *SenderProfileUpdateOne {
	spuo.mutation.ClearAPIKey()
	return spuo
}

// ClearPaymentOrders clears all "payment_orders" edges to the PaymentOrder entity.
func (spuo *SenderProfileUpdateOne) ClearPaymentOrders() *SenderProfileUpdateOne {
	spuo.mutation.ClearPaymentOrders()
	return spuo
}

// RemovePaymentOrderIDs removes the "payment_orders" edge to PaymentOrder entities by IDs.
func (spuo *SenderProfileUpdateOne) RemovePaymentOrderIDs(ids ...uuid.UUID) *SenderProfileUpdateOne {
	spuo.mutation.RemovePaymentOrderIDs(ids...)
	return spuo
}

// RemovePaymentOrders removes "payment_orders" edges to PaymentOrder entities.
func (spuo *SenderProfileUpdateOne) RemovePaymentOrders(p ...*PaymentOrder) *SenderProfileUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return spuo.RemovePaymentOrderIDs(ids...)
}

// Where appends a list predicates to the SenderProfileUpdate builder.
func (spuo *SenderProfileUpdateOne) Where(ps ...predicate.SenderProfile) *SenderProfileUpdateOne {
	spuo.mutation.Where(ps...)
	return spuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spuo *SenderProfileUpdateOne) Select(field string, fields ...string) *SenderProfileUpdateOne {
	spuo.fields = append([]string{field}, fields...)
	return spuo
}

// Save executes the query and returns the updated SenderProfile entity.
func (spuo *SenderProfileUpdateOne) Save(ctx context.Context) (*SenderProfile, error) {
	spuo.defaults()
	return withHooks(ctx, spuo.sqlSave, spuo.mutation, spuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (spuo *SenderProfileUpdateOne) SaveX(ctx context.Context) *SenderProfile {
	node, err := spuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spuo *SenderProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := spuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spuo *SenderProfileUpdateOne) ExecX(ctx context.Context) {
	if err := spuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spuo *SenderProfileUpdateOne) defaults() {
	if _, ok := spuo.mutation.UpdatedAt(); !ok {
		v := senderprofile.UpdateDefaultUpdatedAt()
		spuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spuo *SenderProfileUpdateOne) check() error {
	if _, ok := spuo.mutation.UserID(); spuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SenderProfile.user"`)
	}
	return nil
}

func (spuo *SenderProfileUpdateOne) sqlSave(ctx context.Context) (_node *SenderProfile, err error) {
	if err := spuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(senderprofile.Table, senderprofile.Columns, sqlgraph.NewFieldSpec(senderprofile.FieldID, field.TypeUUID))
	id, ok := spuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SenderProfile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := spuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, senderprofile.FieldID)
		for _, f := range fields {
			if !senderprofile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != senderprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spuo.mutation.WebhookURL(); ok {
		_spec.SetField(senderprofile.FieldWebhookURL, field.TypeString, value)
	}
	if spuo.mutation.WebhookURLCleared() {
		_spec.ClearField(senderprofile.FieldWebhookURL, field.TypeString)
	}
	if value, ok := spuo.mutation.FeePerTokenUnit(); ok {
		_spec.SetField(senderprofile.FieldFeePerTokenUnit, field.TypeFloat64, value)
	}
	if value, ok := spuo.mutation.AddedFeePerTokenUnit(); ok {
		_spec.AddField(senderprofile.FieldFeePerTokenUnit, field.TypeFloat64, value)
	}
	if value, ok := spuo.mutation.FeeAddress(); ok {
		_spec.SetField(senderprofile.FieldFeeAddress, field.TypeString, value)
	}
	if spuo.mutation.FeeAddressCleared() {
		_spec.ClearField(senderprofile.FieldFeeAddress, field.TypeString)
	}
	if value, ok := spuo.mutation.RefundAddress(); ok {
		_spec.SetField(senderprofile.FieldRefundAddress, field.TypeString, value)
	}
	if spuo.mutation.RefundAddressCleared() {
		_spec.ClearField(senderprofile.FieldRefundAddress, field.TypeString)
	}
	if value, ok := spuo.mutation.DomainWhitelist(); ok {
		_spec.SetField(senderprofile.FieldDomainWhitelist, field.TypeJSON, value)
	}
	if value, ok := spuo.mutation.AppendedDomainWhitelist(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, senderprofile.FieldDomainWhitelist, value)
		})
	}
	if value, ok := spuo.mutation.IsActive(); ok {
		_spec.SetField(senderprofile.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := spuo.mutation.UpdatedAt(); ok {
		_spec.SetField(senderprofile.FieldUpdatedAt, field.TypeTime, value)
	}
	if spuo.mutation.APIKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   senderprofile.APIKeyTable,
			Columns: []string{senderprofile.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.APIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   senderprofile.APIKeyTable,
			Columns: []string{senderprofile.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if spuo.mutation.PaymentOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   senderprofile.PaymentOrdersTable,
			Columns: []string{senderprofile.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.RemovedPaymentOrdersIDs(); len(nodes) > 0 && !spuo.mutation.PaymentOrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   senderprofile.PaymentOrdersTable,
			Columns: []string{senderprofile.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.PaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   senderprofile.PaymentOrdersTable,
			Columns: []string{senderprofile.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SenderProfile{config: spuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{senderprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	spuo.mutation.done = true
	return _node, nil
}
