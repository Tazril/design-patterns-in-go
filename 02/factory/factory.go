package factory

import (
	"fmt"
)

type IMobileOS interface {
	Load()
}

// Android subclass
type Android struct {
}

// IOS subclass
type IOS struct {
}

func (os *Android) Load() {
	fmt.Println("Loading Android Operating System")
}

func (os *IOS) Load() {
	fmt.Println("Loading IOS Operating System")
}

// NewMobileOS [Factory Method] Returns Object based on the tag, hiding implementation
func NewMobileOS(tag string) IMobileOS {
	switch tag {
	case "opensource":
		return &Android{}
	case "elite":
		return &IOS{}
	default:
		return nil
	}
}
