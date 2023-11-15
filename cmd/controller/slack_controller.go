package controller

import (
	"encoding/json"
	"net/http"
	"profile-sandbox/internal/model/slack"
	"profile-sandbox/internal/service/slack_service"
)

func Slack(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		_, _ = writer.Write([]byte("Failed to read message" + err.Error()))
		return
	}

	body := slack.Body{
		Text:   request.Form["text"][0],
		UserId: request.Form["user_id"][0],
	}

	responseBlock := slack_service.HandleCommand(body)
	writer.Header().Set("Content-Type", "application/json")

	sr, _ := json.Marshal(responseBlock)
	_, _ = writer.Write(sr)
}
