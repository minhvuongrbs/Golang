package models

import "gopkg.in/mgo.v2/bson"

type UserInfo struct {
	UserID        bson.ObjectId `bson:"_id,omitempty" json:"user_id"`
	Name          string        `bson:"name" json:"name"`
	UserName      string        `bson:"username" json:"username"`
	Avatar        string        `bson:"avatar,omitempty" json:"avatar"`
	Description   [2]string     `bson:"description" json:"description"`
	Language      string        `bson:"language" json:"language"`
	HierarchyName string        `bson:"hierarchy_name" json:"hierarchy_name"`
	Password      string        `bson:"password,omitempty" json:"password"`
	Permission    int           `bson:"permission" json:"permission"`
}
