package service

import (
	"encoding/json"
	"sajuna-blog-backend/internal/config"
	"sajuna-blog-backend/internal/model"
	"sajuna-blog-backend/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
	cfg *config.Config
}

func NewCategoryService(db *gorm.DB, cfg *config.Config) *CategoryService {
	return &CategoryService{
		db: db,
		cfg: cfg,
	}
}

// GetCategories 获取分类列表
func (s *CategoryService) GetCategories() (int, string, string) {
	var categories []model.Category
	if err := s.db.Find(&categories).Error; err != nil {
		return 1, "", "获取分类列表失败"
	}

	categoriesJSON, _ := json.Marshal(categories)
	return 0, string(categoriesJSON), "获取成功"
}

// CreateCategory 创建分类
func (s *CategoryService) CreateCategory(token, name, slug string) (int, string) {
	// 验证token
	claims, err := s.validateToken(token)
	if err != nil {
		return 1, "Token无效"
	}

	// 检查分类名是否已存在
	var existingCategory model.Category
	if err := s.db.Where("name = ? OR slug = ?", name, slug).First(&existingCategory).Error; err == nil {
		return 1, "分类名或slug已存在"
	}

	// 创建分类
	category := model.Category{
		Name: name,
		Slug: slug,
	}

	if err := s.db.Create(&category).Error; err != nil {
		return 1, "分类创建失败"
	}

	return 0, "分类创建成功"
}

// validateToken 验证JWT token
func (s *CategoryService) validateToken(tokenString string) (jwt.MapClaims, error) {
	return utils.ValidateToken(tokenString, s.cfg)
}
