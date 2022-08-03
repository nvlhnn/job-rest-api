package middleware

import (
	"dansmultipro/recruitment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {

			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, _ := jwtService.ValidateToken(authHeader)
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

