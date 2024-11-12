package loader

import (
	"bufio"
	"cw-q3arena/infra"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services"
	"encoding/json"
	"fmt"
	"strings"
)

type LoaderService struct {
	ioAdapter    infra.IoAdapter
	gamesReports []reportmodels.ProcessorReport
	gameProcesor services.GameProcessor
}

func NewLoaderService(
	ioAdapter infra.IoAdapter,
	gameProcessor services.GameProcessor,
) services.LoaderService {
	return &LoaderService{
		ioAdapter:    ioAdapter,
		gamesReports: []reportmodels.ProcessorReport{},
		gameProcesor: gameProcessor,
	}
}

// Load loads the game log and process game-by-game.
// We don't parallelize all the games processing due to being memory-conservative
func (s *LoaderService) Load(path string) (string, error) {
	currentDir, err := s.ioAdapter.Getwd()
	if err != nil {
		return "", err
	}
	realPath := s.ioAdapter.Join(currentDir, path)

	file, err := s.ioAdapter.Open(realPath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	game := []string{}

	scanner := bufio.NewScanner(file)

	// We process game by game to save memory space
	// That's the same reason why we don't load the whole log into memory
	gameId := 1
	for scanner.Scan() {
		text := scanner.Text()

		// using a simple string match here since regex would slow us down for no clear benefit
		if len(text) > 7 && strings.HasPrefix(text[7:], "InitGame:") {
			s.gamesReports = append(s.gamesReports, s.gameProcesor.ProcessGame(fmt.Sprintf("game_%d", gameId), game))

			game = []string{}
			gameId += 1
			continue
		}
		game = append(game, text)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	result := []map[string]map[string]any{}

	for _, report := range s.gamesReports {

		gameData := map[string]map[string]any{
			report.Game: {
				"kill_data":      map[string]any{},
				"rank_data":      []map[string]any{},
				"kills_by_means": map[string]any{},
			},
		}

		if report.KillReport != nil {
			gameData[report.Game]["kill_data"] = report.KillReport[report.Game]
		}

		if report.RankinReport != nil {
			gameData[report.Game]["rank_data"] = report.RankinReport[report.Game]
		}

		if report.DeathCauseReport != nil {
			gameData[report.Game]["kills_by_means"] = report.DeathCauseReport
		}

		result = append(result, gameData)
	}

	report, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(report), nil
}
