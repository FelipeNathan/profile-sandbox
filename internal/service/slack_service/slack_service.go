package slack_service

import (
	"fmt"
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/model/slack"
	"profile-sandbox/internal/service/scope_service"
	"strings"
)

func HandleCommand(body slack.Body) *slack.ResponseBlock {
	request, err := parseBody(body)
	if err != nil {
		return unknownCommandResponseBlock()
	}

	scope, err := scope_service.HandleCommand(request)
	if err != nil {
		return unknownCommandResponseBlock()
	}

	return statusResponseBlock(scope)
}

func parseBody(body slack.Body) (*sandbox.Request, error) {
	params := strings.Split(strings.TrimSpace(body.Text), " ")

	if len(params) != 2 {
		return nil, sandbox.IllegalArguments
	}

	if strings.TrimSpace(params[0]) == "" {
		return nil, sandbox.CommandNotFound
	}

	request := &sandbox.Request{
		Command: sandbox.Command(params[0]),
		Scope:   params[1],
		UserId:  body.UserId,
	}

	if err := request.IsKnownCommand(); err != nil {
		return nil, err
	}

	return request, nil
}

func statusResponseBlock(scope *sandbox.Scope) *slack.ResponseBlock {
	response := fmt.Sprintf("Sandbox %s is %s", scope.Name, scope.Status)
	if strings.TrimSpace(scope.LockedBy) != "" {
		response = response + " by <@" + scope.LockedBy + ">"
	}
	return buildResponseBlock(response)
}

func unknownCommandResponseBlock() *slack.ResponseBlock {
	unknownCommand := "Unknown command, the commands are `status <scope>`, `lock <scope>` and `unlock <scope>`"
	return buildResponseBlock(unknownCommand)
}

func buildResponseBlock(response string) *slack.ResponseBlock {
	return &slack.ResponseBlock{
		Text:         response,
		ResponseType: "in_channel",
		Type:         "section",
	}
}
