package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Feedback struct {
	FeedbackID bson.ObjectId `json:"_id,omitempty"`
	Comment    string        `json:"comment"`
	Rating     int           `json:"rating"`
	CreatedAT  time.Time     `json:"created_at,omitempty"`
}
