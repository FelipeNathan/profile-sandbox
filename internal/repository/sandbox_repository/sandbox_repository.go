package sandbox_repository

import (
	"profile-sandbox/internal/model/sandbox"
	"profile-sandbox/internal/repository/config"
)

func FindBy(scope string) *sandbox.Scope {
	savedScope := &sandbox.Scope{
		Name: scope,
	}
	config.DB.Find(savedScope)
	return savedScope
}

func Save(scope *sandbox.Scope) *sandbox.Scope {
	config.DB.Save(&scope)
	return scope
}

func Remove(scope string) {
	config.DB.Delete(&sandbox.Scope{}, scope)
}

func LoadAll() []*sandbox.Scope {
	var scopes []*sandbox.Scope
	config.DB.Find(&scopes)
	return scopes
}
