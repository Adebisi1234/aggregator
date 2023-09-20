// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/paycrest-protocol/ent/lockpaymentorder"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
	"github.com/paycrest/paycrest-protocol/ent/provisionbucket"
	"github.com/paycrest/paycrest-protocol/ent/token"
	"github.com/shopspring/decimal"
)

// LockPaymentOrderUpdate is the builder for updating LockPaymentOrder entities.
type LockPaymentOrderUpdate struct {
	config
	hooks    []Hook
	mutation *LockPaymentOrderMutation
}

// Where appends a list predicates to the LockPaymentOrderUpdate builder.
func (lpou *LockPaymentOrderUpdate) Where(ps ...predicate.LockPaymentOrder) *LockPaymentOrderUpdate {
	lpou.mutation.Where(ps...)
	return lpou
}

// SetUpdatedAt sets the "updated_at" field.
func (lpou *LockPaymentOrderUpdate) SetUpdatedAt(t time.Time) *LockPaymentOrderUpdate {
	lpou.mutation.SetUpdatedAt(t)
	return lpou
}

// SetOrderID sets the "order_id" field.
func (lpou *LockPaymentOrderUpdate) SetOrderID(s string) *LockPaymentOrderUpdate {
	lpou.mutation.SetOrderID(s)
	return lpou
}

// SetAmount sets the "amount" field.
func (lpou *LockPaymentOrderUpdate) SetAmount(d decimal.Decimal) *LockPaymentOrderUpdate {
	lpou.mutation.ResetAmount()
	lpou.mutation.SetAmount(d)
	return lpou
}

// AddAmount adds d to the "amount" field.
func (lpou *LockPaymentOrderUpdate) AddAmount(d decimal.Decimal) *LockPaymentOrderUpdate {
	lpou.mutation.AddAmount(d)
	return lpou
}

// SetRate sets the "rate" field.
func (lpou *LockPaymentOrderUpdate) SetRate(d decimal.Decimal) *LockPaymentOrderUpdate {
	lpou.mutation.ResetRate()
	lpou.mutation.SetRate(d)
	return lpou
}

// AddRate adds d to the "rate" field.
func (lpou *LockPaymentOrderUpdate) AddRate(d decimal.Decimal) *LockPaymentOrderUpdate {
	lpou.mutation.AddRate(d)
	return lpou
}

