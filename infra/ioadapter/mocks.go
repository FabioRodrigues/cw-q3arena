package ioadapter

import "os"

type Mock struct {
	GetWdFn func() (string, error)
	JoinFn  func(elem ...string) string
	OpenFn  func(name string) (*os.File, error)
}

func (m *Mock) Getwd() (dir string, err error) {
	if m.GetWdFn != nil {
		return m.GetWdFn()
	}
	return "", nil
}

func (m *Mock) Join(elem ...string) string {
	if m.JoinFn != nil {
		return m.JoinFn(elem...)
	}
	return ""
}

func (m *Mock) Open(name string) (*os.File, error) {
	if m.OpenFn != nil {
		return m.OpenFn(name)
	}
	return nil, nil
}
