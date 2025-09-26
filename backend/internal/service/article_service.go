package service

import (
	"encoding/json"
	"errors"
	"sajuna-blog-backend/internal/config"
	"sajuna-blog-backend/internal/database"
	"sajuna-blog-backend/internal/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type ArticleService struct {
	db *gorm.DB
	cfg *config.Config
}

func NewArticleService(db *gorm.DB, cfg *config.Config) *ArticleService {
	return &ArticleService{
		db: db,
		cfg: cfg,
	}
}

// GetArticles 获取文章列表
func (s *ArticleService) GetArticles(page, pageSize int, category, tag string) (int, string, string) {
	var articles []model.Article
	query := s.db.Model(&model.Article{}).Where("status = ?", "published")

	// 分类筛选
	if category != "" {
		query = query.Joins("JOIN categories ON articles.category_id = categories.id").
			Where("categories.slug = ?", category)
	}

	// 标签筛选
	if tag != "" {
		query = query.Joins("JOIN article_tags ON articles.id = article_tags.article_id").
			Joins("JOIN tags ON article_tags.tag_id = tags.id").
			Where("tags.slug = ?", tag)
	}

	// 分页
	offset := (page - 1) * pageSize
	if err := query.Preload("Category").Preload("Tags").
		Order("created_at DESC").
		Offset(offset).Limit(pageSize).Find(&articles).Error; err != nil {
		return 1, "", "获取文章列表失败"
	}

	// 获取总数
	var total int64
	query.Count(&total)

	result := map[string]interface{}{
		"articles": articles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}

	articlesJSON, _ := json.Marshal(result)
	return 0, string(articlesJSON), "获取成功"
}

// GetArticle 获取文章详情
func (s *ArticleService) GetArticle(articleId int) (int, string, string) {
	var article model.Article
	if err := s.db.Preload("Category").Preload("Tags").
		Where("id = ? AND status = ?", articleId, "published").
		First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1, "", "文章不存在"
		}
		return 1, "", "获取文章失败"
	}

	articleJSON, _ := json.Marshal(article)
	return 0, string(articleJSON), "获取成功"
}

// CreateArticle 创建文章
func (s *ArticleService) CreateArticle(token, title, content, summary string, categoryId int, tags string) (int, string) {
	// 验证token
	claims, err := s.validateToken(token)
	if err != nil {
		return 1, "Token无效"
	}

	// 创建文章
	article := model.Article{
		Title:      title,
		Content:    content,
		Summary:    summary,
		Status:     "published",
		CategoryID: uint(categoryId),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.db.Create(&article).Error; err != nil {
		return 1, "文章创建失败"
	}

	// 处理标签
	if tags != "" {
		s.processTags(&article, tags)
	}

	return 0, "文章创建成功"
}

// UpdateArticle 更新文章
func (s *ArticleService) UpdateArticle(token string, articleId int, title, content, summary string, categoryId int, tags string) (int, string) {
	// 验证token
	claims, err := s.validateToken(token)
	if err != nil {
		return 1, "Token无效"
	}

	// 查找文章
	var article model.Article
	if err := s.db.Where("id = ?", articleId).First(&article).Error; err != nil {
		return 1, "文章不存在"
	}

	// 更新文章
	article.Title = title
	article.Content = content
	article.Summary = summary
	article.CategoryID = uint(categoryId)
	article.UpdatedAt = time.Now()

	if err := s.db.Save(&article).Error; err != nil {
		return 1, "文章更新失败"
	}

	// 处理标签
	s.processTags(&article, tags)

	return 0, "文章更新成功"
}

// DeleteArticle 删除文章
func (s *ArticleService) DeleteArticle(token string, articleId int) (int, string) {
	// 验证token
	claims, err := s.validateToken(token)
	if err != nil {
		return 1, "Token无效"
	}

	// 软删除文章
	if err := s.db.Where("id = ?", articleId).Delete(&model.Article{}).Error; err != nil {
		return 1, "文章删除失败"
	}

	return 0, "文章删除成功"
}

// processTags 处理文章标签
func (s *ArticleService) processTags(article *model.Article, tags string) {
	// 清除现有标签关联
	s.db.Where("article_id = ?", article.ID).Delete(&model.ArticleTag{})

	// 解析标签字符串
	tagNames := []string{}
	if tags != "" {
		// 简单的逗号分隔解析
		// 实际项目中可能需要更复杂的解析逻辑
		for _, tagName := range []string{tags} {
			if tagName != "" {
				tagNames = append(tagNames, tagName)
			}
		}
	}

	// 创建或查找标签，并建立关联
	for _, tagName := range tagNames {
		var tag model.Tag
		if err := s.db.Where("name = ?", tagName).First(&tag).Error; err != nil {
			// 标签不存在，创建新标签
			tag = model.Tag{
				Name: tagName,
				Slug: tagName, // 简化处理，实际应该生成slug
			}
			s.db.Create(&tag)
		}

		// 建立文章和标签的关联
		articleTag := model.ArticleTag{
			ArticleID: article.ID,
			TagID:     tag.ID,
		}
		s.db.Create(&articleTag)
	}
}

// validateToken 验证JWT token
func (s *ArticleService) validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
