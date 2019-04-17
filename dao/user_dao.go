package dao

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"welcomerobot-api/models"
)

func GetUserById(id string) (models.User, error) {
	var user models.User
	err := ConnectDatabase().C(UserCollection).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}
func GetUserByName(name string) (models.User, error) {
	var user models.User
	err := ConnectDatabase().C(UserCollection).Find(bson.M{"full_name": name}).One(&user)
	return user, err
}
func DeleteUser(id string) error {
	err := ConnectDatabase().C(UserCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}
func InsertUser(userInfo models.UserInfo) (bool, models.User, error) {
	var user models.User
	if userInfo.Language == "vi" {
		user.LastName, user.FirstName = handleNameByLanguage(userInfo.Name)
	} else {
		user.FirstName, user.LastName = handleNameByLanguage(userInfo.Name)
	}
	user.FullName = userInfo.Name
	user.Username = userInfo.UserName
	user.Password = userInfo.Password
	user.Data.Avatar = userInfo.Avatar
	user.Data.Description = userInfo.Description
	user.Data.HierarchyName = userInfo.HierarchyName
	user.Permission = userInfo.Permission
	if userInfo.UserID.Hex() == "" {
		//insert user
		user.UserID = bson.NewObjectId()
		err := ConnectDatabase().C(UserCollection).Insert(&user)
		return true, user, err
	} else {
		//update user
		user.UserID = userInfo.UserID
		err := ConnectDatabase().C(UserCollection).UpdateId(user.UserID, &user)
		return false, user, err
	}
}
func handleNameByLanguage(fullName string) (string, string) {
	var firstName, lastName string
	if strings.Contains(fullName, " ") {
		index := strings.Index(fullName, " ")
		lastIndex := strings.LastIndex(fullName, " ")
		firstName = fullName[0:index]
		lastName = fullName[lastIndex+1 : (len(fullName))]
	} else {
		firstName = fullName
		lastName = ""
	}
	return firstName, lastName
}
