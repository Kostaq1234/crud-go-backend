package graph

import "github.com/Kostaq1234/graphql/postgres"

type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UserRepo    postgres.UserRepo
}
