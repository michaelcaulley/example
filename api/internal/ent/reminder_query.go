// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"example/internal/ent/predicate"
	"example/internal/ent/reminder"
	"example/internal/ent/todo"
	"example/internal/ent/todoreminder"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReminderQuery is the builder for querying Reminder entities.
type ReminderQuery struct {
	config
	ctx                    *QueryContext
	order                  []reminder.OrderOption
	inters                 []Interceptor
	predicates             []predicate.Reminder
	withTodo               *TodoQuery
	withTodoReminders      *TodoReminderQuery
	loadTotal              []func(context.Context, []*Reminder) error
	modifiers              []func(*sql.Selector)
	withNamedTodo          map[string]*TodoQuery
	withNamedTodoReminders map[string]*TodoReminderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReminderQuery builder.
func (_q *ReminderQuery) Where(ps ...predicate.Reminder) *ReminderQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *ReminderQuery) Limit(limit int) *ReminderQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *ReminderQuery) Offset(offset int) *ReminderQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *ReminderQuery) Unique(unique bool) *ReminderQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *ReminderQuery) Order(o ...reminder.OrderOption) *ReminderQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryTodo chains the current query on the "todo" edge.
func (_q *ReminderQuery) QueryTodo() *TodoQuery {
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
			sqlgraph.From(reminder.Table, reminder.FieldID, selector),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, reminder.TodoTable, reminder.TodoPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTodoReminders chains the current query on the "todo_reminders" edge.
func (_q *ReminderQuery) QueryTodoReminders() *TodoReminderQuery {
	query := (&TodoReminderClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reminder.Table, reminder.FieldID, selector),
			sqlgraph.To(todoreminder.Table, todoreminder.ReminderColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, reminder.TodoRemindersTable, reminder.TodoRemindersColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Reminder entity from the query.
// Returns a *NotFoundError when no Reminder was found.
func (_q *ReminderQuery) First(ctx context.Context) (*Reminder, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reminder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *ReminderQuery) FirstX(ctx context.Context) *Reminder {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Reminder ID from the query.
// Returns a *NotFoundError when no Reminder ID was found.
func (_q *ReminderQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reminder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *ReminderQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Reminder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Reminder entity is found.
// Returns a *NotFoundError when no Reminder entities are found.
func (_q *ReminderQuery) Only(ctx context.Context) (*Reminder, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reminder.Label}
	default:
		return nil, &NotSingularError{reminder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *ReminderQuery) OnlyX(ctx context.Context) *Reminder {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Reminder ID in the query.
// Returns a *NotSingularError when more than one Reminder ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *ReminderQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reminder.Label}
	default:
		err = &NotSingularError{reminder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *ReminderQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Reminders.
func (_q *ReminderQuery) All(ctx context.Context) ([]*Reminder, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Reminder, *ReminderQuery]()
	return withInterceptors[[]*Reminder](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *ReminderQuery) AllX(ctx context.Context) []*Reminder {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Reminder IDs.
func (_q *ReminderQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(reminder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *ReminderQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *ReminderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*ReminderQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *ReminderQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *ReminderQuery) Exist(ctx context.Context) (bool, error) {
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
func (_q *ReminderQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReminderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *ReminderQuery) Clone() *ReminderQuery {
	if _q == nil {
		return nil
	}
	return &ReminderQuery{
		config:            _q.config,
		ctx:               _q.ctx.Clone(),
		order:             append([]reminder.OrderOption{}, _q.order...),
		inters:            append([]Interceptor{}, _q.inters...),
		predicates:        append([]predicate.Reminder{}, _q.predicates...),
		withTodo:          _q.withTodo.Clone(),
		withTodoReminders: _q.withTodoReminders.Clone(),
		// clone intermediate query.
		sql:       _q.sql.Clone(),
		path:      _q.path,
		modifiers: append([]func(*sql.Selector){}, _q.modifiers...),
	}
}

// WithTodo tells the query-builder to eager-load the nodes that are connected to
// the "todo" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *ReminderQuery) WithTodo(opts ...func(*TodoQuery)) *ReminderQuery {
	query := (&TodoClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withTodo = query
	return _q
}

// WithTodoReminders tells the query-builder to eager-load the nodes that are connected to
// the "todo_reminders" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *ReminderQuery) WithTodoReminders(opts ...func(*TodoReminderQuery)) *ReminderQuery {
	query := (&TodoReminderClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withTodoReminders = query
	return _q
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
//	client.Reminder.Query().
//		GroupBy(reminder.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *ReminderQuery) GroupBy(field string, fields ...string) *ReminderGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReminderGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = reminder.Label
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
//	client.Reminder.Query().
//		Select(reminder.FieldCreatedAt).
//		Scan(ctx, &v)
func (_q *ReminderQuery) Select(fields ...string) *ReminderSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &ReminderSelect{ReminderQuery: _q}
	sbuild.label = reminder.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReminderSelect configured with the given aggregations.
func (_q *ReminderQuery) Aggregate(fns ...AggregateFunc) *ReminderSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *ReminderQuery) prepareQuery(ctx context.Context) error {
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
		if !reminder.ValidColumn(f) {
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

func (_q *ReminderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Reminder, error) {
	var (
		nodes       = []*Reminder{}
		_spec       = _q.querySpec()
		loadedTypes = [2]bool{
			_q.withTodo != nil,
			_q.withTodoReminders != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Reminder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Reminder{config: _q.config}
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
	if query := _q.withTodo; query != nil {
		if err := _q.loadTodo(ctx, query, nodes,
			func(n *Reminder) { n.Edges.Todo = []*Todo{} },
			func(n *Reminder, e *Todo) { n.Edges.Todo = append(n.Edges.Todo, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withTodoReminders; query != nil {
		if err := _q.loadTodoReminders(ctx, query, nodes,
			func(n *Reminder) { n.Edges.TodoReminders = []*TodoReminder{} },
			func(n *Reminder, e *TodoReminder) { n.Edges.TodoReminders = append(n.Edges.TodoReminders, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range _q.withNamedTodo {
		if err := _q.loadTodo(ctx, query, nodes,
			func(n *Reminder) { n.appendNamedTodo(name) },
			func(n *Reminder, e *Todo) { n.appendNamedTodo(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range _q.withNamedTodoReminders {
		if err := _q.loadTodoReminders(ctx, query, nodes,
			func(n *Reminder) { n.appendNamedTodoReminders(name) },
			func(n *Reminder, e *TodoReminder) { n.appendNamedTodoReminders(name, e) }); err != nil {
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

func (_q *ReminderQuery) loadTodo(ctx context.Context, query *TodoQuery, nodes []*Reminder, init func(*Reminder), assign func(*Reminder, *Todo)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Reminder)
	nids := make(map[int]map[*Reminder]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(reminder.TodoTable)
		s.Join(joinT).On(s.C(todo.FieldID), joinT.C(reminder.TodoPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(reminder.TodoPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(reminder.TodoPrimaryKey[1]))
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
					nids[inValue] = map[*Reminder]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Todo](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "todo" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (_q *ReminderQuery) loadTodoReminders(ctx context.Context, query *TodoReminderQuery, nodes []*Reminder, init func(*Reminder), assign func(*Reminder, *TodoReminder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Reminder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(todoreminder.FieldReminderID)
	}
	query.Where(predicate.TodoReminder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(reminder.TodoRemindersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReminderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "reminder_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (_q *ReminderQuery) sqlCount(ctx context.Context) (int, error) {
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

func (_q *ReminderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(reminder.Table, reminder.Columns, sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reminder.FieldID)
		for i := range fields {
			if fields[i] != reminder.FieldID {
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

func (_q *ReminderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(reminder.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = reminder.Columns
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
func (_q *ReminderQuery) Modify(modifiers ...func(s *sql.Selector)) *ReminderSelect {
	_q.modifiers = append(_q.modifiers, modifiers...)
	return _q.Select()
}

// WithNamedTodo tells the query-builder to eager-load the nodes that are connected to the "todo"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (_q *ReminderQuery) WithNamedTodo(name string, opts ...func(*TodoQuery)) *ReminderQuery {
	query := (&TodoClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if _q.withNamedTodo == nil {
		_q.withNamedTodo = make(map[string]*TodoQuery)
	}
	_q.withNamedTodo[name] = query
	return _q
}

// WithNamedTodoReminders tells the query-builder to eager-load the nodes that are connected to the "todo_reminders"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (_q *ReminderQuery) WithNamedTodoReminders(name string, opts ...func(*TodoReminderQuery)) *ReminderQuery {
	query := (&TodoReminderClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if _q.withNamedTodoReminders == nil {
		_q.withNamedTodoReminders = make(map[string]*TodoReminderQuery)
	}
	_q.withNamedTodoReminders[name] = query
	return _q
}

// ReminderGroupBy is the group-by builder for Reminder entities.
type ReminderGroupBy struct {
	selector
	build *ReminderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReminderGroupBy) Aggregate(fns ...AggregateFunc) *ReminderGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReminderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, ent.OpQueryGroupBy)
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReminderQuery, *ReminderGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReminderGroupBy) sqlScan(ctx context.Context, root *ReminderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReminderSelect is the builder for selecting fields of Reminder entities.
type ReminderSelect struct {
	*ReminderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReminderSelect) Aggregate(fns ...AggregateFunc) *ReminderSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReminderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, ent.OpQuerySelect)
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReminderQuery, *ReminderSelect](ctx, rs.ReminderQuery, rs, rs.inters, v)
}

func (rs *ReminderSelect) sqlScan(ctx context.Context, root *ReminderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rs *ReminderSelect) Modify(modifiers ...func(s *sql.Selector)) *ReminderSelect {
	rs.modifiers = append(rs.modifiers, modifiers...)
	return rs
}
