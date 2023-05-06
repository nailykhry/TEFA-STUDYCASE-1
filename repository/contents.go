package repository

import (
	"TEFA-STUDYCASE-1/database"
	"TEFA-STUDYCASE-1/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ContentsCollection = "contents"

type ContentsRepository interface {
	UploadContent(content *models.Content) error
	Update(content *models.Content) error
	GetById(id string) (content *models.Content, err error)
	GetAll() (contents []*models.Content, err error)
	Delete(id string) error
}

type contentsRepository struct {
	c *mgo.Collection
}

func NewContentRepository(conn database.Connection) ContentsRepository {
	return &contentsRepository{conn.DB().C(ContentsCollection)}
}

func (r *contentsRepository) UploadContent(content *models.Content) error {
	return r.c.Insert(content)
}

func (r *contentsRepository) Update(content *models.Content) error {
	return r.c.UpdateId(content.Id, content)
}

func (r *contentsRepository) GetById(id string) (content *models.Content, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&content)
	return content, err
}

func (r *contentsRepository) GetAll() (contents []*models.Content, err error) {
	err = r.c.Find(bson.M{}).All(&contents)
	return contents, err
}

func (r *contentsRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}

// func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
// 	err = r.c.Find(bson.M{"email": email}).One(&user)
// 	return user, err
// }
