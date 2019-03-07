package dao

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"welcome_robot/models"
)

const UserCollection = "User"

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := ConnectDatabase().DB(DatabaseName).C(UserCollection).Find(nil).All(&users)
	if err != nil {
		return users, err
	}
	return users, err
}

func InsertUser(username string, password string, avatar string, discription string, language string) {
	var firstName, lastName string
	if language == "vi" {
		firstName, lastName = HandleNameInVi(username)
	} else {
		firstName, lastName = HandleNameInEng(username)
	}
	user := models.User{
		UserID:    bson.NewObjectId(),
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}
	_ = ConnectDatabase().DB(DatabaseName).C(UserCollection).Insert(user)

}

func HandleNameInVi(fullName string) (string, string) {
	var firstName, lastName string
	if strings.Contains(fullName, " ") {
		index := strings.Index(fullName, " ")
		lastIndex := strings.LastIndex(fullName, " ")
		firstName = fullName[0:index]
		lastName = fullName[lastIndex+1 : (len(fullName))]
	}
	return firstName, lastName
}

func HandleNameInEng(fullName string) (string, string) {
	var firstName, lastName string
	if strings.Contains(fullName, " ") {
		index := strings.Index(fullName, " ")
		lastIndex := strings.LastIndex(fullName, " ")
		lastName = fullName[0:index]
		firstName = fullName[lastIndex+1 : (len(fullName))]
	}
	return firstName, lastName
}
