package model

type Book struct {
	BaseModel
	Isbn       string `gorm:"uniqueIndex"`
	Name       string
	PageCount  int
	Count      int
	Author     []Author `gorm:"many2many:book_authors;"`
	CategoryId Category `gorm:"category_id"`
}

type Author struct {
	BaseModel
	Name     string
	LastName string
	Email    string
}
type Category struct {
	BaseModel
	CategoryName string
	Books        []Book `gorm:"foreignKey:book_id"`
}
