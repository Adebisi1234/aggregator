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
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/lockpaymentorder"
	"github.com/paycrest/protocol/ent/network"
	"github.com/paycrest/protocol/ent/paymentorder"
	"github.com/paycrest/protocol/ent/senderordertoken"
	"github.com/paycrest/protocol/ent/token"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tc *TokenCreate) SetCreatedAt(t time.Time) *TokenCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TokenCreate) SetNillableCreatedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TokenCreate) SetUpdatedAt(t time.Time) *TokenCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TokenCreate) SetNillableUpdatedAt(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetSymbol sets the "symbol" field.
func (tc *TokenCreate) SetSymbol(s string) *TokenCreate {
	tc.mutation.SetSymbol(s)
	return tc
}

// SetContractAddress sets the "contract_address" field.
func (tc *TokenCreate) SetContractAddress(s string) *TokenCreate {
	tc.mutation.SetContractAddress(s)
	return tc
}

// SetDecimals sets the "decimals" field.
func (tc *TokenCreate) SetDecimals(i int8) *TokenCreate {
	tc.mutation.SetDecimals(i)
	return tc
}

// SetIsEnabled sets the "is_enabled" field.
func (tc *TokenCreate) SetIsEnabled(b bool) *TokenCreate {
	tc.mutation.SetIsEnabled(b)
	return tc
}

// SetNillableIsEnabled sets the "is_enabled" field if the given value is not nil.
func (tc *TokenCreate) SetNillableIsEnabled(b *bool) *TokenCreate {
	if b != nil {
		tc.SetIsEnabled(*b)
	}
	return tc
}

// SetNetworkID sets the "network" edge to the Network entity by ID.
func (tc *TokenCreate) SetNetworkID(id int) *TokenCreate {
	tc.mutation.SetNetworkID(id)
	return tc
}

// SetNetwork sets the "network" edge to the Network entity.
func (tc *TokenCreate) SetNetwork(n *Network) *TokenCreate {
	return tc.SetNetworkID(n.ID)
}

// AddPaymentOrderIDs adds the "payment_orders" edge to the PaymentOrder entity by IDs.
func (tc *TokenCreate) AddPaymentOrderIDs(ids ...uuid.UUID) *TokenCreate {
	tc.mutation.AddPaymentOrderIDs(ids...)
	return tc
}

// AddPaymentOrders adds the "payment_orders" edges to the PaymentOrder entity.
func (tc *TokenCreate) AddPaymentOrders(p ...*PaymentOrder) *TokenCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddPaymentOrderIDs(ids...)
}

// AddLockPaymentOrderIDs adds the "lock_payment_orders" edge to the LockPaymentOrder entity by IDs.
func (tc *TokenCreate) AddLockPaymentOrderIDs(ids ...uuid.UUID) *TokenCreate {
	tc.mutation.AddLockPaymentOrderIDs(ids...)
	return tc
}

// AddLockPaymentOrders adds the "lock_payment_orders" edges to the LockPaymentOrder entity.
func (tc *TokenCreate) AddLockPaymentOrders(l ...*LockPaymentOrder) *TokenCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return tc.AddLockPaymentOrderIDs(ids...)
}

// AddSenderSettingIDs adds the "sender_settings" edge to the SenderOrderToken entity by IDs.
func (tc *TokenCreate) AddSenderSettingIDs(ids ...int) *TokenCreate {
	tc.mutation.AddSenderSettingIDs(ids...)
	return tc
}

