package main

import (
	"log"
	"os"
	"sajuna-blog-backend/internal/config"
	"sajuna-blog-backend/internal/database"
	"sajuna-blog-backend/internal/handler"
	"sajuna-blog-backend/internal/middleware"
	"sajuna-blog-backend/internal/model"

	"github.com/gin-gonic/gin"
)

// @title Sajuna Blog API
// @version 1.0
// @description 私人博客系统API文档
// @host localhost:8080
// @BasePath /api/v1
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
	)
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	// 创建Gin引擎
	r := gin.Default()

	// 配置代理信任 - 信任Docker网络内的代理
	// 从环境变量读取信任的代理，如果没有设置则使用默认的Docker网络子网
	trustedProxies := os.Getenv("GIN_TRUSTED_PROXIES")
	if trustedProxies == "" {
		// 默认信任Docker Compose创建的自定义网络子网
		trustedProxies = "172.18.0.0/16"
	}
	
	// 设置信任的代理
	if err := r.SetTrustedProxies([]string{trustedProxies}); err != nil {
		log.Printf("设置信任代理失败: %v", err)
	} else {
		log.Printf("已设置信任代理: %s", trustedProxies)
	}

	// 安全中间件
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.RateLimiter())
	r.Use(middleware.RequestSizeLimit(10 * 1024 * 1024)) // 10MB限制
	r.Use(middleware.InputSanitizer())
	
	// 基础中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	// 路由组
	api := r.Group("/api/v1")
	{
		// 健康检查
		api.GET("/health", handler.HealthCheck)
		
		// 用户相关
		users := api.Group("/users")
		{
			users.POST("/register", handler.Register)
			users.POST("/login", handler.Login)
			users.GET("/profile", middleware.AuthRequired(), handler.GetProfile)
		}

		// 文章相关
		articles := api.Group("/articles")
		{
			articles.GET("", handler.GetArticles)
			articles.GET("/:id", handler.GetArticle)
			articles.POST("", middleware.AuthRequired(), handler.CreateArticle)
			articles.PUT("/:id", middleware.AuthRequired(), handler.UpdateArticle)
			articles.DELETE("/:id", middleware.AuthRequired(), handler.DeleteArticle)
		}

		// 分类相关
		categories := api.Group("/categories")
		{
			categories.GET("", handler.GetCategories)
			categories.POST("", middleware.AuthRequired(), handler.CreateCategory)
		}

		// 标签相关
		tags := api.Group("/tags")
		{
			tags.GET("", handler.GetTags)
			tags.POST("", middleware.AuthRequired(), handler.CreateTag)
		}
	}

	// 启动服务器
	log.Printf("服务器启动在端口 %s", cfg.Server.Port)
	log.Fatal(r.Run(":" + cfg.Server.Port))
}



