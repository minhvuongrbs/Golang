package dao

import (
	"gopkg.in/mgo.v2/bson"
	. "welcome_robot/models"
)

const VideoTimeCollection = "VideoTime"

func FindVideoTimeById(id string) (VideoTime, error) {
	var videoTime VideoTime
	err := ConnectDatabase().C(VideoTimeCollection).FindId(bson.ObjectIdHex(id)).One(&videoTime)
	return videoTime, err
}
func FindAllVideoTime() ([]VideoTime, error) {
	var videoTimes []VideoTime
	err := ConnectDatabase().C(VideoTimeCollection).Find(bson.M{}).All(&videoTimes)
	return videoTimes, err
}
func InsertVideoTime(videoTime VideoTime) error {
	err := ConnectDatabase().C(VideoTimeCollection).Insert(&videoTime)
	return err
}
func UpdateVideoTime(videoTime VideoTime) error {
	err := ConnectDatabase().C(VideoTimeCollection).UpdateId(videoTime.VideoTimeID, &videoTime)
	return err
}
func DeleteVideoTime(id string) error {
	err := ConnectDatabase().C(VideoTimeCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}
func DeleteVideoTimeBySessionId(id string) error {
	err := ConnectDatabase().C(VideoTimeCollection).Remove(bson.M{"session_id": bson.ObjectIdHex(id)})
	return err
}
