package models

type UserInfo struct {
	UserID        string    `json:"_id,omitempty"`
	Name          string    `json:"name"`
	Avatar        string    `json:"avatar,omitempty"`
	Description   [2]string `json:"description"`
	Language      string    `json:"language"`
	HierarchyName string    `json:"hierarchy_name"`
	Password      string    `json:"password,omitempty"`
	Permission    int       `json:"permission"`
}
