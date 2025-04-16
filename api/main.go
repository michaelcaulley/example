package main

import (
	"context"
	"database/sql"
	"example/internal/ent"
	"example/internal/graph"
	"fmt"
	"net/http"
	"os"
	"slices"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

// newSQLDB creates a new SQL database connection and returns it.
func newSQLDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("failed newSQLDB: missing DATABASE_URL")
	}

	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed newSQLDB: %w", err)
	}

	poolConfig.MaxConnLifetime = time.Minute * 60
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed newSQLDB: %w", err)
	}
	db := stdlib.OpenDBFromPool(pool)

	if os.Getenv("ENVIRONMENT") == "local" {
		loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
		db = sqldblogger.OpenDriver(dbURL, db.Driver(), loggerAdapter)
	}

	// Check the db connection is working
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed newSQLDB: failed ping: %w", err)
	}

	return db, nil
}

// newEntClient creates a new ent client and returns it.
func newEntClient(db *sql.DB) *ent.Client {
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	db, err := newSQLDB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed app initialization")
	}

	entClient := newEntClient(db)
	defer entClient.Close()

	// Set up our routing
	router := chi.NewRouter()

	// region: GraphQL Server Endpoint and Playground Routes
	srv := graph.New(graph.NewSchema(entClient))
	srv.Use(entgql.Transactioner{TxOpener: entClient}) // Mutations offer a transactional client
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		// The `github.com/gorilla/websocket.Upgrader` is used to handle the transition
		// from an HTTP connection to a WebSocket connection.
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Allow exact match on host.
				origin := r.Header.Get("Origin")
				host := r.Header.Get("Host")
				if origin == "https://"+host {
					return true
				}

				var allowedOrigins []string
				if os.Getenv("ENVIRONMENT") == "local" {
					allowedOrigins = []string{"https://localhost"}
				} else {
					allowedOrigins = []string{}
				}

				if !slices.Contains(allowedOrigins, origin) {
					log.Warn().Str("origin", origin).Msg("rejected websocket connection attempt")
					return false
				} else {
					return true
				}
			},
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	router.Handle("/graphql", srv)
	router.Handle("/graphiql", playground.Handler("GraphQL Playground", "/graphql"))
	// endregion

	// Finally, start the http server to listen and serve
	log.Info().Msg("listening and serving...")
	log.Fatal().Err(http.ListenAndServe(":80", router)).Send()
}
