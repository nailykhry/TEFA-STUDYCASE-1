package models

import (
	"time" //O(1)

	"gopkg.in/mgo.v2/bson" //O(1)
)

type Userchat struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`                //O(1)
	Title     string        `json:"title,omitempty" bson:"title,omitempty"`           //O(1)
	Question  string        `json:"question,omitempty" bson:"question,omitempty"`     //O(1)
	Keyword   string        `json:"keyword,omitempty" bson:"keyword,omitempty"`       //O(1)
	Status    bool          `json:"status,omitempty" bson:"status,omitempty"`         //O(1)
	UserId    bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`       //O(1)
	CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"` //O(1)
	UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at,omitempty"` //O(1)
}
