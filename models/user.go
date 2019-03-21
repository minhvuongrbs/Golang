package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	UserID     bson.ObjectId `bson:"_id,omitempty" json:"user_id"`
	Username   string        `bson:"username,omitempty" json:"username"`
	FirstName  string        `bson:"first_name" json:"first_name"`
	LastName   string        `bson:"last_name,omitempty" json:"last_name"`
	Password   string        `bson:"password,omitempty" json:"password"`
	Data       Data          `bson:"data" json:"data"`
	Permission int           `bson:"permission" json:"permission"`
}
