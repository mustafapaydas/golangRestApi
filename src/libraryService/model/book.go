package model

type Book struct {
	BaseModel
	Isbn      string `gorm:"uniqueIndex"`
	Name      string
	PageCount int
	Count     int
	//Author    []Author `gorm:"many2many:book_authors;"`
}

//type Author struct {
//	BaseModel
//	Name     string
//	LastName string
//	Email    string
//}
