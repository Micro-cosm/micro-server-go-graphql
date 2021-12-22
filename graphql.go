package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"net/http"
	"os"
	"weja.us/micro-cosm/micro-server-go-graphql/graph"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/generated"
)

const defaultPort = "8080"

var (
	rosterSheetId    = "1V8L8Ub1FRKhXo1pLxwxXiBwIz1TWtatqheHh4RPltJ8"
	rosterSheetRange = "Presbies-dev!A2:Q"
)

func main() {
	router := chi.NewRouter()
	// sheetId := os.Getenv("ROSTER_SHEET_ID")
	// sheetRange := os.Getenv("ROSTER_TAB_NAME")
	port := os.Getenv("TARGET_REMOTE_PORT")
	if port == "" {
		port = defaultPort
	}

	graph.Presbies = getSheet(rosterSheetId, rosterSheetRange).Values

	router.Use(cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowCredentials: true, Debug: true}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{Upgrader: websocket.Upgrader{CheckOrigin: nil, ReadBufferSize: 1024, WriteBufferSize: 1024}})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getSheet(spreadsheetId string, readRange string) *sheets.ValueRange {
	ctx := context.Background()
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", ".secrets/credentials.json")
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return resp
}
