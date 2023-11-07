package handler

import (
	"CustomerLabsTest/model"
	"CustomerLabsTest/worker"
	"encoding/json"
	"net/http"
)

func HandleJSONRequest(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	inputData := make(chan map[string]interface{})
	outputData := make(chan *model.OuputData)

	go worker.Worker(inputData, outputData)
	inputData <- body
	close(inputData)

	outputbytes, err := json.Marshal(<-outputData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(outputbytes)
}
