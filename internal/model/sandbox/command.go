package sandbox

import "errors"

var IllegalArguments = errors.New("command was not created properly")
var CommandNotFound = errors.New("command Not Found")

type Command string

const (
	Lock   Command = "lock"
	Unlock         = "unlock"
	Status         = "status"
	Remove         = "remove"
)

var allCommands = map[Command]bool{
	Lock:   true,
	Unlock: true,
	Status: true,
	Remove: true,
}

type Request struct {
	Command Command
	Scope   string
	UserId  string
}

func (r Request) IsKnownCommand() error {
	active, found := allCommands[r.Command]

	if !found || !active {
		return CommandNotFound
	}

	return nil
}
