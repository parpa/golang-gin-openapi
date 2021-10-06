package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OAuthIntrospectionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/health" {
			return
		}
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.String(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}
		// todo check auth
		// todo set token
		// c.Set("token", token)
	}
}
