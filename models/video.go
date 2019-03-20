package models

import "gopkg.in/mgo.v2/bson"

type Video struct {
	VideoID bson.ObjectId `json:"_id,omitempty"`
	Title   string        `json:"title"`
	Image   string        `json:"image"`
	Link    string        `json:"link"`
}
