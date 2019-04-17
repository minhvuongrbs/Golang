package dao

import (
	"gopkg.in/mgo.v2/bson"
	. "welcomerobot-api/models"
)

func GetAllVideos() ([]Video, error) {
	var videos []Video
	err := ConnectDatabase().C(VideoDaoCollection).Find(bson.M{}).All(&videos)
	return videos, err
}
func InsertVideo(video Video) error {
	err := ConnectDatabase().C(VideoDaoCollection).Insert(&video)
	return err
}
func DeleteVideo(id string) error {
	err := ConnectDatabase().C(VideoDaoCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}
