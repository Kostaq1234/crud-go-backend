package domain

import (
	"context"
	"errors"
	"github.com/Kostaq1234/graphql/graph/models"
)

func (r *Domain) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := r.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("email alredy in use")
	}
	_, err = r.UserRepo.GetUserByUsername(input.Username)
	if err != nil {
		return nil, errors.New("username alredy in use")
	}
	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	err = user.HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("meetup not exist")
	}
	//send verification code
	tx, err := r.UserRepo.DB.Begin()
	if err != nil {
		return nil, errors.New("smth went wrong")
	}
	defer tx.Rollback()
	r.UserRepo.CreateUser(tx, user)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	token, err := user.GenToken()
	if err != nil {
		return nil, err
	}
	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
func (r *Domain) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := r.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("email or password not working")
	}
	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, errors.New("email or password not working")
	}
	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("smth not working")
	}
	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
