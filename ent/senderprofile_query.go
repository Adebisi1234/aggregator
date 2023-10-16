// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
	"github.com/paycrest/paycrest-protocol/ent/senderprofile"
	"github.com/paycrest/paycrest-protocol/ent/user"
)

// SenderProfileQuery is the builder for querying SenderProfile entities.
type SenderProfileQuery struct {
	config
	ctx        *QueryContext
	order      []senderprofile.OrderOption
	inters     []Interceptor
	predicates []predicate.SenderProfile
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SenderProfileQuery builder.
func (spq *SenderProfileQuery) Where(ps ...predicate.SenderProfile) *SenderProfileQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *SenderProfileQuery) Limit(limit int) *SenderProfileQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *SenderProfileQuery) Offset(offset int) *SenderProfileQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *SenderProfileQuery) Unique(unique bool) *SenderProfileQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *SenderProfileQuery) Order(o ...senderprofile.OrderOption) *SenderProfileQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryUser chains the current query on the "user" edge.
func (spq *SenderProfileQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(senderprofile.Table, senderprofile.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, senderprofile.UserTable, senderprofile.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SenderProfile entity from the query.
// Returns a *NotFoundError when no SenderProfile was found.
func (spq *SenderProfileQuery) First(ctx context.Context) (*SenderProfile, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{senderprofile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *SenderProfileQuery) FirstX(ctx context.Context) *SenderProfile {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SenderProfile ID from the query.
// Returns a *NotFoundError when no SenderProfile ID was found.
func (spq *SenderProfileQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{senderprofile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *SenderProfileQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SenderProfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SenderProfile entity is found.
// Returns a *NotFoundError when no SenderProfile entities are found.
func (spq *SenderProfileQuery) Only(ctx context.Context) (*SenderProfile, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{senderprofile.Label}
	default:
		return nil, &NotSingularError{senderprofile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *SenderProfileQuery) OnlyX(ctx context.Context) *SenderProfile {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SenderProfile ID in the query.
// Returns a *NotSingularError when more than one SenderProfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *SenderProfileQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{senderprofile.Label}
	default:
		err = &NotSingularError{senderprofile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *SenderProfileQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SenderProfiles.
func (spq *SenderProfileQuery) All(ctx context.Context) ([]*SenderProfile, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SenderProfile, *SenderProfileQuery]()
	return withInterceptors[[]*SenderProfile](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *SenderProfileQuery) AllX(ctx context.Context) []*SenderProfile {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SenderProfile IDs.
func (spq *SenderProfileQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(senderprofile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *SenderProfileQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *SenderProfileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*SenderProfileQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *SenderProfileQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *SenderProfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *SenderProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SenderProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *SenderProfileQuery) Clone() *SenderProfileQuery {
	if spq == nil {
		return nil
	}
	return &SenderProfileQuery{
		config:     spq.config,
		ctx:        spq.ctx.Clone(),
		order:      append([]senderprofile.OrderOption{}, spq.order...),
		inters:     append([]Interceptor{}, spq.inters...),
		predicates: append([]predicate.SenderProfile{}, spq.predicates...),
		withUser:   spq.withUser.Clone(),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *SenderProfileQuery) WithUser(opts ...func(*UserQuery)) *SenderProfileQuery {
	query := (&UserClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withUser = query
	return spq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WebhookURL string `json:"webhook_url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SenderProfile.Query().
//		GroupBy(senderprofile.FieldWebhookURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (spq *SenderProfileQuery) GroupBy(field string, fields ...string) *SenderProfileGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SenderProfileGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = senderprofile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WebhookURL string `json:"webhook_url,omitempty"`
//	}
//
//	client.SenderProfile.Query().
//		Select(senderprofile.FieldWebhookURL).
//		Scan(ctx, &v)
func (spq *SenderProfileQuery) Select(fields ...string) *SenderProfileSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &SenderProfileSelect{SenderProfileQuery: spq}
	sbuild.label = senderprofile.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SenderProfileSelect configured with the given aggregations.
func (spq *SenderProfileQuery) Aggregate(fns ...AggregateFunc) *SenderProfileSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *SenderProfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	for _, f := range spq.ctx.Fields {
		if !senderprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *SenderProfileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SenderProfile, error) {
	var (
		nodes       = []*SenderProfile{}
		withFKs     = spq.withFKs
		_spec       = spq.querySpec()
		loadedTypes = [1]bool{
			spq.withUser != nil,
		}
	)
	if spq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, senderprofile.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SenderProfile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SenderProfile{config: spq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := spq.withUser; query != nil {
		if err := spq.loadUser(ctx, query, nodes, nil,
			func(n *SenderProfile, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (spq *SenderProfileQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*SenderProfile, init func(*SenderProfile), assign func(*SenderProfile, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*SenderProfile)
	for i := range nodes {
		if nodes[i].user_sender_profile == nil {
			continue
		}
		fk := *nodes[i].user_sender_profile
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_sender_profile" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (spq *SenderProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.ctx.Fields
	if len(spq.ctx.Fields) > 0 {
		_spec.Unique = spq.ctx.Unique != nil && *spq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *SenderProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(senderprofile.Table, senderprofile.Columns, sqlgraph.NewFieldSpec(senderprofile.FieldID, field.TypeUUID))
	_spec.From = spq.sql
	if unique := spq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if spq.path != nil {
		_spec.Unique = true
	}
	if fields := spq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, senderprofile.FieldID)
		for i := range fields {
			if fields[i] != senderprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *SenderProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(senderprofile.Table)
	columns := spq.ctx.Fields
	if len(columns) == 0 {
		columns = senderprofile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.ctx.Unique != nil && *spq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SenderProfileGroupBy is the group-by builder for SenderProfile entities.
type SenderProfileGroupBy struct {
	selector
	build *SenderProfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *SenderProfileGroupBy) Aggregate(fns ...AggregateFunc) *SenderProfileGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *SenderProfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SenderProfileQuery, *SenderProfileGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *SenderProfileGroupBy) sqlScan(ctx context.Context, root *SenderProfileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*spgb.flds)+len(spgb.fns))
		for _, f := range *spgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*spgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SenderProfileSelect is the builder for selecting fields of SenderProfile entities.
type SenderProfileSelect struct {
	*SenderProfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *SenderProfileSelect) Aggregate(fns ...AggregateFunc) *SenderProfileSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *SenderProfileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SenderProfileQuery, *SenderProfileSelect](ctx, sps.SenderProfileQuery, sps, sps.inters, v)
}

func (sps *SenderProfileSelect) sqlScan(ctx context.Context, root *SenderProfileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sps.fns))
	for _, fn := range sps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
