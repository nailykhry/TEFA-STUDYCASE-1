package repository

import (
	"TEFA-STUDYCASE-1/database"
	"TEFA-STUDYCASE-1/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UserchatsCollection = "userchats"

type UserchatsRepository interface {
	CreateUserchat(userchat *models.Userchat) error
	UpdateUserchat(userchat *models.Userchat) error
	GetUserchatById(id string) (userchat *models.Userchat, err error)
	GetAllUserchats() (userchats []*models.Userchat, err error)
}

type userchatsRepository struct {
	c *mgo.Collection
}

func NewTUserchatRepository(conn database.Connection) UserchatsRepository {
	return &userchatsRepository{conn.DB().C(UserchatsCollection)}
}

func (r *userchatsRepository) CreateUserchat(userchat *models.Userchat) error {
	return r.c.Insert(userchat)
}

func (r *userchatsRepository) UpdateUserchat(userchat *models.Userchat) error {
	return r.c.UpdateId(userchat.Id, userchat)
}

func (r *userchatsRepository) GetUserchatById(id string) (userchat *models.Userchat, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&userchat)
	return userchat, err
}

func (r *userchatsRepository) GetAllUserchats() (userchats []*models.Userchat, err error) {
	err = r.c.Find(bson.M{}).All(&userchats)
	return userchats, err
}
