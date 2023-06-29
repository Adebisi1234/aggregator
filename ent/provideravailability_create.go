// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/paycrest-protocol/ent/provideravailability"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
)

// ProviderAvailabilityCreate is the builder for creating a ProviderAvailability entity.
type ProviderAvailabilityCreate struct {
	config
	mutation *ProviderAvailabilityMutation
	hooks    []Hook
}

// SetCadence sets the "cadence" field.
func (pac *ProviderAvailabilityCreate) SetCadence(pr provideravailability.Cadence) *ProviderAvailabilityCreate {
	pac.mutation.SetCadence(pr)
	return pac
}

// SetStartTime sets the "start_time" field.
func (pac *ProviderAvailabilityCreate) SetStartTime(t time.Time) *ProviderAvailabilityCreate {
	pac.mutation.SetStartTime(t)
	return pac
}

// SetEndTime sets the "end_time" field.
func (pac *ProviderAvailabilityCreate) SetEndTime(t time.Time) *ProviderAvailabilityCreate {
	pac.mutation.SetEndTime(t)
	return pac
}

// SetProviderID sets the "provider" edge to the ProviderProfile entity by ID.
func (pac *ProviderAvailabilityCreate) SetProviderID(id string) *ProviderAvailabilityCreate {
	pac.mutation.SetProviderID(id)
	return pac
}

// SetNillableProviderID sets the "provider" edge to the ProviderProfile entity by ID if the given value is not nil.
func (pac *ProviderAvailabilityCreate) SetNillableProviderID(id *string) *ProviderAvailabilityCreate {
	if id != nil {
		pac = pac.SetProviderID(*id)
	}
	return pac
}

// SetProvider sets the "provider" edge to the ProviderProfile entity.
func (pac *ProviderAvailabilityCreate) SetProvider(p *ProviderProfile) *ProviderAvailabilityCreate {
	return pac.SetProviderID(p.ID)
}

// Mutation returns the ProviderAvailabilityMutation object of the builder.
func (pac *ProviderAvailabilityCreate) Mutation() *ProviderAvailabilityMutation {
	return pac.mutation
}

// Save creates the ProviderAvailability in the database.
func (pac *ProviderAvailabilityCreate) Save(ctx context.Context) (*ProviderAvailability, error) {
	return withHooks(ctx, pac.sqlSave, pac.mutation, pac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pac *ProviderAvailabilityCreate) SaveX(ctx context.Context) *ProviderAvailability {
	v, err := pac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pac *ProviderAvailabilityCreate) Exec(ctx context.Context) error {
	_, err := pac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pac *ProviderAvailabilityCreate) ExecX(ctx context.Context) {
	if err := pac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pac *ProviderAvailabilityCreate) check() error {
	if _, ok := pac.mutation.Cadence(); !ok {
		return &ValidationError{Name: "cadence", err: errors.New(`ent: missing required field "ProviderAvailability.cadence"`)}
	}
	if v, ok := pac.mutation.Cadence(); ok {
		if err := provideravailability.CadenceValidator(v); err != nil {
			return &ValidationError{Name: "cadence", err: fmt.Errorf(`ent: validator failed for field "ProviderAvailability.cadence": %w`, err)}
		}
	}
	if _, ok := pac.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "ProviderAvailability.start_time"`)}
	}
	if _, ok := pac.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "ProviderAvailability.end_time"`)}
	}
	return nil
}

func (pac *ProviderAvailabilityCreate) sqlSave(ctx context.Context) (*ProviderAvailability, error) {
	if err := pac.check(); err != nil {
		return nil, err
	}
	_node, _spec := pac.createSpec()
	if err := sqlgraph.CreateNode(ctx, pac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pac.mutation.id = &_node.ID
	pac.mutation.done = true
	return _node, nil
}

func (pac *ProviderAvailabilityCreate) createSpec() (*ProviderAvailability, *sqlgraph.CreateSpec) {
	var (
		_node = &ProviderAvailability{config: pac.config}
		_spec = sqlgraph.NewCreateSpec(provideravailability.Table, sqlgraph.NewFieldSpec(provideravailability.FieldID, field.TypeInt))
	)
	if value, ok := pac.mutation.Cadence(); ok {
		_spec.SetField(provideravailability.FieldCadence, field.TypeEnum, value)
		_node.Cadence = value
	}
	if value, ok := pac.mutation.StartTime(); ok {
		_spec.SetField(provideravailability.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := pac.mutation.EndTime(); ok {
		_spec.SetField(provideravailability.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if nodes := pac.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provideravailability.ProviderTable,
			Columns: []string{provideravailability.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provider_profile_availability = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProviderAvailabilityCreateBulk is the builder for creating many ProviderAvailability entities in bulk.
type ProviderAvailabilityCreateBulk struct {
	config
	builders []*ProviderAvailabilityCreate
}

// Save creates the ProviderAvailability entities in the database.
func (pacb *ProviderAvailabilityCreateBulk) Save(ctx context.Context) ([]*ProviderAvailability, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pacb.builders))
	nodes := make([]*ProviderAvailability, len(pacb.builders))
	mutators := make([]Mutator, len(pacb.builders))
	for i := range pacb.builders {
		func(i int, root context.Context) {
			builder := pacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProviderAvailabilityMutation)
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
					_, err = mutators[i+1].Mutate(root, pacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pacb *ProviderAvailabilityCreateBulk) SaveX(ctx context.Context) []*ProviderAvailability {
	v, err := pacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pacb *ProviderAvailabilityCreateBulk) Exec(ctx context.Context) error {
	_, err := pacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pacb *ProviderAvailabilityCreateBulk) ExecX(ctx context.Context) {
	if err := pacb.Exec(ctx); err != nil {
		panic(err)
	}
}
