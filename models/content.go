package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Content struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Title     string        `json:"title" bson:"title"`
	Content   string        `json:"content" bson:"content"`
	User_ID   string        `json:"user_id" bson:"user_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
