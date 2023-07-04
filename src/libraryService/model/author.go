package model

type Author struct {
	BaseModel
	Name     string
	LastName string
	Email    string
	Books    []Book `gorm:"many2many:book_authors;"`
}

func (Author) TableName() string {
	return "tbl_author"
}
