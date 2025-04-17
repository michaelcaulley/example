// Code generated by ent, DO NOT EDIT.

package ent

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Text        string
	ReminderIDs []int
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetText(i.Text)
	if v := i.ReminderIDs; len(v) > 0 {
		m.AddReminderIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name    string
	TodoIDs []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetName(i.Name)
	if v := i.TodoIDs; len(v) > 0 {
		m.AddTodoIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}
