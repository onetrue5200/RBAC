package models

type User struct {
	Id       int
	Username string
	Password string
	IsSuper  int
	RoleId   int
	IsDelete int
}

func (User) TableName() string {
	return "user"
}
