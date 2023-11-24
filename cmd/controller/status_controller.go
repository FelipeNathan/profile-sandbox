package controller

import (
	"html/template"
	"net/http"
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/service/sandbox_service"
	"sort"
)

func Status(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		_, _ = writer.Write([]byte("Failed to read message" + err.Error()))
		return
	}

	scopes := sandbox_service.LoadAll()
	sort.Slice(scopes, func(i, j int) bool {
		return scopes[i].Name < scopes[j].Name
	})
	writePage(writer, scopes, err)
}

func Command(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		_, _ = writer.Write([]byte("Failed to read message" + err.Error()))
		return
	}

	req := &sandbox.Request{
		Command: sandbox.Command(request.Form["command"][0]),
		Scope:   request.Form["scope"][0],
		UserId:  request.Form["user_id"][0],
	}

	_, err = sandbox_service.HandleCommand(req)
	http.Redirect(writer, request, "/status", 302)
}

func writePage(w http.ResponseWriter, scope []*sandbox.Scope, err error) {
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	t, err := template.ParseFiles("./web/index.html")
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = t.Execute(w, scope)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
