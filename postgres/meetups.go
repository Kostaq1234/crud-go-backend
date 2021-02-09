package postgres

import (
	"fmt"
	"github.com/Kostaq1234/graphql/graph/models"
	"github.com/go-pg/pg"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups(filter *models.MeetupFilter, limit, offset *int) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	query := m.DB.Model(&meetups).Order("id")
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ?", fmt.Sprintf("%%s%", *filter.Name))
		}
	}
	if limit != nil {
		query.Limit(*limit)
	}
	if offset != nil {
		query.Offset(*offset)
	}
	err := query.Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()
	return meetup, err
}

func (m *MeetupsRepo) GetById(id string) (*models.Meetup, error) {
	var meetup models.Meetup
	err := m.DB.Model(&meetup).Where("id=?", id).First()
	return &meetup, err

}

func (m *MeetupsRepo) Update(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("id=?", meetup.ID).Update()
	return meetup, err
}

func (m *MeetupsRepo) Delete(meetup *models.Meetup) error {
	_, err := m.DB.Model(meetup).Where("id=?", meetup.ID).Delete()
	return err
}

func (m *MeetupsRepo) GetMeetupsForUser(user *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Where("user_id=?", user.ID).Order("id").Select()
	return meetups, err
}
