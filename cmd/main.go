package main

import (
	"fmt"
	"net/http"
	"profile-sandbox/cmd/controller"

	"github.com/go-chi/chi/v5"
)

const api = "/profile"

func main() {
	r := chi.NewRouter()
	r.Route(api, func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, api+"/status", 302)
		})
		r.Get("/status", controller.Status)
		r.Get("/status/command", controller.Command)
		r.Get("/should_i_activate_ftu", controller.ShouldIActivateFTU)

		r.Handle("/static/*", staticHandler())
	})

	fmt.Println("Running")
	_ = http.ListenAndServe(":8080", r)
}

func staticHandler() http.Handler {
	staticFileHandler := http.FileServer(http.Dir("./web/static/"))
	return http.StripPrefix(api+"/static/", staticFileHandler)
}
