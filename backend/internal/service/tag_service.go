package service

import (
	"encoding/json"
	"sajuna-blog-backend/internal/config"
	"sajuna-blog-backend/internal/model"
	"sajuna-blog-backend/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type TagService struct {
	db *gorm.DB
	cfg *config.Config
}

func NewTagService(db *gorm.DB, cfg *config.Config) *TagService {
	return &TagService{
		db: db,
		cfg: cfg,
	}
}

// GetTags 获取标签列表
func (s *TagService) GetTags() (int, string, string) {
	var tags []model.Tag
	if err := s.db.Find(&tags).Error; err != nil {
		return 1, "", "获取标签列表失败"
	}

	tagsJSON, _ := json.Marshal(tags)
	return 0, string(tagsJSON), "获取成功"
}

// CreateTag 创建标签
func (s *TagService) CreateTag(token, name, slug string) (int, string) {
	// 验证token
	claims, err := s.validateToken(token)
	if err != nil {
		return 1, "Token无效"
	}

	// 检查标签名是否已存在
	var existingTag model.Tag
	if err := s.db.Where("name = ? OR slug = ?", name, slug).First(&existingTag).Error; err == nil {
		return 1, "标签名或slug已存在"
	}

	// 创建标签
	tag := model.Tag{
		Name: name,
		Slug: slug,
	}

	if err := s.db.Create(&tag).Error; err != nil {
		return 1, "标签创建失败"
	}

	return 0, "标签创建成功"
}

// validateToken 验证JWT token
func (s *TagService) validateToken(tokenString string) (jwt.MapClaims, error) {
	return utils.ValidateToken(tokenString, s.cfg)
}
