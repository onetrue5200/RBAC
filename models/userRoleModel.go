package models

type UserRole struct {
	UserId int
	RoleId int
}

func (UserRole) TableName() string {
	return "user_role"
}
