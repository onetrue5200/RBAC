package models

type Access struct {
	Id       int
	Name     string
	Kind     string
	Url      string
	Action   string
	ModuleId int
	IsDelete int
}

func (Access) TableName() string {
	return "access"
}
