// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"example/internal/ent/moderator"
	"example/internal/ent/predicate"
	"example/internal/ent/todo"
	"example/internal/ent/user"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx                     *QueryContext
	order                   []user.OrderOption
	inters                  []Interceptor
	predicates              []predicate.User
	withTodos               *TodoQuery
	withModeratorUsers      *UserQuery
	withModerators          *UserQuery
	withModerator           *ModeratorQuery
	loadTotal               []func(context.Context, []*User) error
	modifiers               []func(*sql.Selector)
	withNamedTodos          map[string]*TodoQuery
	withNamedModeratorUsers map[string]*UserQuery
	withNamedModerators     map[string]*UserQuery
	withNamedModerator      map[string]*ModeratorQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (_q *UserQuery) Where(ps ...predicate.User) *UserQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *UserQuery) Limit(limit int) *UserQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *UserQuery) Offset(offset int) *UserQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *UserQuery) Unique(unique bool) *UserQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryTodos chains the current query on the "todos" edge.
func (_q *UserQuery) QueryTodos() *TodoQuery {
	query := (&TodoClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TodosTable, user.TodosColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryModeratorUsers chains the current query on the "moderator_users" edge.
func (_q *UserQuery) QueryModeratorUsers() *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.ModeratorUsersTable, user.ModeratorUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryModerators chains the current query on the "moderators" edge.
func (_q *UserQuery) QueryModerators() *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ModeratorsTable, user.ModeratorsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryModerator chains the current query on the "moderator" edge.
func (_q *UserQuery) QueryModerator() *ModeratorQuery {
	query := (&ModeratorClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(moderator.Table, moderator.UserColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, user.ModeratorTable, user.ModeratorColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (_q *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *UserQuery) FirstX(ctx context.Context) *User {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (_q *UserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *UserQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (_q *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *UserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *UserQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (_q *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (_q *UserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *UserQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*UserQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *UserQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *UserQuery) Clone() *UserQuery {
	if _q == nil {
		return nil
	}
	return &UserQuery{
		config:             _q.config,
		ctx:                _q.ctx.Clone(),
		order:              append([]user.OrderOption{}, _q.order...),
		inters:             append([]Interceptor{}, _q.inters...),
		predicates:         append([]predicate.User{}, _q.predicates...),
		withTodos:          _q.withTodos.Clone(),
		withModeratorUsers: _q.withModeratorUsers.Clone(),
		withModerators:     _q.withModerators.Clone(),
		withModerator:      _q.withModerator.Clone(),
		// clone intermediate query.
		sql:       _q.sql.Clone(),
		path:      _q.path,
		modifiers: append([]func(*sql.Selector){}, _q.modifiers...),
	}
}

// WithTodos tells the query-builder to eager-load the nodes that are connected to
// the "todos" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *UserQuery) WithTodos(opts ...func(*TodoQuery)) *UserQuery {
	query := (&TodoClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withTodos = query
	return _q
}

// WithModeratorUsers tells the query-builder to eager-load the nodes that are connected to
// the "moderator_users" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *UserQuery) WithModeratorUsers(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withModeratorUsers = query
	return _q
}

// WithModerators tells the query-builder to eager-load the nodes that are connected to
// the "moderators" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *UserQuery) WithModerators(opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withModerators = query
	return _q
}

// WithModerator tells the query-builder to eager-load the nodes that are connected to
// the "moderator" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *UserQuery) WithModerator(opts ...func(*ModeratorQuery)) *UserQuery {
	query := (&ModeratorClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withModerator = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldName).
//		Scan(ctx, &v)
func (_q *UserQuery) Select(fields ...string) *UserSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: _q}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (_q *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	for _, f := range _q.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.sql = prev
	}
	return nil
}

func (_q *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		_spec       = _q.querySpec()
		loadedTypes = [4]bool{
			_q.withTodos != nil,
			_q.withModeratorUsers != nil,
			_q.withModerators != nil,
			_q.withModerator != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: _q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, _q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := _q.withTodos; query != nil {
		if err := _q.loadTodos(ctx, query, nodes,
			func(n *User) { n.Edges.Todos = []*Todo{} },
			func(n *User, e *Todo) { n.Edges.Todos = append(n.Edges.Todos, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withModeratorUsers; query != nil {
		if err := _q.loadModeratorUsers(ctx, query, nodes,
			func(n *User) { n.Edges.ModeratorUsers = []*User{} },
			func(n *User, e *User) { n.Edges.ModeratorUsers = append(n.Edges.ModeratorUsers, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withModerators; query != nil {
		if err := _q.loadModerators(ctx, query, nodes,
			func(n *User) { n.Edges.Moderators = []*User{} },
			func(n *User, e *User) { n.Edges.Moderators = append(n.Edges.Moderators, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withModerator; query != nil {
		if err := _q.loadModerator(ctx, query, nodes,
			func(n *User) { n.Edges.Moderator = []*Moderator{} },
			func(n *User, e *Moderator) { n.Edges.Moderator = append(n.Edges.Moderator, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range _q.withNamedTodos {
		if err := _q.loadTodos(ctx, query, nodes,
			func(n *User) { n.appendNamedTodos(name) },
			func(n *User, e *Todo) { n.appendNamedTodos(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range _q.withNamedModeratorUsers {
		if err := _q.loadModeratorUsers(ctx, query, nodes,
			func(n *User) { n.appendNamedModeratorUsers(name) },
			func(n *User, e *User) { n.appendNamedModeratorUsers(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range _q.withNamedModerators {
		if err := _q.loadModerators(ctx, query, nodes,
			func(n *User) { n.appendNamedModerators(name) },
			func(n *User, e *User) { n.appendNamedModerators(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range _q.withNamedModerator {
		if err := _q.loadModerator(ctx, query, nodes,
			func(n *User) { n.appendNamedModerator(name) },
			func(n *User, e *Moderator) { n.appendNamedModerator(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range _q.loadTotal {
		if err := _q.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *UserQuery) loadTodos(ctx context.Context, query *TodoQuery, nodes []*User, init func(*User), assign func(*User, *Todo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(todo.FieldOwnerID)
	}
	query.Where(predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.TodosColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OwnerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "owner_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (_q *UserQuery) loadModeratorUsers(ctx context.Context, query *UserQuery, nodes []*User, init func(*User), assign func(*User, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*User)
	nids := make(map[int]map[*User]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(user.ModeratorUsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(user.ModeratorUsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(user.ModeratorUsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(user.ModeratorUsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*User]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "moderator_users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (_q *UserQuery) loadModerators(ctx context.Context, query *UserQuery, nodes []*User, init func(*User), assign func(*User, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*User)
	nids := make(map[int]map[*User]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(user.ModeratorsTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(user.ModeratorsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(user.ModeratorsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(user.ModeratorsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*User]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "moderators" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (_q *UserQuery) loadModerator(ctx context.Context, query *ModeratorQuery, nodes []*User, init func(*User), assign func(*User, *Moderator)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(moderator.FieldUserID)
	}
	query.Where(predicate.Moderator(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.ModeratorColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (_q *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := _q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := _q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := _q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := _q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (_q *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range _q.modifiers {
		m(selector)
	}
	for _, p := range _q.predicates {
		p(selector)
	}
	for _, p := range _q.order {
		p(selector)
	}
	if offset := _q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := _q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (_q *UserQuery) Modify(modifiers ...func(s *sql.Selector)) *UserSelect {
	_q.modifiers = append(_q.modifiers, modifiers...)
	return _q.Select()
}

// WithNamedTodos tells the query-builder to eager-load the nodes that are connected to the "todos"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (_q *UserQuery) WithNamedTodos(name string, opts ...func(*TodoQuery)) *UserQuery {
	query := (&TodoClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if _q.withNamedTodos == nil {
		_q.withNamedTodos = make(map[string]*TodoQuery)
	}
	_q.withNamedTodos[name] = query
	return _q
}

// WithNamedModeratorUsers tells the query-builder to eager-load the nodes that are connected to the "moderator_users"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (_q *UserQuery) WithNamedModeratorUsers(name string, opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if _q.withNamedModeratorUsers == nil {
		_q.withNamedModeratorUsers = make(map[string]*UserQuery)
	}
	_q.withNamedModeratorUsers[name] = query
	return _q
}

// WithNamedModerators tells the query-builder to eager-load the nodes that are connected to the "moderators"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (_q *UserQuery) WithNamedModerators(name string, opts ...func(*UserQuery)) *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if _q.withNamedModerators == nil {
		_q.withNamedModerators = make(map[string]*UserQuery)
	}
	_q.withNamedModerators[name] = query
	return _q
}

// WithNamedModerator tells the query-builder to eager-load the nodes that are connected to the "moderator"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (_q *UserQuery) WithNamedModerator(name string, opts ...func(*ModeratorQuery)) *UserQuery {
	query := (&ModeratorClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if _q.withNamedModerator == nil {
		_q.withNamedModerator = make(map[string]*ModeratorQuery)
	}
	_q.withNamedModerator[name] = query
	return _q
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, ent.OpQueryGroupBy)
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, ent.OpQuerySelect)
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (us *UserSelect) Modify(modifiers ...func(s *sql.Selector)) *UserSelect {
	us.modifiers = append(us.modifiers, modifiers...)
	return us
}
