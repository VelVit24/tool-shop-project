package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/VelVit24/projext/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, "not logined")
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&service.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte("key"), nil
			},
		)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, "invalid token")
			return
		}
		id := token.Claims.(*service.Claims).Id
		role := token.Claims.(*service.Claims).Role
		c.Set("user_id", id)
		c.Set("role", role)
		c.Next()
	}
}

func CheckAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok {
			c.JSON(http.StatusUnauthorized, "invalid token")
			return
		}
		if role != "admin" {
			c.JSON(http.StatusForbidden, "forbidden")
			return
		}
		c.Next()
	}
}

var apiLimiter = rate.NewLimiter(rate.Every(5*time.Second), 1)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !apiLimiter.Allow() {
			c.JSON(http.StatusTooManyRequests, "Too Many Requests")
			return
		}
		c.Next()
	}
}
