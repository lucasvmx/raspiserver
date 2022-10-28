package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var filename string = "settings.json"

type RaspiServerSettings struct {
	BuzzerIoPin uint `json:"buzzer_io_pin"`
}

var BuzzerConfig *RaspiServerSettings

func GetConfig() *RaspiServerSettings {
	return BuzzerConfig
}

func ReadSettings() (cfg *RaspiServerSettings) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("[ERROR] Can't open %v: %v", filename, err)
	}

	defer file.Close()

	data, fail := io.ReadAll(file)
	if fail != nil {
		log.Fatalf("[CRITICAL] Failed to read %v: %v", filename, fail)
	}

	fail = json.Unmarshal(data, &cfg)
	if fail != nil {
		log.Fatalf("[CRITICAL] Failed to decode settings file: %v", fail)
	}

	return
}
