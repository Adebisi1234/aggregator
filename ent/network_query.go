// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/paycrest-protocol/ent/network"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
	"github.com/paycrest/paycrest-protocol/ent/token"
)

// NetworkQuery is the builder for querying Network entities.
type NetworkQuery struct {
	config
	ctx        *QueryContext
	order      []network.OrderOption
	inters     []Interceptor
	predicates []predicate.Network
	withTokens *TokenQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NetworkQuery builder.
func (nq *NetworkQuery) Where(ps ...predicate.Network) *NetworkQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NetworkQuery) Limit(limit int) *NetworkQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NetworkQuery) Offset(offset int) *NetworkQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NetworkQuery) Unique(unique bool) *NetworkQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NetworkQuery) Order(o ...network.OrderOption) *NetworkQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryTokens chains the current query on the "tokens" edge.
func (nq *NetworkQuery) QueryTokens() *TokenQuery {
	query := (&TokenClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(network.Table, network.FieldID, selector),
			sqlgraph.To(token.Table, token.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, network.TokensTable, network.TokensColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Network entity from the query.
// Returns a *NotFoundError when no Network was found.
func (nq *NetworkQuery) First(ctx context.Context) (*Network, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{network.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NetworkQuery) FirstX(ctx context.Context) *Network {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Network ID from the query.
// Returns a *NotFoundError when no Network ID was found.
func (nq *NetworkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{network.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NetworkQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Network entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Network entity is found.
// Returns a *NotFoundError when no Network entities are found.
func (nq *NetworkQuery) Only(ctx context.Context) (*Network, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{network.Label}
	default:
		return nil, &NotSingularError{network.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NetworkQuery) OnlyX(ctx context.Context) *Network {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Network ID in the query.
// Returns a *NotSingularError when more than one Network ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NetworkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = &NotSingularError{network.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NetworkQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Networks.
func (nq *NetworkQuery) All(ctx context.Context) ([]*Network, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Network, *NetworkQuery]()
	return withInterceptors[[]*Network](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NetworkQuery) AllX(ctx context.Context) []*Network {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Network IDs.
func (nq *NetworkQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(network.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NetworkQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NetworkQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NetworkQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NetworkQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NetworkQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NetworkQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NetworkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NetworkQuery) Clone() *NetworkQuery {
	if nq == nil {
		return nil
	}
	return &NetworkQuery{
		config:     nq.config,
		ctx:        nq.ctx.Clone(),
		order:      append([]network.OrderOption{}, nq.order...),
		inters:     append([]Interceptor{}, nq.inters...),
		predicates: append([]predicate.Network{}, nq.predicates...),
		withTokens: nq.withTokens.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithTokens tells the query-builder to eager-load the nodes that are connected to
// the "tokens" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NetworkQuery) WithTokens(opts ...func(*TokenQuery)) *NetworkQuery {
	query := (&TokenClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withTokens = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Network.Query().
//		GroupBy(network.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NetworkQuery) GroupBy(field string, fields ...string) *NetworkGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NetworkGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = network.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Network.Query().
//		Select(network.FieldCreatedAt).
//		Scan(ctx, &v)
func (nq *NetworkQuery) Select(fields ...string) *NetworkSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NetworkSelect{NetworkQuery: nq}
	sbuild.label = network.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NetworkSelect configured with the given aggregations.
func (nq *NetworkQuery) Aggregate(fns ...AggregateFunc) *NetworkSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NetworkQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !network.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NetworkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Network, error) {
	var (
		nodes       = []*Network{}
		_spec       = nq.querySpec()
		loadedTypes = [1]bool{
			nq.withTokens != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Network).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Network{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withTokens; query != nil {
		if err := nq.loadTokens(ctx, query, nodes,
			func(n *Network) { n.Edges.Tokens = []*Token{} },
			func(n *Network, e *Token) { n.Edges.Tokens = append(n.Edges.Tokens, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NetworkQuery) loadTokens(ctx context.Context, query *TokenQuery, nodes []*Network, init func(*Network), assign func(*Network, *Token)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Network)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Token(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(network.TokensColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.network_tokens
		if fk == nil {
			return fmt.Errorf(`foreign-key "network_tokens" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "network_tokens" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nq *NetworkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NetworkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(network.Table, network.Columns, sqlgraph.NewFieldSpec(network.FieldID, field.TypeInt))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, network.FieldID)
		for i := range fields {
			if fields[i] != network.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NetworkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(network.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = network.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NetworkGroupBy is the group-by builder for Network entities.
type NetworkGroupBy struct {
	selector
	build *NetworkQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NetworkGroupBy) Aggregate(fns ...AggregateFunc) *NetworkGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NetworkGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NetworkQuery, *NetworkGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NetworkGroupBy) sqlScan(ctx context.Context, root *NetworkQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NetworkSelect is the builder for selecting fields of Network entities.
type NetworkSelect struct {
	*NetworkQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NetworkSelect) Aggregate(fns ...AggregateFunc) *NetworkSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NetworkSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NetworkQuery, *NetworkSelect](ctx, ns.NetworkQuery, ns, ns.inters, v)
}

func (ns *NetworkSelect) sqlScan(ctx context.Context, root *NetworkQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
