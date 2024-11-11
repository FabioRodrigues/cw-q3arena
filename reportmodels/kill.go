package reportmodels

import "cw-q3arena/constants"

// KillReport consolidates all the Kill events for a specific report
type KillReport struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
	killsMap   map[int]int
	playersMap map[int]string
}

func (k *KillReport) AddPlayers(players map[int]string) {
	if k.playersMap == nil {
		k.playersMap = map[int]string{}
	}

	for id, name := range players {
		if name == constants.WorldUsername {
			continue
		}
		k.playersMap[id] = name
	}

	k.Players = []string{}
	for _, name := range k.playersMap {
		k.Players = append(k.Players, name)
	}
}

func (k *KillReport) AddKill(killerId int, killerName string, victimId int) {
	if k.killsMap == nil {
		k.killsMap = map[int]int{}
	}

	k.TotalKills += 1
	if killerName == constants.WorldUsername {
		kills, _ := k.killsMap[victimId]
		k.killsMap[victimId] = kills - 1
		return
	}

	kills, _ := k.killsMap[killerId]
	k.killsMap[killerId] = kills + 1
}

func (k *KillReport) GetKills() map[string]int {
	// We need to have this since world can overkill and a player can be negative.
	// Since we allow parallel events, it's impossible to know total kills of a player before all the events
	// are received
	result := map[string]int{}
	for id, kills := range k.killsMap {
		player := k.playersMap[id]
		if kills < 0 {
			kills = 0
		}

		result[player] = kills
	}

	return result
}
