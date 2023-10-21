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
	"github.com/paycrest/paycrest-protocol/ent/predicate"
	"github.com/paycrest/paycrest-protocol/ent/providerordertoken"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
	"github.com/shopspring/decimal"
)

// ProviderOrderTokenUpdate is the builder for updating ProviderOrderToken entities.
type ProviderOrderTokenUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderOrderTokenMutation
}

// Where appends a list predicates to the ProviderOrderTokenUpdate builder.
func (potu *ProviderOrderTokenUpdate) Where(ps ...predicate.ProviderOrderToken) *ProviderOrderTokenUpdate {
	potu.mutation.Where(ps...)
	return potu
}

// SetUpdatedAt sets the "updated_at" field.
func (potu *ProviderOrderTokenUpdate) SetUpdatedAt(t time.Time) *ProviderOrderTokenUpdate {
	potu.mutation.SetUpdatedAt(t)
	return potu
}

// SetSymbol sets the "symbol" field.
func (potu *ProviderOrderTokenUpdate) SetSymbol(s string) *ProviderOrderTokenUpdate {
	potu.mutation.SetSymbol(s)
	return potu
}

// SetFixedConversionRate sets the "fixed_conversion_rate" field.
func (potu *ProviderOrderTokenUpdate) SetFixedConversionRate(d decimal.Decimal) *ProviderOrderTokenUpdate {
	potu.mutation.ResetFixedConversionRate()
	potu.mutation.SetFixedConversionRate(d)
	return potu
}

// AddFixedConversionRate adds d to the "fixed_conversion_rate" field.
func (potu *ProviderOrderTokenUpdate) AddFixedConversionRate(d decimal.Decimal) *ProviderOrderTokenUpdate {
	potu.mutation.AddFixedConversionRate(d)
	return potu
}

// SetFloatingConversionRate sets the "floating_conversion_rate" field.
func (potu *ProviderOrderTokenUpdate) SetFloatingConversionRate(d decimal.Decimal) *ProviderOrderTokenUpdate {
	potu.mutation.ResetFloatingConversionRate()
	potu.mutation.SetFloatingConversionRate(d)
	return potu
}

// AddFloatingConversionRate adds d to the "floating_conversion_rate" field.
func (potu *ProviderOrderTokenUpdate) AddFloatingConversionRate(d decimal.Decimal) *ProviderOrderTokenUpdate {
	potu.mutation.AddFloatingConversionRate(d)
	return potu
}

// SetConversionRateType sets the "conversion_rate_type" field.
func (potu *ProviderOrderTokenUpdate) SetConversionRateType(prt providerordertoken.ConversionRateType) *ProviderOrderTokenUpdate {
	potu.mutation.SetConversionRateType(prt)
	return potu
}

// SetMaxOrderAmount sets the "max_order_amount" field.
func (potu *ProviderOrderTokenUpdate) SetMaxOrderAmount(d decimal.Decimal) *ProviderOrderTokenUpdate {
	potu.mutation.ResetMaxOrderAmount()
	potu.mutation.SetMaxOrderAmount(d)
	return potu
}

// AddMaxOrderAmount adds d to the "max_order_amount" field.
func (potu *ProviderOrderTokenUpdate) AddMaxOrderAmount(d decimal.Decimal) *ProviderOrderTokenUpdate {
	potu.mutation.AddMaxOrderAmount(d)
	return potu
}

// SetMinOrderAmount sets the "min_order_amount" field.
func (potu *ProviderOrderTokenUpdate) SetMinOrderAmount(d decimal.Decimal) *ProviderOrderTokenUpdate {
	potu.mutation.ResetMinOrderAmount()
	potu.mutation.SetMinOrderAmount(d)
	return potu
}

// AddMinOrderAmount adds d to the "min_order_amount" field.
func (potu *ProviderOrderTokenUpdate) AddMinOrderAmount(d decimal.Decimal) *ProviderOrderTokenUpdate {
	potu.mutation.AddMinOrderAmount(d)
	return potu
}

