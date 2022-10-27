package main

import (
	"log"
	"net/http"
	"raspiserver/buzzer"
	"raspiserver/controller"
)

var mux *http.ServeMux

func configure() {
	buzzer.ConfigureBuzzer(18)
}

func start() {
	log.Printf("[INFO] Configuring I/O")

	// Configures GPIO pins
	configure()

	log.Printf("[INFO] Starting http server")

	// Starts web server
	mux = http.NewServeMux()
	mux.HandleFunc("/lampada", controller.HandleBuzzer)
	http.ListenAndServe("0.0.0.0:2711", mux)
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
