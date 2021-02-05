package postgres

import (
	"github.com/Kostaq1234/graphql/graph/models"
	"github.com/go-pg/pg"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil

}
