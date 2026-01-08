package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // или "http://localhost:5173"
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}
