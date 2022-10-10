package models

type User struct {
	Id       int
	Username string
	Password string
	IsSuper  int
	IsDelete int
	Roles    []Role `gorm:"many2many:user_role;"`
}

func (User) TableName() string {
	return "user"
}
