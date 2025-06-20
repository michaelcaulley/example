package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TodoGroup struct {
	ent.Schema
}

func (TodoGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (TodoGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty().
			MaxLen(255),
	}
}

func (TodoGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type).
			Through("grouped_todos", TodoToTodoGroupAssociation.Type).
			Annotations(
				entgql.RelayConnection(),
				entsql.OnDelete(entsql.Cascade),
			),
	}
}

func (TodoGroup) Indexes() []ent.Index {
	return []ent.Index{}
}

func (TodoGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("todo"),
		entsql.Table("todo_groups"),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		entsql.WithComments(true),
	}
}
