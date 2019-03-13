package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	. "welcome_robot/models"
)

type VideoDAO struct {
	Server string
	Database string
	db *mgo.Database
	collection string
}

var videoDao VideoDAO

func init() {
	videoDao.collection= "Video"
}

func (m *VideoDAO) Connect() {
	session, err :=mgo.Dial(m.Server)
	if err!=nil{
		log.Fatal(err)
	}
	m.db = session.DB(m.Database)
}

func (m *VideoDAO) FindAll() ([]Video, error) {
	var videos []Video
	err := m.db.C(videoDao.collection).Find(bson.M{}).All(&videos)
	return videos, err
}
func (m *VideoDAO) Insert(video Video) error {
	err := m.db.C(videoDao.collection).Insert(&video)
	return err
}
func (m *VideoDAO) Delete(id string) error{
	err:=m.db.C(videoDao.collection).RemoveId(bson.ObjectIdHex(id))
	return err
}