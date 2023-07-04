package model

type Category struct {
	BaseModel
	CategoryName string
	Books        []Book `gorm:"constraint:OnDelete:SET NULL;"`
}

func (Category) TableName() string {
	return "tbl_category"
}
