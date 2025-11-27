package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/arhief32/emp-be/config"
)

type JWTMiddleware struct {
	cfg *config.Config
}

func NewJWTMiddleware(cfg *config.Config) *JWTMiddleware {
	return &JWTMiddleware{cfg: cfg}
}

// Gin middleware factory
func (m *JWTMiddleware) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}
		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// validate alg
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenMalformed
			}
			return []byte(m.cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}
		fmt.Println(claims)
		if sub, ok := claims["sub"].(float64); ok {
			c.Set("user_id", int(sub))
			c.Set("username", claims["usr"])
		}
		c.Next()
	}
}

// simple role middleware helper (local to router file)
// you can move to middleware package if preferred
func (m *JWTMiddleware) RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, _ := c.Get("role")
		if r == nil {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
			return
		}
		rs, ok := r.(string)
		if !ok || rs != role {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}

// shorthand for using middleware in routers
func (m *JWTMiddleware) Gin() gin.HandlerFunc {
	return m.HandlerFunc()
}
