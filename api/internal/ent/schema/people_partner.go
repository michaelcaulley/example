package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PeoplePartner struct {
	ent.Schema
}

func (PeoplePartner) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

func (PeoplePartner) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("people_partner_user_id"),
	}
}

func (PeoplePartner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("people_partner", User.Type).
			Required().
			Unique().
			Field("people_partner_user_id"),
	}
}

func (PeoplePartner) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("people_partner_user_id"),
	}
}

func (PeoplePartner) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("people_partners"),
		entsql.Schema("user"),
		field.ID("user_id", "people_partner_user_id"),
		entsql.WithComments(true),
	}
}
