package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	SessionID bson.ObjectId `bson:"_id,omitempty"`
	HostID bson.ObjectId `bson:"supporter_id"`
	UserID bson.ObjectId `bson:"user_id"`
	IsOrdered bool `bson:"is_ordered"`
	CheckInTime time.Time `bson:"check_in_time,omitempty"`
}
