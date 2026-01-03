package belajargolanggin

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func ResponseSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(200, ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ResponseCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(201, ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ResponseBadRequest(c *gin.Context, message string) {
	c.JSON(400, ApiResponse{
		Success: false,
		Error:   message,
	})
}

func ResponseAuthorized(c *gin.Context, message string) {
	c.JSON(401, ApiResponse{
		Success: false,
		Data:    message,
	})
}

func ResponseNotFound(c *gin.Context, message string) {
	c.JSON(404, ApiResponse{
		Success: false,
		Data:    message,
	})
}

func ResponseInternalServerError(c *gin.Context, message string) {
	c.JSON(500, ApiResponse{
		Success: false,
		Data:    message,
	})
}
