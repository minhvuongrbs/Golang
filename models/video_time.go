package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type VideoTime struct {
	VideoTimeID bson.ObjectId `bson:"_id,omitempty"`
	VideoID     bson.ObjectId `bson:"video_id"`
	TimeStamp   int           `bson:"time_stamp"`
	TimeStart   time.Time     `bson:"time_start"`
	IsPause     bool          `bson:"is_pause"`
}
