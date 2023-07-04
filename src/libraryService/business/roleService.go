package business

import (
	"LibraryApi/src/libraryService/common"
	"LibraryApi/src/libraryService/db"
)

type RoleLogic struct {
	*AbstractService
}

func FindByName(_name []string) *common.Role {
	_role := new(common.Role)
	db.ConnectToDb().Table("tbl_role").Model(&common.Role{}).Preload("Authorizations").Where(" name in ? ", _name).First(&_role)
	println("xxxxxxxxxxxxxxxxxx", _role)
	return _role
}
