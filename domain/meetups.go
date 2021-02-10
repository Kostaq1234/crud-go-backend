package domain

import (
	"context"
	"errors"
	"github.com/Kostaq1234/graphql/graph/models"
	"github.com/Kostaq1234/graphql/middleware"
)

func (r *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return false, ForbidenError
	}
	meetup, err := r.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup does not exist")
	}

	if !meetup.IsOwner(currentUser) {
		return false, ForbidenError
	}
	err = r.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, err
	}
	return true, err

}

func (r *Domain) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, ForbidenError
	}
	meetup, err := r.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}
	if !meetup.IsOwner(currentUser) {
		return nil, ForbidenError
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

func (r *Domain) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errors.New("user not authenticated")
	}

	if len(input.Name) < 3 {
		return nil, errors.New("not long enough")
	}
	if len(input.Description) < 3 {
		return nil, errors.New("not long enough")
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserId:      currentUser.ID,
	}
	return r.MeetupsRepo.CreateMeetup(meetup)
}
