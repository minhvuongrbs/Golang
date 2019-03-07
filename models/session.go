package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	SessionID bson.ObjectId `bson:"_id,omitempty"`
	HostID User `bson:"host_id"`
	UserID User `bson:"user_id"`
	CheckInTime time.Time `bson:"check_in_time"`
	VideoTimeID []VideoTime `bson:"video_time_id"`
}
