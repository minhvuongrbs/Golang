package models

import "gopkg.in/mgo.v2/bson"

type Video struct {
	VideoID bson.ObjectId `bson:"_id,omitempty"`
	Title   string        `bson:"title"`
	Image   string        `bson:"image"`
	Link    string        `bson:"link"`
}
