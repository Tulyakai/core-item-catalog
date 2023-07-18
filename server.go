package main

import (
	"com.ktb.ai.core-item-catalog/graph"
	"com.ktb.ai.core-item-catalog/app/config"
	"com.ktb.ai.core-item-catalog/app/infrastructure/gormrepository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler(repo *gormrepository.GormCatalogRepository) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		GormCatalogRepository: repo,
	}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func main() {
	// Setting up DB connection
	db := config.NewDBInstance(&config.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		Database: "postgres",
		TimeZone: "Asia/Bangkok",
	})
	config.Migrate(db)
	repo := gormrepository.NewGormCatalogRepository(db)

	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler(repo))
	r.GET("/", playgroundHandler())
	r.Run()
}
