package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// SecurityHeaders 添加安全响应头
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 防止XSS攻击
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		
		// 强制HTTPS（生产环境）
		if gin.Mode() == gin.ReleaseMode {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}
		
		// 内容安全策略
		c.Header("Content-Security-Policy", "default-src 'self'")
		
		// 引用者策略
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		
		c.Next()
	}
}

// RateLimiter 请求频率限制
func RateLimiter() gin.HandlerFunc {
	// 创建限流器：每秒10个请求，突发20个
	limiter := rate.NewLimiter(10, 20)
	
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// IPWhitelist IP白名单（可选）
func IPWhitelist(allowedIPs []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		
		// 如果白名单为空，则允许所有IP
		if len(allowedIPs) == 0 {
			c.Next()
			return
		}
		
		for _, ip := range allowedIPs {
			if ip == clientIP {
				c.Next()
				return
			}
		}
		
		c.JSON(http.StatusForbidden, gin.H{
			"error": "IP地址不在允许列表中",
		})
		c.Abort()
	}
}

// RequestSizeLimit 限制请求体大小
func RequestSizeLimit(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength > maxSize {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": "请求体过大",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Timeout 请求超时控制
func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置超时上下文
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// InputSanitizer 输入清理
func InputSanitizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 清理查询参数
		for _, values := range c.Request.URL.Query() {
			for i, value := range values {
				// 移除潜在的恶意字符
				cleaned := strings.TrimSpace(value)
				cleaned = strings.ReplaceAll(cleaned, "<script>", "")
				cleaned = strings.ReplaceAll(cleaned, "</script>", "")
				cleaned = strings.ReplaceAll(cleaned, "javascript:", "")
				cleaned = strings.ReplaceAll(cleaned, "vbscript:", "")
				cleaned = strings.ReplaceAll(cleaned, "onload=", "")
				cleaned = strings.ReplaceAll(cleaned, "onerror=", "")
				
				values[i] = cleaned
			}
		}
		
		c.Next()
	}
}
