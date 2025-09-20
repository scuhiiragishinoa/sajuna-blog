package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	
	// 开发环境允许的源
	config.AllowOrigins = []string{
		"http://localhost:3000", 
		"http://127.0.0.1:3000",
		"http://localhost:80",
		"http://127.0.0.1:80",
	}
	
	// 允许的HTTP方法
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	
	// 允许的请求头
	config.AllowHeaders = []string{
		"Origin", 
		"Content-Type", 
		"Accept", 
		"Authorization", 
		"X-Requested-With",
		"X-Real-IP",
		"X-Forwarded-For",
		"X-Forwarded-Proto",
	}
	
	// 允许凭据（仅在开发环境）
	config.AllowCredentials = true
	
	// 暴露的响应头
	config.ExposeHeaders = []string{"Content-Length", "Content-Type"}
	
	// 预检请求缓存时间（秒）
	config.MaxAge = 12 * 3600

	return cors.New(config)
}



