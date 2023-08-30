package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"r2-fibonacci-matrix/auth"
	"r2-fibonacci-matrix/config"
	matrixhandler "r2-fibonacci-matrix/internal/app/handlers"
	matrixservice "r2-fibonacci-matrix/internal/app/services"
	userhandler "r2-fibonacci-matrix/internal/user/handlers"
	"r2-fibonacci-matrix/internal/user/repositories"
	userservice "r2-fibonacci-matrix/internal/user/services"
	"strconv"
)

const (
	accessTokenSecretKeyEnv          = "ACCESS_TOKEN_SECRET_KEY"
	accessTokenExpirationMinutesEnv  = "ACCESS_TOKEN_EXPIRATION_MINUTES"
	refreshTokenSecretKeyEnv         = "REFRESH_TOKEN_SECRET_KEY"
	refreshTokenExpirationMinutesEnv = "REFRESH_TOKEN_EXPIRATION_MINUTES"
	tokenIssuer                      = "TOKEN_ISSUER"
)

func NewRouter() (*gin.Engine, error) {
	router := gin.Default()

	// read env variables
	envVariables := config.ReadEnvVariables()

	// initialize db connection
	db := config.InitDatabase(envVariables)

	// jwt
	expAccessToken, err := strconv.Atoi(envVariables[accessTokenExpirationMinutesEnv])
	if err != nil {
		log.Fatal().Err(err).Msg("error converting string to int")
	}
	expRefreshToken, err := strconv.Atoi(envVariables[refreshTokenExpirationMinutesEnv])
	if err != nil {
		log.Fatal().Err(err).Msg("error converting string to int")
	}
	jwtService := auth.NewJwtService(envVariables[accessTokenSecretKeyEnv], envVariables[refreshTokenSecretKeyEnv],
		envVariables[tokenIssuer], expAccessToken, expRefreshToken)

	// user
	userRepository := repositories.NewUserRepository(db)
	userService := userservice.NewUserService(userRepository, jwtService)
	userHandler := userhandler.NewUserHandler(userService, jwtService)

	// matrix
	matrixService := matrixservice.NewMatrixService()
	matrixHandler := matrixhandler.NewMatrixHandler(matrixService, userService, jwtService)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	apiRouter := router.Group("/api")

	// user routes
	apiRouter.POST("/register", userHandler.Register)
	apiRouter.POST("/login", userHandler.Login)
	apiRouter.GET("/refresh", userHandler.Refresh)

	// matrix routes
	apiRouter.GET("/matrix", matrixHandler.GenerateSpiralMatrix)

	return router, nil
}
