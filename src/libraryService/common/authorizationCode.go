package common

import "LibraryApi/src/libraryService/model"

type Authorization struct {
	model.BaseModel
	Code  string
	Title string
	Roles []Role `gorm:"many2many:tbl_role_authorizations;"`
}

func (Authorization) TableName() string {
	return "tbl_authorization"
}
