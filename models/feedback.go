package models

import "gopkg.in/mgo.v2/bson"

type Feedback struct {
	FeedbackID bson.ObjectId `bson:"_id,omitempty"`
	Comment    string        `bson:"comment"`
	Rating     int           `bson:"rating"`
}
