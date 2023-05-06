package repository

import (
	"TEFA-STUDYCASE-1/database"
	"TEFA-STUDYCASE-1/models"

	"gopkg.in/mgo.v2"
)

const usersubsCollection = "usersubs"

type UsersubsRepository interface {
	CreateUsersub(usersub *models.Usersub) error
}

type usersubsRepository struct {
	c *mgo.Collection
}

func NewUsersubRepository(conn database.Connection) UsersubsRepository {
	return &usersubsRepository{conn.DB().C(usersubsCollection)}
}

func (r *usersubsRepository) CreateUsersub(usersub *models.Usersub) error {
	return r.c.Insert(usersub)
}
