package lib

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	RosterSheetId    = "1V8L8Ub1FRKhXo1pLxwxXiBwIz1TWtatqheHh4RPltJ8"
	RosterSheetRange = "!A2:P"
	PresbyData       [][]interface{}
)

func QueryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tabAlias := chi.URLParam(r, "tabName")
		log.Printf("?????? tab name--->%s%s<---", tabAlias, RosterSheetRange)
		if tabAlias != "" {
			target := tabAlias + RosterSheetRange
			PresbyData = getSheet(RosterSheetId, target).Values
			ctx := context.WithValue(r.Context(), "presbyData", PresbyData)
			next.ServeHTTP(w, r.WithContext(ctx))
			log.Printf("!!!!! func'ing tab--->%s<---\n\n%v", target, PresbyData)
		}
	})
}

func getSheet(spreadsheetId string, readRange string) *sheets.ValueRange {
	ctx := context.Background()
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", ".secrets/credentials.json")
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	service, err := sheets.NewService(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	log.Printf("?!?!?!?!?!?!?!?!??!?!?!?!!?!?-->%s<-->%s<",
		spreadsheetId,
		readRange)
	response, err := service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return response
}
