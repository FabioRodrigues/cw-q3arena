package ioadapter

import (
	"io"
)

type Mock struct {
	GetWdFn func() (string, error)
	JoinFn  func(elem ...string) string
	OpenFn  func(name string) (io.ReadCloser, error)
}

func (m Mock) Getwd() (dir string, err error) {
	if m.GetWdFn != nil {
		return m.GetWdFn()
	}
	return "", nil
}

func (m Mock) Join(elem ...string) string {
	if m.JoinFn != nil {
		return m.JoinFn(elem...)
	}
	return ""
}

func (m Mock) Open(name string) (io.ReadCloser, error) {
	if m.OpenFn != nil {
		return m.OpenFn(name)
	}
	return nil, nil
}

type MockReadCloser struct {
	ReadFn  func(p []byte) (n int, err error)
	CloseFn func() error
}

func (m *MockReadCloser) Read(p []byte) (int, error) {
	return m.ReadFn(p)
}

func (m *MockReadCloser) Close() error {
	if m.CloseFn != nil {
		return m.CloseFn()
	}
	return nil
}
