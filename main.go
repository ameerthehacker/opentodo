package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"opentodo/config"
	"opentodo/controllers"
	"strconv"
)

func main() {
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
