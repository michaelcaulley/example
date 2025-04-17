data "composite_schema" "app" {
  schema {
    url = "file://db/extensions.pg.hcl"
  }

  schema {
    url = "ent://internal/ent/schema"
  }

}

env "local" {
  src = data.composite_schema.app.url
  dev = "postgres://postgres:password@atlas-db:5432/postgres?sslmode=disable"
}