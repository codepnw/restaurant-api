package middleware

import (
	"fmt"
	"net/http"

	"github.com/codepnw/restaurant-api/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authentication header provides")})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("FirstName", claims.FirstName)
		c.Set("LastName", claims.LastName)
		c.Set("Uid", claims.Uid)

		c.Next()
	}
}
