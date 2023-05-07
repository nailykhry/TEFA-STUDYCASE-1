package repository

import (
	"TEFA-STUDYCASE-1/database" //O(1)
	"TEFA-STUDYCASE-1/models"   //O(1)

	"gopkg.in/mgo.v2" //O(1)
)

const usersubsCollection = "usersubs" //O(1)

type UsersubsRepository interface {
	CreateUsersub(usersub *models.Usersub) error //O(1)
}

type usersubsRepository struct {
	c *mgo.Collection //O(1)
}

func NewUsersubRepository(conn database.Connection) UsersubsRepository {
	return &usersubsRepository{conn.DB().C(usersubsCollection)} //O(1)
}

func (r *usersubsRepository) CreateUsersub(usersub *models.Usersub) error {
	return r.c.Insert(usersub) //O(1)
}
