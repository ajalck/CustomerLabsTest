package handler

import (
	"CustomerLabsTest/model"
	"CustomerLabsTest/worker"
	"bytes"
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

	outputbytes, _ := json.Marshal(<-outputData)

	webhookURL := "https://webhook.site/0f5750a7-7c86-4c41-ac4b-6b261dcf7b62"
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(outputbytes))
	if err != nil {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte("Webhook sent successfully"))
	} else {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte("Webhook request failed with status: " + resp.Status))
	}
}
