package main

import (
	"log"
	"raspiserver/buzzer"
	"raspiserver/controller"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func configure() {
	buzzer.ConfigureBuzzer(18)
}

func start() {
	log.Printf("[INFO] Configuring I/O")

	// Configures GPIO pins
	configure()

	log.Printf("[INFO] Starting http server")

	// Configures routes
	engine = gin.New()
	engine.POST("/buzzer", controller.HandleBuzzerRequest)

	// Starts web server
	engine.Run(":2711")
}

func stop() {
	log.Printf("[INFO] Shutting down ...")
}

func main() {

	// Defer server stopping
	defer stop()

	// Starts server
	start()
}
