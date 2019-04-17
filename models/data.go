package models

type Data struct {
	Avatar        string    `bson:"avatar" json:"avatar"`
	HierarchyName string    `bson:"hierarchy_name" json:"hierarchy_name"`
	Description   [2]string `bson:"description" json:"description"`
}
