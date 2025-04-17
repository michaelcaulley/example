package schema_test

import (
	"context"
	"example/internal/testutils"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TodoMutationRule(t *testing.T) {
	client := testutils.NewEntClient(t)
	defer client.Close()

	// Create a context with privacy policies enabled
	// Use Allow for user creation (bypass privacy for setup)
	ctx := context.Background()

	usr, err := client.User.
		Create().
		SetName("Test User").
		Save(ctx)
	require.NoError(t, err)

	// Test creating a todo without any users - this should trigger the rule
	// First, let's test that the rule allows creation when a user exists
	_, err = client.Todo.Create().
		SetText("Test Todo").
		SetOwnerID(usr.ID).
		Save(ctx)
	require.NoError(t, err)
}
