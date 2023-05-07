package models

import (
	"time" //O(1)

	"gopkg.in/mgo.v2/bson" //O(1)
)

type User struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`                    //O(1)
	Email        string        `json:"email" bson:"email"`               //O(1)
	Nama         string        `json:"nama" bson:"nama"`                 //O(1)
	Nama_Sekolah string        `json:"nama_sekolah" bson:"nama_sekolah"` //O(1)
	Telp         string        `json:"telp" bson:"telp"`                 //O(1)
	Role         string        `json:"role" bson:"role"`                 //O(1)
	Password     string        `json:"password" bson:"password"`         //O(1)
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`     //O(1)
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`     //O(1)
}
