package models

type UserInfo struct {
	UserID        string `bson:"_id,omitempty"`
	Name          string `bson:"name"`
	Avatar        string `bson:"avatar,omitempty"`
	Description   string `bson:"description"`
	Language      string `bson:"language"`
	HierarchyName string `bson:"hierarchy_name"`
	Password      string `bson:"password,omitempty"`
	Permission    int    `bson:"permission"`
}
