package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"
	"strings"

	"weja.us/micro-cosm/micro-server-go-graphql/graph/generated"
	"weja.us/micro-cosm/micro-server-go-graphql/graph/model"
	"weja.us/micro-cosm/micro-server-go-graphql/lib"
)

func (r *queryResolver) Presbies(ctx context.Context) ([]model.Presby, error) {
	buildPresbies := make([]model.Presby, len(lib.PresbyData))
	for cnt, field := range lib.PresbyData {

		log.Printf("presby record  #%d", cnt)

		tmpId, e := strconv.ParseUint(field[0].(string), 10, 64)
		if e != nil {
			log.Printf("FAIL!!! ID")
		}
		buildPresbies[cnt].ID = int(tmpId)
		if field[1] == "TRUE" {
			buildPresbies[cnt].IsActive = true
		} else {
			buildPresbies[cnt].IsActive = false
		}
		buildPresbies[cnt].Last = field[2].(string)
		buildPresbies[cnt].Guests = strings.Split(field[3].(string), ",")
		buildPresbies[cnt].Guestings = strings.Split(field[4].(string), ",")
		buildPresbies[cnt].Hostings = strings.Split(field[5].(string), ",")
		tmpSeats, e := strconv.ParseUint(field[6].(string), 10, 64)
		if e != nil {
			log.Printf("FAIL!!! seats")
		}
		buildPresbies[cnt].Seats = int(tmpSeats)
		tmpSubs, e := strconv.ParseUint(field[7].(string), 10, 64)
		if e != nil {
			log.Printf("FAIL!!! seats")
		}
		buildPresbies[cnt].Subs = int(tmpSubs)
		tmpSteps, e := strconv.ParseUint(field[8].(string), 10, 64)
		if e != nil {
			log.Printf("FAIL!!! seats")
		}
		buildPresbies[cnt].Steps = int(tmpSteps)
		buildPresbies[cnt].Email = field[9].(string)
		buildPresbies[cnt].Home = field[10].(string)
		buildPresbies[cnt].Cell = field[11].(string)
		buildPresbies[cnt].Smail = field[12].(string)
		buildPresbies[cnt].City = field[13].(string)
		buildPresbies[cnt].St = field[14].(string)
		buildPresbies[cnt].Zip = field[15].(string)
		buildPresbies[cnt].Key = buildPresbies[cnt].Last + "-" +
			strconv.Itoa(buildPresbies[cnt].ID) + "-" +
			strconv.Itoa(buildPresbies[cnt].Seats) + "-" +
			strconv.Itoa(len(buildPresbies[cnt].Guests))
	}
	return buildPresbies, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
