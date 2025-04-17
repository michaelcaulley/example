package rule

import (
	"context"

	"example/internal/ent"
	"example/internal/ent/privacy"
)

// TodoCreationRule is a silly rule, but it shows how the mutation client doesn't understand what
// schema other entities are part of.
func TodoCreationRule() privacy.MutationRule {
	return privacy.TodoMutationRuleFunc(func(ctx context.Context, m *ent.TodoMutation) error {

		users, err := m.Client().User.Query().All(ctx)
		if err != nil {
			return privacy.Denyf("error querying users: %v", err)
		}

		if len(users) == 0 {
			return privacy.Denyf("no users found")
		}

		return privacy.Allow
	})
}
