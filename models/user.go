package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	UserID    bson.ObjectId `bson:"_id,omitempty"`
	Username  string        `bson:"username"`
	FirstName string        `bson:"first_name"`
	LastName  string        `bson:"last_name"`
	Password  string        `bson:"password"`
	Data      Data          `bson:"data"`
}
