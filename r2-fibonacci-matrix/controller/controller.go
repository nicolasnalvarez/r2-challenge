package controller

import (
	"net/http"
	"r2-fibonacci-matrix/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	Controller struct {
		fibonacciMatrixService service.FibonacciMatrixService
	}
)

func NewController(fibonacciMatrixService service.FibonacciMatrixService) *Controller {
	return &Controller{
		fibonacciMatrixService: fibonacciMatrixService,
	}
}

func (c *Controller) GenerateSpiralMatrix(ctx *gin.Context) {
	rows := ctx.Query("rows")
	cols := ctx.Query("columns")

	rowsInt, err := strconv.Atoi(rows)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "Invalid value for rows"})
		return
	}

	colsInt, err := strconv.Atoi(cols)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "Invalid value for columns"})
		return
	}

	if rowsInt <= 0 || colsInt <= 0 || rowsInt > 9 || colsInt > 9 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "rows and columns values must be between 0 and 9"})
		return
	}

	response := MatrixResponse{Rows: c.fibonacciMatrixService.GenerateMatrix(rowsInt, colsInt)}

	ctx.JSON(http.StatusOK, response)
}
