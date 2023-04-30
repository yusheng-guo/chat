package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/yushengguo557/chat/internal/common"
)

// GenerateToken 生成 token 对称加密
func GenerateToken(id string) (token string, err error) {
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, common.Claims{
		ID: id,
	}).SignedString([]byte("123"))
	// 使用密钥对 token 进行签名
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}
	return token, nil
}
