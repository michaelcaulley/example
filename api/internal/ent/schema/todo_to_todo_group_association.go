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

type TodoToTodoGroupAssociation struct {
	ent.Schema
}

func (TodoToTodoGroupAssociation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (TodoToTodoGroupAssociation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("todo_id").
			Immutable(),
		field.Int("todo_group_really_really_long_identifier").
			Immutable(),
		field.Int("assignee_id"),
	}
}

func (TodoToTodoGroupAssociation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todo", Todo.Type).
			Required().
			Unique().
			Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("todo_id"),
		edge.To("todo_group", TodoGroup.Type).
			Required().
			Unique().
			Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Field("todo_group_really_really_long_identifier"),
	}
}

func (TodoToTodoGroupAssociation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("todo_id"),
		index.Fields("todo_group_really_really_long_identifier"),
		index.Fields("assignee_id"),
	}
}

func (TodoToTodoGroupAssociation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("todo"),
		entsql.Table("todo_to_todo_group_associations"),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entsql.WithComments(true),
	}
}
