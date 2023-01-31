package business

import "libraryApi/src/libraryService/model"

type ILogic interface {
	Bind() model.BaseModel
	Create() model.BaseModel
	FindByID() model.BaseModel
	Update() model.BaseModel
	Delete()
	GetAll() []model.BaseModel
}
