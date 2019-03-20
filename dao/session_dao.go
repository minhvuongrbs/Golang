package dao

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"welcome_robot/models"
)

const (
	SessionCollection = "Session"
)

func InsertSession(session models.Session) error {
	err := ConnectDatabase().C(SessionCollection).Insert(&session)
	return err
}
func RemoveSession(id string) error {
	if bson.IsObjectIdHex(id) {
		err := ConnectDatabase().C(SessionCollection).RemoveId(bson.ObjectIdHex(id))
		return err
	} else {
		err := errors.New("invalid input to ObjectIdHex ")
		return err
	}
}

func GetSessionById(id string) (models.Session, error) {
	var session models.Session
	if bson.IsObjectIdHex(id) {
		err := ConnectDatabase().C(SessionCollection).FindId(bson.ObjectIdHex(id)).One(&session)
		return session, err
	} else {
		err := errors.New("invalid input to ObjectIdHex ")
		return session, err
	}
}

func GetSessionByUserID(userID string) (models.Session, error) {
	var session models.Session
	if bson.IsObjectIdHex(userID) {
		err := ConnectDatabase().C(SessionCollection).Find(bson.M{"user_id": bson.ObjectIdHex(userID)}).One(&session)
		return session, err
	} else {
		err := errors.New("invalid input to ObjectIdHex ")
		return session, err
	}
}
func GetAllSession() ([]models.Session, error) {
	var sessions []models.Session
	err := ConnectDatabase().C(SessionCollection).Find(nil).All(&sessions)
	return sessions, err
}
