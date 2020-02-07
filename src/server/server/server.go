package main

import (
	catalog "catalog"
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

	datastore := catalog.OpenDatastore(*dbPath)
	datastore.Ping()

	defer datastore.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(catalog.NewExecutableSchema(catalog.Config{Resolvers: &catalog.Resolver{datastore}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
