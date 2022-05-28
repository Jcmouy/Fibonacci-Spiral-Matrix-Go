package httperror

import (
	"github.com/gin-gonic/gin"
	"fibonacci-spiral-matrix-go/internal/api/dto"
)

func NewError(ctx *gin.Context, status int, message string, err error) {
	output := dto.ErrorOutput{
		Code:    status,
		Message: message,
	}
	if err != nil {
		output.Details = []string{err.Error()}
	}
	ctx.JSON(status, output)
}
