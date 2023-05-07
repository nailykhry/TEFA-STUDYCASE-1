package models

import (
	"time" //O(1)

	"gopkg.in/mgo.v2/bson" //O(1)
)

type Usersub struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`                //O(1)
	Plan      string        `json:"plan,omitempty" bson:"plan,omitempty"`             //O(1)
	Price     int           `json:"price,omitempty" bson:"price,omitempty"`           //O(1)
	UserId    bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`       //O(1)
	CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"` //O(1)
}
