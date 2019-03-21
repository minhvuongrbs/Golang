package dao

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	DatabaseHost = "localhost:27017"
	DatabaseName = "welcome_robot"
)

func ConnectDatabase() *mgo.Database {
	session, err := mgo.Dial(DatabaseHost)
	if err != nil {
		log.Fatal(err.Error())
	}
	if session.Ping() != nil {
		log.Fatal(err)
	}
	return session.DB(DatabaseName)
}

//func ConnectDatabase() *mgo.Database {
//	tlsConfig := &tls.Config{}
//
//	dialInfo := &mgo.DialInfo{
//		Addrs: []string{"mongodb+srv://welcome_robot:<welcome_robot>@vuongtran-k4iil.gcp.mongodb.net/test?retryWrites=true"},
//		Database: "welcome_robot",
//		Username: "welcome_robot",
//		Password: "welcome_robot",
//	}
//	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
//		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
//		return conn, err
//	}
//	session, err := mgo.DialWithInfo(dialInfo)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	if session.Ping() != nil {
//		log.Fatal(err)
//	}
//	return session.DB(DatabaseName)
//}
