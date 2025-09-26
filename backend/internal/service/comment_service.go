package service

import (
	"encoding/json"
	"sajuna-blog-backend/internal/config"
	"sajuna-blog-backend/internal/model"
	"sajuna-blog-backend/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
	cfg *config.Config
}

func NewCommentService(db *gorm.DB, cfg *config.Config) *CommentService {
	return &CommentService{
		db: db,
		cfg: cfg,
	}
}

// GetComments 获取文章评论
func (s *CommentService) GetComments(articleId, page, pageSize int) (int, string, string) {
	var comments []model.Comment
	query := s.db.Where("article_id = ?", articleId)

	// 分页
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").
		Offset(offset).Limit(pageSize).Find(&comments).Error; err != nil {
		return 1, "", "获取评论列表失败"
	}

	// 获取总数
	var total int64
	query.Count(&total)

	result := map[string]interface{}{
		"comments": comments,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}

	commentsJSON, _ := json.Marshal(result)
	return 0, string(commentsJSON), "获取成功"
}

// CreateComment 创建评论
func (s *CommentService) CreateComment(articleId int, authorName, authorEmail, content string, parentId int) (int, string) {
	// 创建评论
	comment := model.Comment{
		ArticleID:   uint(articleId),
		AuthorName:  authorName,
		AuthorEmail: authorEmail,
		Content:     content,
		ParentID:    uint(parentId),
		IsPinned:    false,
		CreatedAt:   time.Now(),
	}

	if err := s.db.Create(&comment).Error; err != nil {
		return 1, "评论创建失败"
	}

	return 0, "评论创建成功"
}

// DeleteComment 删除评论
func (s *CommentService) DeleteComment(token string, commentId int) (int, string) {
	// 验证token
	claims, err := s.validateToken(token)
	if err != nil {
		return 1, "Token无效"
	}

	// 删除评论
	if err := s.db.Where("id = ?", commentId).Delete(&model.Comment{}).Error; err != nil {
		return 1, "评论删除失败"
	}

	return 0, "评论删除成功"
}

// validateToken 验证JWT token
func (s *CommentService) validateToken(tokenString string) (jwt.MapClaims, error) {
	return utils.ValidateToken(tokenString, s.cfg)
}
