package schema

import (
	"time"

	"entgo.io/ent/privacy"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type GroupMember struct {
	ent.Schema
}

func (GroupMember) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the GroupMember.
func (GroupMember) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				&entsql.Annotation{
					Default: "CURRENT_TIMESTAMP",
				},
			),
		field.Int("user_id"),
		field.Int("group_id"),
	}
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("group", Group.Type).
			Required().
			Unique().
			Field("group_id").
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
	}
}

func (GroupMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("group_id"),
		index.Fields("user_id"),
	}
}

// Annotations of the GroupMember.
func (GroupMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "group_id"),
		entsql.WithComments(true),
	}
}

func (GroupMember) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
