// Code generated by ent, DO NOT EDIT.

package reminder

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the reminder type in the database.
	Label = "reminder"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTodo holds the string denoting the todo edge name in mutations.
	EdgeTodo = "todo"
	// EdgeTodoReminders holds the string denoting the todo_reminders edge name in mutations.
	EdgeTodoReminders = "todo_reminders"
	// Table holds the table name of the reminder in the database.
	Table = "reminders"
	// TodoTable is the table that holds the todo relation/edge. The primary key declared below.
	TodoTable = "todo_reminders"
	// TodoInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodoInverseTable = "todos"
	// TodoRemindersTable is the table that holds the todo_reminders relation/edge.
	TodoRemindersTable = "todo_reminders"
	// TodoRemindersInverseTable is the table name for the TodoReminder entity.
	// It exists in this package in order to avoid circular dependency with the "todoreminder" package.
	TodoRemindersInverseTable = "todo_reminders"
	// TodoRemindersColumn is the table column denoting the todo_reminders relation/edge.
	TodoRemindersColumn = "reminder_id"
)

// Columns holds all SQL columns for reminder fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// TodoPrimaryKey and TodoColumn2 are the table columns denoting the
	// primary key for the todo relation (M2M).
	TodoPrimaryKey = []string{"todo_id", "reminder_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "example/internal/ent/runtime"
var (
	Hooks [1]ent.Hook
)

// OrderOption defines the ordering options for the Reminder queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTodoCount orders the results by todo count.
func ByTodoCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTodoStep(), opts...)
	}
}

// ByTodo orders the results by todo terms.
func ByTodo(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodoStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTodoRemindersCount orders the results by todo_reminders count.
func ByTodoRemindersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTodoRemindersStep(), opts...)
	}
}

// ByTodoReminders orders the results by todo_reminders terms.
func ByTodoReminders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodoRemindersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTodoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TodoTable, TodoPrimaryKey...),
	)
}
func newTodoRemindersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodoRemindersInverseTable, TodoRemindersColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, TodoRemindersTable, TodoRemindersColumn),
	)
}
