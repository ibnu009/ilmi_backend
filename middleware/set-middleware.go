package middleware

import (
	"ilmi_backend/auth"
	"ilmi_backend/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetMiddlewareAuthentication(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		err := auth.TokenValid(c)
		if err != nil {
			response.GenericJsonResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			return
		}
		next(c)
	}
}
