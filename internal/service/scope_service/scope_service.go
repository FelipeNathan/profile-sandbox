package scope_service

import (
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/repository/sandbox_repository"
	"time"
)

func Lock(scope string, userId string, minutes int) *sandbox.Scope {
	savedScope := Find(scope)
	if savedScope.IsLocked() {
		return savedScope
	}

	duration := time.Minute * time.Duration(minutes)
	savedScope.ToLocked(userId, duration)
	savedScope = sandbox_repository.Save(savedScope)
	return savedScope
}

func Unlock(scope string) *sandbox.Scope {
	savedScope := Find(scope)
	if savedScope.IsUnlocked() {
		return savedScope
	}

	savedScope.ToUnlocked()
	savedScope = sandbox_repository.Save(savedScope)
	return savedScope
}

func Remove(scope string) {
	sandbox_repository.Remove(scope)
}

func Find(scope string) *sandbox.Scope {
	savedScope := sandbox_repository.FindBy(scope)
	if savedScope != nil {
		unlockIfTimedOut(savedScope)
		return savedScope
	}

	return sandbox.NewAvailableScope(scope)
}

func FindAll() []*sandbox.Scope {
	scopes := sandbox_repository.LoadAll()
	for _, scope := range scopes {
		unlockIfTimedOut(scope)
		scope.LoadedAt = time.Now().Format(time.RFC3339)
	}
	return scopes
}

func unlockIfTimedOut(scope *sandbox.Scope) {
	if scope.IsLocked() && time.Now().After(scope.FinishAt) {
		UnlockInstance(scope)
	}
}

func UnlockInstance(scope *sandbox.Scope) {
	scope.ToUnlocked()
	_ = sandbox_repository.Save(scope)
}
