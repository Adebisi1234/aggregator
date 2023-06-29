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
	"github.com/paycrest/paycrest-protocol/ent/apikey"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
	"github.com/paycrest/paycrest-protocol/ent/provideravailability"
	"github.com/paycrest/paycrest-protocol/ent/providerordertoken"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
)

// ProviderProfileUpdate is the builder for updating ProviderProfile entities.
type ProviderProfileUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderProfileMutation
}

// Where appends a list predicates to the ProviderProfileUpdate builder.
func (ppu *ProviderProfileUpdate) Where(ps ...predicate.ProviderProfile) *ProviderProfileUpdate {
	ppu.mutation.Where(ps...)
	return ppu
}

// SetUpdatedAt sets the "updated_at" field.
func (ppu *ProviderProfileUpdate) SetUpdatedAt(t time.Time) *ProviderProfileUpdate {
	ppu.mutation.SetUpdatedAt(t)
	return ppu
}

// SetTradingName sets the "trading_name" field.
func (ppu *ProviderProfileUpdate) SetTradingName(s string) *ProviderProfileUpdate {
	ppu.mutation.SetTradingName(s)
	return ppu
}

// SetCountry sets the "country" field.
func (ppu *ProviderProfileUpdate) SetCountry(s string) *ProviderProfileUpdate {
	ppu.mutation.SetCountry(s)
	return ppu
}

// SetAPIKeyID sets the "api_key" edge to the APIKey entity by ID.
func (ppu *ProviderProfileUpdate) SetAPIKeyID(id uuid.UUID) *ProviderProfileUpdate {
	ppu.mutation.SetAPIKeyID(id)
	return ppu
}

// SetNillableAPIKeyID sets the "api_key" edge to the APIKey entity by ID if the given value is not nil.
func (ppu *ProviderProfileUpdate) SetNillableAPIKeyID(id *uuid.UUID) *ProviderProfileUpdate {
	if id != nil {
		ppu = ppu.SetAPIKeyID(*id)
	}
	return ppu
}

// SetAPIKey sets the "api_key" edge to the APIKey entity.
func (ppu *ProviderProfileUpdate) SetAPIKey(a *APIKey) *ProviderProfileUpdate {
	return ppu.SetAPIKeyID(a.ID)
}

// AddOrderTokenIDs adds the "order_tokens" edge to the ProviderOrderToken entity by IDs.
func (ppu *ProviderProfileUpdate) AddOrderTokenIDs(ids ...int) *ProviderProfileUpdate {
	ppu.mutation.AddOrderTokenIDs(ids...)
	return ppu
}

// AddOrderTokens adds the "order_tokens" edges to the ProviderOrderToken entity.
func (ppu *ProviderProfileUpdate) AddOrderTokens(p ...*ProviderOrderToken) *ProviderProfileUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppu.AddOrderTokenIDs(ids...)
}

// AddAvailabilityIDs adds the "availability" edge to the ProviderAvailability entity by IDs.
func (ppu *ProviderProfileUpdate) AddAvailabilityIDs(ids ...int) *ProviderProfileUpdate {
	ppu.mutation.AddAvailabilityIDs(ids...)
	return ppu
}

// AddAvailability adds the "availability" edges to the ProviderAvailability entity.
func (ppu *ProviderProfileUpdate) AddAvailability(p ...*ProviderAvailability) *ProviderProfileUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppu.AddAvailabilityIDs(ids...)
}

// Mutation returns the ProviderProfileMutation object of the builder.
func (ppu *ProviderProfileUpdate) Mutation() *ProviderProfileMutation {
	return ppu.mutation
}

// ClearAPIKey clears the "api_key" edge to the APIKey entity.
func (ppu *ProviderProfileUpdate) ClearAPIKey() *ProviderProfileUpdate {
	ppu.mutation.ClearAPIKey()
	return ppu
}

// ClearOrderTokens clears all "order_tokens" edges to the ProviderOrderToken entity.
func (ppu *ProviderProfileUpdate) ClearOrderTokens() *ProviderProfileUpdate {
	ppu.mutation.ClearOrderTokens()
	return ppu
}

// RemoveOrderTokenIDs removes the "order_tokens" edge to ProviderOrderToken entities by IDs.
func (ppu *ProviderProfileUpdate) RemoveOrderTokenIDs(ids ...int) *ProviderProfileUpdate {
	ppu.mutation.RemoveOrderTokenIDs(ids...)
	return ppu
}

