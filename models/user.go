package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Email        string        `json:"email" bson:"email"`
	Nama         string        `json:"nama" bson:"nama"`
	Nama_Sekolah string        `json:"nama_sekolah" bson:"nama_sekolah"`
	Telp         string        `json:"telp" bson:"telp"`
	Role         string        `json:"role" bson:"role"`
	Password     string        `json:"password" bson:"password"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`
}
