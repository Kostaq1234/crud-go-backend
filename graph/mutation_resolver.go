package graph

import (
	"context"
	"errors"
	"github.com/Kostaq1234/graphql/graph/generated"
	"github.com/Kostaq1234/graphql/graph/models"
)

var (
	ErrInput = errors.New("input error")
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.Login(ctx, input)
}

func (r *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.Register(ctx, input)
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return r.Domain.DeleteMeetup(ctx, id)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	return r.Domain.UpdateMeetup(ctx, id, input)
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	return r.Domain.CreateMeetup(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
