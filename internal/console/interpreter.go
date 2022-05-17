package console

import (
	"fmt"
)

type Interpreter struct {
	handler *Handler // 处理器
	prompt  string   // 提示符
	text    string   // 输入内容
}

func NewInterpreter() (*Interpreter, error) {
	handler, err := NewHandler()
	if err != nil {
		return nil, err
	}
	return &Interpreter{
		handler: handler,
		prompt:  "bot > ",
		text:    "",
	}, nil
}

func (i *Interpreter) Start() error {
	for {
		fmt.Print(i.prompt)
		_, err := fmt.Scanln(&i.text)
		if err != nil {
			// 空行
			if err.Error() == "unexpected newline" {
				continue
			} else {
				return err
			}
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