// AddSenderSettings adds the "sender_settings" edges to the SenderOrderToken entity.
func (tc *TokenCreate) AddSenderSettings(s ...*SenderOrderToken) *TokenCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddSenderSettingIDs(ids...)
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokenCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokenCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokenCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := token.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := token.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.IsEnabled(); !ok {
		v := token.DefaultIsEnabled
		tc.mutation.SetIsEnabled(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Token.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Token.updated_at"`)}
	}
	if _, ok := tc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "Token.symbol"`)}
	}
	if v, ok := tc.mutation.Symbol(); ok {
		if err := token.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Token.symbol": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ContractAddress(); !ok {
		return &ValidationError{Name: "contract_address", err: errors.New(`ent: missing required field "Token.contract_address"`)}
	}
	if v, ok := tc.mutation.ContractAddress(); ok {
		if err := token.ContractAddressValidator(v); err != nil {
			return &ValidationError{Name: "contract_address", err: fmt.Errorf(`ent: validator failed for field "Token.contract_address": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Decimals(); !ok {
		return &ValidationError{Name: "decimals", err: errors.New(`ent: missing required field "Token.decimals"`)}
	}
	if _, ok := tc.mutation.IsEnabled(); !ok {
		return &ValidationError{Name: "is_enabled", err: errors.New(`ent: missing required field "Token.is_enabled"`)}
	}
	if len(tc.mutation.NetworkIDs()) == 0 {
		return &ValidationError{Name: "network", err: errors.New(`ent: missing required edge "Token.network"`)}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(token.Table, sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(token.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(token.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Symbol(); ok {
		_spec.SetField(token.FieldSymbol, field.TypeString, value)
		_node.Symbol = value
	}
	if value, ok := tc.mutation.ContractAddress(); ok {
		_spec.SetField(token.FieldContractAddress, field.TypeString, value)
		_node.ContractAddress = value
	}
	if value, ok := tc.mutation.Decimals(); ok {
		_spec.SetField(token.FieldDecimals, field.TypeInt8, value)
		_node.Decimals = value
	}
	if value, ok := tc.mutation.IsEnabled(); ok {
		_spec.SetField(token.FieldIsEnabled, field.TypeBool, value)
		_node.IsEnabled = value
	}
	if nodes := tc.mutation.NetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.NetworkTable,
			Columns: []string{token.NetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(network.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.network_tokens = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.PaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.PaymentOrdersTable,
			Columns: []string{token.PaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.LockPaymentOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.LockPaymentOrdersTable,
			Columns: []string{token.LockPaymentOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SenderSettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.SenderSettingsTable,
			Columns: []string{token.SenderSettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(senderordertoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tc *TokenCreate) OnConflict(opts ...sql.ConflictOption) *TokenUpsertOne {
	tc.conflict = opts
	return &TokenUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TokenCreate) OnConflictColumns(columns ...string) *TokenUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertOne{
		create: tc,
	}
}

type (
	// TokenUpsertOne is the builder for "upsert"-ing
	//  one Token node.
	TokenUpsertOne struct {
		create *TokenCreate
	}

	// TokenUpsert is the "OnConflict" setter.
	TokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *TokenUpsert) SetUpdatedAt(v time.Time) *TokenUpsert {
	u.Set(token.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TokenUpsert) UpdateUpdatedAt() *TokenUpsert {
	u.SetExcluded(token.FieldUpdatedAt)
	return u
}

// SetSymbol sets the "symbol" field.
func (u *TokenUpsert) SetSymbol(v string) *TokenUpsert {
	u.Set(token.FieldSymbol, v)
	return u
}

// UpdateSymbol sets the "symbol" field to the value that was provided on create.
func (u *TokenUpsert) UpdateSymbol() *TokenUpsert {
	u.SetExcluded(token.FieldSymbol)
	return u
}

// SetContractAddress sets the "contract_address" field.
func (u *TokenUpsert) SetContractAddress(v string) *TokenUpsert {
	u.Set(token.FieldContractAddress, v)
	return u
}

// UpdateContractAddress sets the "contract_address" field to the value that was provided on create.
func (u *TokenUpsert) UpdateContractAddress() *TokenUpsert {
	u.SetExcluded(token.FieldContractAddress)
	return u
}

// SetDecimals sets the "decimals" field.
func (u *TokenUpsert) SetDecimals(v int8) *TokenUpsert {
	u.Set(token.FieldDecimals, v)
	return u
}

// UpdateDecimals sets the "decimals" field to the value that was provided on create.
func (u *TokenUpsert) UpdateDecimals() *TokenUpsert {
	u.SetExcluded(token.FieldDecimals)
	return u
}

// AddDecimals adds v to the "decimals" field.
func (u *TokenUpsert) AddDecimals(v int8) *TokenUpsert {
	u.Add(token.FieldDecimals, v)
	return u
}

// SetIsEnabled sets the "is_enabled" field.
func (u *TokenUpsert) SetIsEnabled(v bool) *TokenUpsert {
	u.Set(token.FieldIsEnabled, v)
	return u
}

// UpdateIsEnabled sets the "is_enabled" field to the value that was provided on create.
func (u *TokenUpsert) UpdateIsEnabled() *TokenUpsert {
	u.SetExcluded(token.FieldIsEnabled)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TokenUpsertOne) UpdateNewValues() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(token.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TokenUpsertOne) Ignore() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertOne) DoNothing() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreate.OnConflict
// documentation for more info.
func (u *TokenUpsertOne) Update(set func(*TokenUpsert)) *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TokenUpsertOne) SetUpdatedAt(v time.Time) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateUpdatedAt() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSymbol sets the "symbol" field.
func (u *TokenUpsertOne) SetSymbol(v string) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetSymbol(v)
	})
}

// UpdateSymbol sets the "symbol" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateSymbol() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateSymbol()
	})
}

// SetContractAddress sets the "contract_address" field.
func (u *TokenUpsertOne) SetContractAddress(v string) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetContractAddress(v)
	})
}

