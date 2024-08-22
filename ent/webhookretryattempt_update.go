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
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/paycrest/protocol/ent/webhookretryattempt"
)

// WebhookRetryAttemptUpdate is the builder for updating WebhookRetryAttempt entities.
type WebhookRetryAttemptUpdate struct {
	config
	hooks    []Hook
	mutation *WebhookRetryAttemptMutation
}

// Where appends a list predicates to the WebhookRetryAttemptUpdate builder.
func (wrau *WebhookRetryAttemptUpdate) Where(ps ...predicate.WebhookRetryAttempt) *WebhookRetryAttemptUpdate {
	wrau.mutation.Where(ps...)
	return wrau
}

// SetUpdatedAt sets the "updated_at" field.
func (wrau *WebhookRetryAttemptUpdate) SetUpdatedAt(t time.Time) *WebhookRetryAttemptUpdate {
	wrau.mutation.SetUpdatedAt(t)
	return wrau
}

// SetAttemptNumber sets the "attempt_number" field.
func (wrau *WebhookRetryAttemptUpdate) SetAttemptNumber(i int) *WebhookRetryAttemptUpdate {
	wrau.mutation.ResetAttemptNumber()
	wrau.mutation.SetAttemptNumber(i)
	return wrau
}

// SetNillableAttemptNumber sets the "attempt_number" field if the given value is not nil.
func (wrau *WebhookRetryAttemptUpdate) SetNillableAttemptNumber(i *int) *WebhookRetryAttemptUpdate {
	if i != nil {
		wrau.SetAttemptNumber(*i)
	}
	return wrau
}

// AddAttemptNumber adds i to the "attempt_number" field.
func (wrau *WebhookRetryAttemptUpdate) AddAttemptNumber(i int) *WebhookRetryAttemptUpdate {
	wrau.mutation.AddAttemptNumber(i)
	return wrau
}

// SetNextRetryTime sets the "next_retry_time" field.
func (wrau *WebhookRetryAttemptUpdate) SetNextRetryTime(t time.Time) *WebhookRetryAttemptUpdate {
	wrau.mutation.SetNextRetryTime(t)
	return wrau
}

// SetNillableNextRetryTime sets the "next_retry_time" field if the given value is not nil.
func (wrau *WebhookRetryAttemptUpdate) SetNillableNextRetryTime(t *time.Time) *WebhookRetryAttemptUpdate {
	if t != nil {
		wrau.SetNextRetryTime(*t)
	}
	return wrau
}

// SetPayload sets the "payload" field.
func (wrau *WebhookRetryAttemptUpdate) SetPayload(m map[string]interface{}) *WebhookRetryAttemptUpdate {
	wrau.mutation.SetPayload(m)
	return wrau
}

// SetSignature sets the "signature" field.
func (wrau *WebhookRetryAttemptUpdate) SetSignature(s string) *WebhookRetryAttemptUpdate {
	wrau.mutation.SetSignature(s)
	return wrau
}

// SetNillableSignature sets the "signature" field if the given value is not nil.
func (wrau *WebhookRetryAttemptUpdate) SetNillableSignature(s *string) *WebhookRetryAttemptUpdate {
	if s != nil {
		wrau.SetSignature(*s)
	}
	return wrau
}

// ClearSignature clears the value of the "signature" field.
func (wrau *WebhookRetryAttemptUpdate) ClearSignature() *WebhookRetryAttemptUpdate {
	wrau.mutation.ClearSignature()
	return wrau
}

// SetWebhookURL sets the "webhook_url" field.
func (wrau *WebhookRetryAttemptUpdate) SetWebhookURL(s string) *WebhookRetryAttemptUpdate {
	wrau.mutation.SetWebhookURL(s)
	return wrau
}

// SetNillableWebhookURL sets the "webhook_url" field if the given value is not nil.
func (wrau *WebhookRetryAttemptUpdate) SetNillableWebhookURL(s *string) *WebhookRetryAttemptUpdate {
	if s != nil {
		wrau.SetWebhookURL(*s)
	}
	return wrau
}

// SetStatus sets the "status" field.
func (wrau *WebhookRetryAttemptUpdate) SetStatus(w webhookretryattempt.Status) *WebhookRetryAttemptUpdate {
	wrau.mutation.SetStatus(w)
	return wrau
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wrau *WebhookRetryAttemptUpdate) SetNillableStatus(w *webhookretryattempt.Status) *WebhookRetryAttemptUpdate {
	if w != nil {
		wrau.SetStatus(*w)
	}
	return wrau
}

