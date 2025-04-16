package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin implements the ent.Mixin for sharing created_at and updated_at fields.
type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entgql.OrderField("CREATED_AT"),
				&entsql.Annotation{
					Default: "CURRENT_TIMESTAMP",
				},
			),
		field.Time("updated_at").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entgql.OrderField("UPDATED_AT"),
				&entsql.Annotation{
					DefaultExpr: "CURRENT_TIMESTAMP",
				},
			),
	}
}

func (TimeMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		// We use this mutation hook to set the `created_at` and `updated_at` fields instead of
		// using the Default() function in the schema definition because we want to ensure that on
		// creation that the `created_at` and `updated_at` fields are set to the same value.
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				if skip, _ := ctx.Value(timeKey{}).(bool); skip {
					return next.Mutate(ctx, m)
				}

				now := time.Now().Truncate(time.Microsecond) // match postgres timestamp precision
				if m.Op().Is(ent.OpUpdate) || m.Op().Is(ent.OpUpdateOne) {
					err := m.SetField("updated_at", now)
					if err != nil {
						return nil, fmt.Errorf("time mixin hook: %w ", err)
					}
				} else if m.Op().Is(ent.OpCreate) {
					err := m.SetField("created_at", now)
					if err != nil {
						return nil, fmt.Errorf("time mixin hook: %w ", err)
					}
					err = m.SetField("updated_at", now)
					if err != nil {
						return nil, fmt.Errorf("time mixin hook: %w ", err)
					}
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}

type timeKey struct{}

// SkipTime returns a new context that skips the time hook.
func SkipTime(parent context.Context) context.Context {
	return context.WithValue(parent, timeKey{}, true)
}
