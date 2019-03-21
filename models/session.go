package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	SessionID   bson.ObjectId `bson:"_id,omitempty"`
	SupporterID bson.ObjectId `bson:"supporter_id,omitempty"`
	UserID      bson.ObjectId `bson:"user_id"`
	CheckInTime time.Time     `bson:"check_in_time,omitempty"`
}
