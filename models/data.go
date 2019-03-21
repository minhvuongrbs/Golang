package models

type Data struct {
	Avatar        string    `bson:"avatar"`
	HierarchyName string    `bson:"hierarchy_name"`
	Description   [2]string `bson:"description"`
}
