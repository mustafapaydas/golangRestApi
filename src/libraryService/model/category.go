package model

type Category struct {
	BaseModel
	CategoryName string
	Books        []Book `gorm:"foreignKey:Category"`
}

func (Category) TableName() string {
	return "tbl_category"
}
