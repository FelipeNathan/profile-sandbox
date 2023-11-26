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

	http.Handle("/static/", staticHandler())

	fmt.Println("Running")
	http.ListenAndServe(":8080", nil)
}

func staticHandler() http.Handler {
	staticFileHandler := http.FileServer(http.Dir("./web/static/"))
	return http.StripPrefix("/static/", staticFileHandler)
}
