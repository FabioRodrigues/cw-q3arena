package loader

import (
	"bufio"
	"cw-q3arena/infra"
	"cw-q3arena/reportmodels"
	"cw-q3arena/services"
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
) *LoaderService {
	return &LoaderService{
		ioAdapter:    ioAdapter,
		gamesReports: []reportmodels.ProcessorReport{},
		gameProcesor: gameProcessor,
	}
}

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
	// That's the same reason why we don't load the whole log in memory
	gameId := 1
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text[7:], "InitGame:") {
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

	results := []string{}

	for _, report := range s.gamesReports {
		if report.KillReport != "" {
			results = append(results, fmt.Sprintf(`{"%s":%s}`, report.Game, report.KillReport))
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(results, ",\n")), nil
}
