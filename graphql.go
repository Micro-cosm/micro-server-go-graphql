package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"weja.us/micro-cosm/micro-server-go-graphql/graph"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()
	port := os.Getenv("TARGET_REMOTE_PORT")

	if port == "" {
		port = defaultPort
	}

	// router.Use( cors.New( cors.Options{ AllowCredentials: true, Debug: true }).Handler)								// Add CORS to request -- https://github.com/rs/cors
	router.Use(cors.New(cors.Options{}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{Upgrader: websocket.Upgrader{CheckOrigin: nil, ReadBufferSize: 1024, WriteBufferSize: 1024}})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
