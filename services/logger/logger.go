package logger

import (
	"cw-q3arena/services"
	"fmt"
)

// Logger is an extremely simple & silly log implementation, just to give an idea of a logger
type Logger struct{}

func NewLogger() services.Logger {
	return &Logger{}
}

func (l Logger) Info(args ...interface{}) {
	fmt.Printf("[INFO] - %s\n", fmt.Sprint(args...))
}

func (l Logger) Error(args ...interface{}) {
	fmt.Printf("[ERROR] - %s\n", fmt.Sprint(args...))
}
