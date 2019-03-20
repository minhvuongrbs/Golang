package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	UserID     bson.ObjectId `json:"_id,omitempty"`
	Username   string        `json:"username,omitempty"`
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name,omitempty"`
	Password   string        `json:"password,omitempty"`
	Data       Data          `json:"data"`
	Permission int           `json:"permission"`
}
