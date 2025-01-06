package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"graphql-server/client"
	"graphql-server/graph"
	"graphql-server/resolver"
	"net/http"
	"os"
	"shared/logger"
)

func main() {
	port := os.Getenv("PORT")
	client.InitCustomerClient(os.Getenv("CUSTOMER_SERVICE_ADDR"))
	client.InitOrderClient(os.Getenv("ORDER_SERVICE_ADDR"))
	client.InitProductClient(os.Getenv("PRODUCT_SERVICE_ADDR"))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
	// error middleware
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		logger.Logger.WithFields(logrus.Fields{
			"message":    err.Message,
			"path":       err.Path,
			"locations":  err.Locations,
			"extensions": err.Extensions,
			"rule":       err.Rule,
		}).Error()
		return err
	})
	// log query
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		oc := graphql.GetOperationContext(ctx)
		logger.Logger.Infof(oc.RawQuery)
		resp := next(ctx)
		return resp
	})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	logger.Logger.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
