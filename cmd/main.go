package main

import (
	"fmt"
	"net/http"
	"profile-sandbox/cmd/controller"
)

func main() {
	http.HandleFunc("/profile", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/profile/status", 302)
	})
	http.HandleFunc("/profile/slack", controller.Slack)
	http.HandleFunc("/profile/status", controller.Status)
	http.HandleFunc("/profile/status/command", controller.Command)
	http.HandleFunc("/profile/should_i_activate_ftu", controller.ShouldIActivateFTU)

	http.Handle("/profile/static/", staticHandler())

	fmt.Println("Running")
	_ = http.ListenAndServe(":8080", nil)
}

func staticHandler() http.Handler {
	staticFileHandler := http.FileServer(http.Dir("./web/static/"))
	return http.StripPrefix("/profile/static/", staticFileHandler)
}
