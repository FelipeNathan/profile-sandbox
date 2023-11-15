package scope_service

import (
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/repository/sandbox_repository"
)

func HandleCommand(request *sandbox.Request) (*sandbox.Scope, error) {

	if err := request.IsKnownCommand(); err != nil {
		return nil, err
	}
	
	var scope *sandbox.Scope
	switch request.Command {
	case sandbox.Lock:
		scope = Lock(request.Scope, request.UserId)
	case sandbox.Unlock:
		scope = Unlock(request.Scope)
	case sandbox.Remove:
		Remove(request.Scope)
		fallthrough
	case sandbox.Status:
		scope = Find(request.Scope)
	default:
		return nil, sandbox.CommandNotFound
	}

	return scope, nil
}

func LoadAll() []*sandbox.Scope {
	return sandbox_repository.LoadAll()
}
