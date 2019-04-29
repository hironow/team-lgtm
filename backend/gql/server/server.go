package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/hironow/team-lgtm/backend/gql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Handle("/", handler.Playground("TeamLGTM GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: gql.NewResolver()})))

	log.Printf("connect to http://localhost:%s/ for TeamLGTM GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
