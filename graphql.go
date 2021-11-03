

package main

import (
	"github.com/99designs/gqlgen/example/todo"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"weja.us/micro-cosm/micro-server-go-graphql/graph"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port	:= os.Getenv("PORT")
	if port == "" { port = defaultPort }
	srv		:= handler.NewDefaultServer( generated.NewExecutableSchema( generated.Config{ Resolvers: &graph.Resolver{}}))
	srv1	:= handler.New( todo.NewExecutableSchema( todo.New()))
	http.Handle("/", playground.Handler("GraphQL playground", "/query" ))
	http.Handle("/query", srv )
	http.Handle("/query1", srv1 )
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port )
	log.Fatal( http.ListenAndServe( ":" + port, nil ))
}
