package handlers

import (
	"errors"
	"net/http"
	"r2-fibonacci-matrix/auth"
	"r2-fibonacci-matrix/internal/app/dtos"
	"r2-fibonacci-matrix/internal/app/services"
	userservice "r2-fibonacci-matrix/internal/user/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	noAuthHeaderError  = "no authorization header provided"
	authTokenBadFormat = "incorrect Format of Authorization Token"
)

type (
	Handler struct {
		fibonacciMatrixService services.FibonacciMatrixService
		userService            userservice.UserService
		jwtService             auth.JwtService
	}
)

func NewMatrixHandler(fibonacciMatrixService services.FibonacciMatrixService, userService userservice.UserService, jwtService auth.JwtService) *Handler {
	return &Handler{
		fibonacciMatrixService: fibonacciMatrixService,
		userService:            userService,
		jwtService:             jwtService,
	}
}

func (h *Handler) GenerateSpiralMatrix(ctx *gin.Context) {
	if err := h.validateUserToken(ctx); err != nil {
		if err.Error() == noAuthHeaderError {
			ctx.AbortWithStatusJSON(http.StatusForbidden,
				gin.H{"error": noAuthHeaderError})
			return
		} else if err.Error() == authTokenBadFormat {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": authTokenBadFormat})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"error": err.Error()})
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

func (h *Handler) validateUserToken(ctx *gin.Context) error {
	clientToken := ctx.Request.Header.Get("Authorization")
	if clientToken == "" {
		return errors.New(noAuthHeaderError)
	}
	extractedToken := strings.Split(clientToken, "Bearer ")
	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		return errors.New(authTokenBadFormat)
	}

	// Validate the token
	claims, err := h.jwtService.ValidateAccessToken(clientToken)
	if err != nil {
		return err
	}

	if _, err := h.userService.FindUserByEmail(claims.Email); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": "error trying to find user in the db"})
		return err
	}
	return nil
}
