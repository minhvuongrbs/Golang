package dao

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"welcomerobot-api/models"
)

func InsertSession(session models.Session) error {
	err := ConnectDatabase().C(SessionCollection).Insert(&session)
	return err
}
func DeleteSession(id string) error {
	if bson.IsObjectIdHex(id) {
		err := ConnectDatabase().C(SessionCollection).RemoveId(bson.ObjectIdHex(id))
		return err
	} else {
		err := errors.New("invalid input to ObjectIdHex")
		return err
	}
}
func UpdateSession(session models.Session) error {
	err := ConnectDatabase().C(SessionCollection).UpdateId(session.SessionID, &session)
	return err
}
func GetSessionById(id string) (models.Session, error) {
	var session models.Session
	if bson.IsObjectIdHex(id) {
		err := ConnectDatabase().C(SessionCollection).FindId(bson.ObjectIdHex(id)).One(&session)
		return session, err
	} else {
		err := errors.New("invalid input to ObjectIdHex")
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
func GetAllSessions() ([]models.Session, error) {
	var sessions []models.Session
	err := ConnectDatabase().C(SessionCollection).Find(bson.M{}).All(&sessions)
	return sessions, err
}
