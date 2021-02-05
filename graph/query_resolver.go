package graph

import (
	"context"
	"github.com/Kostaq1234/graphql/graph/generated"
	"github.com/Kostaq1234/graphql/graph/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {

	return r.MeetupsRepo.GetMeetups()
}
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
