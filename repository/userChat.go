package repository

import (
	"TEFA-STUDYCASE-1/database" // O(1)
	"TEFA-STUDYCASE-1/models"   // O(1)

	"gopkg.in/mgo.v2"      // O(1)
	"gopkg.in/mgo.v2/bson" // O(1)
)

const UserchatsCollection = "userchats" // O(1)

type UserchatsRepository interface {
	CreateUserchat(userchat *models.Userchat) error                   // O(1)
	UpdateUserchat(userchat *models.Userchat) error                   // O(1)
	GetUserchatById(id string) (userchat *models.Userchat, err error) // O(1)
	GetAllUserchats() (userchats []*models.Userchat, err error)       // O(1)
}

type userchatsRepository struct {
	c *mgo.Collection // O(1)
}

func NewTUserchatRepository(conn database.Connection) UserchatsRepository {
	return &userchatsRepository{conn.DB().C(UserchatsCollection)} // O(1)
}

func (r *userchatsRepository) CreateUserchat(userchat *models.Userchat) error {
	return r.c.Insert(userchat) // O(1)
}

func (r *userchatsRepository) UpdateUserchat(userchat *models.Userchat) error {
	return r.c.UpdateId(userchat.Id, userchat) // O(1)
}

func (r *userchatsRepository) GetUserchatById(id string) (userchat *models.Userchat, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&userchat) // O(1)
	return userchat, err
}

func (r *userchatsRepository) GetAllUserchats() (userchats []*models.Userchat, err error) {
	err = r.c.Find(bson.M{}).All(&userchats) // O(n)
	return userchats, err
}
