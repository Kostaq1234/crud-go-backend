package graph

import "github.com/Kostaq1234/graphql/domain"

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

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
//func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
//	_, err := r.UserRepo.GetUserByEmail(input.Email)
//	if err != nil {
//		return nil, errors.New("email alredy in use")
//	}
//	_, err = r.UserRepo.GetUserByUsername(input.Username)
//	if err != nil {
//		return nil, errors.New("username alredy in use")
//	}
//	user := &models.User{
//		Username:  input.Username,
//		Email:     input.Email,
//		FirstName: input.FirstName,
//		LastName:  input.LastName,
//	}
//	err = user.HashPassword(input.Password)
//	if err != nil {
//		return nil, errors.New("meetup not exist")
//	}
//	//send verification code
//	tx, err := r.UserRepo.DB.Begin()
//	if err != nil {
//		return nil, errors.New("smth went wrong")
//	}
//	defer tx.Rollback()
//	r.UserRepo.CreateUser(tx, user)
//	if err != nil {
//		return nil, err
//	}
//	if err := tx.Commit(); err != nil {
//		return nil, err
//	}
//	token, err := user.GenToken()
//	if err != nil {
//		return nil, err
//	}
//	return &models.AuthResponse{
//		AuthToken: token,
//		User:      user,
//	}, nil
//}
//
//func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
//	panic(fmt.Errorf("not implemented"))
//}
//
//func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
//	return r.MeetupsRepo.GetMeetups(filter, limit, offset)
//}
//
//func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
//	return r.UserRepo.GetUserById(id)
//}
//
//func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
//	return u.MeetupsRepo.GetMeetupsForUser(obj)
//}
//
//// Mutation returns generated.MutationResolver implementation.
//func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
//
//// Query returns generated.QueryResolver implementation.
//func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
//
//// User returns generated.UserResolver implementation.
//func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }
//
//type mutationResolver struct{ *Resolver }
//type queryResolver struct{ *Resolver }
//type userResolver struct{ *Resolver }
//
//// !!! WARNING !!!
//// The code below was going to be deleted when updating resolvers. It has been copied here so you have
//// one last chance to move it out of harms way if you want. There are two reasons this happens:
////  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
////    it when you're done.
////  - You have helper methods in this file. Move them out to keep these resolver files clean.
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
//func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
//	_, err := r.UserRepo.GetUserByEmail(input.Email)
//	if err != nil {
//		return nil, errors.New("email alredy in use")
//	}
//	_, err = r.UserRepo.GetUserByUsername(input.Username)
//	if err != nil {
//		return nil, errors.New("username alredy in use")
//	}
//	user := &models.User{
//		Username:  input.Username,
//		Email:     input.Email,
//		FirstName: input.FirstName,
//		LastName:  input.LastName,
//	}
//	err = user.HashPassword(input.Password)
//	if err != nil {
//		return nil, errors.New("meetup not exist")
//	}
//	//send verification code
//	tx, err := r.UserRepo.DB.Begin()
//	if err != nil {
//		return nil, errors.New("smth went wrong")
//	}
//	defer tx.Rollback()
//	r.UserRepo.CreateUser(tx, user)
//	if err != nil {
//		return nil, err
//	}
//	if err := tx.Commit(); err != nil {
//		return nil, err
//	}
//	token, err := user.GenToken()
//	if err != nil {
//		return nil, err
//	}
//	return &models.AuthResponse{
//		AuthToken: token,
//		User:      user,
//	}, nil
//}
//func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
//	return r.MeetupsRepo.GetMeetups(filter, limit, offset)
//}
//func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
//	return r.UserRepo.GetUserById(id)
//}
//func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
//func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type Resolver struct {
	Domain *domain.Domain
}
