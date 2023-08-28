package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"r2-fibonacci-matrix/controller"
	"r2-fibonacci-matrix/service"
)

func main() {
	router := gin.Default()

	fibonacciMatrixService := service.NewService()
	fibonacciMatrixController := controller.NewController(fibonacciMatrixService)

	router.Use(cors.Default())
	router.GET("/matrix", fibonacciMatrixController.GenerateSpiralMatrix)

	if err := router.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("error starting router")
	}
}
