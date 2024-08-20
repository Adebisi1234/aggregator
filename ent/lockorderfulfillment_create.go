// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/lockorderfulfillment"
	"github.com/paycrest/protocol/ent/lockpaymentorder"
)

// LockOrderFulfillmentCreate is the builder for creating a LockOrderFulfillment entity.
type LockOrderFulfillmentCreate struct {
	config
	mutation *LockOrderFulfillmentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (lofc *LockOrderFulfillmentCreate) SetCreatedAt(t time.Time) *LockOrderFulfillmentCreate {
	lofc.mutation.SetCreatedAt(t)
	return lofc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableCreatedAt(t *time.Time) *LockOrderFulfillmentCreate {
	if t != nil {
		lofc.SetCreatedAt(*t)
	}
	return lofc
}

// SetUpdatedAt sets the "updated_at" field.
func (lofc *LockOrderFulfillmentCreate) SetUpdatedAt(t time.Time) *LockOrderFulfillmentCreate {
	lofc.mutation.SetUpdatedAt(t)
	return lofc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableUpdatedAt(t *time.Time) *LockOrderFulfillmentCreate {
	if t != nil {
		lofc.SetUpdatedAt(*t)
	}
	return lofc
}

// SetTxID sets the "tx_id" field.
func (lofc *LockOrderFulfillmentCreate) SetTxID(s string) *LockOrderFulfillmentCreate {
	lofc.mutation.SetTxID(s)
	return lofc
}

// SetPsp sets the "psp" field.
func (lofc *LockOrderFulfillmentCreate) SetPsp(s string) *LockOrderFulfillmentCreate {
	lofc.mutation.SetPsp(s)
	return lofc
}

// SetNillablePsp sets the "psp" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillablePsp(s *string) *LockOrderFulfillmentCreate {
	if s != nil {
		lofc.SetPsp(*s)
	}
	return lofc
}

// SetValidationStatus sets the "validation_status" field.
func (lofc *LockOrderFulfillmentCreate) SetValidationStatus(ls lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentCreate {
	lofc.mutation.SetValidationStatus(ls)
	return lofc
}

// SetNillableValidationStatus sets the "validation_status" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableValidationStatus(ls *lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentCreate {
	if ls != nil {
		lofc.SetValidationStatus(*ls)
	}
	return lofc
}

// SetValidationError sets the "validation_error" field.
func (lofc *LockOrderFulfillmentCreate) SetValidationError(s string) *LockOrderFulfillmentCreate {
	lofc.mutation.SetValidationError(s)
	return lofc
}

// SetNillableValidationError sets the "validation_error" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableValidationError(s *string) *LockOrderFulfillmentCreate {
	if s != nil {
		lofc.SetValidationError(*s)
	}
	return lofc
}

// SetID sets the "id" field.
func (lofc *LockOrderFulfillmentCreate) SetID(u uuid.UUID) *LockOrderFulfillmentCreate {
	lofc.mutation.SetID(u)
	return lofc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lofc *LockOrderFulfillmentCreate) SetNillableID(u *uuid.UUID) *LockOrderFulfillmentCreate {
	if u != nil {
		lofc.SetID(*u)
	}
	return lofc
}

// SetOrderID sets the "order" edge to the LockPaymentOrder entity by ID.
func (lofc *LockOrderFulfillmentCreate) SetOrderID(id uuid.UUID) *LockOrderFulfillmentCreate {
	lofc.mutation.SetOrderID(id)
	return lofc
}

// SetOrder sets the "order" edge to the LockPaymentOrder entity.
func (lofc *LockOrderFulfillmentCreate) SetOrder(l *LockPaymentOrder) *LockOrderFulfillmentCreate {
	return lofc.SetOrderID(l.ID)
}

// Mutation returns the LockOrderFulfillmentMutation object of the builder.
func (lofc *LockOrderFulfillmentCreate) Mutation() *LockOrderFulfillmentMutation {
	return lofc.mutation
}

// Save creates the LockOrderFulfillment in the database.
func (lofc *LockOrderFulfillmentCreate) Save(ctx context.Context) (*LockOrderFulfillment, error) {
	lofc.defaults()
	return withHooks(ctx, lofc.sqlSave, lofc.mutation, lofc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lofc *LockOrderFulfillmentCreate) SaveX(ctx context.Context) *LockOrderFulfillment {
	v, err := lofc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lofc *LockOrderFulfillmentCreate) Exec(ctx context.Context) error {
	_, err := lofc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lofc *LockOrderFulfillmentCreate) ExecX(ctx context.Context) {
	if err := lofc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lofc *LockOrderFulfillmentCreate) defaults() {
	if _, ok := lofc.mutation.CreatedAt(); !ok {
		v := lockorderfulfillment.DefaultCreatedAt()
		lofc.mutation.SetCreatedAt(v)
	}
	if _, ok := lofc.mutation.UpdatedAt(); !ok {
		v := lockorderfulfillment.DefaultUpdatedAt()
		lofc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lofc.mutation.ValidationStatus(); !ok {
		v := lockorderfulfillment.DefaultValidationStatus
		lofc.mutation.SetValidationStatus(v)
	}
	if _, ok := lofc.mutation.ID(); !ok {
		v := lockorderfulfillment.DefaultID()
		lofc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lofc *LockOrderFulfillmentCreate) check() error {
	if _, ok := lofc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LockOrderFulfillment.created_at"`)}
	}
	if _, ok := lofc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LockOrderFulfillment.updated_at"`)}
	}
	if _, ok := lofc.mutation.TxID(); !ok {
		return &ValidationError{Name: "tx_id", err: errors.New(`ent: missing required field "LockOrderFulfillment.tx_id"`)}
	}
	if _, ok := lofc.mutation.ValidationStatus(); !ok {
		return &ValidationError{Name: "validation_status", err: errors.New(`ent: missing required field "LockOrderFulfillment.validation_status"`)}
	}
	if v, ok := lofc.mutation.ValidationStatus(); ok {
		if err := lockorderfulfillment.ValidationStatusValidator(v); err != nil {
			return &ValidationError{Name: "validation_status", err: fmt.Errorf(`ent: validator failed for field "LockOrderFulfillment.validation_status": %w`, err)}
		}
	}
	if _, ok := lofc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "LockOrderFulfillment.order"`)}
	}
	return nil
}

func (lofc *LockOrderFulfillmentCreate) sqlSave(ctx context.Context) (*LockOrderFulfillment, error) {
	if err := lofc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lofc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lofc.driver, _spec); err != nil {
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
	lofc.mutation.id = &_node.ID
	lofc.mutation.done = true
	return _node, nil
}

func (lofc *LockOrderFulfillmentCreate) createSpec() (*LockOrderFulfillment, *sqlgraph.CreateSpec) {
	var (
		_node = &LockOrderFulfillment{config: lofc.config}
		_spec = sqlgraph.NewCreateSpec(lockorderfulfillment.Table, sqlgraph.NewFieldSpec(lockorderfulfillment.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = lofc.conflict
	if id, ok := lofc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lofc.mutation.CreatedAt(); ok {
		_spec.SetField(lockorderfulfillment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lofc.mutation.UpdatedAt(); ok {
		_spec.SetField(lockorderfulfillment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lofc.mutation.TxID(); ok {
		_spec.SetField(lockorderfulfillment.FieldTxID, field.TypeString, value)
		_node.TxID = value
	}
	if value, ok := lofc.mutation.Psp(); ok {
		_spec.SetField(lockorderfulfillment.FieldPsp, field.TypeString, value)
		_node.Psp = value
	}
	if value, ok := lofc.mutation.ValidationStatus(); ok {
		_spec.SetField(lockorderfulfillment.FieldValidationStatus, field.TypeEnum, value)
		_node.ValidationStatus = value
	}
	if value, ok := lofc.mutation.ValidationError(); ok {
		_spec.SetField(lockorderfulfillment.FieldValidationError, field.TypeString, value)
		_node.ValidationError = value
	}
	if nodes := lofc.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lockorderfulfillment.OrderTable,
			Columns: []string{lockorderfulfillment.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.lock_payment_order_fulfillments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LockOrderFulfillment.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LockOrderFulfillmentUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (lofc *LockOrderFulfillmentCreate) OnConflict(opts ...sql.ConflictOption) *LockOrderFulfillmentUpsertOne {
	lofc.conflict = opts
	return &LockOrderFulfillmentUpsertOne{
		create: lofc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LockOrderFulfillment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lofc *LockOrderFulfillmentCreate) OnConflictColumns(columns ...string) *LockOrderFulfillmentUpsertOne {
	lofc.conflict = append(lofc.conflict, sql.ConflictColumns(columns...))
	return &LockOrderFulfillmentUpsertOne{
		create: lofc,
	}
}

type (
	// LockOrderFulfillmentUpsertOne is the builder for "upsert"-ing
	//  one LockOrderFulfillment node.
	LockOrderFulfillmentUpsertOne struct {
		create *LockOrderFulfillmentCreate
	}

	// LockOrderFulfillmentUpsert is the "OnConflict" setter.
	LockOrderFulfillmentUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *LockOrderFulfillmentUpsert) SetUpdatedAt(v time.Time) *LockOrderFulfillmentUpsert {
	u.Set(lockorderfulfillment.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsert) UpdateUpdatedAt() *LockOrderFulfillmentUpsert {
	u.SetExcluded(lockorderfulfillment.FieldUpdatedAt)
	return u
}

// SetTxID sets the "tx_id" field.
func (u *LockOrderFulfillmentUpsert) SetTxID(v string) *LockOrderFulfillmentUpsert {
	u.Set(lockorderfulfillment.FieldTxID, v)
	return u
}

// UpdateTxID sets the "tx_id" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsert) UpdateTxID() *LockOrderFulfillmentUpsert {
	u.SetExcluded(lockorderfulfillment.FieldTxID)
	return u
}

// SetPsp sets the "psp" field.
func (u *LockOrderFulfillmentUpsert) SetPsp(v string) *LockOrderFulfillmentUpsert {
	u.Set(lockorderfulfillment.FieldPsp, v)
	return u
}

// UpdatePsp sets the "psp" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsert) UpdatePsp() *LockOrderFulfillmentUpsert {
	u.SetExcluded(lockorderfulfillment.FieldPsp)
	return u
}

// ClearPsp clears the value of the "psp" field.
func (u *LockOrderFulfillmentUpsert) ClearPsp() *LockOrderFulfillmentUpsert {
	u.SetNull(lockorderfulfillment.FieldPsp)
	return u
}

// SetValidationStatus sets the "validation_status" field.
func (u *LockOrderFulfillmentUpsert) SetValidationStatus(v lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentUpsert {
	u.Set(lockorderfulfillment.FieldValidationStatus, v)
	return u
}

// UpdateValidationStatus sets the "validation_status" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsert) UpdateValidationStatus() *LockOrderFulfillmentUpsert {
	u.SetExcluded(lockorderfulfillment.FieldValidationStatus)
	return u
}

// SetValidationError sets the "validation_error" field.
func (u *LockOrderFulfillmentUpsert) SetValidationError(v string) *LockOrderFulfillmentUpsert {
	u.Set(lockorderfulfillment.FieldValidationError, v)
	return u
}

// UpdateValidationError sets the "validation_error" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsert) UpdateValidationError() *LockOrderFulfillmentUpsert {
	u.SetExcluded(lockorderfulfillment.FieldValidationError)
	return u
}

// ClearValidationError clears the value of the "validation_error" field.
func (u *LockOrderFulfillmentUpsert) ClearValidationError() *LockOrderFulfillmentUpsert {
	u.SetNull(lockorderfulfillment.FieldValidationError)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.LockOrderFulfillment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lockorderfulfillment.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LockOrderFulfillmentUpsertOne) UpdateNewValues() *LockOrderFulfillmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(lockorderfulfillment.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(lockorderfulfillment.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LockOrderFulfillment.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LockOrderFulfillmentUpsertOne) Ignore() *LockOrderFulfillmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LockOrderFulfillmentUpsertOne) DoNothing() *LockOrderFulfillmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LockOrderFulfillmentCreate.OnConflict
// documentation for more info.
func (u *LockOrderFulfillmentUpsertOne) Update(set func(*LockOrderFulfillmentUpsert)) *LockOrderFulfillmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LockOrderFulfillmentUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LockOrderFulfillmentUpsertOne) SetUpdatedAt(v time.Time) *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertOne) UpdateUpdatedAt() *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTxID sets the "tx_id" field.
func (u *LockOrderFulfillmentUpsertOne) SetTxID(v string) *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetTxID(v)
	})
}

// UpdateTxID sets the "tx_id" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertOne) UpdateTxID() *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdateTxID()
	})
}

// SetPsp sets the "psp" field.
func (u *LockOrderFulfillmentUpsertOne) SetPsp(v string) *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetPsp(v)
	})
}

// UpdatePsp sets the "psp" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertOne) UpdatePsp() *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdatePsp()
	})
}

// ClearPsp clears the value of the "psp" field.
func (u *LockOrderFulfillmentUpsertOne) ClearPsp() *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.ClearPsp()
	})
}

// SetValidationStatus sets the "validation_status" field.
func (u *LockOrderFulfillmentUpsertOne) SetValidationStatus(v lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetValidationStatus(v)
	})
}

// UpdateValidationStatus sets the "validation_status" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertOne) UpdateValidationStatus() *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdateValidationStatus()
	})
}

// SetValidationError sets the "validation_error" field.
func (u *LockOrderFulfillmentUpsertOne) SetValidationError(v string) *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetValidationError(v)
	})
}

// UpdateValidationError sets the "validation_error" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertOne) UpdateValidationError() *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdateValidationError()
	})
}

// ClearValidationError clears the value of the "validation_error" field.
func (u *LockOrderFulfillmentUpsertOne) ClearValidationError() *LockOrderFulfillmentUpsertOne {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.ClearValidationError()
	})
}

// Exec executes the query.
func (u *LockOrderFulfillmentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LockOrderFulfillmentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LockOrderFulfillmentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LockOrderFulfillmentUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: LockOrderFulfillmentUpsertOne.ID is not supported by MySQL driver. Use LockOrderFulfillmentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LockOrderFulfillmentUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LockOrderFulfillmentCreateBulk is the builder for creating many LockOrderFulfillment entities in bulk.
type LockOrderFulfillmentCreateBulk struct {
	config
	builders []*LockOrderFulfillmentCreate
	conflict []sql.ConflictOption
}

// Save creates the LockOrderFulfillment entities in the database.
func (lofcb *LockOrderFulfillmentCreateBulk) Save(ctx context.Context) ([]*LockOrderFulfillment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lofcb.builders))
	nodes := make([]*LockOrderFulfillment, len(lofcb.builders))
	mutators := make([]Mutator, len(lofcb.builders))
	for i := range lofcb.builders {
		func(i int, root context.Context) {
			builder := lofcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LockOrderFulfillmentMutation)
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
					_, err = mutators[i+1].Mutate(root, lofcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lofcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lofcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lofcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lofcb *LockOrderFulfillmentCreateBulk) SaveX(ctx context.Context) []*LockOrderFulfillment {
	v, err := lofcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lofcb *LockOrderFulfillmentCreateBulk) Exec(ctx context.Context) error {
	_, err := lofcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lofcb *LockOrderFulfillmentCreateBulk) ExecX(ctx context.Context) {
	if err := lofcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LockOrderFulfillment.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LockOrderFulfillmentUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (lofcb *LockOrderFulfillmentCreateBulk) OnConflict(opts ...sql.ConflictOption) *LockOrderFulfillmentUpsertBulk {
	lofcb.conflict = opts
	return &LockOrderFulfillmentUpsertBulk{
		create: lofcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LockOrderFulfillment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lofcb *LockOrderFulfillmentCreateBulk) OnConflictColumns(columns ...string) *LockOrderFulfillmentUpsertBulk {
	lofcb.conflict = append(lofcb.conflict, sql.ConflictColumns(columns...))
	return &LockOrderFulfillmentUpsertBulk{
		create: lofcb,
	}
}

// LockOrderFulfillmentUpsertBulk is the builder for "upsert"-ing
// a bulk of LockOrderFulfillment nodes.
type LockOrderFulfillmentUpsertBulk struct {
	create *LockOrderFulfillmentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LockOrderFulfillment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(lockorderfulfillment.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LockOrderFulfillmentUpsertBulk) UpdateNewValues() *LockOrderFulfillmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(lockorderfulfillment.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(lockorderfulfillment.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LockOrderFulfillment.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LockOrderFulfillmentUpsertBulk) Ignore() *LockOrderFulfillmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LockOrderFulfillmentUpsertBulk) DoNothing() *LockOrderFulfillmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LockOrderFulfillmentCreateBulk.OnConflict
// documentation for more info.
func (u *LockOrderFulfillmentUpsertBulk) Update(set func(*LockOrderFulfillmentUpsert)) *LockOrderFulfillmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LockOrderFulfillmentUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LockOrderFulfillmentUpsertBulk) SetUpdatedAt(v time.Time) *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertBulk) UpdateUpdatedAt() *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTxID sets the "tx_id" field.
func (u *LockOrderFulfillmentUpsertBulk) SetTxID(v string) *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetTxID(v)
	})
}

// UpdateTxID sets the "tx_id" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertBulk) UpdateTxID() *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdateTxID()
	})
}

// SetPsp sets the "psp" field.
func (u *LockOrderFulfillmentUpsertBulk) SetPsp(v string) *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetPsp(v)
	})
}

// UpdatePsp sets the "psp" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertBulk) UpdatePsp() *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdatePsp()
	})
}

// ClearPsp clears the value of the "psp" field.
func (u *LockOrderFulfillmentUpsertBulk) ClearPsp() *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.ClearPsp()
	})
}

// SetValidationStatus sets the "validation_status" field.
func (u *LockOrderFulfillmentUpsertBulk) SetValidationStatus(v lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetValidationStatus(v)
	})
}

// UpdateValidationStatus sets the "validation_status" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertBulk) UpdateValidationStatus() *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdateValidationStatus()
	})
}

// SetValidationError sets the "validation_error" field.
func (u *LockOrderFulfillmentUpsertBulk) SetValidationError(v string) *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.SetValidationError(v)
	})
}

// UpdateValidationError sets the "validation_error" field to the value that was provided on create.
func (u *LockOrderFulfillmentUpsertBulk) UpdateValidationError() *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.UpdateValidationError()
	})
}

// ClearValidationError clears the value of the "validation_error" field.
func (u *LockOrderFulfillmentUpsertBulk) ClearValidationError() *LockOrderFulfillmentUpsertBulk {
	return u.Update(func(s *LockOrderFulfillmentUpsert) {
		s.ClearValidationError()
	})
}

// Exec executes the query.
func (u *LockOrderFulfillmentUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LockOrderFulfillmentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LockOrderFulfillmentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LockOrderFulfillmentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
