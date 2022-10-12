package models

type Role struct {
	Id       int
	Name     string
	IsDelete int
	Accesses []Access `gorm:"many2many:role_access;"`
}

func (Role) TableName() string {
	return "role"
}
