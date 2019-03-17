package dao

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"welcome_robot/models"
)

const UserCollection = "User"

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := ConnectDatabase().DB(DatabaseName).C(UserCollection).Find(nil).Select(bson.M{"permission":3}).All(&users)
	if err != nil {
		return users, err
	}
	return users, err
}

func  RemoveUser(id string) error{
	err:=ConnectDatabase().DB(DatabaseName).C(UserCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}

//func InsertUser(username string, password string, avatar string, discription string, language string) {
//	var firstName, lastName string
//	if language == "vi" {
//		firstName, lastName = HandleNameInVi(username)
//	} else {
//		firstName, lastName = HandleNameInEng(username)
//	}
//	user := models.User{
//		UserID:    bson.NewObjectId(),
//		Username:  username,
//		FirstName: firstName,
//		LastName:  lastName,
//		Password:  password,
//	}
//	_ = ConnectDatabase().DB(DatabaseName).C(UserCollection).Insert(user)
//
//}
func InsertUser(userInfor models.UserInfo) (models.User, error) {
	var user models.User
	user.UserID = bson.NewObjectId()
	if userInfor.Language== "vi" {
		user.FirstName, user.LastName= HandleNameInVi(userInfor.Name)
	} else {
		user.FirstName, user.LastName= HandleNameInEng(userInfor.Name)
	}
	user.Username =userInfor.Name
	user.Password = userInfor.Password
	user.Data.Avatar= userInfor.Avatar
	user.Data.Description =userInfor.Description
	user.Data.HierarchyName = userInfor.HierarchyName
	user.Permission= userInfor.Permission
	err :=ConnectDatabase().DB(DatabaseName).C(UserCollection).Insert(&user)
	return user,err
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
