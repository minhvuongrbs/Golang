package dao

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"welcome_robot/models"
)

const UserCollection = "User"

func FindUserById(id string) (models.User, error) {
	var user models.User
	err := ConnectDatabase().C(UserCollection).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func FindUserByName(name string) (models.User, error) {
	var user models.User
	err := ConnectDatabase().C(UserCollection).Find(bson.M{"full_name": name}).One(&user)
	return user, err
}

func RemoveUser(id string) error {
	err := ConnectDatabase().C(UserCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}

func InsertUser(userInfo models.UserInfo) (bool, models.User, error) {
	var user models.User
	if userInfo.Language == "vi" {
		user.FirstName, user.LastName = HandleNameInVi(userInfo.Name)
	} else {
		user.FirstName, user.LastName = HandleNameInEng(userInfo.Name)
	}
	user.FullName = userInfo.Name
	user.Username = userInfo.UserName
	user.Password = userInfo.Password
	user.Data.Avatar = userInfo.Avatar
	user.Data.Description = userInfo.Description
	user.Data.HierarchyName = userInfo.HierarchyName
	user.Permission = userInfo.Permission
	if userInfo.UserID.Hex() == "" {
		//insert user=> true
		user.UserID = bson.NewObjectId()
		err := ConnectDatabase().C(UserCollection).Insert(&user)
		return true, user, err
	} else {
		//update user=>false
		user.UserID = userInfo.UserID
		err := ConnectDatabase().C(UserCollection).UpdateId(user.UserID, &user)
		return false, user, err
	}
}

func HandleNameInVi(fullName string) (string, string) {
	var firstName, lastName string
	if strings.Contains(fullName, " ") {
		index := strings.Index(fullName, " ")
		lastIndex := strings.LastIndex(fullName, " ")
		firstName = fullName[0:index]
		lastName = fullName[lastIndex+1 : (len(fullName))]
	} else {
		lastName = fullName
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
	} else {
		firstName = fullName
	}
	return firstName, lastName
}
