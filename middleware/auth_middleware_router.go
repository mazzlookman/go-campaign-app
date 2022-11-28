package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-campaign-app/helper"
	"go-campaign-app/service"
	"net/http"
	"strings"
)

func NewJWTAuthMiddleware(auth JWTAuth, service service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.Contains(header, "Hii") {
			helper.ErrorUnauthorized(c, http.StatusUnauthorized)
			return
		}

		splitHeader := strings.Split(header, " ")
		token := splitHeader[1]

		validateToken, err := auth.ValidateToken(token)
		if err != nil {
			helper.ErrorUnauthorized(c, http.StatusUnauthorized)
			return
		}

		claim, ok := validateToken.Claims.(jwt.MapClaims)
		if !ok || !validateToken.Valid {
			helper.ErrorUnauthorized(c, http.StatusUnauthorized)
			return
		}

		userId := int(claim["user_id"].(float64))

		findById, err := service.FindById(userId)
		if err != nil {
			helper.ErrorUnauthorized(c, http.StatusUnauthorized)
			return
		}

		c.Set("currentUser", findById.Id)
	}
}
