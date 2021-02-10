package domain

import (
	"errors"
	"github.com/Kostaq1234/graphql/graph/models"
	"github.com/Kostaq1234/graphql/postgres"
)

var (
	ForbidenError = errors.New("unauthorised")
)

type Domain struct {
	UserRepo    postgres.UserRepo
	MeetupsRepo postgres.MeetupsRepo
}

func NewDomain(userRepo postgres.UserRepo, meetupsRepo postgres.MeetupsRepo) *Domain {
	return &Domain{UserRepo: userRepo, MeetupsRepo: meetupsRepo}
}

type Owner interface {
	IsOwner(user *models.User) bool
}

func checkOwnership(o Owner, user *models.User) bool {
	return o.IsOwner(user)
}
