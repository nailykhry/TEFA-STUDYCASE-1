package repository

import (
	"TEFA-STUDYCASE-1/database"
	"TEFA-STUDYCASE-1/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users" //O(1)

type UsersRepository interface {
	Save(user *models.User) error                           //O(1)
	Update(user *models.User) error                         //O(1)
	GetById(id string) (user *models.User, err error)       //O(1)
	GetByEmail(email string) (user *models.User, err error) //O(1)
	GetAll() (users []*models.User, err error)              //O(1)
	Delete(id string) error                                 //O(1)
}

type usersRepository struct {
	c *mgo.Collection //O(1)
}

func NewUsersRepository(conn database.Connection) UsersRepository {
	return &usersRepository{conn.DB().C(UsersCollection)} //O(1)
}

func (r *usersRepository) Save(user *models.User) error {
	return r.c.Insert(user) //O(1)
}

func (r *usersRepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user) //O(1)
}

func (r *usersRepository) GetById(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user) //O(1)
	return user, err
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	err = r.c.Find(bson.M{"email": email}).One(&user) //O(1)
	return user, err
}

func (r *usersRepository) GetAll() (users []*models.User, err error) {
	err = r.c.Find(bson.M{}).All(&users) //O(n)
	return users, err
}

func (r *usersRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id)) //O(1)
}
