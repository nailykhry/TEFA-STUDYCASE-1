package repository

import (
	"TEFA-STUDYCASE-1/database"
	"TEFA-STUDYCASE-1/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ContentsCollection = "contents" //O(1)

type ContentsRepository interface {
	UploadContent(content *models.Content) error            //O(1)
	Update(content *models.Content) error                   //O(1)
	GetById(id string) (content *models.Content, err error) //O(1)
	GetAll() (contents []*models.Content, err error)        //O(1)
	Delete(id string) error                                 //O(1)
}

type contentsRepository struct {
	c *mgo.Collection //O(1)
}

func NewContentRepository(conn database.Connection) ContentsRepository {
	return &contentsRepository{conn.DB().C(ContentsCollection)} //O(1)
}

func (r *contentsRepository) UploadContent(content *models.Content) error {
	return r.c.Insert(content) //O(1)
}

func (r *contentsRepository) Update(content *models.Content) error {
	return r.c.UpdateId(content.Id, content) //O(1)
}

func (r *contentsRepository) GetById(id string) (content *models.Content, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&content) //O(1)
	return content, err                                  //O(1)
}

func (r *contentsRepository) GetAll() (contents []*models.Content, err error) {
	err = r.c.Find(bson.M{}).All(&contents) //O(n)
	return contents, err                    //O(1)
}

func (r *contentsRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id)) //O(1)
}
