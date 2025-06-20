package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Todo struct {
	ent.Schema
}

func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty().
			MaxLen(255),
		field.Time("done_at").
			Optional().
			Nillable().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput),
			),
		field.Int("owner_id").
			Immutable().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput),
			),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("todos").
			Field("owner_id").
			Unique().
			Required().
			Immutable().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput),
			),
		edge.To("reminders", Reminder.Type).
			Through("todo_reminders", TodoReminder.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("groups", TodoGroup.Type).
			Ref("todos").
			Annotations(
				entgql.RelayConnection(),
				entsql.OnDelete(entsql.Cascade),
			).
			Through("grouped_todos", TodoToTodoGroupAssociation.Type),
	}
}

func (Todo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}

// Annotations of the Todo.
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("todo"),
		entsql.Table("todos"),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entsql.WithComments(true),
	}
}
