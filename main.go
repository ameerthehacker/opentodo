package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"opentodo/config"
	"opentodo/controllers"
	"opentodo/db"
	"strconv"
)

func main() {
	// Load the config from .env if possible
	_ = godotenv.Load()
	envPort := config.GetConfig(config.Config{
		Name:         "HTTP_PORT",
		DefaultValue: "8000",
	})

	// Connect to postgres
	dbConfig := db.GetDBConfig()
	db, err := db.Connect(dbConfig)
	defer db.Close()

	if err != nil {
		log.Fatalf("Unable to connect to db: %v", err)
	} else {
		fmt.Printf(
			"Connected to postgres {server=%s, dbName=%s, user=%s, sslMode=%s}",
			dbConfig.Host, dbConfig.DBName, dbConfig.User, dbConfig.SSLMode)
	}

	port, err := strconv.ParseInt(envPort, 10, 32)

	if err != nil {
		log.Panicln("Invalid port number", envPort)
	}

	router := gin.Default()

	controllers.TodosController(router)

	router.Run(fmt.Sprintf(":%d", port))
}
