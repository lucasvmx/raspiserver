package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"raspiserver/buzzer"
	"raspiserver/model"
)

func logRequest(res http.ResponseWriter, req *http.Request, statusCode *int) {
	log.Printf("[INFO] [%v] %v %v: %v", &res, req.Method, req.URL.Path, *statusCode)
}

func readBody(req *http.Request) []byte {
	defer req.Body.Close()
	data, fail := io.ReadAll(req.Body)
	if fail != nil {
		log.Printf("[ERROR] Can't read body: %v", fail)
		return nil
	}

	return data
}

func HandleBuzzer(writer http.ResponseWriter, req *http.Request) {
	var statusCode = http.StatusOK
	var message model.Buzzer

	defer logRequest(writer, req, &statusCode)

	// Decode payload
	payload := readBody(req)

	err := json.Unmarshal(payload, &message)
	if err != nil {
		log.Printf("[ERROR] Can't decode payload: %v", err)
		statusCode = http.StatusBadRequest
	}

	if message.TempoSegundos <= 0 || message.Estado != 0 && message.Estado != 1 {
		statusCode = http.StatusBadRequest
	}

	if statusCode == http.StatusOK {
		go func() {
			buzzer.BeepBuzzer(18, uint(message.TempoSegundos)*1000, uint(message.QuantidadeVezes))
		}()
	}

	writer.WriteHeader(statusCode)
}
