// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"example/internal/ent/peoplepartner"
	"example/internal/ent/predicate"
	"example/internal/ent/user"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"example/internal/ent/internal"
)

// PeoplePartnerUpdate is the builder for updating PeoplePartner entities.
type PeoplePartnerUpdate struct {
	config
	hooks     []Hook
	mutation  *PeoplePartnerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PeoplePartnerUpdate builder.
func (_u *PeoplePartnerUpdate) Where(ps ...predicate.PeoplePartner) *PeoplePartnerUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetUserID sets the "user_id" field.
func (_u *PeoplePartnerUpdate) SetUserID(v int) *PeoplePartnerUpdate {
	_u.mutation.SetUserID(v)
	return _u
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (_u *PeoplePartnerUpdate) SetNillableUserID(v *int) *PeoplePartnerUpdate {
	if v != nil {
		_u.SetUserID(*v)
	}
	return _u
}

// SetPeoplePartnerUserID sets the "people_partner_user_id" field.
func (_u *PeoplePartnerUpdate) SetPeoplePartnerUserID(v int) *PeoplePartnerUpdate {
	_u.mutation.SetPeoplePartnerUserID(v)
	return _u
}

// SetNillablePeoplePartnerUserID sets the "people_partner_user_id" field if the given value is not nil.
func (_u *PeoplePartnerUpdate) SetNillablePeoplePartnerUserID(v *int) *PeoplePartnerUpdate {
	if v != nil {
		_u.SetPeoplePartnerUserID(*v)
	}
	return _u
}

// SetUser sets the "user" edge to the User entity.
func (_u *PeoplePartnerUpdate) SetUser(v *User) *PeoplePartnerUpdate {
	return _u.SetUserID(v.ID)
}

// SetPeoplePartnerID sets the "people_partner" edge to the User entity by ID.
func (_u *PeoplePartnerUpdate) SetPeoplePartnerID(id int) *PeoplePartnerUpdate {
	_u.mutation.SetPeoplePartnerID(id)
	return _u
}

// SetPeoplePartner sets the "people_partner" edge to the User entity.
func (_u *PeoplePartnerUpdate) SetPeoplePartner(v *User) *PeoplePartnerUpdate {
	return _u.SetPeoplePartnerID(v.ID)
}

// Mutation returns the PeoplePartnerMutation object of the builder.
func (_u *PeoplePartnerUpdate) Mutation() *PeoplePartnerMutation {
	return _u.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (_u *PeoplePartnerUpdate) ClearUser() *PeoplePartnerUpdate {
	_u.mutation.ClearUser()
	return _u
}

// ClearPeoplePartner clears the "people_partner" edge to the User entity.
func (_u *PeoplePartnerUpdate) ClearPeoplePartner() *PeoplePartnerUpdate {
	_u.mutation.ClearPeoplePartner()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *PeoplePartnerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *PeoplePartnerUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *PeoplePartnerUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *PeoplePartnerUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *PeoplePartnerUpdate) check() error {
	if _u.mutation.UserCleared() && len(_u.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PeoplePartner.user"`)
	}
	if _u.mutation.PeoplePartnerCleared() && len(_u.mutation.PeoplePartnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PeoplePartner.people_partner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *PeoplePartnerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PeoplePartnerUpdate {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *PeoplePartnerUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(peoplepartner.Table, peoplepartner.Columns, sqlgraph.NewFieldSpec(peoplepartner.FieldUserID, field.TypeInt), sqlgraph.NewFieldSpec(peoplepartner.FieldPeoplePartnerUserID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if _u.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   peoplepartner.UserTable,
			Columns: []string{peoplepartner.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.PeoplePartner
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   peoplepartner.UserTable,
			Columns: []string{peoplepartner.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.PeoplePartner
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.PeoplePartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   peoplepartner.PeoplePartnerTable,
			Columns: []string{peoplepartner.PeoplePartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.PeoplePartner
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.PeoplePartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   peoplepartner.PeoplePartnerTable,
			Columns: []string{peoplepartner.PeoplePartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.PeoplePartner
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = _u.schemaConfig.PeoplePartner
	ctx = internal.NewSchemaConfigContext(ctx, _u.schemaConfig)
	_spec.AddModifiers(_u.modifiers...)
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{peoplepartner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// PeoplePartnerUpdateOne is the builder for updating a single PeoplePartner entity.
type PeoplePartnerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PeoplePartnerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUserID sets the "user_id" field.
func (_u *PeoplePartnerUpdateOne) SetUserID(v int) *PeoplePartnerUpdateOne {
	_u.mutation.SetUserID(v)
	return _u
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (_u *PeoplePartnerUpdateOne) SetNillableUserID(v *int) *PeoplePartnerUpdateOne {
	if v != nil {
		_u.SetUserID(*v)
	}
	return _u
}

// SetPeoplePartnerUserID sets the "people_partner_user_id" field.
func (_u *PeoplePartnerUpdateOne) SetPeoplePartnerUserID(v int) *PeoplePartnerUpdateOne {
	_u.mutation.SetPeoplePartnerUserID(v)
	return _u
}

// SetNillablePeoplePartnerUserID sets the "people_partner_user_id" field if the given value is not nil.
func (_u *PeoplePartnerUpdateOne) SetNillablePeoplePartnerUserID(v *int) *PeoplePartnerUpdateOne {
	if v != nil {
		_u.SetPeoplePartnerUserID(*v)
	}
	return _u
}

// SetUser sets the "user" edge to the User entity.
func (_u *PeoplePartnerUpdateOne) SetUser(v *User) *PeoplePartnerUpdateOne {
	return _u.SetUserID(v.ID)
}

// SetPeoplePartnerID sets the "people_partner" edge to the User entity by ID.
func (_u *PeoplePartnerUpdateOne) SetPeoplePartnerID(id int) *PeoplePartnerUpdateOne {
	_u.mutation.SetPeoplePartnerID(id)
	return _u
}

// SetPeoplePartner sets the "people_partner" edge to the User entity.
func (_u *PeoplePartnerUpdateOne) SetPeoplePartner(v *User) *PeoplePartnerUpdateOne {
	return _u.SetPeoplePartnerID(v.ID)
}

// Mutation returns the PeoplePartnerMutation object of the builder.
func (_u *PeoplePartnerUpdateOne) Mutation() *PeoplePartnerMutation {
	return _u.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (_u *PeoplePartnerUpdateOne) ClearUser() *PeoplePartnerUpdateOne {
	_u.mutation.ClearUser()
	return _u
}

// ClearPeoplePartner clears the "people_partner" edge to the User entity.
func (_u *PeoplePartnerUpdateOne) ClearPeoplePartner() *PeoplePartnerUpdateOne {
	_u.mutation.ClearPeoplePartner()
	return _u
}

// Where appends a list predicates to the PeoplePartnerUpdate builder.
func (_u *PeoplePartnerUpdateOne) Where(ps ...predicate.PeoplePartner) *PeoplePartnerUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *PeoplePartnerUpdateOne) Select(field string, fields ...string) *PeoplePartnerUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated PeoplePartner entity.
func (_u *PeoplePartnerUpdateOne) Save(ctx context.Context) (*PeoplePartner, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *PeoplePartnerUpdateOne) SaveX(ctx context.Context) *PeoplePartner {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *PeoplePartnerUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *PeoplePartnerUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *PeoplePartnerUpdateOne) check() error {
	if _u.mutation.UserCleared() && len(_u.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PeoplePartner.user"`)
	}
	if _u.mutation.PeoplePartnerCleared() && len(_u.mutation.PeoplePartnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PeoplePartner.people_partner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *PeoplePartnerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PeoplePartnerUpdateOne {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *PeoplePartnerUpdateOne) sqlSave(ctx context.Context) (_node *PeoplePartner, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(peoplepartner.Table, peoplepartner.Columns, sqlgraph.NewFieldSpec(peoplepartner.FieldUserID, field.TypeInt), sqlgraph.NewFieldSpec(peoplepartner.FieldPeoplePartnerUserID, field.TypeInt))
	if id, ok := _u.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New(`ent: missing "PeoplePartner.user_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := _u.mutation.PeoplePartnerUserID(); !ok {
		return nil, &ValidationError{Name: "people_partner_user_id", err: errors.New(`ent: missing "PeoplePartner.people_partner_user_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !peoplepartner.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if _u.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   peoplepartner.UserTable,
			Columns: []string{peoplepartner.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.PeoplePartner
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   peoplepartner.UserTable,
			Columns: []string{peoplepartner.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.PeoplePartner
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.PeoplePartnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   peoplepartner.PeoplePartnerTable,
			Columns: []string{peoplepartner.PeoplePartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.PeoplePartner
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.PeoplePartnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   peoplepartner.PeoplePartnerTable,
			Columns: []string{peoplepartner.PeoplePartnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.PeoplePartner
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = _u.schemaConfig.PeoplePartner
	ctx = internal.NewSchemaConfigContext(ctx, _u.schemaConfig)
	_spec.AddModifiers(_u.modifiers...)
	_node = &PeoplePartner{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{peoplepartner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
