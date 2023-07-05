package model

type Book struct {
	BaseModel
	Isbn      string `gorm:"uniqueIndex,index:book_isbn,default:NULL",json:"isbn,omitempty"`
	Name      string `gorm:"index:book_name,default:NULL",json:"name,omitempty"`
	PageCount int    `gorm:"default:NULL",json:"pageCount,omitempty"`
	Count     int    `gorm:"default:NULL",json:"Count,omitempty"`
	Category  Category
	//Author    []Author `gorm:"many2many:tbl_book_author_relation;,default:NULL"`
	//Category  Category `gorm:"column:category_id"`
}

func (Book) TableName() string {
	return "tbl_book"
}
