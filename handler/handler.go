package handler

import (
	"CustomerLabsTest/model"
	"CustomerLabsTest/worker"
	"bytes"
	"encoding/json"
	"fmt"
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
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Webhook sent successfully")
	} else {
		fmt.Println("Webhook request failed with status:", resp.Status)
	}
	w.WriteHeader(200)
}
