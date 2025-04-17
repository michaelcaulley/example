// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"example/internal/ent/reminder"
	"example/internal/ent/todo"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReminderCreate is the builder for creating a Reminder entity.
type ReminderCreate struct {
	config
	mutation *ReminderMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (_c *ReminderCreate) SetCreatedAt(v time.Time) *ReminderCreate {
	_c.mutation.SetCreatedAt(v)
	return _c
}

// SetUpdatedAt sets the "updated_at" field.
func (_c *ReminderCreate) SetUpdatedAt(v time.Time) *ReminderCreate {
	_c.mutation.SetUpdatedAt(v)
	return _c
}

// AddTodoIDs adds the "todo" edge to the Todo entity by IDs.
func (_c *ReminderCreate) AddTodoIDs(ids ...int) *ReminderCreate {
	_c.mutation.AddTodoIDs(ids...)
	return _c
}

// AddTodo adds the "todo" edges to the Todo entity.
func (_c *ReminderCreate) AddTodo(v ...*Todo) *ReminderCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddTodoIDs(ids...)
}

// Mutation returns the ReminderMutation object of the builder.
func (_c *ReminderCreate) Mutation() *ReminderMutation {
	return _c.mutation
}

// Save creates the Reminder in the database.
func (_c *ReminderCreate) Save(ctx context.Context) (*Reminder, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *ReminderCreate) SaveX(ctx context.Context) *Reminder {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *ReminderCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *ReminderCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *ReminderCreate) check() error {
	if _, ok := _c.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Reminder.created_at"`)}
	}
	return nil
}

func (_c *ReminderCreate) sqlSave(ctx context.Context) (*Reminder, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *ReminderCreate) createSpec() (*Reminder, *sqlgraph.CreateSpec) {
	var (
		_node = &Reminder{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(reminder.Table, sqlgraph.NewFieldSpec(reminder.FieldID, field.TypeInt))
	)
	_spec.Schema = _c.schemaConfig.Reminder
	_spec.OnConflict = _c.conflict
	if value, ok := _c.mutation.CreatedAt(); ok {
		_spec.SetField(reminder.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := _c.mutation.UpdatedAt(); ok {
		_spec.SetField(reminder.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := _c.mutation.TodoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   reminder.TodoTable,
			Columns: reminder.TodoPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _c.schemaConfig.TodoReminder
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Reminder.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReminderUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (_c *ReminderCreate) OnConflict(opts ...sql.ConflictOption) *ReminderUpsertOne {
	_c.conflict = opts
	return &ReminderUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Reminder.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *ReminderCreate) OnConflictColumns(columns ...string) *ReminderUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &ReminderUpsertOne{
		create: _c,
	}
}

type (
	// ReminderUpsertOne is the builder for "upsert"-ing
	//  one Reminder node.
	ReminderUpsertOne struct {
		create *ReminderCreate
	}

	// ReminderUpsert is the "OnConflict" setter.
	ReminderUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ReminderUpsert) SetUpdatedAt(v time.Time) *ReminderUpsert {
	u.Set(reminder.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ReminderUpsert) UpdateUpdatedAt() *ReminderUpsert {
	u.SetExcluded(reminder.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Reminder.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ReminderUpsertOne) UpdateNewValues() *ReminderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(reminder.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Reminder.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ReminderUpsertOne) Ignore() *ReminderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReminderUpsertOne) DoNothing() *ReminderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReminderCreate.OnConflict
// documentation for more info.
func (u *ReminderUpsertOne) Update(set func(*ReminderUpsert)) *ReminderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReminderUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ReminderUpsertOne) SetUpdatedAt(v time.Time) *ReminderUpsertOne {
	return u.Update(func(s *ReminderUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ReminderUpsertOne) UpdateUpdatedAt() *ReminderUpsertOne {
	return u.Update(func(s *ReminderUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ReminderUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReminderCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReminderUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ReminderUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ReminderUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ReminderCreateBulk is the builder for creating many Reminder entities in bulk.
type ReminderCreateBulk struct {
	config
	err      error
	builders []*ReminderCreate
	conflict []sql.ConflictOption
}

// Save creates the Reminder entities in the database.
func (_c *ReminderCreateBulk) Save(ctx context.Context) ([]*Reminder, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Reminder, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReminderMutation)
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
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = _c.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *ReminderCreateBulk) SaveX(ctx context.Context) []*Reminder {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *ReminderCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *ReminderCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Reminder.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReminderUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (_c *ReminderCreateBulk) OnConflict(opts ...sql.ConflictOption) *ReminderUpsertBulk {
	_c.conflict = opts
	return &ReminderUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Reminder.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *ReminderCreateBulk) OnConflictColumns(columns ...string) *ReminderUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &ReminderUpsertBulk{
		create: _c,
	}
}

// ReminderUpsertBulk is the builder for "upsert"-ing
// a bulk of Reminder nodes.
type ReminderUpsertBulk struct {
	create *ReminderCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Reminder.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ReminderUpsertBulk) UpdateNewValues() *ReminderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(reminder.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Reminder.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ReminderUpsertBulk) Ignore() *ReminderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReminderUpsertBulk) DoNothing() *ReminderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReminderCreateBulk.OnConflict
// documentation for more info.
func (u *ReminderUpsertBulk) Update(set func(*ReminderUpsert)) *ReminderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReminderUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ReminderUpsertBulk) SetUpdatedAt(v time.Time) *ReminderUpsertBulk {
	return u.Update(func(s *ReminderUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ReminderUpsertBulk) UpdateUpdatedAt() *ReminderUpsertBulk {
	return u.Update(func(s *ReminderUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ReminderUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ReminderCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReminderCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReminderUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
