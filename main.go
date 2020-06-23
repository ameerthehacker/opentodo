package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"opentodo/config"
	"opentodo/controllers"
	"strconv"
)

func main() {
	// Load the config from .env if possible
	_ = godotenv.Load()
	envPort := config.GetConfig(config.Config{
		Name:         "HTTP_PORT",
		DefaultValue: "8000",
	})
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

	// Connect to postgres
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", envDBHost, envDBUser, envDBName, envSSLMode, envDBPassword)
	db, err := gorm.Open("postgres", connectionString)
	defer db.Close()

	if err != nil {
		log.Fatalf("Unable to connect to db: %v", err)
	} else {
		fmt.Printf("Connected to postgres {server=%s, dbName=%s, user=%s, sslMode=%s}", envDBHost, envDBName, envDBUser, envSSLMode)
	}

	port, err := strconv.ParseInt(envPort, 10, 32)

	if err != nil {
		log.Panicln("Invalid port number", envPort)
	}

	router := gin.Default()

	controllers.TodosController(router)

	router.Run(fmt.Sprintf(":%d", port))
}
