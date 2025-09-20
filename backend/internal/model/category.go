package model

import (
	"time"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"size:50;not null;uniqueIndex"`
	Slug        string         `json:"slug" gorm:"size:50;not null;uniqueIndex"`
	Description string         `json:"description" gorm:"type:text"`
	Color       string         `json:"color" gorm:"size:20"`
	Sort        int            `json:"sort" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联
	Articles []Article `json:"articles,omitempty" gorm:"foreignKey:CategoryID"`
}


