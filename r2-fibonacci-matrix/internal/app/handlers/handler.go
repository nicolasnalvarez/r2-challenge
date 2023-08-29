package handlers

import (
	"net/http"
	"r2-fibonacci-matrix/internal/app/dtos"
	"r2-fibonacci-matrix/internal/app/services"
	userservice "r2-fibonacci-matrix/internal/user/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	Handler struct {
		fibonacciMatrixService services.FibonacciMatrixService
		userService            userservice.UserService
	}
)

func NewMatrixHandler(fibonacciMatrixService services.FibonacciMatrixService, userService userservice.UserService) *Handler {
	return &Handler{
		fibonacciMatrixService: fibonacciMatrixService,
		userService:            userService,
	}
}

func (h *Handler) GenerateSpiralMatrix(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": "error trying to get email from request"})
		return
	}
	if _, err := h.userService.FindUserByEmail(email.(string)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": "error trying to find user in the db"})
		return
	}

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

	response := dtos.MatrixResponse{Rows: h.fibonacciMatrixService.GenerateMatrix(rowsInt, colsInt)}

	ctx.JSON(http.StatusOK, response)
}
