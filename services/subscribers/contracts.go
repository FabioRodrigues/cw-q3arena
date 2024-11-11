package subscribers

type Subscriber interface {
	Receive(gameId string, data any)
	GetData(gameId string) (map[string]any, error)
}
