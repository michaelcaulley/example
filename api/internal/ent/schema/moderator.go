package schema

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Moderator struct {
	ent.Schema
}

func (Moderator) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (Moderator) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("moderator_user_id"),
	}
}

func (Moderator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("moderator", User.Type).
			Required().
			Unique().
			Field("moderator_user_id"),
	}
}

func (Moderator) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("moderator_user_id"),
	}
}

func (Moderator) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "moderator_user_id"),
		entsql.WithComments(true),
		entsql.Schema("user"),
	}
}
