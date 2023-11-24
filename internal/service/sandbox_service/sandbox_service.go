package sandbox_service

import (
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/service/scope_service"
)

func HandleCommand(request *sandbox.Request) (*sandbox.Scope, error) {

	if err := request.IsKnownCommand(); err != nil {
		return nil, err
	}

	var scope *sandbox.Scope
	switch request.Command {
	case sandbox.Lock:
		scope = scope_service.Lock(request.Scope, request.UserId)
	case sandbox.Unlock:
		scope = scope_service.Unlock(request.Scope)
	case sandbox.Remove:
		scope_service.Remove(request.Scope)
		fallthrough
	case sandbox.Status:
		scope = scope_service.Find(request.Scope)
	default:
		return nil, sandbox.CommandNotFound
	}

	return scope, nil
}

func LoadAll() []*sandbox.Scope {
	return scope_service.FindAll()
}
