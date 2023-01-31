package model

import (
	"time"
)

type BaseModel struct {
	//gorm.Model
	Id          uint      `json:"id" gorm:"primaryKey"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	CreatedBy   string
	UpdateDate  time.Time `gorm:"autoUpdateTime"`
	UpdatedBy   string
}
