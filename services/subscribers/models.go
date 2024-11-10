package subscribers

type KillSubscriberModel struct {
	Players map[int]string
}

type KillModel struct {
	KillerId int
	VictimId int
}
