package postgres

import (
	"fmt"
	"github.com/Kostaq1234/graphql/graph/models"
	"github.com/go-pg/pg"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUserByField(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v =?", field), value).First()
	return &user, err

}

func (u *UserRepo) GetUserById(id string) (*models.User, error) {
	return u.GetUserByField("id", id)

}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}
func (u *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	return u.GetUserByField("username", username)
}
func (u *UserRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(&user).Returning("*").Insert()
	return user, err

}
