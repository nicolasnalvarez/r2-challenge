package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"r2-fibonacci-matrix/config"
	matrixhandler "r2-fibonacci-matrix/internal/app/handlers"
	matrixservice "r2-fibonacci-matrix/internal/app/services"
	userhandler "r2-fibonacci-matrix/internal/user/handlers"
	"r2-fibonacci-matrix/internal/user/repositories"
	userservice "r2-fibonacci-matrix/internal/user/services"
)

func NewRouter() (*gin.Engine, error) {
	router := gin.Default()

	// read env variables
	envVariables := config.ReadEnvVariables()

	// initialize db connection
	db := config.InitDatabase(envVariables)

	// user
	userRepository := repositories.NewUserRepository(db)
	userService := userservice.NewUserService(userRepository)
	userHandler := userhandler.NewUserHandler(userService)

	matrixService := matrixservice.NewMatrixService()
	matrixHandler := matrixhandler.NewMatrixHandler(matrixService)

	router.Use(cors.Default())
	apiRouter := router.Group("/api")

	// user routes
	apiRouter.POST("/register", userHandler.Register)
	apiRouter.POST("/login", userHandler.Login)

	// matrix routes
	apiRouter.GET("/matrix", matrixHandler.GenerateSpiralMatrix)

	return router, nil
}
