package console

import (
	"errors"
	"strings"
	. "github.com/Local-Micro/Uni-Bot-Framework/internal/console/commands"
)

type Handler struct {
	commands map[string]Command
}

func NewHandler() (*Handler, error) {
	commands := make(map[string]Command)
	exit, err := NewExit()
	if err != nil {
		return nil, err
	}
	commands["exit"] = exit
	return &Handler{
		commands: commands,
	}, nil
}

func (h *Handler)Start(text string) error {
	args := strings.Fields(text)
	if value, ok := h.commands[args[0]]; ok {
		err := value.Start(args)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Unknown command!")
	}
	return nil
}
