package core

import (
	"plugin"
)

type DriverManager struct {
	drivers map[string]Driver
}

func NewDriverManager() (*DriverManager, error) {
	return &DriverManager{
		plugins: make(map[string]Plugin),
	}, nil
}
