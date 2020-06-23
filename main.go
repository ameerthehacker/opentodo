package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	port, err := strconv.ParseInt(envPort, 10, 32)

	if err != nil {
		log.Panicln("Invalid port number", envPort)
	}

	router := gin.Default()

	controllers.TodosController(router)

	router.Run(fmt.Sprintf(":%d", port))
}
