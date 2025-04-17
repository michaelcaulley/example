data "composite_schema" "app" {
  schema {
    url = "file://db/extensions.pg.hcl"
  }

  schema {
    url = "ent://internal/ent/schema?globalid=static&hash-symbols=1"
  }
}

lint {
  non_linear {
    error = true
  }
  naming {
    error   = true
    match   = "^[a-z0-9](?:[a-z0-9_]*[a-z0-9])?$"
    message = "must be lowercase, cannot start/end with underscore"
  }
}

env "local" {
  src = data.composite_schema.app.url
  dev = "postgres://postgres:password@atlas-db:5432/postgres?sslmode=disable"
}
