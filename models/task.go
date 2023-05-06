package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title,omitempty" bson:"title,omitempty"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	ContentId   bson.ObjectId `json:"content_id,omitempty" bson:"content_id,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