// SetAddresses sets the "addresses" field.
func (potu *ProviderOrderTokenUpdate) SetAddresses(s []struct {
	Address string "json:\"address\""
	Network string "json:\"network\""
}) *ProviderOrderTokenUpdate {
	potu.mutation.SetAddresses(s)
	return potu
}

// AppendAddresses appends s to the "addresses" field.
func (potu *ProviderOrderTokenUpdate) AppendAddresses(s []struct {
	Address string "json:\"address\""
	Network string "json:\"network\""
}) *ProviderOrderTokenUpdate {
	potu.mutation.AppendAddresses(s)
	return potu
}

// SetProviderID sets the "provider" edge to the ProviderProfile entity by ID.
func (potu *ProviderOrderTokenUpdate) SetProviderID(id string) *ProviderOrderTokenUpdate {
	potu.mutation.SetProviderID(id)
	return potu
}

// SetNillableProviderID sets the "provider" edge to the ProviderProfile entity by ID if the given value is not nil.
func (potu *ProviderOrderTokenUpdate) SetNillableProviderID(id *string) *ProviderOrderTokenUpdate {
	if id != nil {
		potu = potu.SetProviderID(*id)
	}
	return potu
}

// SetProvider sets the "provider" edge to the ProviderProfile entity.
func (potu *ProviderOrderTokenUpdate) SetProvider(p *ProviderProfile) *ProviderOrderTokenUpdate {
	return potu.SetProviderID(p.ID)
}

// Mutation returns the ProviderOrderTokenMutation object of the builder.
func (potu *ProviderOrderTokenUpdate) Mutation() *ProviderOrderTokenMutation {
	return potu.mutation
}

