package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Usersub struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Plan      string        `json:"plan,omitempty" bson:"plan,omitempty"`
	Price     int           `json:"price,omitempty" bson:"price,omitempty"`
	UserId    bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
