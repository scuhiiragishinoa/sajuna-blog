package model

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string         `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Password  string         `json:"-" gorm:"size:255;not null"`
	Nickname  string         `json:"nickname" gorm:"size:50"`
	Avatar    string         `json:"avatar" gorm:"size:255"`
	Bio       string         `json:"bio" gorm:"type:text"`
	Role      string         `json:"role" gorm:"size:20;default:'user'"`
	Status    string         `json:"status" gorm:"size:20;default:'active'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联
	Articles  []Article  `json:"articles,omitempty" gorm:"foreignKey:UserID"`
	Comments  []Comment  `json:"comments,omitempty" gorm:"foreignKey:UserID"`
}