// ClearProvider clears the "provider" edge to the ProviderProfile entity.
func (potu *ProviderOrderTokenUpdate) ClearProvider() *ProviderOrderTokenUpdate {
	potu.mutation.ClearProvider()
	return potu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (potu *ProviderOrderTokenUpdate) Save(ctx context.Context) (int, error) {
	potu.defaults()
	return withHooks(ctx, potu.sqlSave, potu.mutation, potu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (potu *ProviderOrderTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := potu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (potu *ProviderOrderTokenUpdate) Exec(ctx context.Context) error {
	_, err := potu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (potu *ProviderOrderTokenUpdate) ExecX(ctx context.Context) {
	if err := potu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (potu *ProviderOrderTokenUpdate) defaults() {
	if _, ok := potu.mutation.UpdatedAt(); !ok {
		v := providerordertoken.UpdateDefaultUpdatedAt()
		potu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (potu *ProviderOrderTokenUpdate) check() error {
	if v, ok := potu.mutation.ConversionRateType(); ok {
		if err := providerordertoken.ConversionRateTypeValidator(v); err != nil {
			return &ValidationError{Name: "conversion_rate_type", err: fmt.Errorf(`ent: validator failed for field "ProviderOrderToken.conversion_rate_type": %w`, err)}
		}
	}
	return nil
}

func (potu *ProviderOrderTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := potu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(providerordertoken.Table, providerordertoken.Columns, sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt))
	if ps := potu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := potu.mutation.UpdatedAt(); ok {
		_spec.SetField(providerordertoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := potu.mutation.Symbol(); ok {
		_spec.SetField(providerordertoken.FieldSymbol, field.TypeString, value)
	}
	if value, ok := potu.mutation.FixedConversionRate(); ok {
		_spec.SetField(providerordertoken.FieldFixedConversionRate, field.TypeFloat64, value)
	}
	if value, ok := potu.mutation.AddedFixedConversionRate(); ok {
		_spec.AddField(providerordertoken.FieldFixedConversionRate, field.TypeFloat64, value)
	}
	if value, ok := potu.mutation.FloatingConversionRate(); ok {
		_spec.SetField(providerordertoken.FieldFloatingConversionRate, field.TypeFloat64, value)
	}
	if value, ok := potu.mutation.AddedFloatingConversionRate(); ok {
		_spec.AddField(providerordertoken.FieldFloatingConversionRate, field.TypeFloat64, value)
	}
	if value, ok := potu.mutation.ConversionRateType(); ok {
		_spec.SetField(providerordertoken.FieldConversionRateType, field.TypeEnum, value)
	}
	if value, ok := potu.mutation.MaxOrderAmount(); ok {
		_spec.SetField(providerordertoken.FieldMaxOrderAmount, field.TypeFloat64, value)
	}
	if value, ok := potu.mutation.AddedMaxOrderAmount(); ok {
		_spec.AddField(providerordertoken.FieldMaxOrderAmount, field.TypeFloat64, value)
	}
	if value, ok := potu.mutation.MinOrderAmount(); ok {
		_spec.SetField(providerordertoken.FieldMinOrderAmount, field.TypeFloat64, value)
	}
	if value, ok := potu.mutation.AddedMinOrderAmount(); ok {
		_spec.AddField(providerordertoken.FieldMinOrderAmount, field.TypeFloat64, value)
	}
	if value, ok := potu.mutation.Addresses(); ok {
		_spec.SetField(providerordertoken.FieldAddresses, field.TypeJSON, value)
	}
	if value, ok := potu.mutation.AppendedAddresses(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, providerordertoken.FieldAddresses, value)
		})
	}
	if potu.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerordertoken.ProviderTable,
			Columns: []string{providerordertoken.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := potu.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerordertoken.ProviderTable,
			Columns: []string{providerordertoken.ProviderColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, potu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providerordertoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	potu.mutation.done = true
	return n, nil
}

// ProviderOrderTokenUpdateOne is the builder for updating a single ProviderOrderToken entity.
type ProviderOrderTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderOrderTokenMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (potuo *ProviderOrderTokenUpdateOne) SetUpdatedAt(t time.Time) *ProviderOrderTokenUpdateOne {
	potuo.mutation.SetUpdatedAt(t)
	return potuo
}

// SetSymbol sets the "symbol" field.
func (potuo *ProviderOrderTokenUpdateOne) SetSymbol(s string) *ProviderOrderTokenUpdateOne {
	potuo.mutation.SetSymbol(s)
	return potuo
}

// SetFixedConversionRate sets the "fixed_conversion_rate" field.
func (potuo *ProviderOrderTokenUpdateOne) SetFixedConversionRate(d decimal.Decimal) *ProviderOrderTokenUpdateOne {
	potuo.mutation.ResetFixedConversionRate()
	potuo.mutation.SetFixedConversionRate(d)
	return potuo
}

// AddFixedConversionRate adds d to the "fixed_conversion_rate" field.
func (potuo *ProviderOrderTokenUpdateOne) AddFixedConversionRate(d decimal.Decimal) *ProviderOrderTokenUpdateOne {
	potuo.mutation.AddFixedConversionRate(d)
	return potuo
}

// SetFloatingConversionRate sets the "floating_conversion_rate" field.
func (potuo *ProviderOrderTokenUpdateOne) SetFloatingConversionRate(d decimal.Decimal) *ProviderOrderTokenUpdateOne {
	potuo.mutation.ResetFloatingConversionRate()
	potuo.mutation.SetFloatingConversionRate(d)
	return potuo
}

// AddFloatingConversionRate adds d to the "floating_conversion_rate" field.
func (potuo *ProviderOrderTokenUpdateOne) AddFloatingConversionRate(d decimal.Decimal) *ProviderOrderTokenUpdateOne {
	potuo.mutation.AddFloatingConversionRate(d)
	return potuo
}

// SetConversionRateType sets the "conversion_rate_type" field.
func (potuo *ProviderOrderTokenUpdateOne) SetConversionRateType(prt providerordertoken.ConversionRateType) *ProviderOrderTokenUpdateOne {
	potuo.mutation.SetConversionRateType(prt)
	return potuo
}

// SetMaxOrderAmount sets the "max_order_amount" field.
func (potuo *ProviderOrderTokenUpdateOne) SetMaxOrderAmount(d decimal.Decimal) *ProviderOrderTokenUpdateOne {
	potuo.mutation.ResetMaxOrderAmount()
	potuo.mutation.SetMaxOrderAmount(d)
	return potuo
}

// AddMaxOrderAmount adds d to the "max_order_amount" field.
func (potuo *ProviderOrderTokenUpdateOne) AddMaxOrderAmount(d decimal.Decimal) *ProviderOrderTokenUpdateOne {
	potuo.mutation.AddMaxOrderAmount(d)
	return potuo
}

// SetMinOrderAmount sets the "min_order_amount" field.
func (potuo *ProviderOrderTokenUpdateOne) SetMinOrderAmount(d decimal.Decimal) *ProviderOrderTokenUpdateOne {
	potuo.mutation.ResetMinOrderAmount()
	potuo.mutation.SetMinOrderAmount(d)
	return potuo
}

// AddMinOrderAmount adds d to the "min_order_amount" field.
func (potuo *ProviderOrderTokenUpdateOne) AddMinOrderAmount(d decimal.Decimal) *ProviderOrderTokenUpdateOne {
	potuo.mutation.AddMinOrderAmount(d)
	return potuo
}

// SetAddresses sets the "addresses" field.
func (potuo *ProviderOrderTokenUpdateOne) SetAddresses(s []struct {
	Address string "json:\"address\""
	Network string "json:\"network\""
}) *ProviderOrderTokenUpdateOne {
	potuo.mutation.SetAddresses(s)
	return potuo
}

// AppendAddresses appends s to the "addresses" field.
func (potuo *ProviderOrderTokenUpdateOne) AppendAddresses(s []struct {
	Address string "json:\"address\""
	Network string "json:\"network\""
}) *ProviderOrderTokenUpdateOne {
	potuo.mutation.AppendAddresses(s)
	return potuo
}

// SetProviderID sets the "provider" edge to the ProviderProfile entity by ID.
func (potuo *ProviderOrderTokenUpdateOne) SetProviderID(id string) *ProviderOrderTokenUpdateOne {
	potuo.mutation.SetProviderID(id)
	return potuo
}

// SetNillableProviderID sets the "provider" edge to the ProviderProfile entity by ID if the given value is not nil.
func (potuo *ProviderOrderTokenUpdateOne) SetNillableProviderID(id *string) *ProviderOrderTokenUpdateOne {
	if id != nil {
		potuo = potuo.SetProviderID(*id)
	}
	return potuo
}

// SetProvider sets the "provider" edge to the ProviderProfile entity.
func (potuo *ProviderOrderTokenUpdateOne) SetProvider(p *ProviderProfile) *ProviderOrderTokenUpdateOne {
	return potuo.SetProviderID(p.ID)
}

// Mutation returns the ProviderOrderTokenMutation object of the builder.
func (potuo *ProviderOrderTokenUpdateOne) Mutation() *ProviderOrderTokenMutation {
	return potuo.mutation
}

// ClearProvider clears the "provider" edge to the ProviderProfile entity.
func (potuo *ProviderOrderTokenUpdateOne) ClearProvider() *ProviderOrderTokenUpdateOne {
	potuo.mutation.ClearProvider()
	return potuo
}

// Where appends a list predicates to the ProviderOrderTokenUpdate builder.
func (potuo *ProviderOrderTokenUpdateOne) Where(ps ...predicate.ProviderOrderToken) *ProviderOrderTokenUpdateOne {
	potuo.mutation.Where(ps...)
	return potuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (potuo *ProviderOrderTokenUpdateOne) Select(field string, fields ...string) *ProviderOrderTokenUpdateOne {
	potuo.fields = append([]string{field}, fields...)
	return potuo
}

// Save executes the query and returns the updated ProviderOrderToken entity.
func (potuo *ProviderOrderTokenUpdateOne) Save(ctx context.Context) (*ProviderOrderToken, error) {
	potuo.defaults()
	return withHooks(ctx, potuo.sqlSave, potuo.mutation, potuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (potuo *ProviderOrderTokenUpdateOne) SaveX(ctx context.Context) *ProviderOrderToken {
	node, err := potuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (potuo *ProviderOrderTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := potuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (potuo *ProviderOrderTokenUpdateOne) ExecX(ctx context.Context) {
	if err := potuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (potuo *ProviderOrderTokenUpdateOne) defaults() {
	if _, ok := potuo.mutation.UpdatedAt(); !ok {
		v := providerordertoken.UpdateDefaultUpdatedAt()
		potuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (potuo *ProviderOrderTokenUpdateOne) check() error {
	if v, ok := potuo.mutation.ConversionRateType(); ok {
		if err := providerordertoken.ConversionRateTypeValidator(v); err != nil {
			return &ValidationError{Name: "conversion_rate_type", err: fmt.Errorf(`ent: validator failed for field "ProviderOrderToken.conversion_rate_type": %w`, err)}
		}
	}
	return nil
}

func (potuo *ProviderOrderTokenUpdateOne) sqlSave(ctx context.Context) (_node *ProviderOrderToken, err error) {
	if err := potuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(providerordertoken.Table, providerordertoken.Columns, sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt))
	id, ok := potuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProviderOrderToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := potuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, providerordertoken.FieldID)
		for _, f := range fields {
			if !providerordertoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != providerordertoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := potuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := potuo.mutation.UpdatedAt(); ok {
		_spec.SetField(providerordertoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := potuo.mutation.Symbol(); ok {
		_spec.SetField(providerordertoken.FieldSymbol, field.TypeString, value)
	}
	if value, ok := potuo.mutation.FixedConversionRate(); ok {
		_spec.SetField(providerordertoken.FieldFixedConversionRate, field.TypeFloat64, value)
	}
	if value, ok := potuo.mutation.AddedFixedConversionRate(); ok {
		_spec.AddField(providerordertoken.FieldFixedConversionRate, field.TypeFloat64, value)
	}
	if value, ok := potuo.mutation.FloatingConversionRate(); ok {
		_spec.SetField(providerordertoken.FieldFloatingConversionRate, field.TypeFloat64, value)
	}
	if value, ok := potuo.mutation.AddedFloatingConversionRate(); ok {
		_spec.AddField(providerordertoken.FieldFloatingConversionRate, field.TypeFloat64, value)
	}
	if value, ok := potuo.mutation.ConversionRateType(); ok {
		_spec.SetField(providerordertoken.FieldConversionRateType, field.TypeEnum, value)
	}
	if value, ok := potuo.mutation.MaxOrderAmount(); ok {
		_spec.SetField(providerordertoken.FieldMaxOrderAmount, field.TypeFloat64, value)
	}
	if value, ok := potuo.mutation.AddedMaxOrderAmount(); ok {
		_spec.AddField(providerordertoken.FieldMaxOrderAmount, field.TypeFloat64, value)
	}
	if value, ok := potuo.mutation.MinOrderAmount(); ok {
		_spec.SetField(providerordertoken.FieldMinOrderAmount, field.TypeFloat64, value)
	}
	if value, ok := potuo.mutation.AddedMinOrderAmount(); ok {
		_spec.AddField(providerordertoken.FieldMinOrderAmount, field.TypeFloat64, value)
	}
	if value, ok := potuo.mutation.Addresses(); ok {
		_spec.SetField(providerordertoken.FieldAddresses, field.TypeJSON, value)
	}
	if value, ok := potuo.mutation.AppendedAddresses(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, providerordertoken.FieldAddresses, value)
		})
	}
	if potuo.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerordertoken.ProviderTable,
			Columns: []string{providerordertoken.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := potuo.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerordertoken.ProviderTable,
			Columns: []string{providerordertoken.ProviderColumn},
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
	_node = &ProviderOrderToken{config: potuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, potuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providerordertoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	potuo.mutation.done = true
	return _node, nil
}
