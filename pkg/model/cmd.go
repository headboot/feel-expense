package model

import (
	"errors"

	"github.com/headboot/feel-expense/pkg"
)

type Command uint8

const (
	Add     Command = 1
	List    Command = 2
	Summary Command = 3
	Delete  Command = 4
	Help    Command = 5
)

var commandsMap = map[string]Command{
	"add":     Add,
	"list":    List,
	"summary": Summary,
	"delete":  Delete,
	"help":    Help,
}

func FromStringToCommand(sCmd string) Command{
	cmd, ok := commandsMap[sCmd]
	if !ok  {
		err := errors.New("Unknown command")
		pkg.ExitWithError(err)
	}
	return cmd
}
