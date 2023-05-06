package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Userchat struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title,omitempty" bson:"title,omitempty"`
	Question  string        `json:"question,omitempty" bson:"question,omitempty"`
	Keyword   string        `json:"keyword,omitempty" bson:"keyword,omitempty"`
	Status    bool          `json:"status,omitempty" bson:"status,omitempty"`
	UserId    bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
