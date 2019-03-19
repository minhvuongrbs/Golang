package dao

import (
	"gopkg.in/mgo.v2/bson"
	. "welcome_robot/models"
)

const videoDaoCollection = "Video"

func FindAllVideo() ([]Video, error) {
	var videos []Video
	err := ConnectDatabase().C(videoDaoCollection).Find(bson.M{}).All(&videos)
	return videos, err
}
func InsertVideo(video Video) error {
	err := ConnectDatabase().C(videoDaoCollection).Insert(&video)
	return err
}
func DeleteVideo(id string) error {
	err := ConnectDatabase().C(videoDaoCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}
