package graph

import (
	"context"
	"github.com/Kostaq1234/graphql/graph/models"
)

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {

	return r.UserRepo.GetUserById(obj.UserId)
}
func (r *Resolver) Meetup() *meetupResolver {

	return &meetupResolver{r}
}

//func (r *Resolver) Meetup() MeetupResolver{
//	return &meetupResolver{r}
//}
