package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Feedback struct {
	FeedbackID bson.ObjectId `bson:"_id,omitempty" json:"feedback_id"`
	Comment    string        `bson:"comment" json:"comment"`
	Rating     int           `bson:"rating" json:"rating"`
	CreatedAT  time.Time     `bson:"created_at,omitempty" json:"created_at"`
}
