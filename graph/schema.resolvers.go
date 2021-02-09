package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/Kostaq1234/graphql/postgres"
)

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UserRepo    postgres.UserRepo
}
