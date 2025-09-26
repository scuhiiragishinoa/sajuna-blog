package main

import (
	"context"
	"github.com/TarsCloud/TarsGo/tars"
)

// BlogServer 定义tars服务接口
type BlogServer struct {
}

// AddServant 添加服务实现
func (s *BlogServer) AddServant(imp BlogServerImpInterface, obj string) {
	tars.AddServant(s, imp, obj)
}

// BlogServerImpInterface 定义服务实现接口
type BlogServerImpInterface interface {
	// 用户服务
	Register(ctx context.Context, username, password, email string, message *string) (int32, error)
	Login(ctx context.Context, username, password string, token, message *string) (int32, error)
	GetProfile(ctx context.Context, token string, userInfo, message *string) (int32, error)
	
	// 文章服务
	GetArticles(ctx context.Context, page, pageSize int32, category, tag string, articles, message *string) (int32, error)
	GetArticle(ctx context.Context, articleId int32, article, message *string) (int32, error)
	CreateArticle(ctx context.Context, token, title, content, summary string, categoryId int32, tags string, message *string) (int32, error)
	UpdateArticle(ctx context.Context, token string, articleId int32, title, content, summary string, categoryId int32, tags string, message *string) (int32, error)
	DeleteArticle(ctx context.Context, token string, articleId int32, message *string) (int32, error)
	
	// 分类服务
	GetCategories(ctx context.Context, categories, message *string) (int32, error)
	CreateCategory(ctx context.Context, token, name, slug string, message *string) (int32, error)
	
	// 标签服务
	GetTags(ctx context.Context, tags, message *string) (int32, error)
	CreateTag(ctx context.Context, token, name, slug string, message *string) (int32, error)
	
	// 评论服务
	GetComments(ctx context.Context, articleId, page, pageSize int32, comments, message *string) (int32, error)
	CreateComment(ctx context.Context, articleId int32, authorName, authorEmail, content string, parentId int32, message *string) (int32, error)
	DeleteComment(ctx context.Context, token string, commentId int32, message *string) (int32, error)
}
