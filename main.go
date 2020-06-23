package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"opentodo/controllers"
	"os"
	"strconv"
)

func main() {
	envPort := os.Getenv("HTTP_PORT")

	if len(envPort) == 0 {
		envPort = "8080"
	}

	port, err := strconv.ParseInt(envPort, 10, 32)

	if err != nil {
		log.Panicln("Invalid port number", envPort)
	}

	router := gin.Default()

	controllers.TodosController(router)

	router.Run(fmt.Sprintf(":%d", port))
}
