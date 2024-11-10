package subscribers

type Subscriber interface {
	Receive(gameId string, data any)
	GetSerializedReport(gameId string) (string, error)
}
