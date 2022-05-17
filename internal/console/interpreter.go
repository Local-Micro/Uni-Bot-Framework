package console

import (
	"fmt"
	"strings"
)

type Interpreter struct {
	handler *Handler
	prompt string  // 提示符
	text string  // 输入内容
}

func NewInterpreter() (*Interpreter, error) {
	handler, err := NewHandler()
	if err != nil {
		return nil, err
	}
	return &Interpreter{
		handler: handler,
		prompt: "bot > ",
		text: "",
	}, nil
}

func (i *Interpreter)Start() error {
	for {
		fmt.Print(i.prompt)
		fmt.Scanln(&i.text)
		if len(strings.Fields(i.text)) == 0 {
			i.text = ""
			continue
		} else {
			err := i.handler.Start(i.text)
			if err != nil {
				if err.Error() == "Unknown command!" {
					continue
				} else {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}
