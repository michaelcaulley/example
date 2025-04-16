//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
)

func main() {
	gqlExtension, err := entgql.NewExtension(
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./internal/graph/ent.graphql"),
		entgql.WithWhereInputs(true),
		entgql.WithNodeDescriptor(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(gqlExtension),
		entc.FeatureNames(
			"privacy",
			"entql",
			"schema/snapshot",
			"sql/globalid",
			"sql/execquery",
			"sql/upsert",
			"sql/modifier",
		),
	}
	if err := entc.Generate("./internal/ent/schema", opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
