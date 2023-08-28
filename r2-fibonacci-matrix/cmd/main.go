package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"r2-fibonacci-matrix/http"
)

func main() {
	router, err := http.NewRouter()
	if err != nil {
		log.Fatal().Err(err).Msg("error creating router")
		os.Exit(1)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("error starting router")
		os.Exit(1)
	}
}
