package graph

import (
	"context"
	"errors"
	"github.com/Kostaq1234/graphql/graph/generated"
	"github.com/Kostaq1234/graphql/graph/model"
	"github.com/Kostaq1234/graphql/graph/models"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("not long enough")
	}
	if len(input.Description) < 3 {
		return nil, errors.New("not long enough")
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserId:      "1",
	}
	return r.MeetupsRepo.CreateMeetup(meetup)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
