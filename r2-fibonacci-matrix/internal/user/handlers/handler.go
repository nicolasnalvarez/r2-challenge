package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"r2-fibonacci-matrix/internal/user/dtos"
	"r2-fibonacci-matrix/internal/user/services"
)

type (
	Handler struct {
		userService services.UserService
	}
)

func NewUserHandler(userService services.UserService) *Handler {
	return &Handler{
		userService: userService,
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
