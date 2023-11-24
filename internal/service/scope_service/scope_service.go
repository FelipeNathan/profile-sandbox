package scope_service

import (
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/repository/sandbox_repository"
	"time"
)

var unlockTimers = map[string]*time.Timer{}

func Lock(scope string, userId string) *sandbox.Scope {
	savedScope := Find(scope)
	if savedScope.IsLocked() {
		return savedScope
	}

	duration := time.Hour
	savedScope.ToLocked(userId, duration)
	savedScope = sandbox_repository.Save(savedScope)
	go lockTimeout(scope, duration)
	return savedScope
}

func Unlock(scope string) *sandbox.Scope {
	savedScope := Find(scope)
	if savedScope.IsUnlocked() {
		return savedScope
	}

	savedScope.ToUnlocked()
	savedScope = sandbox_repository.Save(savedScope)
	go cancelTimeout(scope)
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
		scope = Unlock(scope.Name)
	}
}

func lockTimeout(scope string, t time.Duration) {
	timeout := time.NewTimer(t)
	unlockTimers[scope] = timeout
	<-timeout.C
	Unlock(scope)
}

func cancelTimeout(scope string) {
	timeout, found := unlockTimers[scope]
	if !found {
		return
	}

	if !timeout.Stop() {
		<-timeout.C
	}
}
