package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/kubeql/kubeql/internal/gql"
)

var app *App

type App struct {
	schema graphql.ExecutableSchema
}

func NewApp() *App {
	app = &App{
		schema: gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}}),
	}
	return app
}
