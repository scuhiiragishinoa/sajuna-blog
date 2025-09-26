package model

import (
	"time"
	"gorm.io/gorm"
)

type Article struct {
	ID          uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;size:200;not null"`
	Content     string         `json:"content" gorm:"type:longtext;not null"`
	Summary     string         `json:"summary" gorm:"type:text"`
	Excerpt     string         `json:"excerpt" gorm:"type:text"`
	CoverImage  string         `json:"cover_image" gorm:"size:255"`
	Status      string         `json:"status" gorm:"size:20;default:'draft'"` // draft, published, archived
	ViewCount   int            `json:"view_count" gorm:"default:0"`
	LikeCount   int            `json:"like_count" gorm:"default:0"`
	IsTop       bool           `json:"is_top" gorm:"default:false"`
	PublishedAt *time.Time     `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 外键
	UserID     uint64 `json:"user_id" gorm:"not null"`
	CategoryID uint64 `json:"category_id"`
	
	// 关联
	User      User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category  Category   `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Tags      []Tag      `json:"tags,omitempty" gorm:"many2many:article_tags;"`
	Comments  []Comment  `json:"comments,omitempty" gorm:"foreignKey:ArticleID"`
}

// ArticleTag 文章标签关联表
type ArticleTag struct {
	ID        uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	ArticleID uint64 `json:"article_id" gorm:"not null"`
	TagID     uint64 `json:"tag_id" gorm:"not null"`
	
	// 关联
	Article Article `json:"article,omitempty" gorm:"foreignKey:ArticleID"`
	Tag     Tag     `json:"tag,omitempty" gorm:"foreignKey:TagID"`
}


