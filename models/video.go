package models

import "gopkg.in/mgo.v2/bson"

type Video struct {
	VideoID bson.ObjectId `bson:"_id,omitempty" json:"video_id"`
	Title   string        `bson:"title" json:"title"`
	Image   string        `bson:"image" json:"image"`
	Link    string        `bson:"link" json:"link"`
}
