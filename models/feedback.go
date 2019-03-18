package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Feedback struct {
	FeedbackID bson.ObjectId `bson:"_id,omitempty"`
	Comment    string        `bson:"comment"`
	Rating     int           `bson:"rating"`
	CreatedAT  time.Time     `bson:"created_at,omitempty"`
}
