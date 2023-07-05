package business

import (
	"LibraryApi/src/libraryService/db"
	"fmt"
	"gorm.io/gorm"
)

type ServiceInterface interface {
	Create(e any)
	Update(e any)
	Delete(e any)
	FindByName(e any)
	FindById(e any)
}

type AbstractService struct {
	ServiceInterface
}

var _abstractService = new(AbstractService)

func (s *AbstractService) Create(model any) *gorm.DB {
	return db.ConnectToDb().Create(model)

}

func (s *AbstractService) Update(e any)     { fmt.Println("update") }
func (s *AbstractService) FindByName(e any) { db.ConnectToDb() }

func (s *AbstractService) Delete(e any) {
	fmt.Println("delete")
}
