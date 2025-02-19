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
	"github.com/paycrest/aggregator/ent/predicate"
	"github.com/paycrest/aggregator/ent/providerrating"
	"github.com/shopspring/decimal"
)

// ProviderRatingUpdate is the builder for updating ProviderRating entities.
type ProviderRatingUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderRatingMutation
}

// Where appends a list predicates to the ProviderRatingUpdate builder.
func (pru *ProviderRatingUpdate) Where(ps ...predicate.ProviderRating) *ProviderRatingUpdate {
	pru.mutation.Where(ps...)
	return pru
}

// SetUpdatedAt sets the "updated_at" field.
func (pru *ProviderRatingUpdate) SetUpdatedAt(t time.Time) *ProviderRatingUpdate {
	pru.mutation.SetUpdatedAt(t)
	return pru
}

// SetTrustScore sets the "trust_score" field.
func (pru *ProviderRatingUpdate) SetTrustScore(d decimal.Decimal) *ProviderRatingUpdate {
	pru.mutation.ResetTrustScore()
	pru.mutation.SetTrustScore(d)
	return pru
}

// SetNillableTrustScore sets the "trust_score" field if the given value is not nil.
func (pru *ProviderRatingUpdate) SetNillableTrustScore(d *decimal.Decimal) *ProviderRatingUpdate {
	if d != nil {
		pru.SetTrustScore(*d)
	}
	return pru
}

// AddTrustScore adds d to the "trust_score" field.
func (pru *ProviderRatingUpdate) AddTrustScore(d decimal.Decimal) *ProviderRatingUpdate {
	pru.mutation.AddTrustScore(d)
	return pru
}

// Mutation returns the ProviderRatingMutation object of the builder.
func (pru *ProviderRatingUpdate) Mutation() *ProviderRatingMutation {
	return pru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pru *ProviderRatingUpdate) Save(ctx context.Context) (int, error) {
	pru.defaults()
	return withHooks(ctx, pru.sqlSave, pru.mutation, pru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pru *ProviderRatingUpdate) SaveX(ctx context.Context) int {
	affected, err := pru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pru *ProviderRatingUpdate) Exec(ctx context.Context) error {
	_, err := pru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pru *ProviderRatingUpdate) ExecX(ctx context.Context) {
	if err := pru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pru *ProviderRatingUpdate) defaults() {
	if _, ok := pru.mutation.UpdatedAt(); !ok {
		v := providerrating.UpdateDefaultUpdatedAt()
		pru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pru *ProviderRatingUpdate) check() error {
	if pru.mutation.ProviderProfileCleared() && len(pru.mutation.ProviderProfileIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ProviderRating.provider_profile"`)
	}
	return nil
}

func (pru *ProviderRatingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(providerrating.Table, providerrating.Columns, sqlgraph.NewFieldSpec(providerrating.FieldID, field.TypeInt))
	if ps := pru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pru.mutation.UpdatedAt(); ok {
		_spec.SetField(providerrating.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pru.mutation.TrustScore(); ok {
		_spec.SetField(providerrating.FieldTrustScore, field.TypeFloat64, value)
	}
	if value, ok := pru.mutation.AddedTrustScore(); ok {
		_spec.AddField(providerrating.FieldTrustScore, field.TypeFloat64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providerrating.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pru.mutation.done = true
	return n, nil
}

// ProviderRatingUpdateOne is the builder for updating a single ProviderRating entity.
type ProviderRatingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderRatingMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pruo *ProviderRatingUpdateOne) SetUpdatedAt(t time.Time) *ProviderRatingUpdateOne {
	pruo.mutation.SetUpdatedAt(t)
	return pruo
}

// SetTrustScore sets the "trust_score" field.
func (pruo *ProviderRatingUpdateOne) SetTrustScore(d decimal.Decimal) *ProviderRatingUpdateOne {
	pruo.mutation.ResetTrustScore()
	pruo.mutation.SetTrustScore(d)
	return pruo
}

// SetNillableTrustScore sets the "trust_score" field if the given value is not nil.
func (pruo *ProviderRatingUpdateOne) SetNillableTrustScore(d *decimal.Decimal) *ProviderRatingUpdateOne {
	if d != nil {
		pruo.SetTrustScore(*d)
	}
	return pruo
}

// AddTrustScore adds d to the "trust_score" field.
func (pruo *ProviderRatingUpdateOne) AddTrustScore(d decimal.Decimal) *ProviderRatingUpdateOne {
	pruo.mutation.AddTrustScore(d)
	return pruo
}

// Mutation returns the ProviderRatingMutation object of the builder.
func (pruo *ProviderRatingUpdateOne) Mutation() *ProviderRatingMutation {
	return pruo.mutation
}

// Where appends a list predicates to the ProviderRatingUpdate builder.
func (pruo *ProviderRatingUpdateOne) Where(ps ...predicate.ProviderRating) *ProviderRatingUpdateOne {
	pruo.mutation.Where(ps...)
	return pruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pruo *ProviderRatingUpdateOne) Select(field string, fields ...string) *ProviderRatingUpdateOne {
	pruo.fields = append([]string{field}, fields...)
	return pruo
}

// Save executes the query and returns the updated ProviderRating entity.
func (pruo *ProviderRatingUpdateOne) Save(ctx context.Context) (*ProviderRating, error) {
	pruo.defaults()
	return withHooks(ctx, pruo.sqlSave, pruo.mutation, pruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pruo *ProviderRatingUpdateOne) SaveX(ctx context.Context) *ProviderRating {
	node, err := pruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pruo *ProviderRatingUpdateOne) Exec(ctx context.Context) error {
	_, err := pruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pruo *ProviderRatingUpdateOne) ExecX(ctx context.Context) {
	if err := pruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pruo *ProviderRatingUpdateOne) defaults() {
	if _, ok := pruo.mutation.UpdatedAt(); !ok {
		v := providerrating.UpdateDefaultUpdatedAt()
		pruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pruo *ProviderRatingUpdateOne) check() error {
	if pruo.mutation.ProviderProfileCleared() && len(pruo.mutation.ProviderProfileIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ProviderRating.provider_profile"`)
	}
	return nil
}

func (pruo *ProviderRatingUpdateOne) sqlSave(ctx context.Context) (_node *ProviderRating, err error) {
	if err := pruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(providerrating.Table, providerrating.Columns, sqlgraph.NewFieldSpec(providerrating.FieldID, field.TypeInt))
	id, ok := pruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProviderRating.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, providerrating.FieldID)
		for _, f := range fields {
			if !providerrating.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != providerrating.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pruo.mutation.UpdatedAt(); ok {
		_spec.SetField(providerrating.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pruo.mutation.TrustScore(); ok {
		_spec.SetField(providerrating.FieldTrustScore, field.TypeFloat64, value)
	}
	if value, ok := pruo.mutation.AddedTrustScore(); ok {
		_spec.AddField(providerrating.FieldTrustScore, field.TypeFloat64, value)
	}
	_node = &ProviderRating{config: pruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providerrating.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pruo.mutation.done = true
	return _node, nil
}
