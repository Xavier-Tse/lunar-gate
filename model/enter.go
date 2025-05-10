package model

import (
	"gorm.io/gorm"
	"time"
)

type LunarModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Page struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Sort  string `form:"sort"`
	Key   string `form:"key"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" url:"id"`
}

type IDListRequest struct {
	IDList []uint `json:"idList"`
}
