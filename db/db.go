package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"opentodo/config"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetDBConfig() DBConfig {
	envDBHost := config.GetConfig(config.Config{
		Name:         "DB_HOST",
		DefaultValue: "localhost",
	})
	envDBUser := config.GetConfig(config.Config{
		Name:         "DB_USER",
		DefaultValue: "postgres",
	})
	envDBPassword := config.GetConfig(config.Config{
		Name:         "DB_PASSWORD",
		DefaultValue: "",
	})
	envDBName := config.GetConfig(config.Config{
		Name:         "DB_NAME",
		DefaultValue: "opentodo",
	})
	envSSLMode := config.GetConfig(config.Config{
		Name:         "SSL_MODE",
		DefaultValue: "require",
	})

	return DBConfig{
		Host:     envDBHost,
		User:     envDBUser,
		Password: envDBPassword,
		DBName:   envDBName,
		SSLMode:  envSSLMode,
	}
}

// Connects to postgres bases on environment setting and returns a DB instance
func Connect(dbConfig DBConfig) (*gorm.DB, error) {

	// Connect to postgres
	connectionString := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		dbConfig.Host, dbConfig.User, dbConfig.DBName, dbConfig.SSLMode, dbConfig.SSLMode)

	return gorm.Open("postgres", connectionString)
}
