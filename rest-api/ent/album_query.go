// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/depromeet/everybody-backend/rest-api/ent/album"
	"github.com/depromeet/everybody-backend/rest-api/ent/predicate"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

// AlbumQuery is the builder for querying Album entities.
type AlbumQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Album
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AlbumQuery builder.
func (aq *AlbumQuery) Where(ps ...predicate.Album) *AlbumQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AlbumQuery) Limit(limit int) *AlbumQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AlbumQuery) Offset(offset int) *AlbumQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AlbumQuery) Unique(unique bool) *AlbumQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AlbumQuery) Order(o ...OrderFunc) *AlbumQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUser chains the current query on the "user" edge.
func (aq *AlbumQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, album.UserTable, album.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Album entity from the query.
// Returns a *NotFoundError when no Album was found.
func (aq *AlbumQuery) First(ctx context.Context) (*Album, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{album.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AlbumQuery) FirstX(ctx context.Context) *Album {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Album ID from the query.
// Returns a *NotFoundError when no Album ID was found.
func (aq *AlbumQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{album.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AlbumQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Album entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Album entity is not found.
// Returns a *NotFoundError when no Album entities are found.
func (aq *AlbumQuery) Only(ctx context.Context) (*Album, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{album.Label}
	default:
		return nil, &NotSingularError{album.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AlbumQuery) OnlyX(ctx context.Context) *Album {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Album ID in the query.
// Returns a *NotSingularError when exactly one Album ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aq *AlbumQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = &NotSingularError{album.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AlbumQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Albums.
func (aq *AlbumQuery) All(ctx context.Context) ([]*Album, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AlbumQuery) AllX(ctx context.Context) []*Album {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Album IDs.
func (aq *AlbumQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(album.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AlbumQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AlbumQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AlbumQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AlbumQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AlbumQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AlbumQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AlbumQuery) Clone() *AlbumQuery {
	if aq == nil {
		return nil
	}
	return &AlbumQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]OrderFunc{}, aq.order...),
		predicates: append([]predicate.Album{}, aq.predicates...),
		withUser:   aq.withUser.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AlbumQuery) WithUser(opts ...func(*UserQuery)) *AlbumQuery {
	query := &UserQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FolderName string `json:"folder_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Album.Query().
//		GroupBy(album.FieldFolderName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AlbumQuery) GroupBy(field string, fields ...string) *AlbumGroupBy {
	group := &AlbumGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FolderName string `json:"folder_name,omitempty"`
//	}
//
//	client.Album.Query().
//		Select(album.FieldFolderName).
//		Scan(ctx, &v)
//
func (aq *AlbumQuery) Select(fields ...string) *AlbumSelect {
	aq.fields = append(aq.fields, fields...)
	return &AlbumSelect{AlbumQuery: aq}
}

func (aq *AlbumQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !album.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AlbumQuery) sqlAll(ctx context.Context) ([]*Album, error) {
	var (
		nodes       = []*Album{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withUser != nil,
		}
	)
	if aq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, album.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Album{config: aq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withUser; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Album)
		for i := range nodes {
			if nodes[i].user_album == nil {
				continue
			}
			fk := *nodes[i].user_album
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_album" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (aq *AlbumQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AlbumQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *AlbumQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   album.Table,
			Columns: album.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: album.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, album.FieldID)
		for i := range fields {
			if fields[i] != album.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AlbumQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(album.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = album.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AlbumGroupBy is the group-by builder for Album entities.
type AlbumGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AlbumGroupBy) Aggregate(fns ...AggregateFunc) *AlbumGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AlbumGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *AlbumGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AlbumGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AlbumGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *AlbumGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AlbumGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = agb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = fmt.Errorf("ent: AlbumGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (agb *AlbumGroupBy) StringX(ctx context.Context) string {
	v, err := agb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AlbumGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AlbumGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *AlbumGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AlbumGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = agb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = fmt.Errorf("ent: AlbumGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (agb *AlbumGroupBy) IntX(ctx context.Context) int {
	v, err := agb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AlbumGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AlbumGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *AlbumGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AlbumGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = agb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = fmt.Errorf("ent: AlbumGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (agb *AlbumGroupBy) Float64X(ctx context.Context) float64 {
	v, err := agb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *AlbumGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AlbumGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *AlbumGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *AlbumGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = agb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = fmt.Errorf("ent: AlbumGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (agb *AlbumGroupBy) BoolX(ctx context.Context) bool {
	v, err := agb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *AlbumGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !album.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AlbumGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AlbumSelect is the builder for selecting fields of Album entities.
type AlbumSelect struct {
	*AlbumQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AlbumSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AlbumQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *AlbumSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (as *AlbumSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AlbumSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *AlbumSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (as *AlbumSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = as.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = fmt.Errorf("ent: AlbumSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (as *AlbumSelect) StringX(ctx context.Context) string {
	v, err := as.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (as *AlbumSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AlbumSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *AlbumSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (as *AlbumSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = as.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = fmt.Errorf("ent: AlbumSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (as *AlbumSelect) IntX(ctx context.Context) int {
	v, err := as.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (as *AlbumSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AlbumSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *AlbumSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (as *AlbumSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = as.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = fmt.Errorf("ent: AlbumSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (as *AlbumSelect) Float64X(ctx context.Context) float64 {
	v, err := as.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (as *AlbumSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AlbumSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *AlbumSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (as *AlbumSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = as.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{album.Label}
	default:
		err = fmt.Errorf("ent: AlbumSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (as *AlbumSelect) BoolX(ctx context.Context) bool {
	v, err := as.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *AlbumSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
