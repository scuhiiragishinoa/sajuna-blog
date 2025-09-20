package model

import (
	"time"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	Status    string         `json:"status" gorm:"size:20;default:'pending'"` // pending, approved, rejected
	IP        string         `json:"ip" gorm:"size:45"`
	UserAgent string         `json:"user_agent" gorm:"size:255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 外键
	ArticleID uint64 `json:"article_id" gorm:"not null"`
	UserID    uint64 `json:"user_id"`
	ParentID  *uint64 `json:"parent_id"` // 支持回复评论
	
	// 关联
	Article Article `json:"article,omitempty" gorm:"foreignKey:ArticleID"`
	User    User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Parent  *Comment `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies []Comment `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
}


