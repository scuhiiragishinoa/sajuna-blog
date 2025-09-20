package handler

import (
	"net/http"
	"strconv"

	"sajuna-blog-backend/internal/database"
	"sajuna-blog-backend/internal/model"

	"github.com/gin-gonic/gin"
)

type CreateArticleRequest struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Excerpt     string   `json:"excerpt"`
	CoverImage  string   `json:"cover_image"`
	CategoryID  uint64   `json:"category_id"`
	TagIDs      []uint64 `json:"tag_ids"`
	Status      string   `json:"status"`
}

type UpdateArticleRequest struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Excerpt     string   `json:"excerpt"`
	CoverImage  string   `json:"cover_image"`
	CategoryID  uint64   `json:"category_id"`
	TagIDs      []uint64 `json:"tag_ids"`
	Status      string   `json:"status"`
	IsTop       bool     `json:"is_top"`
}

// GetArticles 获取文章列表
// @Summary 获取文章列表
// @Description 分页获取文章列表
// @Tags 文章
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param status query string false "文章状态" default(published)
// @Param category_id query int false "分类ID"
// @Success 200 {object} map[string]interface{}
// @Router /articles [get]
func GetArticles(c *gin.Context) {
	db := database.GetDB()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.DefaultQuery("status", "published")
	categoryID := c.Query("category_id")

	offset := (page - 1) * pageSize

	query := db.Model(&model.Article{}).Where("status = ?", status)

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	var articles []model.Article
	var total int64

	query.Count(&total)
	query.Preload("User").Preload("Category").Preload("Tags").
		Order("created_at DESC").
		Offset(offset).Limit(pageSize).Find(&articles)

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    total,
		"page":     page,
		"page_size": pageSize,
	})
}

// GetArticle 获取单篇文章
// @Summary 获取单篇文章
// @Description 根据ID获取文章详情
// @Tags 文章
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} map[string]interface{}
// @Router /articles/{id} [get]
func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	db := database.GetDB()
	var article model.Article

	if err := db.Preload("User").Preload("Category").Preload("Tags").Preload("Comments").
		First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 增加浏览量
	db.Model(&article).Update("view_count", article.ViewCount+1)

	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

// CreateArticle 创建文章
// @Summary 创建文章
// @Description 创建新文章
// @Tags 文章
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateArticleRequest true "文章信息"
// @Success 200 {object} map[string]interface{}
// @Router /articles [post]
func CreateArticle(c *gin.Context) {
	var req CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	db := database.GetDB()

	article := model.Article{
		Title:      req.Title,
		Content:    req.Content,
		Excerpt:    req.Excerpt,
		CoverImage: req.CoverImage,
		Status:     req.Status,
		UserID:     userID.(uint64),
		CategoryID: uint64(req.CategoryID),
	}

	if err := db.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章创建失败"})
		return
	}

	// 关联标签
	if len(req.TagIDs) > 0 {
		var tags []model.Tag
		db.Where("id IN ?", req.TagIDs).Find(&tags)
		db.Model(&article).Association("Tags").Append(tags)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "文章创建成功",
		"article": article,
	})
}

// UpdateArticle 更新文章
// @Summary 更新文章
// @Description 更新文章信息
// @Tags 文章
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "文章ID"
// @Param request body UpdateArticleRequest true "文章信息"
// @Success 200 {object} map[string]interface{}
// @Router /articles/{id} [put]
func UpdateArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var req UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	db := database.GetDB()

	var article model.Article
	if err := db.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 检查权限
	if article.UserID != userID.(uint64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限修改此文章"})
		return
	}

	// 更新文章
	updates := map[string]interface{}{
		"title":       req.Title,
		"content":     req.Content,
		"excerpt":     req.Excerpt,
		"cover_image": req.CoverImage,
		"status":      req.Status,
		"is_top":      req.IsTop,
		"category_id": req.CategoryID,
	}

	if err := db.Model(&article).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章更新失败"})
		return
	}

	// 更新标签关联
	if len(req.TagIDs) > 0 {
		var tags []model.Tag
		db.Where("id IN ?", req.TagIDs).Find(&tags)
		db.Model(&article).Association("Tags").Replace(tags)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "文章更新成功",
		"article": article,
	})
}

// DeleteArticle 删除文章
// @Summary 删除文章
// @Description 删除文章
// @Tags 文章
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "文章ID"
// @Success 200 {object} map[string]interface{}
// @Router /articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	userID, _ := c.Get("user_id")
	db := database.GetDB()

	var article model.Article
	if err := db.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 检查权限
	if article.UserID != userID.(uint64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限删除此文章"})
		return
	}

	if err := db.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "文章删除成功",
	})
}



