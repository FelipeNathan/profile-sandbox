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

	savedScope.ToLocked(userId)
	savedScope = sandbox_repository.Save(savedScope)
	go lockTimeout(scope, time.Hour)
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
		return savedScope
	}

	return sandbox.NewAvailableScope(scope)
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
