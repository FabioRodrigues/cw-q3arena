package logger

type Mock struct {
	InfoFn  func(args ...interface{})
	ErrorFn func(args ...interface{})
}

func (m Mock) Info(args ...interface{}) {
	if m.InfoFn != nil {
		m.InfoFn(args...)
	}
}

func (m Mock) Error(args ...interface{}) {
	if m.ErrorFn != nil {
		m.ErrorFn(args...)
	}
}
