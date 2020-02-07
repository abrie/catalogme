package main

import (
	"catalog/datastore"
	"catalog/graph"
	"catalog/resolver"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	dbPath := flag.String("datastore", "", "Path to the datastore db.")
	flag.Parse()
	if *dbPath == "" {
		flag.PrintDefaults()
		return
	}

	store := datastore.OpenDatastore(*dbPath)
	store.Ping()

	defer store.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{store}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
