package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/hironow/team-lgtm/backend/config"
	"github.com/hironow/team-lgtm/backend/gql"
)

func main() {
	env, err := config.ReadFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read env vars: %s\n", err)
		os.Exit(1)
	}

	var dsClient *datastore.Client
	if env.UseDatastore {
		var err error
		dsClient, err = datastore.NewClient(context.Background(), env.DatastoreProjectID)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	resolver, err := gql.NewResolver(dsClient)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := chi.NewRouter()
	router.Handle("/", handler.Playground("TeamLGTM GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: resolver})))

	log.Printf("connect to http://localhost:%d/ for TeamLGTM GraphQL playground", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.Port), router))
}
