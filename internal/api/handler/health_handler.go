package handler

import (
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/gin-gonic/gin"
	"fibonacci-spiral-matrix-go/internal/api/dto"
)

type HealthHandler struct{}

func (h *HealthHandler) Register(apiRouteGroup *gin.RouterGroup) {
	apiRouteGroup.GET("/status", h.Status)
}

// Status Health check godoc
// @Summary Check api status
// @Description Get api pulse status
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthOutput
// @Router /api/status [get]
func (h *HealthHandler) Status(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("HealthCheck")
	defer span.Finish()
	ctx.JSON(http.StatusOK, dto.HealthOutput{
		Status: "pass",
	})
}
