// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/paycrest-protocol/ent/providerordertoken"
	"github.com/paycrest/paycrest-protocol/ent/providerordertokenaddress"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
	"github.com/shopspring/decimal"
)

// ProviderOrderTokenCreate is the builder for creating a ProviderOrderToken entity.
type ProviderOrderTokenCreate struct {
	config
	mutation *ProviderOrderTokenMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (potc *ProviderOrderTokenCreate) SetCreatedAt(t time.Time) *ProviderOrderTokenCreate {
	potc.mutation.SetCreatedAt(t)
	return potc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (potc *ProviderOrderTokenCreate) SetNillableCreatedAt(t *time.Time) *ProviderOrderTokenCreate {
	if t != nil {
		potc.SetCreatedAt(*t)
	}
	return potc
}

// SetUpdatedAt sets the "updated_at" field.
func (potc *ProviderOrderTokenCreate) SetUpdatedAt(t time.Time) *ProviderOrderTokenCreate {
	potc.mutation.SetUpdatedAt(t)
	return potc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (potc *ProviderOrderTokenCreate) SetNillableUpdatedAt(t *time.Time) *ProviderOrderTokenCreate {
	if t != nil {
		potc.SetUpdatedAt(*t)
	}
	return potc
}

// SetName sets the "name" field.
func (potc *ProviderOrderTokenCreate) SetName(pr providerordertoken.Name) *ProviderOrderTokenCreate {
	potc.mutation.SetName(pr)
	return potc
}

// SetFixedConversionRate sets the "fixed_conversion_rate" field.
func (potc *ProviderOrderTokenCreate) SetFixedConversionRate(d decimal.Decimal) *ProviderOrderTokenCreate {
	potc.mutation.SetFixedConversionRate(d)
	return potc
}

// SetFloatingConversionRate sets the "floating_conversion_rate" field.
func (potc *ProviderOrderTokenCreate) SetFloatingConversionRate(d decimal.Decimal) *ProviderOrderTokenCreate {
	potc.mutation.SetFloatingConversionRate(d)
	return potc
}

// SetConversionRateType sets the "conversion_rate_type" field.
func (potc *ProviderOrderTokenCreate) SetConversionRateType(prt providerordertoken.ConversionRateType) *ProviderOrderTokenCreate {
	potc.mutation.SetConversionRateType(prt)
	return potc
}

// SetMaxOrderAmount sets the "max_order_amount" field.
func (potc *ProviderOrderTokenCreate) SetMaxOrderAmount(d decimal.Decimal) *ProviderOrderTokenCreate {
	potc.mutation.SetMaxOrderAmount(d)
	return potc
}

// SetMinOrderAmount sets the "min_order_amount" field.
func (potc *ProviderOrderTokenCreate) SetMinOrderAmount(d decimal.Decimal) *ProviderOrderTokenCreate {
	potc.mutation.SetMinOrderAmount(d)
	return potc
}

// SetProviderID sets the "provider" edge to the ProviderProfile entity by ID.
func (potc *ProviderOrderTokenCreate) SetProviderID(id string) *ProviderOrderTokenCreate {
	potc.mutation.SetProviderID(id)
	return potc
}

// SetNillableProviderID sets the "provider" edge to the ProviderProfile entity by ID if the given value is not nil.
func (potc *ProviderOrderTokenCreate) SetNillableProviderID(id *string) *ProviderOrderTokenCreate {
	if id != nil {
		potc = potc.SetProviderID(*id)
	}
	return potc
}

// SetProvider sets the "provider" edge to the ProviderProfile entity.
func (potc *ProviderOrderTokenCreate) SetProvider(p *ProviderProfile) *ProviderOrderTokenCreate {
	return potc.SetProviderID(p.ID)
}

// AddAddressIDs adds the "addresses" edge to the ProviderOrderTokenAddress entity by IDs.
func (potc *ProviderOrderTokenCreate) AddAddressIDs(ids ...int) *ProviderOrderTokenCreate {
	potc.mutation.AddAddressIDs(ids...)
	return potc
}

// AddAddresses adds the "addresses" edges to the ProviderOrderTokenAddress entity.
func (potc *ProviderOrderTokenCreate) AddAddresses(p ...*ProviderOrderTokenAddress) *ProviderOrderTokenCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return potc.AddAddressIDs(ids...)
}

// Mutation returns the ProviderOrderTokenMutation object of the builder.
func (potc *ProviderOrderTokenCreate) Mutation() *ProviderOrderTokenMutation {
	return potc.mutation
}

// Save creates the ProviderOrderToken in the database.
func (potc *ProviderOrderTokenCreate) Save(ctx context.Context) (*ProviderOrderToken, error) {
	potc.defaults()
	return withHooks(ctx, potc.sqlSave, potc.mutation, potc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (potc *ProviderOrderTokenCreate) SaveX(ctx context.Context) *ProviderOrderToken {
	v, err := potc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (potc *ProviderOrderTokenCreate) Exec(ctx context.Context) error {
	_, err := potc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (potc *ProviderOrderTokenCreate) ExecX(ctx context.Context) {
	if err := potc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (potc *ProviderOrderTokenCreate) defaults() {
	if _, ok := potc.mutation.CreatedAt(); !ok {
		v := providerordertoken.DefaultCreatedAt()
		potc.mutation.SetCreatedAt(v)
	}
	if _, ok := potc.mutation.UpdatedAt(); !ok {
		v := providerordertoken.DefaultUpdatedAt()
		potc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (potc *ProviderOrderTokenCreate) check() error {
	if _, ok := potc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProviderOrderToken.created_at"`)}
	}
	if _, ok := potc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProviderOrderToken.updated_at"`)}
	}
	if _, ok := potc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ProviderOrderToken.name"`)}
	}
	if v, ok := potc.mutation.Name(); ok {
		if err := providerordertoken.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ProviderOrderToken.name": %w`, err)}
		}
	}
	if _, ok := potc.mutation.FixedConversionRate(); !ok {
		return &ValidationError{Name: "fixed_conversion_rate", err: errors.New(`ent: missing required field "ProviderOrderToken.fixed_conversion_rate"`)}
	}
	if _, ok := potc.mutation.FloatingConversionRate(); !ok {
		return &ValidationError{Name: "floating_conversion_rate", err: errors.New(`ent: missing required field "ProviderOrderToken.floating_conversion_rate"`)}
	}
	if _, ok := potc.mutation.ConversionRateType(); !ok {
		return &ValidationError{Name: "conversion_rate_type", err: errors.New(`ent: missing required field "ProviderOrderToken.conversion_rate_type"`)}
	}
	if v, ok := potc.mutation.ConversionRateType(); ok {
		if err := providerordertoken.ConversionRateTypeValidator(v); err != nil {
			return &ValidationError{Name: "conversion_rate_type", err: fmt.Errorf(`ent: validator failed for field "ProviderOrderToken.conversion_rate_type": %w`, err)}
		}
	}
	if _, ok := potc.mutation.MaxOrderAmount(); !ok {
		return &ValidationError{Name: "max_order_amount", err: errors.New(`ent: missing required field "ProviderOrderToken.max_order_amount"`)}
	}
	if _, ok := potc.mutation.MinOrderAmount(); !ok {
		return &ValidationError{Name: "min_order_amount", err: errors.New(`ent: missing required field "ProviderOrderToken.min_order_amount"`)}
	}
	return nil
}

func (potc *ProviderOrderTokenCreate) sqlSave(ctx context.Context) (*ProviderOrderToken, error) {
	if err := potc.check(); err != nil {
		return nil, err
	}
	_node, _spec := potc.createSpec()
	if err := sqlgraph.CreateNode(ctx, potc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	potc.mutation.id = &_node.ID
	potc.mutation.done = true
	return _node, nil
}

func (potc *ProviderOrderTokenCreate) createSpec() (*ProviderOrderToken, *sqlgraph.CreateSpec) {
	var (
		_node = &ProviderOrderToken{config: potc.config}
		_spec = sqlgraph.NewCreateSpec(providerordertoken.Table, sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt))
	)
	if value, ok := potc.mutation.CreatedAt(); ok {
		_spec.SetField(providerordertoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := potc.mutation.UpdatedAt(); ok {
		_spec.SetField(providerordertoken.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := potc.mutation.Name(); ok {
		_spec.SetField(providerordertoken.FieldName, field.TypeEnum, value)
		_node.Name = value
	}
	if value, ok := potc.mutation.FixedConversionRate(); ok {
		_spec.SetField(providerordertoken.FieldFixedConversionRate, field.TypeFloat64, value)
		_node.FixedConversionRate = value
	}
	if value, ok := potc.mutation.FloatingConversionRate(); ok {
		_spec.SetField(providerordertoken.FieldFloatingConversionRate, field.TypeFloat64, value)
		_node.FloatingConversionRate = value
	}
	if value, ok := potc.mutation.ConversionRateType(); ok {
		_spec.SetField(providerordertoken.FieldConversionRateType, field.TypeEnum, value)
		_node.ConversionRateType = value
	}
	if value, ok := potc.mutation.MaxOrderAmount(); ok {
		_spec.SetField(providerordertoken.FieldMaxOrderAmount, field.TypeString, value)
		_node.MaxOrderAmount = value
	}
	if value, ok := potc.mutation.MinOrderAmount(); ok {
		_spec.SetField(providerordertoken.FieldMinOrderAmount, field.TypeString, value)
		_node.MinOrderAmount = value
	}
	if nodes := potc.mutation.ProviderIDs(); len(nodes) > 0 {
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
		_node.provider_profile_order_tokens = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := potc.mutation.AddressesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerordertoken.AddressesTable,
			Columns: []string{providerordertoken.AddressesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerordertokenaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProviderOrderTokenCreateBulk is the builder for creating many ProviderOrderToken entities in bulk.
type ProviderOrderTokenCreateBulk struct {
	config
	builders []*ProviderOrderTokenCreate
}

// Save creates the ProviderOrderToken entities in the database.
func (potcb *ProviderOrderTokenCreateBulk) Save(ctx context.Context) ([]*ProviderOrderToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(potcb.builders))
	nodes := make([]*ProviderOrderToken, len(potcb.builders))
	mutators := make([]Mutator, len(potcb.builders))
	for i := range potcb.builders {
		func(i int, root context.Context) {
			builder := potcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProviderOrderTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, potcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, potcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, potcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (potcb *ProviderOrderTokenCreateBulk) SaveX(ctx context.Context) []*ProviderOrderToken {
	v, err := potcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (potcb *ProviderOrderTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := potcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (potcb *ProviderOrderTokenCreateBulk) ExecX(ctx context.Context) {
	if err := potcb.Exec(ctx); err != nil {
		panic(err)
	}
}
