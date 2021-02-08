package graph

import (
	"context"
	"github.com/Kostaq1234/graphql/graph/models"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() *userResolver {

	return &userResolver{r}
}
func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return u.MeetupsRepo.GetMeetupsForUser(obj)
}
