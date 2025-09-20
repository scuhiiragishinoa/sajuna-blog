package model

import (
	"time"
	"gorm.io/gorm"
)

type Tag struct {
	ID        uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"size:30;not null;uniqueIndex"`
	Slug      string         `json:"slug" gorm:"size:30;not null;uniqueIndex"`
	Color     string         `json:"color" gorm:"size:20"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联
	Articles []Article `json:"articles,omitempty" gorm:"many2many:article_tags;"`
}


