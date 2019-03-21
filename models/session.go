package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	SessionID   bson.ObjectId `bson:"_id,omitempty" json:"session_id"`
	SupporterID bson.ObjectId `bson:"supporter_id,omitempty" json:"supporter_id"`
	UserID      bson.ObjectId `bson:"user_id" json:"user_id"`
	CheckInTime time.Time     `bson:"check_in_time,omitempty" json:"check_in_time"`
}
