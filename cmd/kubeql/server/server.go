package server

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/kubeql/kubeql/internal/gql"
)

func runServer(a *App, listen string) error {
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}})))

	return http.ListenAndServe(listen, nil)
}