// RemoveOrderTokens removes "order_tokens" edges to ProviderOrderToken entities.
func (ppu *ProviderProfileUpdate) RemoveOrderTokens(p ...*ProviderOrderToken) *ProviderProfileUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppu.RemoveOrderTokenIDs(ids...)
}

// ClearAvailability clears all "availability" edges to the ProviderAvailability entity.
func (ppu *ProviderProfileUpdate) ClearAvailability() *ProviderProfileUpdate {
	ppu.mutation.ClearAvailability()
	return ppu
}

// RemoveAvailabilityIDs removes the "availability" edge to ProviderAvailability entities by IDs.
func (ppu *ProviderProfileUpdate) RemoveAvailabilityIDs(ids ...int) *ProviderProfileUpdate {
	ppu.mutation.RemoveAvailabilityIDs(ids...)
	return ppu
}

// RemoveAvailability removes "availability" edges to ProviderAvailability entities.
func (ppu *ProviderProfileUpdate) RemoveAvailability(p ...*ProviderAvailability) *ProviderProfileUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppu.RemoveAvailabilityIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ppu *ProviderProfileUpdate) Save(ctx context.Context) (int, error) {
	ppu.defaults()
	return withHooks(ctx, ppu.sqlSave, ppu.mutation, ppu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppu *ProviderProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := ppu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ppu *ProviderProfileUpdate) Exec(ctx context.Context) error {
	_, err := ppu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppu *ProviderProfileUpdate) ExecX(ctx context.Context) {
	if err := ppu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ppu *ProviderProfileUpdate) defaults() {
	if _, ok := ppu.mutation.UpdatedAt(); !ok {
		v := providerprofile.UpdateDefaultUpdatedAt()
		ppu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppu *ProviderProfileUpdate) check() error {
	if v, ok := ppu.mutation.TradingName(); ok {
		if err := providerprofile.TradingNameValidator(v); err != nil {
			return &ValidationError{Name: "trading_name", err: fmt.Errorf(`ent: validator failed for field "ProviderProfile.trading_name": %w`, err)}
		}
	}
	if v, ok := ppu.mutation.Country(); ok {
		if err := providerprofile.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "ProviderProfile.country": %w`, err)}
		}
	}
	return nil
}

func (ppu *ProviderProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ppu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(providerprofile.Table, providerprofile.Columns, sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString))
	if ps := ppu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ppu.mutation.UpdatedAt(); ok {
		_spec.SetField(providerprofile.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ppu.mutation.TradingName(); ok {
		_spec.SetField(providerprofile.FieldTradingName, field.TypeString, value)
	}
	if value, ok := ppu.mutation.Country(); ok {
		_spec.SetField(providerprofile.FieldCountry, field.TypeString, value)
	}
	if ppu.mutation.APIKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerprofile.APIKeyTable,
			Columns: []string{providerprofile.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.APIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerprofile.APIKeyTable,
			Columns: []string{providerprofile.APIKeyColumn},
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
	if ppu.mutation.OrderTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.OrderTokensTable,
			Columns: []string{providerprofile.OrderTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.RemovedOrderTokensIDs(); len(nodes) > 0 && !ppu.mutation.OrderTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.OrderTokensTable,
			Columns: []string{providerprofile.OrderTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.OrderTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.OrderTokensTable,
			Columns: []string{providerprofile.OrderTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ppu.mutation.AvailabilityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.AvailabilityTable,
			Columns: []string{providerprofile.AvailabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provideravailability.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.RemovedAvailabilityIDs(); len(nodes) > 0 && !ppu.mutation.AvailabilityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.AvailabilityTable,
			Columns: []string{providerprofile.AvailabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provideravailability.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.AvailabilityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.AvailabilityTable,
			Columns: []string{providerprofile.AvailabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provideravailability.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ppu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providerprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ppu.mutation.done = true
	return n, nil
}

// ProviderProfileUpdateOne is the builder for updating a single ProviderProfile entity.
type ProviderProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderProfileMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ppuo *ProviderProfileUpdateOne) SetUpdatedAt(t time.Time) *ProviderProfileUpdateOne {
	ppuo.mutation.SetUpdatedAt(t)
	return ppuo
}

// SetTradingName sets the "trading_name" field.
func (ppuo *ProviderProfileUpdateOne) SetTradingName(s string) *ProviderProfileUpdateOne {
	ppuo.mutation.SetTradingName(s)
	return ppuo
}

// SetCountry sets the "country" field.
func (ppuo *ProviderProfileUpdateOne) SetCountry(s string) *ProviderProfileUpdateOne {
	ppuo.mutation.SetCountry(s)
	return ppuo
}

// SetAPIKeyID sets the "api_key" edge to the APIKey entity by ID.
func (ppuo *ProviderProfileUpdateOne) SetAPIKeyID(id uuid.UUID) *ProviderProfileUpdateOne {
	ppuo.mutation.SetAPIKeyID(id)
	return ppuo
}

// SetNillableAPIKeyID sets the "api_key" edge to the APIKey entity by ID if the given value is not nil.
func (ppuo *ProviderProfileUpdateOne) SetNillableAPIKeyID(id *uuid.UUID) *ProviderProfileUpdateOne {
	if id != nil {
		ppuo = ppuo.SetAPIKeyID(*id)
	}
	return ppuo
}

// SetAPIKey sets the "api_key" edge to the APIKey entity.
func (ppuo *ProviderProfileUpdateOne) SetAPIKey(a *APIKey) *ProviderProfileUpdateOne {
	return ppuo.SetAPIKeyID(a.ID)
}

// AddOrderTokenIDs adds the "order_tokens" edge to the ProviderOrderToken entity by IDs.
func (ppuo *ProviderProfileUpdateOne) AddOrderTokenIDs(ids ...int) *ProviderProfileUpdateOne {
	ppuo.mutation.AddOrderTokenIDs(ids...)
	return ppuo
}

// AddOrderTokens adds the "order_tokens" edges to the ProviderOrderToken entity.
func (ppuo *ProviderProfileUpdateOne) AddOrderTokens(p ...*ProviderOrderToken) *ProviderProfileUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppuo.AddOrderTokenIDs(ids...)
}

// AddAvailabilityIDs adds the "availability" edge to the ProviderAvailability entity by IDs.
func (ppuo *ProviderProfileUpdateOne) AddAvailabilityIDs(ids ...int) *ProviderProfileUpdateOne {
	ppuo.mutation.AddAvailabilityIDs(ids...)
	return ppuo
}

// AddAvailability adds the "availability" edges to the ProviderAvailability entity.
func (ppuo *ProviderProfileUpdateOne) AddAvailability(p ...*ProviderAvailability) *ProviderProfileUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppuo.AddAvailabilityIDs(ids...)
}

// Mutation returns the ProviderProfileMutation object of the builder.
func (ppuo *ProviderProfileUpdateOne) Mutation() *ProviderProfileMutation {
	return ppuo.mutation
}

// ClearAPIKey clears the "api_key" edge to the APIKey entity.
func (ppuo *ProviderProfileUpdateOne) ClearAPIKey() *ProviderProfileUpdateOne {
	ppuo.mutation.ClearAPIKey()
	return ppuo
}

// ClearOrderTokens clears all "order_tokens" edges to the ProviderOrderToken entity.
func (ppuo *ProviderProfileUpdateOne) ClearOrderTokens() *ProviderProfileUpdateOne {
	ppuo.mutation.ClearOrderTokens()
	return ppuo
}

// RemoveOrderTokenIDs removes the "order_tokens" edge to ProviderOrderToken entities by IDs.
func (ppuo *ProviderProfileUpdateOne) RemoveOrderTokenIDs(ids ...int) *ProviderProfileUpdateOne {
	ppuo.mutation.RemoveOrderTokenIDs(ids...)
	return ppuo
}

// RemoveOrderTokens removes "order_tokens" edges to ProviderOrderToken entities.
func (ppuo *ProviderProfileUpdateOne) RemoveOrderTokens(p ...*ProviderOrderToken) *ProviderProfileUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppuo.RemoveOrderTokenIDs(ids...)
}

// ClearAvailability clears all "availability" edges to the ProviderAvailability entity.
func (ppuo *ProviderProfileUpdateOne) ClearAvailability() *ProviderProfileUpdateOne {
	ppuo.mutation.ClearAvailability()
	return ppuo
}

// RemoveAvailabilityIDs removes the "availability" edge to ProviderAvailability entities by IDs.
func (ppuo *ProviderProfileUpdateOne) RemoveAvailabilityIDs(ids ...int) *ProviderProfileUpdateOne {
	ppuo.mutation.RemoveAvailabilityIDs(ids...)
	return ppuo
}

// RemoveAvailability removes "availability" edges to ProviderAvailability entities.
func (ppuo *ProviderProfileUpdateOne) RemoveAvailability(p ...*ProviderAvailability) *ProviderProfileUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppuo.RemoveAvailabilityIDs(ids...)
}

// Where appends a list predicates to the ProviderProfileUpdate builder.
func (ppuo *ProviderProfileUpdateOne) Where(ps ...predicate.ProviderProfile) *ProviderProfileUpdateOne {
	ppuo.mutation.Where(ps...)
	return ppuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ppuo *ProviderProfileUpdateOne) Select(field string, fields ...string) *ProviderProfileUpdateOne {
	ppuo.fields = append([]string{field}, fields...)
	return ppuo
}

// Save executes the query and returns the updated ProviderProfile entity.
func (ppuo *ProviderProfileUpdateOne) Save(ctx context.Context) (*ProviderProfile, error) {
	ppuo.defaults()
	return withHooks(ctx, ppuo.sqlSave, ppuo.mutation, ppuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppuo *ProviderProfileUpdateOne) SaveX(ctx context.Context) *ProviderProfile {
	node, err := ppuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ppuo *ProviderProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := ppuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppuo *ProviderProfileUpdateOne) ExecX(ctx context.Context) {
	if err := ppuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ppuo *ProviderProfileUpdateOne) defaults() {
	if _, ok := ppuo.mutation.UpdatedAt(); !ok {
		v := providerprofile.UpdateDefaultUpdatedAt()
		ppuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppuo *ProviderProfileUpdateOne) check() error {
	if v, ok := ppuo.mutation.TradingName(); ok {
		if err := providerprofile.TradingNameValidator(v); err != nil {
			return &ValidationError{Name: "trading_name", err: fmt.Errorf(`ent: validator failed for field "ProviderProfile.trading_name": %w`, err)}
		}
	}
	if v, ok := ppuo.mutation.Country(); ok {
		if err := providerprofile.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "ProviderProfile.country": %w`, err)}
		}
	}
	return nil
}

