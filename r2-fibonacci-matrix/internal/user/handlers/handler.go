package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"r2-fibonacci-matrix/auth"
	"r2-fibonacci-matrix/internal/user/dtos"
	"r2-fibonacci-matrix/internal/user/services"
)

const refreshAccessTokenError = "could not refresh access token"

type (
	Handler struct {
		userService services.UserService
		jwtService  auth.JwtService
	}
)

func NewUserHandler(userService services.UserService, jwtService auth.JwtService) *Handler {
	return &Handler{
		userService: userService,
		jwtService:  jwtService,
	}
}

// Register is a function that handles user register
func (h Handler) Register(ctx *gin.Context) {
	var registerRequest dtos.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		log.Error().Err(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid input values inserted",
		})
		return
	}
	err = h.userService.Save(registerRequest)
	if err != nil {
		log.Error().Err(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error trying to create user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

// Login is a function that handles user login
func (h Handler) Login(ctx *gin.Context) {
	var loginRequest dtos.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		log.Error().Err(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid input values inserted",
		})
		return
	}

	loginResponse, err := h.userService.Login(loginRequest)
	if err != nil {
		log.Error().Err(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "there was a problem trying to login user",
		})
		return
	}

	ctx.JSON(200, loginResponse)
}

// Refresh is a function that handles users token refresh
func (h Handler) Refresh(ctx *gin.Context) {
	refreshToken := ctx.Query("refresh_token")
	if refreshToken == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": refreshAccessTokenError,
		})
		return
	}
	claims, err := h.jwtService.ValidateRefreshToken(refreshToken)
	if err != nil {
		log.Error().Err(err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": refreshAccessTokenError,
		})
		return
	}

	newAccessToken, err := h.jwtService.GenerateToken(claims.Email)
	if err != nil {
		log.Error().Err(err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": refreshAccessTokenError,
		})
		return
	}

	ctx.JSON(200, dtos.LoginResponse{
		Token:        newAccessToken,
		RefreshToken: refreshToken,
	})
}