// Mutation returns the WebhookRetryAttemptMutation object of the builder.
func (wrau *WebhookRetryAttemptUpdate) Mutation() *WebhookRetryAttemptMutation {
	return wrau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wrau *WebhookRetryAttemptUpdate) Save(ctx context.Context) (int, error) {
	wrau.defaults()
	return withHooks(ctx, wrau.sqlSave, wrau.mutation, wrau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wrau *WebhookRetryAttemptUpdate) SaveX(ctx context.Context) int {
	affected, err := wrau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wrau *WebhookRetryAttemptUpdate) Exec(ctx context.Context) error {
	_, err := wrau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrau *WebhookRetryAttemptUpdate) ExecX(ctx context.Context) {
	if err := wrau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wrau *WebhookRetryAttemptUpdate) defaults() {
	if _, ok := wrau.mutation.UpdatedAt(); !ok {
		v := webhookretryattempt.UpdateDefaultUpdatedAt()
		wrau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wrau *WebhookRetryAttemptUpdate) check() error {
	if v, ok := wrau.mutation.Status(); ok {
		if err := webhookretryattempt.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "WebhookRetryAttempt.status": %w`, err)}
		}
	}
	return nil
}

func (wrau *WebhookRetryAttemptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wrau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(webhookretryattempt.Table, webhookretryattempt.Columns, sqlgraph.NewFieldSpec(webhookretryattempt.FieldID, field.TypeInt))
	if ps := wrau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wrau.mutation.UpdatedAt(); ok {
		_spec.SetField(webhookretryattempt.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wrau.mutation.AttemptNumber(); ok {
		_spec.SetField(webhookretryattempt.FieldAttemptNumber, field.TypeInt, value)
	}
	if value, ok := wrau.mutation.AddedAttemptNumber(); ok {
		_spec.AddField(webhookretryattempt.FieldAttemptNumber, field.TypeInt, value)
	}
	if value, ok := wrau.mutation.NextRetryTime(); ok {
		_spec.SetField(webhookretryattempt.FieldNextRetryTime, field.TypeTime, value)
	}
	if value, ok := wrau.mutation.Payload(); ok {
		_spec.SetField(webhookretryattempt.FieldPayload, field.TypeJSON, value)
	}
	if value, ok := wrau.mutation.Signature(); ok {
		_spec.SetField(webhookretryattempt.FieldSignature, field.TypeString, value)
	}
	if wrau.mutation.SignatureCleared() {
		_spec.ClearField(webhookretryattempt.FieldSignature, field.TypeString)
	}
	if value, ok := wrau.mutation.WebhookURL(); ok {
		_spec.SetField(webhookretryattempt.FieldWebhookURL, field.TypeString, value)
	}
	if value, ok := wrau.mutation.Status(); ok {
		_spec.SetField(webhookretryattempt.FieldStatus, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wrau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{webhookretryattempt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wrau.mutation.done = true
	return n, nil
}

// WebhookRetryAttemptUpdateOne is the builder for updating a single WebhookRetryAttempt entity.
type WebhookRetryAttemptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WebhookRetryAttemptMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (wrauo *WebhookRetryAttemptUpdateOne) SetUpdatedAt(t time.Time) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.SetUpdatedAt(t)
	return wrauo
}

// SetAttemptNumber sets the "attempt_number" field.
func (wrauo *WebhookRetryAttemptUpdateOne) SetAttemptNumber(i int) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.ResetAttemptNumber()
	wrauo.mutation.SetAttemptNumber(i)
	return wrauo
}

// SetNillableAttemptNumber sets the "attempt_number" field if the given value is not nil.
func (wrauo *WebhookRetryAttemptUpdateOne) SetNillableAttemptNumber(i *int) *WebhookRetryAttemptUpdateOne {
	if i != nil {
		wrauo.SetAttemptNumber(*i)
	}
	return wrauo
}

// AddAttemptNumber adds i to the "attempt_number" field.
func (wrauo *WebhookRetryAttemptUpdateOne) AddAttemptNumber(i int) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.AddAttemptNumber(i)
	return wrauo
}

// SetNextRetryTime sets the "next_retry_time" field.
func (wrauo *WebhookRetryAttemptUpdateOne) SetNextRetryTime(t time.Time) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.SetNextRetryTime(t)
	return wrauo
}

// SetNillableNextRetryTime sets the "next_retry_time" field if the given value is not nil.
func (wrauo *WebhookRetryAttemptUpdateOne) SetNillableNextRetryTime(t *time.Time) *WebhookRetryAttemptUpdateOne {
	if t != nil {
		wrauo.SetNextRetryTime(*t)
	}
	return wrauo
}

// SetPayload sets the "payload" field.
func (wrauo *WebhookRetryAttemptUpdateOne) SetPayload(m map[string]interface{}) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.SetPayload(m)
	return wrauo
}

// SetSignature sets the "signature" field.
func (wrauo *WebhookRetryAttemptUpdateOne) SetSignature(s string) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.SetSignature(s)
	return wrauo
}

// SetNillableSignature sets the "signature" field if the given value is not nil.
func (wrauo *WebhookRetryAttemptUpdateOne) SetNillableSignature(s *string) *WebhookRetryAttemptUpdateOne {
	if s != nil {
		wrauo.SetSignature(*s)
	}
	return wrauo
}

// ClearSignature clears the value of the "signature" field.
func (wrauo *WebhookRetryAttemptUpdateOne) ClearSignature() *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.ClearSignature()
	return wrauo
}

// SetWebhookURL sets the "webhook_url" field.
func (wrauo *WebhookRetryAttemptUpdateOne) SetWebhookURL(s string) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.SetWebhookURL(s)
	return wrauo
}

// SetNillableWebhookURL sets the "webhook_url" field if the given value is not nil.
func (wrauo *WebhookRetryAttemptUpdateOne) SetNillableWebhookURL(s *string) *WebhookRetryAttemptUpdateOne {
	if s != nil {
		wrauo.SetWebhookURL(*s)
	}
	return wrauo
}

// SetStatus sets the "status" field.
func (wrauo *WebhookRetryAttemptUpdateOne) SetStatus(w webhookretryattempt.Status) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.SetStatus(w)
	return wrauo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wrauo *WebhookRetryAttemptUpdateOne) SetNillableStatus(w *webhookretryattempt.Status) *WebhookRetryAttemptUpdateOne {
	if w != nil {
		wrauo.SetStatus(*w)
	}
	return wrauo
}

// Mutation returns the WebhookRetryAttemptMutation object of the builder.
func (wrauo *WebhookRetryAttemptUpdateOne) Mutation() *WebhookRetryAttemptMutation {
	return wrauo.mutation
}

// Where appends a list predicates to the WebhookRetryAttemptUpdate builder.
func (wrauo *WebhookRetryAttemptUpdateOne) Where(ps ...predicate.WebhookRetryAttempt) *WebhookRetryAttemptUpdateOne {
	wrauo.mutation.Where(ps...)
	return wrauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wrauo *WebhookRetryAttemptUpdateOne) Select(field string, fields ...string) *WebhookRetryAttemptUpdateOne {
	wrauo.fields = append([]string{field}, fields...)
	return wrauo
}

// Save executes the query and returns the updated WebhookRetryAttempt entity.
func (wrauo *WebhookRetryAttemptUpdateOne) Save(ctx context.Context) (*WebhookRetryAttempt, error) {
	wrauo.defaults()
	return withHooks(ctx, wrauo.sqlSave, wrauo.mutation, wrauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wrauo *WebhookRetryAttemptUpdateOne) SaveX(ctx context.Context) *WebhookRetryAttempt {
	node, err := wrauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wrauo *WebhookRetryAttemptUpdateOne) Exec(ctx context.Context) error {
	_, err := wrauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wrauo *WebhookRetryAttemptUpdateOne) ExecX(ctx context.Context) {
	if err := wrauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wrauo *WebhookRetryAttemptUpdateOne) defaults() {
	if _, ok := wrauo.mutation.UpdatedAt(); !ok {
		v := webhookretryattempt.UpdateDefaultUpdatedAt()
		wrauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wrauo *WebhookRetryAttemptUpdateOne) check() error {
	if v, ok := wrauo.mutation.Status(); ok {
		if err := webhookretryattempt.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "WebhookRetryAttempt.status": %w`, err)}
		}
	}
	return nil
}

