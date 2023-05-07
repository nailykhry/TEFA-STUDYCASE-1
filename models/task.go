package models

import (
	"time" //O(1)

	"gopkg.in/mgo.v2/bson" //O(1)
)

type Task struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`                  //O(1)
	Title       string        `json:"title,omitempty" bson:"title,omitempty"`             //O(1)
	Description string        `json:"description,omitempty" bson:"description,omitempty"` //O(1)
	ContentId   bson.ObjectId `json:"content_id,omitempty" bson:"content_id,omitempty"`   //O(1)
	CreatedAt   time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"`   //O(1)
	UpdatedAt   time.Time     `json:"updated_at,omitempty" bson:"updated_at,omitempty"`   //O(1)
}
