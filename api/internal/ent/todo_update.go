// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"example/internal/ent/predicate"
	"example/internal/ent/reminder"
	"example/internal/ent/todo"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"example/internal/ent/internal"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks     []Hook
	mutation  *TodoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TodoUpdate builder.
func (_u *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetText sets the "text" field.
func (_u *TodoUpdate) SetText(v string) *TodoUpdate {
	_u.mutation.SetText(v)
	return _u
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_u *TodoUpdate) SetNillableText(v *string) *TodoUpdate {
	if v != nil {
		_u.SetText(*v)
	}
	return _u
}

// SetDoneAt sets the "done_at" field.
func (_u *TodoUpdate) SetDoneAt(v time.Time) *TodoUpdate {
	_u.mutation.SetDoneAt(v)
	return _u
}

// SetNillableDoneAt sets the "done_at" field if the given value is not nil.
func (_u *TodoUpdate) SetNillableDoneAt(v *time.Time) *TodoUpdate {
	if v != nil {
		_u.SetDoneAt(*v)
	}
	return _u
}

// ClearDoneAt clears the value of the "done_at" field.
func (_u *TodoUpdate) ClearDoneAt() *TodoUpdate {
	_u.mutation.ClearDoneAt()
	return _u
}

// AddReminderIDs adds the "reminders" edge to the Reminder entity by IDs.
func (_u *TodoUpdate) AddReminderIDs(ids ...int) *TodoUpdate {
	_u.mutation.AddReminderIDs(ids...)
	return _u
}

// AddReminders adds the "reminders" edges to the Reminder entity.
func (_u *TodoUpdate) AddReminders(v ...*Reminder) *TodoUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddReminderIDs(ids...)
}

// Mutation returns the TodoMutation object of the builder.
func (_u *TodoUpdate) Mutation() *TodoMutation {
	return _u.mutation
}

// ClearReminders clears all "reminders" edges to the Reminder entity.
func (_u *TodoUpdate) ClearReminders() *TodoUpdate {
	_u.mutation.ClearReminders()
	return _u
}

// RemoveReminderIDs removes the "reminders" edge to Reminder entities by IDs.
func (_u *TodoUpdate) RemoveReminderIDs(ids ...int) *TodoUpdate {
	_u.mutation.RemoveReminderIDs(ids...)
	return _u
}

// RemoveReminders removes "reminders" edges to Reminder entities.
func (_u *TodoUpdate) RemoveReminders(v ...*Reminder) *TodoUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveReminderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *TodoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *TodoUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *TodoUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *TodoUpdate) check() error {
	if v, ok := _u.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	if _u.mutation.OwnerCleared() && len(_u.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Todo.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *TodoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TodoUpdate {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *TodoUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Text(); ok {
		_spec.SetField(todo.FieldText, field.TypeString, value)
	}
	if value, ok := _u.mutation.DoneAt(); ok {
		_spec.SetField(todo.FieldDoneAt, field.TypeTime, value)
	}
	if _u.mutation.DoneAtCleared() {
		_spec.ClearField(todo.FieldDoneAt, field.TypeTime)
	}
	if _u.mutation.RemindersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.RemindersTable,
			Columns: todo.RemindersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.TodoReminder
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedRemindersIDs(); len(nodes) > 0 && !_u.mutation.RemindersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.RemindersTable,
			Columns: todo.RemindersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.TodoReminder
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemindersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.RemindersTable,
			Columns: todo.RemindersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.TodoReminder
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = _u.schemaConfig.Todo
	ctx = internal.NewSchemaConfigContext(ctx, _u.schemaConfig)
	_spec.AddModifiers(_u.modifiers...)
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TodoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetText sets the "text" field.
func (_u *TodoUpdateOne) SetText(v string) *TodoUpdateOne {
	_u.mutation.SetText(v)
	return _u
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_u *TodoUpdateOne) SetNillableText(v *string) *TodoUpdateOne {
	if v != nil {
		_u.SetText(*v)
	}
	return _u
}

// SetDoneAt sets the "done_at" field.
func (_u *TodoUpdateOne) SetDoneAt(v time.Time) *TodoUpdateOne {
	_u.mutation.SetDoneAt(v)
	return _u
}

// SetNillableDoneAt sets the "done_at" field if the given value is not nil.
func (_u *TodoUpdateOne) SetNillableDoneAt(v *time.Time) *TodoUpdateOne {
	if v != nil {
		_u.SetDoneAt(*v)
	}
	return _u
}

// ClearDoneAt clears the value of the "done_at" field.
func (_u *TodoUpdateOne) ClearDoneAt() *TodoUpdateOne {
	_u.mutation.ClearDoneAt()
	return _u
}

// AddReminderIDs adds the "reminders" edge to the Reminder entity by IDs.
func (_u *TodoUpdateOne) AddReminderIDs(ids ...int) *TodoUpdateOne {
	_u.mutation.AddReminderIDs(ids...)
	return _u
}

// AddReminders adds the "reminders" edges to the Reminder entity.
func (_u *TodoUpdateOne) AddReminders(v ...*Reminder) *TodoUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddReminderIDs(ids...)
}

// Mutation returns the TodoMutation object of the builder.
func (_u *TodoUpdateOne) Mutation() *TodoMutation {
	return _u.mutation
}

// ClearReminders clears all "reminders" edges to the Reminder entity.
func (_u *TodoUpdateOne) ClearReminders() *TodoUpdateOne {
	_u.mutation.ClearReminders()
	return _u
}

// RemoveReminderIDs removes the "reminders" edge to Reminder entities by IDs.
func (_u *TodoUpdateOne) RemoveReminderIDs(ids ...int) *TodoUpdateOne {
	_u.mutation.RemoveReminderIDs(ids...)
	return _u
}

// RemoveReminders removes "reminders" edges to Reminder entities.
func (_u *TodoUpdateOne) RemoveReminders(v ...*Reminder) *TodoUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveReminderIDs(ids...)
}

// Where appends a list predicates to the TodoUpdate builder.
func (_u *TodoUpdateOne) Where(ps ...predicate.Todo) *TodoUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Todo entity.
func (_u *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *TodoUpdateOne) check() error {
	if v, ok := _u.mutation.Text(); ok {
		if err := todo.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Todo.text": %w`, err)}
		}
	}
	if _u.mutation.OwnerCleared() && len(_u.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Todo.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *TodoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TodoUpdateOne {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Text(); ok {
		_spec.SetField(todo.FieldText, field.TypeString, value)
	}
	if value, ok := _u.mutation.DoneAt(); ok {
		_spec.SetField(todo.FieldDoneAt, field.TypeTime, value)
	}
	if _u.mutation.DoneAtCleared() {
		_spec.ClearField(todo.FieldDoneAt, field.TypeTime)
	}
	if _u.mutation.RemindersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.RemindersTable,
			Columns: todo.RemindersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.TodoReminder
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedRemindersIDs(); len(nodes) > 0 && !_u.mutation.RemindersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.RemindersTable,
			Columns: todo.RemindersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.TodoReminder
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemindersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   todo.RemindersTable,
			Columns: todo.RemindersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.TodoReminder
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = _u.schemaConfig.Todo
	ctx = internal.NewSchemaConfigContext(ctx, _u.schemaConfig)
	_spec.AddModifiers(_u.modifiers...)
	_node = &Todo{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
