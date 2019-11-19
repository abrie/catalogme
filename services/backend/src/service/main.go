package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"service/storage"
)

func main() {
	port := flag.Int("p", 0, "port to serve on")
	directory := flag.String("d", "", "the directory to store json")
	flag.Parse()

	if *port == 0 || *directory == "" {
		flag.PrintDefaults()
		panic("Missing arguments.")
	}

	storage := &storage.Storage{Directory: *directory}

	r := chi.NewRouter()
	r.Use(NewCorsHandler())
	r.Get("/data.json", storage.GetHandler)
	//r.Post("/", storage.PostHandler)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("—Backend serving %s via HTTP on %s\n", storage.Datafile(), addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func NewCorsHandler() func(http.Handler) http.Handler {
	options := cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}

	cors := cors.New(options)
	return cors.Handler
}
