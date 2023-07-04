package common

type Role struct {
	Id             int
	Name           string
	Active         bool
	Authorizations []Authorization `gorm:"many2many:tbl_role_authorizations;"`
}

func (Role) TableName() string {
	return "tbl_role"
}
