package main

import (
	"fmt"
	"net/http"
	"profile-sandbox/cmd/controller"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/status", 302)
	})
	http.HandleFunc("/slack", controller.Slack)
	http.HandleFunc("/status", controller.Status)
	http.HandleFunc("/status/command", controller.Command)
	http.HandleFunc("/should_i_activate_ftu", controller.ShouldIActivateFTU)

	http.Handle("/static/", staticHandler())

	fmt.Println("Running")
	_ = http.ListenAndServe(":8080", nil)
}

func staticHandler() http.Handler {
	staticFileHandler := http.FileServer(http.Dir("./web/static/"))
	return http.StripPrefix("/static/", staticFileHandler)
}
