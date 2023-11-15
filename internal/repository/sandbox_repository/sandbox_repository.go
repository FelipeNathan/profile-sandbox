package sandbox_repository

import (
	"profile-sandbox/internal/model/sandbox"
)

var scopesStore = map[string]*sandbox.Scope{
	"auth-my-profile": {
		Name:     "auth-my-profile",
		Status:   sandbox.Available,
		LockedBy: "",
	},
	"auth-profile-automation-be": {
		Name:     "auth-profile-automation-be",
		Status:   sandbox.Available,
		LockedBy: "",
	},
	"user-identity-me": {
		Name:     "user-identity-me",
		Status:   sandbox.Available,
		LockedBy: "",
	},
	"user-identity-fe": {
		Name:     "user-identity-fe",
		Status:   sandbox.Available,
		LockedBy: "",
	},
	"my-data-fe": {
		Name:     "my-data-fe",
		Status:   sandbox.Available,
		LockedBy: "",
	},
}

func FindBy(scope string) *sandbox.Scope {
	savedScope, found := scopesStore[scope]
	if !found {
		return nil
	}
	return savedScope
}

func Save(scope *sandbox.Scope) *sandbox.Scope {
	scopesStore[scope.Name] = scope
	return scope
}

func Remove(scope string) {
	delete(scopesStore, scope)
}

func LoadAll() []*sandbox.Scope {
	var scopes []*sandbox.Scope
	for _, stored := range scopesStore {
		scopes = append(scopes, stored)
	}
	return scopes
}
