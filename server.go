package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fico308/graphql/graph"
	"github.com/fico308/graphql/graph/auth"
	"github.com/go-chi/chi"
)


func main() {
	router := chi.NewRouter()

	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	fmt.Println("Server start")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
