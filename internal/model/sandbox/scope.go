package sandbox

import "time"

type StatusType string

const (
	Available StatusType = "Available"
	Locked               = "Locked"
)

type Scope struct {
	Name     string
	Status   StatusType
	LockedBy string
	LockedAt string
	FinishAt string
	LoadedAt string
}

func (scope *Scope) ToLocked(lockedBy string, duration time.Duration) {
	scope.LockedBy = lockedBy
	scope.Status = Locked

	lockedAt := time.Now()
	scope.LockedAt = lockedAt.Format(time.RFC3339)
	scope.FinishAt = lockedAt.Add(duration).Format(time.RFC3339)
}

func (scope *Scope) ToUnlocked() {
	scope.LockedBy = ""
	scope.Status = Available
}

func (scope *Scope) IsLocked() bool {
	return scope.Status == Locked
}

func (scope *Scope) IsUnlocked() bool {
	return !scope.IsLocked()
}

func NewAvailableScope(name string) *Scope {
	return newSandboxScope(name, Available, "")
}

func newSandboxScope(name string, status StatusType, userId string) *Scope {
	return &Scope{
		Name:     name,
		Status:   status,
		LockedBy: userId,
	}
}
