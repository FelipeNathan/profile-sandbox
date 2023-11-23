package scope_service

import (
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/repository/sandbox_repository"
	"time"
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

	scopes := sandbox_repository.LoadAll()
	for _, scope := range scopes {
		scope.LoadedAt = time.Now().Format(time.RFC3339)
	}
	return sandbox_repository.LoadAll()
}