func (wrauo *WebhookRetryAttemptUpdateOne) sqlSave(ctx context.Context) (_node *WebhookRetryAttempt, err error) {
	if err := wrauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(webhookretryattempt.Table, webhookretryattempt.Columns, sqlgraph.NewFieldSpec(webhookretryattempt.FieldID, field.TypeInt))
	id, ok := wrauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WebhookRetryAttempt.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wrauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webhookretryattempt.FieldID)
		for _, f := range fields {
			if !webhookretryattempt.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != webhookretryattempt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wrauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wrauo.mutation.UpdatedAt(); ok {
		_spec.SetField(webhookretryattempt.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wrauo.mutation.AttemptNumber(); ok {
		_spec.SetField(webhookretryattempt.FieldAttemptNumber, field.TypeInt, value)
	}
	if value, ok := wrauo.mutation.AddedAttemptNumber(); ok {
		_spec.AddField(webhookretryattempt.FieldAttemptNumber, field.TypeInt, value)
	}
	if value, ok := wrauo.mutation.NextRetryTime(); ok {
		_spec.SetField(webhookretryattempt.FieldNextRetryTime, field.TypeTime, value)
	}
	if value, ok := wrauo.mutation.Payload(); ok {
		_spec.SetField(webhookretryattempt.FieldPayload, field.TypeJSON, value)
	}
	if value, ok := wrauo.mutation.Signature(); ok {
		_spec.SetField(webhookretryattempt.FieldSignature, field.TypeString, value)
	}
	if wrauo.mutation.SignatureCleared() {
		_spec.ClearField(webhookretryattempt.FieldSignature, field.TypeString)
	}
	if value, ok := wrauo.mutation.WebhookURL(); ok {
		_spec.SetField(webhookretryattempt.FieldWebhookURL, field.TypeString, value)
	}
	if value, ok := wrauo.mutation.Status(); ok {
		_spec.SetField(webhookretryattempt.FieldStatus, field.TypeEnum, value)
	}
	_node = &WebhookRetryAttempt{config: wrauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wrauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{webhookretryattempt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wrauo.mutation.done = true
	return _node, nil
}
