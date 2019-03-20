package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type VideoTime struct {
	VideoTimeID bson.ObjectId `json:"_id,omitempty"`
	VideoID     bson.ObjectId `json:"video_id"`
	TimeStamp   int           `json:"timestamp"`
	TimeStart   time.Time     `json:"start_time"`
	IsPause     bool          `json:"is_pause"`
	SessionID   bson.ObjectId `json:"session_id"`
}
