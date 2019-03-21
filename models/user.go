package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	UserID     bson.ObjectId `bson:"_id,omitempty"`
	Username   string        `bson:"username,omitempty"`
	FirstName  string        `bson:"first_name"`
	LastName   string        `bson:"last_name,omitempty"`
	Password   string        `bson:"password,omitempty"`
	Data       Data          `bson:"data"`
	Permission int           `bson:"permission"`
}
