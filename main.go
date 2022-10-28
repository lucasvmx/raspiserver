package main

import (
	"log"
	"raspiserver/buzzer"
	"raspiserver/controller"
	"raspiserver/utils"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func configure() {
	utils.ReadSettings()
	pin := utils.GetConfig().BuzzerIoPin
	buzzer.ConfigureBuzzer(pin)
	log.Printf("[INFO] Configured buzzer on pin %v", pin)
}

func start() {
	log.Printf("[INFO] Configuring I/O")

	// Configures GPIO pins
	configure()

	log.Printf("[INFO] Starting http server")

	// Configures routes
	engine = gin.New()
	gin.SetMode(gin.ReleaseMode)
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
