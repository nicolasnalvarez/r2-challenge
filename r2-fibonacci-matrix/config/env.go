package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// ReadEnvVariables Read the environment variables from the .env file
func ReadEnvVariables() map[string]string {
	config, err := godotenv.Read("config/.env")
	if err != nil {
		log.Fatal().Msgf("error reading .env file: %v", err)
	}
	return config
}
