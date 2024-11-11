package loader

type Mock struct {
	LoadFn func(path string) (string, error)
}

func (m Mock) Load(path string) (string, error) {
	if m.LoadFn != nil {
		return m.LoadFn(path)
	}

	return "", nil
}
