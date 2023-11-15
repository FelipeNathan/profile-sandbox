package sandbox

type StatusType string

const (
	Available StatusType = "Available"
	Locked               = "Locked"
)

type Scope struct {
	Name     string     `json:"name"`
	Status   StatusType `json:"status"`
	LockedBy string     `json:"lockedBy"`
}

func (scope *Scope) ToLocked(lockedBy string) {
	scope.LockedBy = lockedBy
	scope.Status = Locked
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
