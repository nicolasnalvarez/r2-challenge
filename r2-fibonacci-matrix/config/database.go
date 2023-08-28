package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDatabase creates a mysql db connection and stores it in the GlobalDB variable
// It reads the environment variables from the .env file and uses them to create the connection
// It returns the db connection and an error if the connection fails
func InitDatabase() (*gorm.DB, error) {
	// Read the environment variables from the .env file
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal().Msg("Error reading .env file")
	}
	// Create the data source name (DSN) using the environment variables
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DATABASE_HOST"],
		config["DB_DATABASE"],
	)
	// Create the connection and store it in the GlobalDB variable
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
