package graph

import (
	"context"
	"github.com/Kostaq1234/graphql/graph/generated"
	"github.com/Kostaq1234/graphql/graph/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {

	return r.MeetupsRepo.GetMeetups(filter, limit, offset)
}
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UserRepo.GetUserById(id)
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
