package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"r2-fibonacci-matrix/internal/user/entities"
)

const connectionError = "error setting up DB connection: %s"

// InitDatabase creates a mysql db connection
func InitDatabase(envVariables map[string]string) *gorm.DB {
	// Create the data source name (DSN) using the environment variables
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		envVariables["DB_USERNAME"],
		envVariables["DB_PASSWORD"],
		envVariables["DATABASE_HOST"],
		envVariables["DB_DATABASE"],
	)
	// Create the connection and store it in the GlobalDB variable
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fatalConnection(err)
	}

	createSchema := db.Exec("CREATE SCHEMA IF NOT EXISTS r2")
	if createSchema.Error != nil {
		fatalConnection(createSchema.Error)
	}

	setSchema := db.Exec("USE r2")
	if setSchema.Error != nil {
		fatalConnection(setSchema.Error)
	}

	if err = db.AutoMigrate(entities.User{}); err != nil {
		fatalConnection(err)
	}

	return db
}

func fatalConnection(err error) {
	log.Fatal().Err(err).Msg(fmt.Sprintf(connectionError, err.Error()))
}
