package service

import (
	"encoding/json"
	"errors"
	"sajuna-blog-backend/internal/config"
	"sajuna-blog-backend/internal/database"
	"sajuna-blog-backend/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
	cfg *config.Config
}

func NewUserService(db *gorm.DB, cfg *config.Config) *UserService {
	return &UserService{
		db: db,
		cfg: cfg,
	}
}

// Register 用户注册
func (s *UserService) Register(username, password, email string) (int, string) {
	// 检查用户名是否已存在
	var existingUser model.User
	if err := s.db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return 1, "用户名已存在"
	}

	// 检查邮箱是否已存在
	if err := s.db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return 1, "邮箱已存在"
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 1, "密码加密失败"
	}

	// 创建用户
	user := model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Email:        email,
		CreatedAt:    time.Now(),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return 1, "用户创建失败"
	}

	return 0, "注册成功"
}

// Login 用户登录
func (s *UserService) Login(username, password string) (int, string, string) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1, "", "用户名或密码错误"
		}
		return 1, "", "登录失败"
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return 1, "", "用户名或密码错误"
	}

	// 生成JWT token
	token, err := s.generateToken(user.ID, user.Username)
	if err != nil {
		return 1, "", "Token生成失败"
	}

	return 0, token, "登录成功"
}

// GetProfile 获取用户信息
func (s *UserService) GetProfile(token string) (int, string, string) {
	// 验证token
	claims, err := s.validateToken(token)
	if err != nil {
		return 1, "", "Token无效"
	}

	var user model.User
	if err := s.db.Where("id = ?", claims["user_id"]).First(&user).Error; err != nil {
		return 1, "", "用户不存在"
	}

	// 返回用户信息（不包含密码）
	userInfo := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"created_at": user.CreatedAt,
	}

	userInfoJSON, _ := json.Marshal(userInfo)
	return 0, string(userInfoJSON), "获取成功"
}

// generateToken 生成JWT token
func (s *UserService) generateToken(userID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(s.cfg.JWT.Expire)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWT.Secret))
}

// validateToken 验证JWT token
func (s *UserService) validateToken(tokenString string) (jwt.MapClaims, error) {
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
