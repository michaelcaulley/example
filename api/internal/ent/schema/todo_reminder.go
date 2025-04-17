package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TodoReminder struct {
	ent.Schema
}

func (TodoReminder) Fields() []ent.Field {
	return []ent.Field{
		field.Int("todo_id"),
		field.Int("reminder_id"),
	}
}

func (TodoReminder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todo", Todo.Type).
			Unique().
			Required().
			Field("todo_id").
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("reminder", Reminder.Type).
			Unique().
			Required().
			Field("reminder_id").
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (TodoReminder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("todo_id", "reminder_id"),
		schema.Comment("A join table holding the relationships of todos to reminders"),
		entsql.WithComments(true),
		// entsql.Schema("todo"),
	}
}

func (TodoReminder) Indexes() []ent.Index {
	return []ent.Index{}
}
