package infra

import "os"

type IoAdapter interface {
	Getwd() (dir string, err error)
	Join(elem ...string) string
	Open(name string) (*os.File, error)
}
