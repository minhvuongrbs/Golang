package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	. "welcome_robot/models"
)

type VideoTimeDAO struct {
	Server     string
	Database   string
	db         *mgo.Database
	Collection string
}

var videoTimeDAO VideoTimeDAO

func init() {
	videoTimeDAO.Collection = "VideoTime"
}

func (m *VideoTimeDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	m.db = session.DB(m.Database)
}

func (m *VideoTimeDAO) FindById(id string) (VideoTime, error) {
	var videoTime VideoTime
	err := m.db.C(videoTimeDAO.Collection).Find(bson.M{}).One(&videoTime)
	return videoTime, err
}
func (m *VideoTimeDAO) FindAll() ([]VideoTime, error) {
	var videoTimes []VideoTime
	err := m.db.C(videoTimeDAO.Collection).Find(bson.M{}).All(&videoTimes)
	return videoTimes, err
}
func (m *VideoTimeDAO) Insert(videoTime VideoTime) error {
	err := m.db.C(videoTimeDAO.Collection).Insert(&videoTime)
	return err
}
func (m *VideoTimeDAO) Update(videoTime VideoTime) error {
	err := m.db.C(m.Collection).UpdateId(videoTime.VideoTimeID, &videoTime)
	return err
}
func (m *VideoTimeDAO) Delete(id string) error{
	err:=m.db.C(m.Collection).RemoveId(id)
	return err
}

