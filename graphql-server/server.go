package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/joho/godotenv/autoload"
	"graphql-server/client"
	"graphql-server/graph"
	"net/http"
	"os"
	"shared/logger"
)

func main() {
	port := os.Getenv("PORT")
	client.InitCustomerClient(os.Getenv("CUSTOMER_SERVICE_ADDR"))
	client.InitOrderClient(os.Getenv("ORDER_SERVICE_ADDR"))
	client.InitProductClient(os.Getenv("PRODUCT_SERVICE_ADDR"))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	logger.Logger.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
