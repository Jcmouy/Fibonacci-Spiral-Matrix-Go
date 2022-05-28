package handler

import (
	"fibonacci-spiral-matrix-go/internal/api/middlewares"
	"fibonacci-spiral-matrix-go/internal/core/helper"
	"fibonacci-spiral-matrix-go/internal/core/interfaces/service"
	"github.com/pkg/errors"
	"net/http"

	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"

	"fibonacci-spiral-matrix-go/internal/api/dto"
	"fibonacci-spiral-matrix-go/internal/api/httperror"
	"fibonacci-spiral-matrix-go/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

type FiboSpiralMatrixHandler struct {
	FiboSpiralMatrixService service.FiboSpiralMatrixService
}

func NewFiboSpiralMatrixHandler(fiboSpiralMatrixService service.FiboSpiralMatrixService) FiboSpiralMatrixHandler {
	return FiboSpiralMatrixHandler{FiboSpiralMatrixService: fiboSpiralMatrixService}
}

func (fbh *FiboSpiralMatrixHandler) Register(apiRouteGroup *gin.RouterGroup) {
	router := apiRouteGroup.Group("/spiral").Use(middlewares.Authorization())
	{
		router.GET("", fbh.GetSpiralMatrix)
	}
}

// GetSpiralMatrix godoc
// @Summary Fetch a spiral matrix
// @Description Get matrix by row and column
// @Tags matrix
// @Accept json
// @Produce json
// @Param matrix body dto.MatrixInput true "Get spiral matrix"
// @Success 200 {object} dto.FibonacciSpiralMatrixDto
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Router /api/user/spiral [get]
func (fbh *FiboSpiralMatrixHandler) GetSpiralMatrix(ctx *gin.Context) {
	if err := helper.CheckUserRole(ctx, "USER"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	span := opentracing.GlobalTracer().StartSpan("FiboSpiralMatrixHandler-GetSpiralMatrix")
	spanContext := opentracing.ContextWithSpan(ctx.Request.Context(), span)
	defer span.Finish()

	rows := ctx.Query("rows")
	cols := ctx.Query("cols")

	matrixInput := dto.MatrixInput{
		Row:    rows,
		Column: cols,
	}

	if err := validator.Validate(matrixInput); err != nil {
		log.Error(errors.Wrap(err, "validation error"))
		httperror.NewError(ctx, http.StatusBadRequest, "Validation error", err)
		return
	}

	spiralMatrix, err := fbh.FiboSpiralMatrixService.GetSpiralMatrix(spanContext, matrixInput.ToModel())

	if err != nil {
		log.Error(errors.Wrap(err, "item create error"))
		httperror.NewError(ctx, http.StatusUnprocessableEntity, "The item failed to create.", err)
		return
	}

	fibonacciSpiralMatrix := dto.FibonacciSpiralMatrixDto{}

	ctx.JSON(http.StatusOK, fibonacciSpiralMatrix.FromModel(spiralMatrix))
}
