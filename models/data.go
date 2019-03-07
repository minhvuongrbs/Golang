package models

type Data struct {
	Avatar      string    `bson:"avatar"`
	HierarchyID Hierarchy `bson:"hierarchy_id"`
	Description string    `bson:"description"`
}
