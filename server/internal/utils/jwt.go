// Package utils 工具函数包
// JWT工具函数，提供token生成和解析功能
package utils

import (
	"errors"
	"time"

	"zxm_ai_admin/server/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

// InitJWT 初始化JWT密钥
func InitJWT() {
	cfg := config.GetConfig()
	if cfg != nil {
		jwtSecret = []byte(cfg.JWT.Secret)
	}
}

// Claims JWT声明
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID uint, username string) (string, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return "", errors.New("配置未初始化")
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(cfg.JWT.ExpireHours) * time.Hour)

	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			Issuer:    "zxm-ai-admin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析JWT token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}

