// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/lockorderfulfillment"
	"github.com/paycrest/protocol/ent/lockpaymentorder"
	"github.com/paycrest/protocol/ent/providerprofile"
	"github.com/paycrest/protocol/ent/provisionbucket"
	"github.com/paycrest/protocol/ent/token"
	"github.com/shopspring/decimal"
)

// LockPaymentOrderCreate is the builder for creating a LockPaymentOrder entity.
type LockPaymentOrderCreate struct {
	config
	mutation *LockPaymentOrderMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lpoc *LockPaymentOrderCreate) SetCreatedAt(t time.Time) *LockPaymentOrderCreate {
	lpoc.mutation.SetCreatedAt(t)
	return lpoc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableCreatedAt(t *time.Time) *LockPaymentOrderCreate {
	if t != nil {
		lpoc.SetCreatedAt(*t)
	}
	return lpoc
}

// SetUpdatedAt sets the "updated_at" field.
func (lpoc *LockPaymentOrderCreate) SetUpdatedAt(t time.Time) *LockPaymentOrderCreate {
	lpoc.mutation.SetUpdatedAt(t)
	return lpoc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableUpdatedAt(t *time.Time) *LockPaymentOrderCreate {
	if t != nil {
		lpoc.SetUpdatedAt(*t)
	}
	return lpoc
}

// SetGatewayID sets the "gateway_id" field.
func (lpoc *LockPaymentOrderCreate) SetGatewayID(s string) *LockPaymentOrderCreate {
	lpoc.mutation.SetGatewayID(s)
	return lpoc
}

// SetNillableGatewayID sets the "gateway_id" field if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableGatewayID(s *string) *LockPaymentOrderCreate {
	if s != nil {
		lpoc.SetGatewayID(*s)
	}
	return lpoc
}

// SetAmount sets the "amount" field.
func (lpoc *LockPaymentOrderCreate) SetAmount(d decimal.Decimal) *LockPaymentOrderCreate {
	lpoc.mutation.SetAmount(d)
	return lpoc
}

// SetRate sets the "rate" field.
func (lpoc *LockPaymentOrderCreate) SetRate(d decimal.Decimal) *LockPaymentOrderCreate {
	lpoc.mutation.SetRate(d)
	return lpoc
}

// SetOrderPercent sets the "order_percent" field.
func (lpoc *LockPaymentOrderCreate) SetOrderPercent(d decimal.Decimal) *LockPaymentOrderCreate {
	lpoc.mutation.SetOrderPercent(d)
	return lpoc
}

// SetTxHash sets the "tx_hash" field.
func (lpoc *LockPaymentOrderCreate) SetTxHash(s string) *LockPaymentOrderCreate {
	lpoc.mutation.SetTxHash(s)
	return lpoc
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableTxHash(s *string) *LockPaymentOrderCreate {
	if s != nil {
		lpoc.SetTxHash(*s)
	}
	return lpoc
}

// SetStatus sets the "status" field.
func (lpoc *LockPaymentOrderCreate) SetStatus(l lockpaymentorder.Status) *LockPaymentOrderCreate {
	lpoc.mutation.SetStatus(l)
	return lpoc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableStatus(l *lockpaymentorder.Status) *LockPaymentOrderCreate {
	if l != nil {
		lpoc.SetStatus(*l)
	}
	return lpoc
}

// SetBlockNumber sets the "block_number" field.
func (lpoc *LockPaymentOrderCreate) SetBlockNumber(i int64) *LockPaymentOrderCreate {
	lpoc.mutation.SetBlockNumber(i)
	return lpoc
}

// SetInstitution sets the "institution" field.
func (lpoc *LockPaymentOrderCreate) SetInstitution(s string) *LockPaymentOrderCreate {
	lpoc.mutation.SetInstitution(s)
	return lpoc
}

// SetAccountIdentifier sets the "account_identifier" field.
func (lpoc *LockPaymentOrderCreate) SetAccountIdentifier(s string) *LockPaymentOrderCreate {
	lpoc.mutation.SetAccountIdentifier(s)
	return lpoc
}

// SetAccountName sets the "account_name" field.
func (lpoc *LockPaymentOrderCreate) SetAccountName(s string) *LockPaymentOrderCreate {
	lpoc.mutation.SetAccountName(s)
	return lpoc
}

// SetMemo sets the "memo" field.
func (lpoc *LockPaymentOrderCreate) SetMemo(s string) *LockPaymentOrderCreate {
	lpoc.mutation.SetMemo(s)
	return lpoc
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableMemo(s *string) *LockPaymentOrderCreate {
	if s != nil {
		lpoc.SetMemo(*s)
	}
	return lpoc
}

// SetCancellationCount sets the "cancellation_count" field.
func (lpoc *LockPaymentOrderCreate) SetCancellationCount(i int) *LockPaymentOrderCreate {
	lpoc.mutation.SetCancellationCount(i)
	return lpoc
}

// SetNillableCancellationCount sets the "cancellation_count" field if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableCancellationCount(i *int) *LockPaymentOrderCreate {
	if i != nil {
		lpoc.SetCancellationCount(*i)
	}
	return lpoc
}

// SetCancellationReasons sets the "cancellation_reasons" field.
func (lpoc *LockPaymentOrderCreate) SetCancellationReasons(s []string) *LockPaymentOrderCreate {
	lpoc.mutation.SetCancellationReasons(s)
	return lpoc
}

// SetID sets the "id" field.
func (lpoc *LockPaymentOrderCreate) SetID(u uuid.UUID) *LockPaymentOrderCreate {
	lpoc.mutation.SetID(u)
	return lpoc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableID(u *uuid.UUID) *LockPaymentOrderCreate {
	if u != nil {
		lpoc.SetID(*u)
	}
	return lpoc
}

// SetTokenID sets the "token" edge to the Token entity by ID.
func (lpoc *LockPaymentOrderCreate) SetTokenID(id int) *LockPaymentOrderCreate {
	lpoc.mutation.SetTokenID(id)
	return lpoc
}

// SetToken sets the "token" edge to the Token entity.
func (lpoc *LockPaymentOrderCreate) SetToken(t *Token) *LockPaymentOrderCreate {
	return lpoc.SetTokenID(t.ID)
}

// SetProvisionBucketID sets the "provision_bucket" edge to the ProvisionBucket entity by ID.
func (lpoc *LockPaymentOrderCreate) SetProvisionBucketID(id int) *LockPaymentOrderCreate {
	lpoc.mutation.SetProvisionBucketID(id)
	return lpoc
}

// SetNillableProvisionBucketID sets the "provision_bucket" edge to the ProvisionBucket entity by ID if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableProvisionBucketID(id *int) *LockPaymentOrderCreate {
	if id != nil {
		lpoc = lpoc.SetProvisionBucketID(*id)
	}
	return lpoc
}

// SetProvisionBucket sets the "provision_bucket" edge to the ProvisionBucket entity.
func (lpoc *LockPaymentOrderCreate) SetProvisionBucket(p *ProvisionBucket) *LockPaymentOrderCreate {
	return lpoc.SetProvisionBucketID(p.ID)
}

// SetProviderID sets the "provider" edge to the ProviderProfile entity by ID.
func (lpoc *LockPaymentOrderCreate) SetProviderID(id string) *LockPaymentOrderCreate {
	lpoc.mutation.SetProviderID(id)
	return lpoc
}

// SetNillableProviderID sets the "provider" edge to the ProviderProfile entity by ID if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableProviderID(id *string) *LockPaymentOrderCreate {
	if id != nil {
		lpoc = lpoc.SetProviderID(*id)
	}
	return lpoc
}

// SetProvider sets the "provider" edge to the ProviderProfile entity.
func (lpoc *LockPaymentOrderCreate) SetProvider(p *ProviderProfile) *LockPaymentOrderCreate {
	return lpoc.SetProviderID(p.ID)
}

// SetFulfillmentID sets the "fulfillment" edge to the LockOrderFulfillment entity by ID.
func (lpoc *LockPaymentOrderCreate) SetFulfillmentID(id uuid.UUID) *LockPaymentOrderCreate {
	lpoc.mutation.SetFulfillmentID(id)
	return lpoc
}

// SetNillableFulfillmentID sets the "fulfillment" edge to the LockOrderFulfillment entity by ID if the given value is not nil.
func (lpoc *LockPaymentOrderCreate) SetNillableFulfillmentID(id *uuid.UUID) *LockPaymentOrderCreate {
	if id != nil {
		lpoc = lpoc.SetFulfillmentID(*id)
	}
	return lpoc
}

// SetFulfillment sets the "fulfillment" edge to the LockOrderFulfillment entity.
func (lpoc *LockPaymentOrderCreate) SetFulfillment(l *LockOrderFulfillment) *LockPaymentOrderCreate {
	return lpoc.SetFulfillmentID(l.ID)
}

// Mutation returns the LockPaymentOrderMutation object of the builder.
func (lpoc *LockPaymentOrderCreate) Mutation() *LockPaymentOrderMutation {
	return lpoc.mutation
}

// Save creates the LockPaymentOrder in the database.
func (lpoc *LockPaymentOrderCreate) Save(ctx context.Context) (*LockPaymentOrder, error) {
	lpoc.defaults()
	return withHooks(ctx, lpoc.sqlSave, lpoc.mutation, lpoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lpoc *LockPaymentOrderCreate) SaveX(ctx context.Context) *LockPaymentOrder {
	v, err := lpoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lpoc *LockPaymentOrderCreate) Exec(ctx context.Context) error {
	_, err := lpoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpoc *LockPaymentOrderCreate) ExecX(ctx context.Context) {
	if err := lpoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lpoc *LockPaymentOrderCreate) defaults() {
	if _, ok := lpoc.mutation.CreatedAt(); !ok {
		v := lockpaymentorder.DefaultCreatedAt()
		lpoc.mutation.SetCreatedAt(v)
	}
	if _, ok := lpoc.mutation.UpdatedAt(); !ok {
		v := lockpaymentorder.DefaultUpdatedAt()
		lpoc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lpoc.mutation.GatewayID(); !ok {
		v := lockpaymentorder.DefaultGatewayID
		lpoc.mutation.SetGatewayID(v)
	}
	if _, ok := lpoc.mutation.Status(); !ok {
		v := lockpaymentorder.DefaultStatus
		lpoc.mutation.SetStatus(v)
	}
	if _, ok := lpoc.mutation.CancellationCount(); !ok {
		v := lockpaymentorder.DefaultCancellationCount
		lpoc.mutation.SetCancellationCount(v)
	}
	if _, ok := lpoc.mutation.CancellationReasons(); !ok {
		v := lockpaymentorder.DefaultCancellationReasons
		lpoc.mutation.SetCancellationReasons(v)
	}
	if _, ok := lpoc.mutation.ID(); !ok {
		v := lockpaymentorder.DefaultID()
		lpoc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lpoc *LockPaymentOrderCreate) check() error {
	if _, ok := lpoc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LockPaymentOrder.created_at"`)}
	}
	if _, ok := lpoc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LockPaymentOrder.updated_at"`)}
	}
	if _, ok := lpoc.mutation.GatewayID(); !ok {
		return &ValidationError{Name: "gateway_id", err: errors.New(`ent: missing required field "LockPaymentOrder.gateway_id"`)}
	}
	if _, ok := lpoc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "LockPaymentOrder.amount"`)}
	}
	if _, ok := lpoc.mutation.Rate(); !ok {
		return &ValidationError{Name: "rate", err: errors.New(`ent: missing required field "LockPaymentOrder.rate"`)}
	}
	if _, ok := lpoc.mutation.OrderPercent(); !ok {
		return &ValidationError{Name: "order_percent", err: errors.New(`ent: missing required field "LockPaymentOrder.order_percent"`)}
	}
	if v, ok := lpoc.mutation.TxHash(); ok {
		if err := lockpaymentorder.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf(`ent: validator failed for field "LockPaymentOrder.tx_hash": %w`, err)}
		}
	}
	if _, ok := lpoc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "LockPaymentOrder.status"`)}
	}
	if v, ok := lpoc.mutation.Status(); ok {
		if err := lockpaymentorder.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "LockPaymentOrder.status": %w`, err)}
		}
	}
	if _, ok := lpoc.mutation.BlockNumber(); !ok {
		return &ValidationError{Name: "block_number", err: errors.New(`ent: missing required field "LockPaymentOrder.block_number"`)}
	}
	if _, ok := lpoc.mutation.Institution(); !ok {
		return &ValidationError{Name: "institution", err: errors.New(`ent: missing required field "LockPaymentOrder.institution"`)}
	}
	if _, ok := lpoc.mutation.AccountIdentifier(); !ok {
		return &ValidationError{Name: "account_identifier", err: errors.New(`ent: missing required field "LockPaymentOrder.account_identifier"`)}
	}
	if _, ok := lpoc.mutation.AccountName(); !ok {
		return &ValidationError{Name: "account_name", err: errors.New(`ent: missing required field "LockPaymentOrder.account_name"`)}
	}
	if _, ok := lpoc.mutation.CancellationCount(); !ok {
		return &ValidationError{Name: "cancellation_count", err: errors.New(`ent: missing required field "LockPaymentOrder.cancellation_count"`)}
	}
	if _, ok := lpoc.mutation.CancellationReasons(); !ok {
		return &ValidationError{Name: "cancellation_reasons", err: errors.New(`ent: missing required field "LockPaymentOrder.cancellation_reasons"`)}
	}
	if _, ok := lpoc.mutation.TokenID(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required edge "LockPaymentOrder.token"`)}
	}
	return nil
}

