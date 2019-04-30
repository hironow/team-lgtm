package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/hironow/team-lgtm/backend/gql"
)

const defaultPort = "8080"

const useDatastore = false

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	var dsClient *datastore.Client
	if useDatastore {
		var err error
		ctx := context.Background()
		dsClient, err = datastore.NewClient(ctx, "project-id")
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	resolver, err := gql.NewResolver(dsClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Handle("/", handler.Playground("TeamLGTM GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: resolver})))

	log.Printf("connect to http://localhost:%s/ for TeamLGTM GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
