package entities

type Kill struct {
	Timestamp  string
	KillerId   int
	VictimId   int
	MethodId   int
	KillerName string
	VictimName string
	MethodName string
}
