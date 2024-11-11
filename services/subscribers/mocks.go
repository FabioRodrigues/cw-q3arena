package subscribers

type Mock struct {
	ReceiveFn func(gameId string, data any)
	GetDataFn func(gameId string) (map[string]any, error)
}

func (m Mock) Receive(gameId string, data any) {
	if m.ReceiveFn != nil {
		m.ReceiveFn(gameId, data)
	}
}

func (m Mock) GetData(gameId string) (map[string]any, error) {
	if m.GetDataFn != nil {
		return m.GetDataFn(gameId)
	}
	return make(map[string]any), nil
}