func (lpoc *LockPaymentOrderCreate) sqlSave(ctx context.Context) (*LockPaymentOrder, error) {
	if err := lpoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lpoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lpoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	lpoc.mutation.id = &_node.ID
	lpoc.mutation.done = true
	return _node, nil
}

func (lpoc *LockPaymentOrderCreate) createSpec() (*LockPaymentOrder, *sqlgraph.CreateSpec) {
	var (
		_node = &LockPaymentOrder{config: lpoc.config}
		_spec = sqlgraph.NewCreateSpec(lockpaymentorder.Table, sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID))
	)
	if id, ok := lpoc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lpoc.mutation.CreatedAt(); ok {
		_spec.SetField(lockpaymentorder.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lpoc.mutation.UpdatedAt(); ok {
		_spec.SetField(lockpaymentorder.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lpoc.mutation.GatewayID(); ok {
		_spec.SetField(lockpaymentorder.FieldGatewayID, field.TypeString, value)
		_node.GatewayID = value
	}
	if value, ok := lpoc.mutation.Amount(); ok {
		_spec.SetField(lockpaymentorder.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := lpoc.mutation.Rate(); ok {
		_spec.SetField(lockpaymentorder.FieldRate, field.TypeFloat64, value)
		_node.Rate = value
	}
	if value, ok := lpoc.mutation.OrderPercent(); ok {
		_spec.SetField(lockpaymentorder.FieldOrderPercent, field.TypeFloat64, value)
		_node.OrderPercent = value
	}
	if value, ok := lpoc.mutation.TxHash(); ok {
		_spec.SetField(lockpaymentorder.FieldTxHash, field.TypeString, value)
		_node.TxHash = value
	}
	if value, ok := lpoc.mutation.Status(); ok {
		_spec.SetField(lockpaymentorder.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := lpoc.mutation.BlockNumber(); ok {
		_spec.SetField(lockpaymentorder.FieldBlockNumber, field.TypeInt64, value)
		_node.BlockNumber = value
	}
	if value, ok := lpoc.mutation.Institution(); ok {
		_spec.SetField(lockpaymentorder.FieldInstitution, field.TypeString, value)
		_node.Institution = value
	}
	if value, ok := lpoc.mutation.AccountIdentifier(); ok {
		_spec.SetField(lockpaymentorder.FieldAccountIdentifier, field.TypeString, value)
		_node.AccountIdentifier = value
	}
	if value, ok := lpoc.mutation.AccountName(); ok {
		_spec.SetField(lockpaymentorder.FieldAccountName, field.TypeString, value)
		_node.AccountName = value
	}
	if value, ok := lpoc.mutation.Memo(); ok {
		_spec.SetField(lockpaymentorder.FieldMemo, field.TypeString, value)
		_node.Memo = value
	}
	if value, ok := lpoc.mutation.CancellationCount(); ok {
		_spec.SetField(lockpaymentorder.FieldCancellationCount, field.TypeInt, value)
		_node.CancellationCount = value
	}
	if value, ok := lpoc.mutation.CancellationReasons(); ok {
		_spec.SetField(lockpaymentorder.FieldCancellationReasons, field.TypeJSON, value)
		_node.CancellationReasons = value
	}
	if nodes := lpoc.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lockpaymentorder.TokenTable,
			Columns: []string{lockpaymentorder.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.token_lock_payment_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lpoc.mutation.ProvisionBucketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lockpaymentorder.ProvisionBucketTable,
			Columns: []string{lockpaymentorder.ProvisionBucketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provisionbucket.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provision_bucket_lock_payment_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lpoc.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lockpaymentorder.ProviderTable,
			Columns: []string{lockpaymentorder.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provider_profile_assigned_orders = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lpoc.mutation.FulfillmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   lockpaymentorder.FulfillmentTable,
			Columns: []string{lockpaymentorder.FulfillmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockorderfulfillment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LockPaymentOrderCreateBulk is the builder for creating many LockPaymentOrder entities in bulk.
type LockPaymentOrderCreateBulk struct {
	config
	builders []*LockPaymentOrderCreate
}

// Save creates the LockPaymentOrder entities in the database.
func (lpocb *LockPaymentOrderCreateBulk) Save(ctx context.Context) ([]*LockPaymentOrder, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lpocb.builders))
	nodes := make([]*LockPaymentOrder, len(lpocb.builders))
	mutators := make([]Mutator, len(lpocb.builders))
	for i := range lpocb.builders {
		func(i int, root context.Context) {
			builder := lpocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LockPaymentOrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lpocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lpocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lpocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lpocb *LockPaymentOrderCreateBulk) SaveX(ctx context.Context) []*LockPaymentOrder {
	v, err := lpocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lpocb *LockPaymentOrderCreateBulk) Exec(ctx context.Context) error {
	_, err := lpocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpocb *LockPaymentOrderCreateBulk) ExecX(ctx context.Context) {
	if err := lpocb.Exec(ctx); err != nil {
		panic(err)
	}
}
