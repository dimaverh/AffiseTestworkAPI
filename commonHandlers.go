package main

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/common/log"
)

// Answer contains template of Response-data.
type Answer map[string]interface{}

// Meta contains template of Response-metadata,
type Meta map[string]interface{}

// sendSuccessfulAnswer sends success response.
func sendSuccessfulAnswer(w http.ResponseWriter, status int, result interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(status)

	answer := Answer{
		"meta":     Meta{"status": status},
		"response": result,
	}

	res, err := json.Marshal(answer)
	if err != nil {
		log.Errorln(err)
	}

	_, err = w.Write(res)
	if err != nil {
		log.Errorln(err)
	}
}

// sendErrorAnswer sends a response with error.
func sendErrorAnswer(w http.ResponseWriter, errorCode int, message string) {

	log.Errorln(message)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusBadRequest)

	answer := Answer{
		"meta": Meta{
			"errorCode": errorCode,
			"message":   message,
		},
	}

	res, err := json.Marshal(answer)
	if err != nil {
		log.Errorln(err)
	}

	_, err = w.Write(res)
	if err != nil {
		log.Errorln(err)
	}
}
