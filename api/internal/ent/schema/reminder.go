package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
	return []ent.Edge{}
}

func (Reminder) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations of the Reminder.
func (Reminder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entsql.WithComments(true),
		// entsql.Schema("todo"),
	}
}
