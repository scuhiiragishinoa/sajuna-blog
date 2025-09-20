package handler

import (
	"net/http"

	"sajuna-blog-backend/internal/database"
	"sajuna-blog-backend/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateTagRequest struct {
	Name  string `json:"name" binding:"required"`
	Slug  string `json:"slug" binding:"required"`
	Color string `json:"color"`
}

// GetTags 获取标签列表
// @Summary 获取标签列表
// @Description 获取所有标签
// @Tags 标签
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /tags [get]
func GetTags(c *gin.Context) {
	db := database.GetDB()

	var tags []model.Tag
	db.Order("created_at DESC").Find(&tags)

	c.JSON(http.StatusOK, gin.H{
		"tags": tags,
	})
}

// CreateTag 创建标签
// @Summary 创建标签
// @Description 创建新标签
// @Tags 标签
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateTagRequest true "标签信息"
// @Success 200 {object} map[string]interface{}
// @Router /tags [post]
func CreateTag(c *gin.Context) {
	var req CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	// 检查标签名是否已存在
	var existingTag model.Tag
	if err := db.Where("name = ? OR slug = ?", req.Name, req.Slug).First(&existingTag).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "标签名或别名已存在"})
		return
	}

	tag := model.Tag{
		Name:  req.Name,
		Slug:  req.Slug,
		Color: req.Color,
	}

	if err := db.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "标签创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "标签创建成功",
		"tag":     tag,
	})
}



