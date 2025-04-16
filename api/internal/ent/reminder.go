// Code generated by ent, DO NOT EDIT.

package ent

import (
	"example/internal/ent/reminder"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Reminder is the model entity for the Reminder schema.
type Reminder struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reminder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reminder.FieldID:
			values[i] = new(sql.NullInt64)
		case reminder.FieldCreatedAt, reminder.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reminder fields.
func (_m *Reminder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reminder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case reminder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				_m.CreatedAt = value.Time
			}
		case reminder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				_m.UpdatedAt = value.Time
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Reminder.
// This includes values selected through modifiers, order, etc.
func (_m *Reminder) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// Update returns a builder for updating this Reminder.
// Note that you need to call Reminder.Unwrap() before calling this method if this Reminder
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Reminder) Update() *ReminderUpdateOne {
	return NewReminderClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Reminder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Reminder) Unwrap() *Reminder {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reminder is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Reminder) String() string {
	var builder strings.Builder
	builder.WriteString("Reminder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(_m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(_m.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Reminders is a parsable slice of Reminder.
type Reminders []*Reminder
