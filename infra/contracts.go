package infra

import (
	"io"
)

type IoAdapter interface {
	Getwd() (dir string, err error)
	Join(elem ...string) string
	Open(name string) (io.ReadCloser, error)
}
