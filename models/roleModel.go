package models

type Role struct {
	Id       int
	Name     string
	IsDelete int
}

func (Role) TableName() string {
	return "role"
}
