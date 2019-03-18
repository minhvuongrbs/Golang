package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type VideoTime struct {
	VideoTimeID bson.ObjectId `bson:"_id,omitempty"`
	VideoID     bson.ObjectId `bson:"video_id"`
	TimeStamp   int           `bson:"timestamp"`
	TimeStart   time.Time     `bson:"start_time"`
	IsPause     bool          `bson:"is_pause"`
	SessionID   bson.ObjectId `bson:"session_id"`

}
