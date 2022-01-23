package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi"
	"github.com/gookit/color"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/willf/pad"
	"weja.us/micro-cosm/micro-server-go-graphql/graph"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/generated"
	"weja.us/micro-cosm/micro-server-go-graphql/lib"
)

var (
	Red               = color.FgRed.Render // Black = color.FgBlack.Render // WhiteOnRed = color.Style{color.FgLightWhite, color.BgRed}.Render // BlackOnGreen = color.Style{color.FgBlack, color.BgGreen}.Render // Blue = color.FgBlue.Render // Green = color.FgGreen.Render // Grey = color.FgDarkGray.Render // Yellow = color.FgYellow.Render // BlackOnGray = color.Style{color.FgBlack, color.BgGray}.Render // renderColor = Black // LogWin = BlackOnGreen(" √ ") // LogLose = WhiteOnRed(" X ")
	IsLocal           bool
	IsLocalFlagSet    bool
	IsLocalPortSet    bool
	IsLocalFlagString string
	LocalPort         string
)

const defaultPort = "8080"

func main() {
	debug()
	r := chi.NewRouter()
	IsLocalFlagString, IsLocalFlagSet = os.LookupEnv("IS_LOCAL")
	IsLocal, _ = strconv.ParseBool(IsLocalFlagString)
	port := os.Getenv("REMOTE_PORT")
	if port == "" {
		port = defaultPort
	}

	r.Use(cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

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

	LocalPort, IsLocalPortSet = os.LookupEnv("LOCAL_PORT")
	if IsLocalPortSet && IsLocalFlagSet && IsLocal {
		log.Printf("verification link: http://localhost:%s", LocalPort)
	} else {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func debug() {
	log.Printf("OPERATING ENVIRONMENT -- Empty values below require an 'environment' entry in docker-compose.yml before use.")
	log.Printf(pad.Right("", 120, "="))
	_, isAcdSet := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !isAcdSet && IsLocal {
		log.Fatalf(Red("'FdAdc' or 'GOOGLE_APPLICATION_CREDENTIALS' must be set to continue")) // Google's Application Default Credentials(ADC)
	} else {
		log.Printf("GOOGLE_APPLICATION_CREDENTIALS%s %s\n", pad.Right(" ", 4, "."), os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	}
	log.Printf("SERVICE%s %s\n", pad.Right(" ", 27, "."), os.Getenv("SERVICE"))
	log.Printf("EXECUTABLE%s %s\n", pad.Right(" ", 24, "."), os.Getenv("EXECUTABLE"))
	log.Printf("DEBUG%s %s\n", pad.Right(" ", 29, "."), os.Getenv("DEBUG"))
	log.Printf("LOGS%s %s\n", pad.Right(" ", 30, "."), os.Getenv("LOGS"))
	log.Printf("LOCAL_PORT%s %s\n", pad.Right(" ", 24, "."), os.Getenv("LOCAL_PORT"))
	log.Printf("REMOTE_PORT%s %s\n", pad.Right(" ", 23, "."), os.Getenv("REMOTE_PORT"))
	log.Printf("ROUTE_BASE%s %s\n", pad.Right(" ", 24, "."), os.Getenv("ROUTE_BASE"))
	log.Printf("TZ%s %s\n", pad.Right(" ", 32, "."), os.Getenv("TZ"))
	log.Printf("IS_DEBUG%s %s\n", pad.Right(" ", 26, "."), os.Getenv("IS_DEBUG"))
	log.Printf("IS_TEST%s %s\n", pad.Right(" ", 27, "."), os.Getenv("IS_TEST"))
	log.Printf("IS_LOCAL%s %s\n", pad.Right(" ", 26, "."), os.Getenv("IS_LOCAL"))
	log.Printf("IS_REMOTE%s %s\n", pad.Right(" ", 25, "."), os.Getenv("IS_REMOTE"))
	log.Printf("IMAGE_URL%s %s\n", pad.Right(" ", 25, "."), os.Getenv("IMAGE_URL"))
	log.Printf("CONTAINER%s %s\n", pad.Right(" ", 25, "."), os.Getenv("CONTAINER"))
	log.Printf("REPO%s %s\n", pad.Right(" ", 30, "."), os.Getenv("REPO"))
	log.Printf("ALIAS%s %s\n", pad.Right(" ", 29, "."), os.Getenv("TARGET_ALIAS"))
	log.Printf("IMAGE_TAG%s %s\n", pad.Right(" ", 25, "."), os.Getenv("TARGET_IMAGE_TAG"))
	log.Printf("LOG_LEVEL%s %s\n", pad.Right(" ", 25, "."), os.Getenv("TARGET_LOG_LEVEL"))
	log.Printf("PROJECT_ID%s %s\n", pad.Right(" ", 24, "."), os.Getenv("TARGET_PROJECT_ID"))

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
