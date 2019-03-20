package models

type Data struct {
	Avatar        string    `json:"avatar"`
	HierarchyName string    `json:"hierarchy_name"`
	Description   [2]string `json:"description"`
}
