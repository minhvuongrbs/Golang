package dao

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"welcome_robot/models"
)

const (
	DatabaseHost      = "localhost:27017"
	DatabaseName      = "welcome_robot"
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
	err := ConnectDatabase().DB(DatabaseName).C(SessionCollection).Insert(&session)
	return err
}
func RemoveSession(id string) error  {
	err:= ConnectDatabase().DB(DatabaseName).C(SessionCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}

func GetSessionById(id string) (models.Session, error) {
	var session models.Session
	err := ConnectDatabase().DB(DatabaseName).C(SessionCollection).FindId(bson.ObjectIdHex(id)).One(&session)
	return session, err
}

func GetAllSession() ([]models.Session, error) {
	var sessions []models.Session
	query := []bson.M{
		//bson.M{
		//	"$match": bson.M{
		//		"_id": bson.ObjectIdHex("56b9df0c1e930a99cb2c33e9")}},
		bson.M{"$lookup": bson.M{
			"from": "user",
			"localField": "user_id",
			"foreignField": "_id",
			"as": "user"}},
	}
	pipe := ConnectDatabase().DB(DatabaseName).C(SessionCollection).Pipe(query)
	resp := []bson.M{}
	err := pipe.All(&resp)
	if err != nil {
		fmt.Println("Errored: %#v \n", err)
	}
	log.Print(resp)
	fmt.Println(resp)
	//err:=ConnectDatabase().DB(DatabaseName).C(SessionCollection).Find()
	return sessions, err
}

//func GetAllSessions() ([]models.Session, error) {
//	collection := ConnectDatabase().DB(DatabaseName).C(SessionCollection)
//	var sessions []models.Session
//	err := collection.Find(nil).All(&sessions)
//	if err != nil {
//		return sessions, err
//	}
//	return sessions, err
//}

func GetSessionByUserID(userID string) (models.Session, error) {
	var session models.Session
	err := ConnectDatabase().DB(DatabaseName).C(SessionCollection).Find(bson.M{"user_id":bson.ObjectIdHex(userID)}).One(&session)
	return session, err
}
