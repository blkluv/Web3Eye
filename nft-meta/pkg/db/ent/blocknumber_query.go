// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/blocknumber"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/predicate"
)

// BlockNumberQuery is the builder for querying BlockNumber entities.
type BlockNumberQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BlockNumber
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlockNumberQuery builder.
func (bnq *BlockNumberQuery) Where(ps ...predicate.BlockNumber) *BlockNumberQuery {
	bnq.predicates = append(bnq.predicates, ps...)
	return bnq
}

// Limit adds a limit step to the query.
func (bnq *BlockNumberQuery) Limit(limit int) *BlockNumberQuery {
	bnq.limit = &limit
	return bnq
}

// Offset adds an offset step to the query.
func (bnq *BlockNumberQuery) Offset(offset int) *BlockNumberQuery {
	bnq.offset = &offset
	return bnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bnq *BlockNumberQuery) Unique(unique bool) *BlockNumberQuery {
	bnq.unique = &unique
	return bnq
}

// Order adds an order step to the query.
func (bnq *BlockNumberQuery) Order(o ...OrderFunc) *BlockNumberQuery {
	bnq.order = append(bnq.order, o...)
	return bnq
}

// First returns the first BlockNumber entity from the query.
// Returns a *NotFoundError when no BlockNumber was found.
func (bnq *BlockNumberQuery) First(ctx context.Context) (*BlockNumber, error) {
	nodes, err := bnq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blocknumber.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bnq *BlockNumberQuery) FirstX(ctx context.Context) *BlockNumber {
	node, err := bnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BlockNumber ID from the query.
// Returns a *NotFoundError when no BlockNumber ID was found.
func (bnq *BlockNumberQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bnq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blocknumber.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bnq *BlockNumberQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := bnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BlockNumber entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BlockNumber entity is found.
// Returns a *NotFoundError when no BlockNumber entities are found.
func (bnq *BlockNumberQuery) Only(ctx context.Context) (*BlockNumber, error) {
	nodes, err := bnq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blocknumber.Label}
	default:
		return nil, &NotSingularError{blocknumber.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bnq *BlockNumberQuery) OnlyX(ctx context.Context) *BlockNumber {
	node, err := bnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BlockNumber ID in the query.
// Returns a *NotSingularError when more than one BlockNumber ID is found.
// Returns a *NotFoundError when no entities are found.
func (bnq *BlockNumberQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bnq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blocknumber.Label}
	default:
		err = &NotSingularError{blocknumber.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bnq *BlockNumberQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := bnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlockNumbers.
func (bnq *BlockNumberQuery) All(ctx context.Context) ([]*BlockNumber, error) {
	if err := bnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bnq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bnq *BlockNumberQuery) AllX(ctx context.Context) []*BlockNumber {
	nodes, err := bnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BlockNumber IDs.
func (bnq *BlockNumberQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := bnq.Select(blocknumber.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bnq *BlockNumberQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := bnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bnq *BlockNumberQuery) Count(ctx context.Context) (int, error) {
	if err := bnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bnq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bnq *BlockNumberQuery) CountX(ctx context.Context) int {
	count, err := bnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bnq *BlockNumberQuery) Exist(ctx context.Context) (bool, error) {
	if err := bnq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bnq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bnq *BlockNumberQuery) ExistX(ctx context.Context) bool {
	exist, err := bnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlockNumberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bnq *BlockNumberQuery) Clone() *BlockNumberQuery {
	if bnq == nil {
		return nil
	}
	return &BlockNumberQuery{
		config:     bnq.config,
		limit:      bnq.limit,
		offset:     bnq.offset,
		order:      append([]OrderFunc{}, bnq.order...),
		predicates: append([]predicate.BlockNumber{}, bnq.predicates...),
		// clone intermediate query.
		sql:    bnq.sql.Clone(),
		path:   bnq.path,
		unique: bnq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BlockNumber.Query().
//		GroupBy(blocknumber.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bnq *BlockNumberQuery) GroupBy(field string, fields ...string) *BlockNumberGroupBy {
	grbuild := &BlockNumberGroupBy{config: bnq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bnq.sqlQuery(ctx), nil
	}
	grbuild.label = blocknumber.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.BlockNumber.Query().
//		Select(blocknumber.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (bnq *BlockNumberQuery) Select(fields ...string) *BlockNumberSelect {
	bnq.fields = append(bnq.fields, fields...)
	selbuild := &BlockNumberSelect{BlockNumberQuery: bnq}
	selbuild.label = blocknumber.Label
	selbuild.flds, selbuild.scan = &bnq.fields, selbuild.Scan
	return selbuild
}

func (bnq *BlockNumberQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bnq.fields {
		if !blocknumber.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bnq.path != nil {
		prev, err := bnq.path(ctx)
		if err != nil {
			return err
		}
		bnq.sql = prev
	}
	if blocknumber.Policy == nil {
		return errors.New("ent: uninitialized blocknumber.Policy (forgotten import ent/runtime?)")
	}
	if err := blocknumber.Policy.EvalQuery(ctx, bnq); err != nil {
		return err
	}
	return nil
}

func (bnq *BlockNumberQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BlockNumber, error) {
	var (
		nodes = []*BlockNumber{}
		_spec = bnq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BlockNumber).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BlockNumber{config: bnq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(bnq.modifiers) > 0 {
		_spec.Modifiers = bnq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bnq *BlockNumberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bnq.querySpec()
	if len(bnq.modifiers) > 0 {
		_spec.Modifiers = bnq.modifiers
	}
	_spec.Node.Columns = bnq.fields
	if len(bnq.fields) > 0 {
		_spec.Unique = bnq.unique != nil && *bnq.unique
	}
	return sqlgraph.CountNodes(ctx, bnq.driver, _spec)
}

func (bnq *BlockNumberQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bnq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bnq *BlockNumberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blocknumber.Table,
			Columns: blocknumber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blocknumber.FieldID,
			},
		},
		From:   bnq.sql,
		Unique: true,
	}
	if unique := bnq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bnq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blocknumber.FieldID)
		for i := range fields {
			if fields[i] != blocknumber.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bnq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bnq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bnq *BlockNumberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bnq.driver.Dialect())
	t1 := builder.Table(blocknumber.Table)
	columns := bnq.fields
	if len(columns) == 0 {
		columns = blocknumber.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bnq.sql != nil {
		selector = bnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bnq.unique != nil && *bnq.unique {
		selector.Distinct()
	}
	for _, m := range bnq.modifiers {
		m(selector)
	}
	for _, p := range bnq.predicates {
		p(selector)
	}
	for _, p := range bnq.order {
		p(selector)
	}
	if offset := bnq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bnq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (bnq *BlockNumberQuery) ForUpdate(opts ...sql.LockOption) *BlockNumberQuery {
	if bnq.driver.Dialect() == dialect.Postgres {
		bnq.Unique(false)
	}
	bnq.modifiers = append(bnq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return bnq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (bnq *BlockNumberQuery) ForShare(opts ...sql.LockOption) *BlockNumberQuery {
	if bnq.driver.Dialect() == dialect.Postgres {
		bnq.Unique(false)
	}
	bnq.modifiers = append(bnq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return bnq
}

// BlockNumberGroupBy is the group-by builder for BlockNumber entities.
type BlockNumberGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bngb *BlockNumberGroupBy) Aggregate(fns ...AggregateFunc) *BlockNumberGroupBy {
	bngb.fns = append(bngb.fns, fns...)
	return bngb
}

// Scan applies the group-by query and scans the result into the given value.
func (bngb *BlockNumberGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bngb.path(ctx)
	if err != nil {
		return err
	}
	bngb.sql = query
	return bngb.sqlScan(ctx, v)
}

func (bngb *BlockNumberGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bngb.fields {
		if !blocknumber.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bngb *BlockNumberGroupBy) sqlQuery() *sql.Selector {
	selector := bngb.sql.Select()
	aggregation := make([]string, 0, len(bngb.fns))
	for _, fn := range bngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bngb.fields)+len(bngb.fns))
		for _, f := range bngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bngb.fields...)...)
}

// BlockNumberSelect is the builder for selecting fields of BlockNumber entities.
type BlockNumberSelect struct {
	*BlockNumberQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bns *BlockNumberSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bns.prepareQuery(ctx); err != nil {
		return err
	}
	bns.sql = bns.BlockNumberQuery.sqlQuery(ctx)
	return bns.sqlScan(ctx, v)
}

func (bns *BlockNumberSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bns.sql.Query()
	if err := bns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
