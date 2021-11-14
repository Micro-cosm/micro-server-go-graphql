package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/generated"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/model"
)

func (r *queryResolver) Presbies(ctx context.Context) ([]model.Presby, error) {
	sheetId := "1V8L8Ub1FRKhXo1pLxwxXiBwIz1TWtatqheHh4RPltJ8"
	sheetRange := "Presbies!A2:Q"
	presbies := getSheet(sheetId, sheetRange).Values
	newPresbies := make([]model.Presby, len(presbies))

	for cnt, field := range presbies {
		log.Printf("presby record  #%d", cnt)
		tmpId, e := strconv.ParseUint(field[0].(string), 10, 64)
		if e != nil {
			log.Printf("FAIL!!! ID")
		}
		newPresbies[cnt].ID = int(tmpId)
		log.Printf("\tfield #%d -- value: %d", 0, tmpId)

		if field[1] == "TRUE" {
			newPresbies[cnt].IsActive = true
		} else {
			newPresbies[cnt].IsActive = false
		}
		log.Printf("\tfield #%d -- value: %v", 1, field[1])
		newPresbies[cnt].Last = field[2].(string)
		log.Printf("\tfield #%d -- value: %v", 2, field[2].(string))

		guestsRaw := field[3].(string)
		guests := strings.Split(guestsRaw, ",")
		newPresbies[cnt].Guests = guests
		log.Printf("\tfield #%d -- value: %v", 3, guests)

		guestingsRaw := field[4].(string)
		guestings := strings.Split(guestingsRaw, ",")
		newPresbies[cnt].Guestings = guestings
		log.Printf("\tfield #%d -- value: %v", 4, guestings)

		hostingsRaw := field[5].(string)
		hostings := strings.Split(hostingsRaw, ",")
		newPresbies[cnt].Hostings = hostings
		log.Printf("\tfield #%d -- value: %v", 5, hostings)

		tmpSeats, e := strconv.ParseUint(field[6].(string), 10, 64)
		if e != nil {
			log.Printf("FAIL!!! seats")
		}
		newPresbies[cnt].Seats = int(tmpSeats)
		log.Printf("\tfield #%d -- value: %v", 6, field[6].(string))

		tmpUnknown1, e := strconv.ParseUint(field[7].(string), 10, 64)
		if e != nil {
			log.Printf("FAIL!!! seats")
		}
		newPresbies[cnt].Unknown1 = int(tmpUnknown1)
		log.Printf("\tfield #%d -- value: %d", 7, tmpUnknown1)

		tmpUnknown2, e := strconv.ParseUint(field[8].(string), 10, 64)
		if e != nil {
			log.Printf("FAIL!!! seats")
		}
		newPresbies[cnt].Unknown2 = int(tmpUnknown2)
		log.Printf("\tfield #%d -- value: %d", 7, tmpUnknown2)

		newPresbies[cnt].Email = field[9].(string)
		log.Printf("\tfield #%d -- value: %v", 9, field[9].(string))
		newPresbies[cnt].Home = field[10].(string)
		log.Printf("\tfield #%d -- value: %v", 10, field[10].(string))
		newPresbies[cnt].Cell = field[11].(string)
		log.Printf("\tfield #%d -- value: %v", 11, field[11].(string))
		newPresbies[cnt].Smail = field[12].(string)
		log.Printf("\tfield #%d -- value: %v", 12, field[12].(string))
		newPresbies[cnt].City = field[13].(string)
		log.Printf("\tfield #%d -- value: %v", 13, field[13].(string))
		newPresbies[cnt].St = field[14].(string)
		log.Printf("\tfield #%d -- value: %v", 14, field[14].(string))
		newPresbies[cnt].Zip = field[15].(string)
		log.Printf("\tfield #%d -- value: %v", 15, field[15].(string))
		newPresbies[cnt].Mmail = field[16].(string)
		log.Printf("\tfield #%d -- value: %v", 16, field[16].(string))

		newPresbies[cnt].Key = newPresbies[cnt].Last + "-" + strconv.Itoa(newPresbies[cnt].ID) + "-" + strconv.Itoa(newPresbies[cnt].Seats) + "-" + strconv.Itoa(len(newPresbies[cnt].Guests))
	}
	return newPresbies, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func getSheet(spreadsheetId string, readRange string) *sheets.ValueRange {
	ctx := context.Background()
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
