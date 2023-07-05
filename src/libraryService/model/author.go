package model

type Author struct {
	BaseModel
	Name     string
	LastName string
	Email    string
	//Books    []Book `gorm:"many2many:tbl_book_author_relation;"`
}

func (Author) TableName() string {
	return "tbl_author"
}
