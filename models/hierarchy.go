package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Hierarchy struct {
	HierarchyID   bson.ObjectId `bson:"_id,omitempty"`
	HierarchyName string        `bson:"hierarchy_name"`
	PermissionID  int           `bson:"permission_id"`
}
