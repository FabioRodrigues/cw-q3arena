package logger

import (
	"cw-q3arena/services"
	"fmt"
)

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
