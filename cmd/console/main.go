package main

import (
	. "github.com/Local-Micro/Uni-Bot-Framework/internal/console"
)

func main() {
	console, err := NewInterpreter()
	if err != nil {
		panic(err)
	}
	err = console.Start()
	if err != nil {
		panic(err)
	}
}
