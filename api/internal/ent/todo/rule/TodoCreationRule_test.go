package rule_test

import (
	"context"
	"example/internal/ent/privacy"
	"example/internal/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoCreationRule(t *testing.T) {
	// Use PostgreSQL test client instead of SQLite
	client := testutils.NewEntClient(t)
	defer client.Close()

	ctx := context.Background()

	t.Run("rule denies todo creation when no users exist", func(t *testing.T) {
		// Try to create a todo when NO users exist
		// This should FAIL because TodoCreationRule checks that at least one user exists
		_, err := client.Todo.Create().
			SetText("Test Todo").
			SetOwnerID(999).
			Save(ctx)

		// The rule should deny this with "no users found" error
		require.Error(t, err)
		assert.Contains(t, err.Error(), "no users found",
			"TodoCreationRule should deny when no users exist")
	})

	t.Run("rule allows todo creation when user exists", func(t *testing.T) {
		// Create a user using Allow context to bypass any user creation rules
		ctxAllow := privacy.DecisionContext(ctx, privacy.Allow)
		user, err := client.User.Create().
			SetName("Test User").
			Save(ctxAllow)
		require.NoError(t, err)

		// Now create the todo - this should PASS the TodoCreationRule since a user exists
		todo, err := client.Todo.Create().
			SetText("Test Todo").
			SetOwnerID(user.ID).
			Save(ctx)
		require.NoError(t, err)
		require.NotNil(t, todo)
		assert.Equal(t, "Test Todo", todo.Text)
	})
}
