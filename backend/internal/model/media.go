package model

import (
	"time"
	"gorm.io/gorm"
)

type Media struct {
	ID        uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Filename  string         `json:"filename" gorm:"size:255;not null"`
	OriginalName string      `json:"original_name" gorm:"size:255;not null"`
	Path      string         `json:"path" gorm:"size:500;not null"`
	URL       string         `json:"url" gorm:"size:500;not null"`
	MimeType  string         `json:"mime_type" gorm:"size:100"`
	Size      int64          `json:"size"`
	Width     int            `json:"width"`
	Height    int            `json:"height"`
	Alt       string         `json:"alt" gorm:"size:255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 外键
	UserID uint64 `json:"user_id" gorm:"not null"`
	
	// 关联
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}


