package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/yushengguo557/chat/internal/common"
)

// JWTAuthMiddleware 个中间件，用于验证 JWT token
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		tokenString := c.GetHeader("Authorization")
		fmt.Println("Authorization", tokenString)
		if tokenString == "" { // 如果请求头中没有 token，则返回错误信息
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}
		tokenString = strings.Split(tokenString, " ")[1]
		// 解析 token
		token, err := jwt.ParseWithClaims(tokenString, &common.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("123"), nil
		})
		if !token.Valid { // token不合法 格式错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			// c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		if err != nil { // 解析失败
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				fmt.Println(err)
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*common.Claims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ok})
			c.Abort()
			return
		}
		// 将用户id存储到上下文中，方便后续的处理
		c.Set("id", claims.ID)
		c.Next()
	}
}
