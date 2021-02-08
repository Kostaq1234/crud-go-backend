package graph

import (
	"context"
	"errors"
	"github.com/Kostaq1234/graphql/graph/generated"
	"github.com/Kostaq1234/graphql/graph/model"
	"github.com/Kostaq1234/graphql/graph/models"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup does not exist")
	}
	err = r.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, err
	}
	return true, err

}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := r.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}
	didUpdate := false
	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}
	if !didUpdate {
		return nil, err
	}
	meetup, err = r.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, err
	}
	return meetup, nil
}

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
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
