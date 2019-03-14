package dao

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"welcome_robot/models"
)

const (
	DatabaseHost = "localhost:27017"
	DatabaseName = "welcome_robot"
	SessionCollection = "Session"
)

func ConnectDatabase() *mgo.Session {
	session, err := mgo.Dial(DatabaseHost)
	if err != nil {
		log.Fatal(err.Error())
	}
	if session.Ping() != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected")
	return session
}
func InsertSession(session models.Session) error {
	err:=ConnectDatabase().DB(DatabaseName).C(SessionCollection).Insert(&session)
	return err
}

func GetAllSessions() ([]models.Session, error) {
	collection := ConnectDatabase().DB(DatabaseName).C(SessionCollection)
	var sessions []models.Session
	err := collection.Find(nil).All(&sessions)
	if err != nil {
		return sessions, err
	}
	return sessions, err
}

func GetUserByUserID(userID string) (models.User, error) {
	var user models.User
	err := ConnectDatabase().DB(DatabaseName).C(SessionCollection).FindId(bson.ObjectIdHex(userID)).One(&user)
	return user, err
}
