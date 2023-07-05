package common

type Role struct {
	Id             int
	Name           string
	Active         bool            `gorm:"column:is_active"`
	Authorizations []Authorization `gorm:"many2many:tbl_role_authorization_code_relation"`
}

func (Role) TableName() string {
	return "tbl_role"
}
