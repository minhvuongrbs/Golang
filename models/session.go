package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	SessionID   bson.ObjectId `json:"_id,omitempty"`
	SupporterID bson.ObjectId `json:"supporter_id,omitempty"`
	UserID      bson.ObjectId `json:"user_id"`
	CheckInTime time.Time     `json:"check_in_time,omitempty"`
}
