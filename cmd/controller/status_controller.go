package controller

import (
	"html/template"
	"math/rand"
	"net/http"
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/service/sandbox_service"
	"sort"
	"strconv"
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

	minutes, err := strconv.Atoi(request.Form["minutes"][0])
	if err != nil {
		minutes = 10
	}

	req := &sandbox.Request{
		Command: sandbox.Command(request.Form["command"][0]),
		Scope:   request.Form["scope"][0],
		UserId:  request.Form["user_id"][0],
		Minutes: minutes,
	}

	_, _ = sandbox_service.HandleCommand(req)
	http.Redirect(writer, request, "/profile/status", 302)
}

func writePage(w http.ResponseWriter, scope []*sandbox.Scope, err error) {
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	t, err := template.ParseFiles("./web/template/index.html")
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

func ShouldIActivateFTU(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles("./web/template/ftu.html")
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	noLinks := []string{
		"https://cdn.ponly.com/wp-content/uploads/No-Memes-1.jpg",
		"https://www.liveabout.com/thmb/dheNUTdHLRC2RbyLDr6EfI1AaFI=/750x0/filters:no_upscale():max_bytes(150000):strip_icc():format(webp)/RageFaceNo-5ae79bcd3128340037321cb4.jpg",
		"https://cdn.ponly.com/wp-content/uploads/No-Memes-5.jpg",
		"https://cdn.ponly.com/wp-content/uploads/No-Memes-4.jpg",
		"https://cdn.ponly.com/wp-content/uploads/No-Memes-3.jpg",
		"https://cdn.ponly.com/wp-content/uploads/No-Memes-2.jpg",
		"https://cdn.ponly.com/wp-content/uploads/No-Memes-1.jpg",
		"https://cdn.ponly.com/wp-content/uploads/No-Memes-15.jpg",
	}

	index := rand.Intn(len(noLinks))
	s := struct {
		Link string
	}{
		Link: noLinks[index],
	}

	err = t.Execute(w, s)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
