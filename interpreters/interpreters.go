package interpreters

import (
	"DummyAlerts/messages"
	"fmt"
)

type Interpreter interface {
	Interpret([]byte) (*messages.Message, error)
}

func GetInterpreter(name string) (Interpreter, error) {
	switch name {
	case "xo":
		return NewXOInterpreter(), nil
	default:
		return nil, fmt.Errorf("no matching interpreter found")
	}
}
