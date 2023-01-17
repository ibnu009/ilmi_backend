package response

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Status  uint16      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GenericJsonResponse(c *gin.Context, statusCode uint16, message string, data interface{}) {
	res := JsonResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	c.JSON(int(statusCode), res)
}

func JSONLOGIN(c *gin.Context, statusCode int, message string, id uint32, email string, token string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"id":      id,
		"email":   email,
		"token":   token,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": err,
	})
}
