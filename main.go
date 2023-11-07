package main

import (
	"CustomerLabsTest/handler"
	"log"
	"net/http"
)

func main() {
	const port string = ":8080"

	http.HandleFunc("/", handler.HandleJSONRequest)
	log.Println("Server Started Listening at",port)
	http.ListenAndServe(port, nil)
}
