package ioadapter

import (
	"cw-q3arena/infra"
	"os"
	"path/filepath"
)

type IOAdapter struct {
}

func NewIOAdapter() infra.IoAdapter {
	return IOAdapter{}
}

func (a IOAdapter) Getwd() (dir string, err error) {
	return os.Getwd()
}

func (a IOAdapter) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func (a IOAdapter) Open(name string) (*os.File, error) {
	return os.Open(name)
}
