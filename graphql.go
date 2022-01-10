package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"weja.us/micro-cosm/micro-server-go-graphql/graph"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/generated"
	"weja.us/micro-cosm/micro-server-go-graphql/lib"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	port := os.Getenv("TARGET_REMOTE_PORT")
	if port == "" {
		port = defaultPort
	}

	r.Use(cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	r.Route("/tab", func(r chi.Router) {
		r.Route("/{tabName}", func(r chi.Router) {
			r.Use(lib.QueryCtx)
			r.Route("/query", func(r chi.Router) {
				srv := handler.NewDefaultServer(
					generated.NewExecutableSchema(
						generated.Config{Resolvers: &graph.Resolver{}},
					),
				)
				srv.AddTransport(
					&transport.Websocket{
						Upgrader: websocket.Upgrader{
							CheckOrigin:     nil,
							ReadBufferSize:  1024,
							WriteBufferSize: 1024,
						},
					},
				)
				log.Printf("!!!! query away !!!!")
				r.Handle("/", srv)
			})
		})
	})
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// r.HandleFunc("/favicon.ico", faviconHandler)
// r.Handle("/", playground.Handler("GraphQL playground", "/query"))
// r.Route("/query", func(r chi.Router) { // r.Handle("/query", srv)
// 	Presbies = getSheet(RosterSheetId, RosterDefaultTab+RosterSheetRange).Values
// 	r.Handle("/", srv)
// })
// r.Handle("/", playground.Handler("GraphQL playground", "/tab/Presbies/query"))
// func faviconHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("!!!!!!!!!!!!!!! ico, i-co, i, co !!!!!!!!!!!!!!!!!")
// 	http.ServeFile(w, r, "favicon.ico")
// }
