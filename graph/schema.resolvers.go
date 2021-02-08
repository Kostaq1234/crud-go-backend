package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/Kostaq1234/graphql/postgres"
)

//func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
//	if len(input.Name) < 3 {
//		return nil, errors.New("not long enough")
//	}
//	if len(input.Description) < 3 {
//		return nil, errors.New("not long enough")
//	}
//	meetup := &models.Meetup{
//		Name:        input.Name,
//		Description: input.Description,
//		UserId:      "1",
//	}
//	return r.MeetupsRepo.CreateMeetup(meetup)
//}
//
//func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*models.Meetup, error) {
//	meetup, err := r.MeetupsRepo.GetById(id)
//	if err != nil || meetup == nil {
//		return nil, errors.New("meetup does not exist")
//	}
//	didUpdate := false
//	if input.Name != nil {
//		if len(*input.Name) < 3 {
//			return nil, errors.New("name is not long enough")
//		}
//		meetup.Name = *input.Name
//		didUpdate = true
//	}
//	if input.Description != nil {
//		if len(*input.Description) < 3 {
//			return nil, errors.New("name is not long enough")
//		}
//		meetup.Description = *input.Description
//		didUpdate = true
//	}
//	if !didUpdate {
//		return nil, err
//	}
//	meetup, err = r.MeetupsRepo.Update(meetup)
//	if err != nil {
//		return nil, err
//	}
//	return meetup, nil
//}
//
//func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
//	meetup, err := r.MeetupsRepo.GetById(id)
//	if err != nil || meetup == nil {
//		return false, errors.New("meetup does not exist")
//	}
//	err = r.MeetupsRepo.Delete(meetup)
//	if err != nil {
//		return false, err
//	}
//	return true, err
//}
//
//func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
//	return r.MeetupsRepo.GetMeetups()
//}
//
//func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
//	panic(fmt.Errorf("not implemented"))
//}
//
//// Mutation returns generated.MutationResolver implementation.
//func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
//
//// Query returns generated.QueryResolver implementation.
//func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

//type mutationResolver struct{ *Resolver }
//type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UserRepo    postgres.UserRepo
}
