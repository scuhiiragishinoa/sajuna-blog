package main

import (
	"context"
	"log"
	"sajuna-blog-backend/internal/config"
	"sajuna-blog-backend/internal/database"
	"sajuna-blog-backend/internal/model"
	"sajuna-blog-backend/internal/service"

	"github.com/TarsCloud/TarsGo/tars"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	db := database.Init(cfg.Database.DSN)
	
	// 自动迁移数据库表
	err := db.AutoMigrate(
		&model.User{},
		&model.Article{},
		&model.Category{},
		&model.Tag{},
		&model.Comment{},
		&model.Media{},
		&model.ArticleTag{},
	)
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	// 创建服务实例
	userService := service.NewUserService(db, cfg)
	articleService := service.NewArticleService(db, cfg)
	categoryService := service.NewCategoryService(db, cfg)
	tagService := service.NewTagService(db, cfg)
	commentService := service.NewCommentService(db, cfg)

	// 注册tars服务
	imp := &BlogServerImp{
		userService:     userService,
		articleService:  articleService,
		categoryService: categoryService,
		tagService:     tagService,
		commentService:  commentService,
	}

	// 启动tars服务
	app := new(BlogServer)
	app.AddServant(imp, cfg.Server.Port)
	tars.Run()
}

// BlogServerImp 实现tars接口
type BlogServerImp struct {
	userService     *service.UserService
	articleService  *service.ArticleService
	categoryService *service.CategoryService
	tagService     *service.TagService
	commentService  *service.CommentService
}

// 用户服务方法
func (imp *BlogServerImp) Register(ctx context.Context, username, password, email string, message *string) (int32, error) {
	code, msg := imp.userService.Register(username, password, email)
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) Login(ctx context.Context, username, password string, token, message *string) (int32, error) {
	code, tok, msg := imp.userService.Login(username, password)
	*token = tok
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) GetProfile(ctx context.Context, token string, userInfo, message *string) (int32, error) {
	code, info, msg := imp.userService.GetProfile(token)
	*userInfo = info
	*message = msg
	return int32(code), nil
}

// 文章服务方法
func (imp *BlogServerImp) GetArticles(ctx context.Context, page, pageSize int32, category, tag string, articles, message *string) (int32, error) {
	code, arts, msg := imp.articleService.GetArticles(int(page), int(pageSize), category, tag)
	*articles = arts
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) GetArticle(ctx context.Context, articleId int32, article, message *string) (int32, error) {
	code, art, msg := imp.articleService.GetArticle(int(articleId))
	*article = art
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) CreateArticle(ctx context.Context, token, title, content, summary string, categoryId int32, tags string, message *string) (int32, error) {
	code, msg := imp.articleService.CreateArticle(token, title, content, summary, int(categoryId), tags)
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) UpdateArticle(ctx context.Context, token string, articleId int32, title, content, summary string, categoryId int32, tags string, message *string) (int32, error) {
	code, msg := imp.articleService.UpdateArticle(token, int(articleId), title, content, summary, int(categoryId), tags)
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) DeleteArticle(ctx context.Context, token string, articleId int32, message *string) (int32, error) {
	code, msg := imp.articleService.DeleteArticle(token, int(articleId))
	*message = msg
	return int32(code), nil
}

// 分类服务方法
func (imp *BlogServerImp) GetCategories(ctx context.Context, categories, message *string) (int32, error) {
	code, cats, msg := imp.categoryService.GetCategories()
	*categories = cats
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) CreateCategory(ctx context.Context, token, name, slug string, message *string) (int32, error) {
	code, msg := imp.categoryService.CreateCategory(token, name, slug)
	*message = msg
	return int32(code), nil
}

// 标签服务方法
func (imp *BlogServerImp) GetTags(ctx context.Context, tags, message *string) (int32, error) {
	code, tgs, msg := imp.tagService.GetTags()
	*tags = tgs
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) CreateTag(ctx context.Context, token, name, slug string, message *string) (int32, error) {
	code, msg := imp.tagService.CreateTag(token, name, slug)
	*message = msg
	return int32(code), nil
}

// 评论服务方法
func (imp *BlogServerImp) GetComments(ctx context.Context, articleId, page, pageSize int32, comments, message *string) (int32, error) {
	code, comms, msg := imp.commentService.GetComments(int(articleId), int(page), int(pageSize))
	*comments = comms
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) CreateComment(ctx context.Context, articleId int32, authorName, authorEmail, content string, parentId int32, message *string) (int32, error) {
	code, msg := imp.commentService.CreateComment(int(articleId), authorName, authorEmail, content, int(parentId))
	*message = msg
	return int32(code), nil
}

func (imp *BlogServerImp) DeleteComment(ctx context.Context, token string, commentId int32, message *string) (int32, error) {
	code, msg := imp.commentService.DeleteComment(token, int(commentId))
	*message = msg
	return int32(code), nil
}
