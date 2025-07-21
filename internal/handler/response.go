package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, logger *slog.Logger, statusCode int, message string) {
	logger.Error(fmt.Sprintf("%v %s", statusCode, message))
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
