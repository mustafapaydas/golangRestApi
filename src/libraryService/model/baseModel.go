package model

import (
	"time"
)

type BaseModel struct {
	//gorm.Model
	Id          uint      `json:"id" gorm:"primaryKey"`
	CreatedDate time.Time `gorm:"<-:create"`
	CreatedBy   string
	UpdateDate  time.Time
	UpdatedBy   string
}
