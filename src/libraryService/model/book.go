package model

type Book struct {
	BaseModel
	Isbn      string   `gorm:"uniqueIndex,index:book_isbn,default:NULL"`
	Name      string   `gorm:"index:book_name,default:NULL"`
	PageCount int      `gorm:"default:NULL"`
	Count     int      `gorm:"default:NULL"`
	Author    []Author `gorm:"many2many:tbl_book_author_relation;,default:NULL"`
	Category  Category `gorm:"column:category_id"`
}

func (Book) TableName() string {
	return "tbl_book"
}
