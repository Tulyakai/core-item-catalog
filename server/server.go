package server

import (
	"log"
	"net/http"

	"com.ktb.ai.core-item-catalog/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// StartGraphQLServer starts the GraphQL server
func StartGraphQLServer() {
	// Create a new GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// Serve the GraphQL API endpoint
	http.Handle("/query", srv)

	// Serve the GraphQL Playground (optional)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	// Start the server on port 8080
	log.Println("GraphQL server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
