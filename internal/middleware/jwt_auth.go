package middleware

import (
	common2 "WoodInspection/internal/product/auth/common"
	"WoodInspection/internal/product/auth/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleWare JWT 认证中间件
func AuthMiddleWare(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string
		
		// 优先从 Cookie 获取 token
		cookieToken, err := c.Cookie("access_token")
		if err == nil && cookieToken != "" {
			tokenString = cookieToken
		} else {
			// 如果 Cookie 中没有，则从 Authorization Header 获取
			authHeader := c.Request.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				common2.Error(c, common2.CodeUnauthorized, "请提供有效的令牌")
				c.Abort()
				return
			}
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}
		token, err := jwt.ParseWithClaims(tokenString, &service.Claims{}, func(token *jwt.Token) (any, error) {
			return secret, nil
		})

		if err != nil {
			common2.Error(c, common2.CodeInvalidToken, "令牌无效或已过期")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*service.Claims)
		if !ok || !token.Valid {
			common2.Error(c, common2.CodeInvalidToken, "令牌无效")
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Set("username", claims.Name)

		c.Next()
	}
}