// SetTxHash sets the "tx_hash" field.
func (lpou *LockPaymentOrderUpdate) SetTxHash(s string) *LockPaymentOrderUpdate {
	lpou.mutation.SetTxHash(s)
	return lpou
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (lpou *LockPaymentOrderUpdate) SetNillableTxHash(s *string) *LockPaymentOrderUpdate {
	if s != nil {
		lpou.SetTxHash(*s)
	}
	return lpou
}

// ClearTxHash clears the value of the "tx_hash" field.
func (lpou *LockPaymentOrderUpdate) ClearTxHash() *LockPaymentOrderUpdate {
	lpou.mutation.ClearTxHash()
	return lpou
}

// SetStatus sets the "status" field.
func (lpou *LockPaymentOrderUpdate) SetStatus(l lockpaymentorder.Status) *LockPaymentOrderUpdate {
	lpou.mutation.SetStatus(l)
	return lpou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (lpou *LockPaymentOrderUpdate) SetNillableStatus(l *lockpaymentorder.Status) *LockPaymentOrderUpdate {
	if l != nil {
		lpou.SetStatus(*l)
	}
	return lpou
}

// SetBlockNumber sets the "block_number" field.
func (lpou *LockPaymentOrderUpdate) SetBlockNumber(i int64) *LockPaymentOrderUpdate {
	lpou.mutation.ResetBlockNumber()
	lpou.mutation.SetBlockNumber(i)
	return lpou
}

// AddBlockNumber adds i to the "block_number" field.
func (lpou *LockPaymentOrderUpdate) AddBlockNumber(i int64) *LockPaymentOrderUpdate {
	lpou.mutation.AddBlockNumber(i)
	return lpou
}

// SetInstitution sets the "institution" field.
func (lpou *LockPaymentOrderUpdate) SetInstitution(s string) *LockPaymentOrderUpdate {
	lpou.mutation.SetInstitution(s)
	return lpou
}

// SetAccountIdentifier sets the "account_identifier" field.
func (lpou *LockPaymentOrderUpdate) SetAccountIdentifier(s string) *LockPaymentOrderUpdate {
	lpou.mutation.SetAccountIdentifier(s)
	return lpou
}

// SetAccountName sets the "account_name" field.
func (lpou *LockPaymentOrderUpdate) SetAccountName(s string) *LockPaymentOrderUpdate {
	lpou.mutation.SetAccountName(s)
	return lpou
}

// SetTokenID sets the "token" edge to the Token entity by ID.
func (lpou *LockPaymentOrderUpdate) SetTokenID(id int) *LockPaymentOrderUpdate {
	lpou.mutation.SetTokenID(id)
	return lpou
}

// SetToken sets the "token" edge to the Token entity.
func (lpou *LockPaymentOrderUpdate) SetToken(t *Token) *LockPaymentOrderUpdate {
	return lpou.SetTokenID(t.ID)
}

// SetProvisionBucketID sets the "provision_bucket" edge to the ProvisionBucket entity by ID.
func (lpou *LockPaymentOrderUpdate) SetProvisionBucketID(id int) *LockPaymentOrderUpdate {
	lpou.mutation.SetProvisionBucketID(id)
	return lpou
}

// SetNillableProvisionBucketID sets the "provision_bucket" edge to the ProvisionBucket entity by ID if the given value is not nil.
func (lpou *LockPaymentOrderUpdate) SetNillableProvisionBucketID(id *int) *LockPaymentOrderUpdate {
	if id != nil {
		lpou = lpou.SetProvisionBucketID(*id)
	}
	return lpou
}

// SetProvisionBucket sets the "provision_bucket" edge to the ProvisionBucket entity.
func (lpou *LockPaymentOrderUpdate) SetProvisionBucket(p *ProvisionBucket) *LockPaymentOrderUpdate {
	return lpou.SetProvisionBucketID(p.ID)
}

// SetProviderID sets the "provider" edge to the ProviderProfile entity by ID.
func (lpou *LockPaymentOrderUpdate) SetProviderID(id string) *LockPaymentOrderUpdate {
	lpou.mutation.SetProviderID(id)
	return lpou
}

// SetNillableProviderID sets the "provider" edge to the ProviderProfile entity by ID if the given value is not nil.
func (lpou *LockPaymentOrderUpdate) SetNillableProviderID(id *string) *LockPaymentOrderUpdate {
	if id != nil {
		lpou = lpou.SetProviderID(*id)
	}
	return lpou
}

// SetProvider sets the "provider" edge to the ProviderProfile entity.
func (lpou *LockPaymentOrderUpdate) SetProvider(p *ProviderProfile) *LockPaymentOrderUpdate {
	return lpou.SetProviderID(p.ID)
}

// Mutation returns the LockPaymentOrderMutation object of the builder.
func (lpou *LockPaymentOrderUpdate) Mutation() *LockPaymentOrderMutation {
	return lpou.mutation
}

// ClearToken clears the "token" edge to the Token entity.
func (lpou *LockPaymentOrderUpdate) ClearToken() *LockPaymentOrderUpdate {
	lpou.mutation.ClearToken()
	return lpou
}

// ClearProvisionBucket clears the "provision_bucket" edge to the ProvisionBucket entity.
func (lpou *LockPaymentOrderUpdate) ClearProvisionBucket() *LockPaymentOrderUpdate {
	lpou.mutation.ClearProvisionBucket()
	return lpou
}

// ClearProvider clears the "provider" edge to the ProviderProfile entity.
func (lpou *LockPaymentOrderUpdate) ClearProvider() *LockPaymentOrderUpdate {
	lpou.mutation.ClearProvider()
	return lpou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lpou *LockPaymentOrderUpdate) Save(ctx context.Context) (int, error) {
	lpou.defaults()
	return withHooks(ctx, lpou.sqlSave, lpou.mutation, lpou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lpou *LockPaymentOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := lpou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lpou *LockPaymentOrderUpdate) Exec(ctx context.Context) error {
	_, err := lpou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpou *LockPaymentOrderUpdate) ExecX(ctx context.Context) {
	if err := lpou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lpou *LockPaymentOrderUpdate) defaults() {
	if _, ok := lpou.mutation.UpdatedAt(); !ok {
		v := lockpaymentorder.UpdateDefaultUpdatedAt()
		lpou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lpou *LockPaymentOrderUpdate) check() error {
	if v, ok := lpou.mutation.TxHash(); ok {
		if err := lockpaymentorder.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf(`ent: validator failed for field "LockPaymentOrder.tx_hash": %w`, err)}
		}
	}
	if v, ok := lpou.mutation.Status(); ok {
		if err := lockpaymentorder.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "LockPaymentOrder.status": %w`, err)}
		}
	}
	if _, ok := lpou.mutation.TokenID(); lpou.mutation.TokenCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LockPaymentOrder.token"`)
	}
	return nil
}

func (lpou *LockPaymentOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lpou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lockpaymentorder.Table, lockpaymentorder.Columns, sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID))
	if ps := lpou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lpou.mutation.UpdatedAt(); ok {
		_spec.SetField(lockpaymentorder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lpou.mutation.OrderID(); ok {
		_spec.SetField(lockpaymentorder.FieldOrderID, field.TypeString, value)
	}
	if value, ok := lpou.mutation.Amount(); ok {
		_spec.SetField(lockpaymentorder.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := lpou.mutation.AddedAmount(); ok {
		_spec.AddField(lockpaymentorder.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := lpou.mutation.Rate(); ok {
		_spec.SetField(lockpaymentorder.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := lpou.mutation.AddedRate(); ok {
		_spec.AddField(lockpaymentorder.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := lpou.mutation.TxHash(); ok {
		_spec.SetField(lockpaymentorder.FieldTxHash, field.TypeString, value)
	}
	if lpou.mutation.TxHashCleared() {
		_spec.ClearField(lockpaymentorder.FieldTxHash, field.TypeString)
	}
	if value, ok := lpou.mutation.Status(); ok {
		_spec.SetField(lockpaymentorder.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := lpou.mutation.BlockNumber(); ok {
		_spec.SetField(lockpaymentorder.FieldBlockNumber, field.TypeInt64, value)
	}
	if value, ok := lpou.mutation.AddedBlockNumber(); ok {
		_spec.AddField(lockpaymentorder.FieldBlockNumber, field.TypeInt64, value)
	}
	if value, ok := lpou.mutation.Institution(); ok {
		_spec.SetField(lockpaymentorder.FieldInstitution, field.TypeString, value)
	}
	if value, ok := lpou.mutation.AccountIdentifier(); ok {
		_spec.SetField(lockpaymentorder.FieldAccountIdentifier, field.TypeString, value)
	}
	if value, ok := lpou.mutation.AccountName(); ok {
		_spec.SetField(lockpaymentorder.FieldAccountName, field.TypeString, value)
	}
	if lpou.mutation.TokenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpou.mutation.TokenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lpou.mutation.ProvisionBucketCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpou.mutation.ProvisionBucketIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lpou.mutation.ProviderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpou.mutation.ProviderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lpou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lockpaymentorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lpou.mutation.done = true
	return n, nil
}

// LockPaymentOrderUpdateOne is the builder for updating a single LockPaymentOrder entity.
type LockPaymentOrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LockPaymentOrderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (lpouo *LockPaymentOrderUpdateOne) SetUpdatedAt(t time.Time) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetUpdatedAt(t)
	return lpouo
}

// SetOrderID sets the "order_id" field.
func (lpouo *LockPaymentOrderUpdateOne) SetOrderID(s string) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetOrderID(s)
	return lpouo
}

// SetAmount sets the "amount" field.
func (lpouo *LockPaymentOrderUpdateOne) SetAmount(d decimal.Decimal) *LockPaymentOrderUpdateOne {
	lpouo.mutation.ResetAmount()
	lpouo.mutation.SetAmount(d)
	return lpouo
}

// AddAmount adds d to the "amount" field.
func (lpouo *LockPaymentOrderUpdateOne) AddAmount(d decimal.Decimal) *LockPaymentOrderUpdateOne {
	lpouo.mutation.AddAmount(d)
	return lpouo
}

// SetRate sets the "rate" field.
func (lpouo *LockPaymentOrderUpdateOne) SetRate(d decimal.Decimal) *LockPaymentOrderUpdateOne {
	lpouo.mutation.ResetRate()
	lpouo.mutation.SetRate(d)
	return lpouo
}

// AddRate adds d to the "rate" field.
func (lpouo *LockPaymentOrderUpdateOne) AddRate(d decimal.Decimal) *LockPaymentOrderUpdateOne {
	lpouo.mutation.AddRate(d)
	return lpouo
}

// SetTxHash sets the "tx_hash" field.
func (lpouo *LockPaymentOrderUpdateOne) SetTxHash(s string) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetTxHash(s)
	return lpouo
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (lpouo *LockPaymentOrderUpdateOne) SetNillableTxHash(s *string) *LockPaymentOrderUpdateOne {
	if s != nil {
		lpouo.SetTxHash(*s)
	}
	return lpouo
}

// ClearTxHash clears the value of the "tx_hash" field.
func (lpouo *LockPaymentOrderUpdateOne) ClearTxHash() *LockPaymentOrderUpdateOne {
	lpouo.mutation.ClearTxHash()
	return lpouo
}

// SetStatus sets the "status" field.
func (lpouo *LockPaymentOrderUpdateOne) SetStatus(l lockpaymentorder.Status) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetStatus(l)
	return lpouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (lpouo *LockPaymentOrderUpdateOne) SetNillableStatus(l *lockpaymentorder.Status) *LockPaymentOrderUpdateOne {
	if l != nil {
		lpouo.SetStatus(*l)
	}
	return lpouo
}

// SetBlockNumber sets the "block_number" field.
func (lpouo *LockPaymentOrderUpdateOne) SetBlockNumber(i int64) *LockPaymentOrderUpdateOne {
	lpouo.mutation.ResetBlockNumber()
	lpouo.mutation.SetBlockNumber(i)
	return lpouo
}

// AddBlockNumber adds i to the "block_number" field.
func (lpouo *LockPaymentOrderUpdateOne) AddBlockNumber(i int64) *LockPaymentOrderUpdateOne {
	lpouo.mutation.AddBlockNumber(i)
	return lpouo
}

// SetInstitution sets the "institution" field.
func (lpouo *LockPaymentOrderUpdateOne) SetInstitution(s string) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetInstitution(s)
	return lpouo
}

// SetAccountIdentifier sets the "account_identifier" field.
func (lpouo *LockPaymentOrderUpdateOne) SetAccountIdentifier(s string) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetAccountIdentifier(s)
	return lpouo
}

// SetAccountName sets the "account_name" field.
func (lpouo *LockPaymentOrderUpdateOne) SetAccountName(s string) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetAccountName(s)
	return lpouo
}

// SetTokenID sets the "token" edge to the Token entity by ID.
func (lpouo *LockPaymentOrderUpdateOne) SetTokenID(id int) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetTokenID(id)
	return lpouo
}

// SetToken sets the "token" edge to the Token entity.
func (lpouo *LockPaymentOrderUpdateOne) SetToken(t *Token) *LockPaymentOrderUpdateOne {
	return lpouo.SetTokenID(t.ID)
}

// SetProvisionBucketID sets the "provision_bucket" edge to the ProvisionBucket entity by ID.
func (lpouo *LockPaymentOrderUpdateOne) SetProvisionBucketID(id int) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetProvisionBucketID(id)
	return lpouo
}

// SetNillableProvisionBucketID sets the "provision_bucket" edge to the ProvisionBucket entity by ID if the given value is not nil.
func (lpouo *LockPaymentOrderUpdateOne) SetNillableProvisionBucketID(id *int) *LockPaymentOrderUpdateOne {
	if id != nil {
		lpouo = lpouo.SetProvisionBucketID(*id)
	}
	return lpouo
}

// SetProvisionBucket sets the "provision_bucket" edge to the ProvisionBucket entity.
func (lpouo *LockPaymentOrderUpdateOne) SetProvisionBucket(p *ProvisionBucket) *LockPaymentOrderUpdateOne {
	return lpouo.SetProvisionBucketID(p.ID)
}

// SetProviderID sets the "provider" edge to the ProviderProfile entity by ID.
func (lpouo *LockPaymentOrderUpdateOne) SetProviderID(id string) *LockPaymentOrderUpdateOne {
	lpouo.mutation.SetProviderID(id)
	return lpouo
}

// SetNillableProviderID sets the "provider" edge to the ProviderProfile entity by ID if the given value is not nil.
func (lpouo *LockPaymentOrderUpdateOne) SetNillableProviderID(id *string) *LockPaymentOrderUpdateOne {
	if id != nil {
		lpouo = lpouo.SetProviderID(*id)
	}
	return lpouo
}

// SetProvider sets the "provider" edge to the ProviderProfile entity.
func (lpouo *LockPaymentOrderUpdateOne) SetProvider(p *ProviderProfile) *LockPaymentOrderUpdateOne {
	return lpouo.SetProviderID(p.ID)
}

// Mutation returns the LockPaymentOrderMutation object of the builder.
func (lpouo *LockPaymentOrderUpdateOne) Mutation() *LockPaymentOrderMutation {
	return lpouo.mutation
}

// ClearToken clears the "token" edge to the Token entity.
func (lpouo *LockPaymentOrderUpdateOne) ClearToken() *LockPaymentOrderUpdateOne {
	lpouo.mutation.ClearToken()
	return lpouo
}

// ClearProvisionBucket clears the "provision_bucket" edge to the ProvisionBucket entity.
func (lpouo *LockPaymentOrderUpdateOne) ClearProvisionBucket() *LockPaymentOrderUpdateOne {
	lpouo.mutation.ClearProvisionBucket()
	return lpouo
}

// ClearProvider clears the "provider" edge to the ProviderProfile entity.
func (lpouo *LockPaymentOrderUpdateOne) ClearProvider() *LockPaymentOrderUpdateOne {
	lpouo.mutation.ClearProvider()
	return lpouo
}

// Where appends a list predicates to the LockPaymentOrderUpdate builder.
func (lpouo *LockPaymentOrderUpdateOne) Where(ps ...predicate.LockPaymentOrder) *LockPaymentOrderUpdateOne {
	lpouo.mutation.Where(ps...)
	return lpouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lpouo *LockPaymentOrderUpdateOne) Select(field string, fields ...string) *LockPaymentOrderUpdateOne {
	lpouo.fields = append([]string{field}, fields...)
	return lpouo
}

// Save executes the query and returns the updated LockPaymentOrder entity.
func (lpouo *LockPaymentOrderUpdateOne) Save(ctx context.Context) (*LockPaymentOrder, error) {
	lpouo.defaults()
	return withHooks(ctx, lpouo.sqlSave, lpouo.mutation, lpouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lpouo *LockPaymentOrderUpdateOne) SaveX(ctx context.Context) *LockPaymentOrder {
	node, err := lpouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lpouo *LockPaymentOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := lpouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpouo *LockPaymentOrderUpdateOne) ExecX(ctx context.Context) {
	if err := lpouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lpouo *LockPaymentOrderUpdateOne) defaults() {
	if _, ok := lpouo.mutation.UpdatedAt(); !ok {
		v := lockpaymentorder.UpdateDefaultUpdatedAt()
		lpouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lpouo *LockPaymentOrderUpdateOne) check() error {
	if v, ok := lpouo.mutation.TxHash(); ok {
		if err := lockpaymentorder.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf(`ent: validator failed for field "LockPaymentOrder.tx_hash": %w`, err)}
		}
	}
	if v, ok := lpouo.mutation.Status(); ok {
		if err := lockpaymentorder.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "LockPaymentOrder.status": %w`, err)}
		}
	}
	if _, ok := lpouo.mutation.TokenID(); lpouo.mutation.TokenCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LockPaymentOrder.token"`)
	}
	return nil
}

func (lpouo *LockPaymentOrderUpdateOne) sqlSave(ctx context.Context) (_node *LockPaymentOrder, err error) {
	if err := lpouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lockpaymentorder.Table, lockpaymentorder.Columns, sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID))
	id, ok := lpouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LockPaymentOrder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lpouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lockpaymentorder.FieldID)
		for _, f := range fields {
			if !lockpaymentorder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lockpaymentorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lpouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lpouo.mutation.UpdatedAt(); ok {
		_spec.SetField(lockpaymentorder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lpouo.mutation.OrderID(); ok {
		_spec.SetField(lockpaymentorder.FieldOrderID, field.TypeString, value)
	}
	if value, ok := lpouo.mutation.Amount(); ok {
		_spec.SetField(lockpaymentorder.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := lpouo.mutation.AddedAmount(); ok {
		_spec.AddField(lockpaymentorder.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := lpouo.mutation.Rate(); ok {
		_spec.SetField(lockpaymentorder.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := lpouo.mutation.AddedRate(); ok {
		_spec.AddField(lockpaymentorder.FieldRate, field.TypeFloat64, value)
	}
	if value, ok := lpouo.mutation.TxHash(); ok {
		_spec.SetField(lockpaymentorder.FieldTxHash, field.TypeString, value)
	}
	if lpouo.mutation.TxHashCleared() {
		_spec.ClearField(lockpaymentorder.FieldTxHash, field.TypeString)
	}
	if value, ok := lpouo.mutation.Status(); ok {
		_spec.SetField(lockpaymentorder.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := lpouo.mutation.BlockNumber(); ok {
		_spec.SetField(lockpaymentorder.FieldBlockNumber, field.TypeInt64, value)
	}
	if value, ok := lpouo.mutation.AddedBlockNumber(); ok {
		_spec.AddField(lockpaymentorder.FieldBlockNumber, field.TypeInt64, value)
	}
	if value, ok := lpouo.mutation.Institution(); ok {
		_spec.SetField(lockpaymentorder.FieldInstitution, field.TypeString, value)
	}
	if value, ok := lpouo.mutation.AccountIdentifier(); ok {
		_spec.SetField(lockpaymentorder.FieldAccountIdentifier, field.TypeString, value)
	}
	if value, ok := lpouo.mutation.AccountName(); ok {
		_spec.SetField(lockpaymentorder.FieldAccountName, field.TypeString, value)
	}
	if lpouo.mutation.TokenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpouo.mutation.TokenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lpouo.mutation.ProvisionBucketCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpouo.mutation.ProvisionBucketIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lpouo.mutation.ProviderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpouo.mutation.ProviderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LockPaymentOrder{config: lpouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lpouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lockpaymentorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lpouo.mutation.done = true
	return _node, nil
}
