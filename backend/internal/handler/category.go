package handler

import (
	"net/http"

	"sajuna-blog-backend/internal/database"
	"sajuna-blog-backend/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Sort        int    `json:"sort"`
}

// GetCategories 获取分类列表
// @Summary 获取分类列表
// @Description 获取所有分类
// @Tags 分类
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	db := database.GetDB()

	var categories []model.Category
	db.Order("sort ASC, created_at DESC").Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

// CreateCategory 创建分类
// @Summary 创建分类
// @Description 创建新分类
// @Tags 分类
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateCategoryRequest true "分类信息"
// @Success 200 {object} map[string]interface{}
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	// 检查分类名是否已存在
	var existingCategory model.Category
	if err := db.Where("name = ? OR slug = ?", req.Name, req.Slug).First(&existingCategory).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "分类名或别名已存在"})
		return
	}

	category := model.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Color:       req.Color,
		Sort:        req.Sort,
	}

	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "分类创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "分类创建成功",
		"category": category,
	})
}



