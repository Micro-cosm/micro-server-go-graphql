package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/generated"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented -- %v -- %v", ctx, input))
}

func (r *queryResolver) Presbies(ctx context.Context) ([]model.Presby, error) {
	sheetId := "1V8L8Ub1FRKhXo1pLxwxXiBwIz1TWtatqheHh4RPltJ8"
	sheetRange := "Presbies!A2:Q"
	presbies := getSheet(sheetId, sheetRange).Values
	newPresbies := make([]model.Presby, len(presbies))

	for cnt, row := range presbies {
		log.Printf("presby record  #%d", cnt)
		newPresbies[cnt].ID = row[0].(string)
		log.Printf("\tfield #%d -- value: %s", 0, row[0])
		newPresbies[cnt].IsActive = row[1].(string)
		log.Printf("\tfield #%d -- value: %v", 1, row[1])
		newPresbies[cnt].Last = row[2].(string)
		log.Printf("\tfield #%d -- value: %v", 2, row[2].(string))
		newPresbies[cnt].Guests = row[3].(string)
		log.Printf("\tfield #%d -- value: %v", 3, row[3].(string))
		newPresbies[cnt].Guestings = row[4].(string)
		log.Printf("\tfield #%d -- value: %v", 4, row[4].(string))
		newPresbies[cnt].Hostings = row[5].(string)
		log.Printf("\tfield #%d -- value: %v", 5, row[5].(string))
		newPresbies[cnt].Seats = row[6].(string)
		log.Printf("\tfield #%d -- value: %v", 6, row[6].(string))
		newPresbies[cnt].Unknown1 = row[7].(string)
		log.Printf("\tfield #%d -- value: %v", 7, row[7].(string))
		newPresbies[cnt].Unknown2 = row[8].(string)
		log.Printf("\tfield #%d -- value: %v", 8, row[8].(string))
		newPresbies[cnt].Email = row[9].(string)
		log.Printf("\tfield #%d -- value: %v", 9, row[9].(string))
		newPresbies[cnt].Home = row[10].(string)
		log.Printf("\tfield #%d -- value: %v", 10, row[10].(string))
		newPresbies[cnt].Cell = row[11].(string)
		log.Printf("\tfield #%d -- value: %v", 11, row[11].(string))
		newPresbies[cnt].Smail = row[12].(string)
		log.Printf("\tfield #%d -- value: %v", 12, row[12].(string))
		newPresbies[cnt].City = row[13].(string)
		log.Printf("\tfield #%d -- value: %v", 13, row[13].(string))
		newPresbies[cnt].St = row[14].(string)
		log.Printf("\tfield #%d -- value: %v", 14, row[14].(string))
		newPresbies[cnt].Zip = row[15].(string)
		log.Printf("\tfield #%d -- value: %v", 15, row[15].(string))
		newPresbies[cnt].Mmail = row[16].(string)
		log.Printf("\tfield #%d -- value: %v", 16, row[16].(string))
	}
	return newPresbies, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]model.Todo, error) {
	panic(fmt.Errorf("not implemented -- %v", ctx))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func getSheet(spreadsheetId string, readRange string) *sheets.ValueRange {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(".secrets/credentials.json"))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	return resp
}
