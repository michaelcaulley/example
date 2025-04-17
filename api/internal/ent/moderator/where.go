// Code generated by ent, DO NOT EDIT.

package moderator

import (
	"example/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"example/internal/ent/internal"
)

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldEQ(FieldUserID, v))
}

// ModeratorUserID applies equality check predicate on the "moderator_user_id" field. It's identical to ModeratorUserIDEQ.
func ModeratorUserID(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldEQ(FieldModeratorUserID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Moderator {
	return predicate.Moderator(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Moderator {
	return predicate.Moderator(sql.FieldNotIn(FieldUserID, vs...))
}

// ModeratorUserIDEQ applies the EQ predicate on the "moderator_user_id" field.
func ModeratorUserIDEQ(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldEQ(FieldModeratorUserID, v))
}

// ModeratorUserIDNEQ applies the NEQ predicate on the "moderator_user_id" field.
func ModeratorUserIDNEQ(v int) predicate.Moderator {
	return predicate.Moderator(sql.FieldNEQ(FieldModeratorUserID, v))
}

// ModeratorUserIDIn applies the In predicate on the "moderator_user_id" field.
func ModeratorUserIDIn(vs ...int) predicate.Moderator {
	return predicate.Moderator(sql.FieldIn(FieldModeratorUserID, vs...))
}

// ModeratorUserIDNotIn applies the NotIn predicate on the "moderator_user_id" field.
func ModeratorUserIDNotIn(vs ...int) predicate.Moderator {
	return predicate.Moderator(sql.FieldNotIn(FieldModeratorUserID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Moderator {
	return predicate.Moderator(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Moderator
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Moderator {
	return predicate.Moderator(func(s *sql.Selector) {
		step := newUserStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Moderator
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasModerator applies the HasEdge predicate on the "moderator" edge.
func HasModerator() predicate.Moderator {
	return predicate.Moderator(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ModeratorColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ModeratorTable, ModeratorColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Moderator
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModeratorWith applies the HasEdge predicate on the "moderator" edge with a given conditions (other predicates).
func HasModeratorWith(preds ...predicate.User) predicate.Moderator {
	return predicate.Moderator(func(s *sql.Selector) {
		step := newModeratorStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Moderator
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Moderator) predicate.Moderator {
	return predicate.Moderator(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Moderator) predicate.Moderator {
	return predicate.Moderator(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Moderator) predicate.Moderator {
	return predicate.Moderator(sql.NotPredicates(p))
}
