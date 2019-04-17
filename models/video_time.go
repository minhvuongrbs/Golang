package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type VideoTime struct {
	VideoTimeID bson.ObjectId `bson:"_id,omitempty" json:"video_time_id"`
	VideoID     bson.ObjectId `bson:"video_id" json:"video_id"`
	TimeStamp   int           `bson:"timestamp" json:"time_stamp"`
	TimeStart   time.Time     `bson:"start_time" json:"time_start"`
	IsPause     bool          `bson:"is_pause" json:"is_pause"`
	SessionID   bson.ObjectId `bson:"session_id" json:"session_id"`
}
