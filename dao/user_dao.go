package dao

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"welcome_robot/models"
)

const UserCollection = "User"

func GetAllVisitors() ([]models.User, error) {
	var users []models.User
	err := ConnectDatabase().C(UserCollection).Find(bson.M{"permission": 3}).All(&users)
	if err != nil {
		return users, err
	}
	return users, err
}
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

func InsertUser(userInfo models.UserInfo) (models.User, error) {
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
		user.UserID = bson.NewObjectId()
		err := ConnectDatabase().C(UserCollection).Insert(&user)
		return user, err
	} else {
		user.UserID = userInfo.UserID
		err := ConnectDatabase().C(UserCollection).UpdateId(user.UserID, &user)
		return user, err
	}
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
