package main

import (
	"backend/graph/generated"
	"backend/graph/resolvers"
	"backend/infrastructure/db"
	"backend/pkg/registry"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	dbClient := db.InitDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	auth := registry.AuthServer(dbClient)
	organization := registry.OrganizationServer(dbClient)
	resolver := resolvers.NewResolver(auth, organization)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	srv.AddTransport(&transport.Websocket{})

	mux := http.NewServeMux()
	mux.Handle("/query", resolvers.Middleware(srv, resolver))
	mux.Handle("/play", playground.Handler("GraphQL playground", "/query"))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PATCH", "PUT", "DELETE", "HEAD"},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"X-Requested-With", "Content-Type",
			"Authorization", "Accept-Encoding", "Host", "Origin", "Accept",
		},
		Debug: false,
	})
	muxHandler := c.Handler(mux)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, muxHandler))
}