func (ppuo *ProviderProfileUpdateOne) sqlSave(ctx context.Context) (_node *ProviderProfile, err error) {
	if err := ppuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(providerprofile.Table, providerprofile.Columns, sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString))
	id, ok := ppuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProviderProfile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ppuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, providerprofile.FieldID)
		for _, f := range fields {
			if !providerprofile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != providerprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ppuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ppuo.mutation.UpdatedAt(); ok {
		_spec.SetField(providerprofile.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ppuo.mutation.TradingName(); ok {
		_spec.SetField(providerprofile.FieldTradingName, field.TypeString, value)
	}
	if value, ok := ppuo.mutation.Country(); ok {
		_spec.SetField(providerprofile.FieldCountry, field.TypeString, value)
	}
	if ppuo.mutation.APIKeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerprofile.APIKeyTable,
			Columns: []string{providerprofile.APIKeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.APIKeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   providerprofile.APIKeyTable,
			Columns: []string{providerprofile.APIKeyColumn},
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
	if ppuo.mutation.OrderTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.OrderTokensTable,
			Columns: []string{providerprofile.OrderTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.RemovedOrderTokensIDs(); len(nodes) > 0 && !ppuo.mutation.OrderTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.OrderTokensTable,
			Columns: []string{providerprofile.OrderTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.OrderTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.OrderTokensTable,
			Columns: []string{providerprofile.OrderTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(providerordertoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ppuo.mutation.AvailabilityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.AvailabilityTable,
			Columns: []string{providerprofile.AvailabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provideravailability.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.RemovedAvailabilityIDs(); len(nodes) > 0 && !ppuo.mutation.AvailabilityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.AvailabilityTable,
			Columns: []string{providerprofile.AvailabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provideravailability.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.AvailabilityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   providerprofile.AvailabilityTable,
			Columns: []string{providerprofile.AvailabilityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(provideravailability.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProviderProfile{config: ppuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ppuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{providerprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ppuo.mutation.done = true
	return _node, nil
}
