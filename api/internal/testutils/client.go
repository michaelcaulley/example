package testutils

import (
	"database/sql"
	"example/internal/ent"
	"fmt"
	"os"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/russellhaering/gosaml2/uuid"

	// required by schema hooks.
	_ "example/internal/ent/runtime"
)

var (
	TestDbPort = os.Getenv("TEST_DB_PORT")
)

// NewEntClient creates a new ent client that can be used in tests.
// The client is connected to a new test database that is created from a postgres template.
func NewEntClient(t *testing.T) *ent.Client {
	t.Setenv("TZ", "UTC")
	port := TestDbPort
	if port == "" {
		port = "5431"
	}

	// Create a test database from the template
	dbURI := fmt.Sprintf("postgres://postgres:password@localhost:%s/postgres", port)
	postgresDB, err := sql.Open("pgx", dbURI)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	testDBName := fmt.Sprintf("test_db_%s", strings.ReplaceAll(uuid.NewV4().String(), "-", ""))
	_, err = postgresDB.Exec(fmt.Sprintf("CREATE DATABASE %s TEMPLATE api_test_template;", testDBName))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Create the ent client using the test database
	templateDBURI := fmt.Sprintf("postgres://postgres:password@localhost:%s/%s", port, testDBName)
	testDB, err := sql.Open("pgx", templateDBURI)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	drv := entsql.OpenDB(dialect.Postgres, testDB)
	client := ent.NewClient(ent.Driver(drv))
	t.Cleanup(func() {
		postgresDB.Close()
		testDB.Close()
		client.Close()
	})

	return client
}
