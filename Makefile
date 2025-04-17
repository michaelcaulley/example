# Initialize the project for local development.
init: down _rebuild migrate
	docker compose up -d

# Stand up all containers and networks.
up:
	docker compose up -d

# Tear down all containers and networks.
down:
	docker compose down

# Run the code generation steps.
generate:
	docker compose run --rm go-runner go generate .
	docker compose restart api

atlas-login:
	@echo "⚠️ Because atlas-login uses docker, you MUST copy+paste the auth code into the terminal."
	docker compose run --rm atlas login
	@docker compose down atlas-db

atlas-whoami:
	@echo "📘 If you are not logged in to Atlas, run 'make atlas-login'."
	docker compose run --rm atlas whoami
	@docker compose down atlas-db -v || true
	@docker compose down atlas-db -v

# Create a migration outside the context of ent. It should not conflict with the ent schema.
manual-migration:
	@read -p "✏️ Migration name (ex. sync_job_profile_edges): " name; \
	docker compose run --rm atlas migrate new $$name --dir "file://migrations" || true # continue on error
	@docker compose down atlas-db -v

# Create a migration syncing the ent schema to the database.
migration: _migration hash
_migration:
	@echo "📘 Make sure you are logged in to Atlas before running this command. Run 'make atlas-whoami' to check."
	@read -p "✏️  Migration name (ex. add_user_middle_name_field): " name; \
	docker compose run --rm atlas migrate diff $$name \
                                    --env local || true # continue on error
	@docker compose down atlas-db -v

# Apply all migrations to the database.
migrate:
	docker compose run --rm atlas migrate apply \
		--url "postgres://example_user:password@db:5432/example?sslmode=disable" || true # continue on error
	@docker compose down atlas-db -v

# Recalculate the atlas migration hash.
hash:
	docker compose run --rm atlas migrate hash

# DO NOT CALL DIRECTLY. This is used by the init command to rebuild the containers.
_rebuild:
	docker rm -f example-api example-db
	docker volume rm -f example_database example_go_cache
	docker compose up --build --force-recreate --no-deps -d --wait