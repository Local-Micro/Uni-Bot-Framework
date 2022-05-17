package commands

import (
	"os"
)

type Exit struct {
}

func NewExit() (*Exit, error) {
	return &Exit{}, nil
}

func (e *Exit) Start(args []string) error {
	os.Exit(0)
	return nil
}
