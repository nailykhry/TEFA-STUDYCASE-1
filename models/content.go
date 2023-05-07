package models

import (
	"time" //O(1)

	"gopkg.in/mgo.v2/bson" //O(1)
)

type Content struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`                //O(1)
	Title     string        `json:"title" bson:"title"`           //O(1)
	Content   string        `json:"content" bson:"content"`       //O(1)
	User_ID   string        `json:"user_id" bson:"user_id"`       //O(1)
	CreatedAt time.Time     `json:"created_at" bson:"created_at"` //O(1)
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"` //O(1)
}
