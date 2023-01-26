package main

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"graphqlPlayground/graph"
	"graphqlPlayground/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	// Load the .env first
	err := godotenv.Load()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error loading .env file"))
	}

	// Get DB set up
	db, err := graph.Connect()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error connecting to DB"))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