// UpdateContractAddress sets the "contract_address" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateContractAddress() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateContractAddress()
	})
}

// SetDecimals sets the "decimals" field.
func (u *TokenUpsertOne) SetDecimals(v int8) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetDecimals(v)
	})
}

// AddDecimals adds v to the "decimals" field.
func (u *TokenUpsertOne) AddDecimals(v int8) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.AddDecimals(v)
	})
}

// UpdateDecimals sets the "decimals" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateDecimals() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateDecimals()
	})
}

// SetIsEnabled sets the "is_enabled" field.
func (u *TokenUpsertOne) SetIsEnabled(v bool) *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.SetIsEnabled(v)
	})
}

// UpdateIsEnabled sets the "is_enabled" field to the value that was provided on create.
func (u *TokenUpsertOne) UpdateIsEnabled() *TokenUpsertOne {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateIsEnabled()
	})
}

// Exec executes the query.
func (u *TokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TokenUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TokenUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TokenCreateBulk is the builder for creating many Token entities in bulk.
type TokenCreateBulk struct {
	config
	err      error
	builders []*TokenCreate
	conflict []sql.ConflictOption
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokenCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tcb *TokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *TokenUpsertBulk {
	tcb.conflict = opts
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TokenCreateBulk) OnConflictColumns(columns ...string) *TokenUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// TokenUpsertBulk is the builder for "upsert"-ing
// a bulk of Token nodes.
type TokenUpsertBulk struct {
	create *TokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TokenUpsertBulk) UpdateNewValues() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(token.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TokenUpsertBulk) Ignore() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertBulk) DoNothing() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreateBulk.OnConflict
// documentation for more info.
func (u *TokenUpsertBulk) Update(set func(*TokenUpsert)) *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TokenUpsertBulk) SetUpdatedAt(v time.Time) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateUpdatedAt() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSymbol sets the "symbol" field.
func (u *TokenUpsertBulk) SetSymbol(v string) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetSymbol(v)
	})
}

// UpdateSymbol sets the "symbol" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateSymbol() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateSymbol()
	})
}

// SetContractAddress sets the "contract_address" field.
func (u *TokenUpsertBulk) SetContractAddress(v string) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetContractAddress(v)
	})
}

// UpdateContractAddress sets the "contract_address" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateContractAddress() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateContractAddress()
	})
}

// SetDecimals sets the "decimals" field.
func (u *TokenUpsertBulk) SetDecimals(v int8) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetDecimals(v)
	})
}

// AddDecimals adds v to the "decimals" field.
func (u *TokenUpsertBulk) AddDecimals(v int8) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.AddDecimals(v)
	})
}

// UpdateDecimals sets the "decimals" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateDecimals() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateDecimals()
	})
}

// SetIsEnabled sets the "is_enabled" field.
func (u *TokenUpsertBulk) SetIsEnabled(v bool) *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.SetIsEnabled(v)
	})
}

// UpdateIsEnabled sets the "is_enabled" field to the value that was provided on create.
func (u *TokenUpsertBulk) UpdateIsEnabled() *TokenUpsertBulk {
	return u.Update(func(s *TokenUpsert) {
		s.UpdateIsEnabled()
	})
}

// Exec executes the query.
func (u *TokenUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
