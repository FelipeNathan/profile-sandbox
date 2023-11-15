package main

import (
	"fmt"
	"net/http"
	"profile-sandbox/cmd/controller"
)

func main() {
	http.HandleFunc("/slack", controller.Slack)
	http.HandleFunc("/status", controller.Status)
	http.HandleFunc("/status/command", controller.Command)

	fmt.Println("Running")
	http.ListenAndServe(":8080", nil)
}
