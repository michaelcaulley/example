package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type Reminder struct {
	ent.Schema
}

func (Reminder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Reminder.
func (Reminder) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Reminder.
func (Reminder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("todo", Todo.Type).
			Ref("reminders").
			Through("todo_reminders", TodoReminder.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Reminder) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations of the Reminder.
func (Reminder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("reminders"),
		schema.Comment("Reminder for a user to take action."),
		entsql.WithComments(true),
		entsql.Schema("todo"),
	}
}
