package retrosheet

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Player a player code and their full name
type Player struct {
	code string
	name string
}

// PitchResult string identifying what occurred for a single pitch
type PitchResult string

const (
	LookingStrike  PitchResult = "C"
	SwingingStrile PitchResult = "S"
	FoulBall       PitchResult = "F"
	InPlay         PitchResult = "X"
)

type Position int

const (
	Catcher     Position = 1
	Pitcher     Position = 2
	FirstBase   Position = 3
	SecondBase  Position = 4
	ThirdBase   Position = 5
	Shortstop   Position = 6
	LeftField   Position = 7
	CenterField Position = 8
	RightField  Position = 9
)

type AtBatResult string

// PlayerAtBat Defines what occurred during an at-bat
type PlayerAtBat struct {
	pitcher     Player
	batter      Player
	pitches     []PitchResult
	result      AtBatResult
	inning      int
	atBatNumber int
}

const (
	UmpHome       string = "umphopme"
	UmpFirstBase  string = "ump1b"
	UmpSecondBase string = "ump2b"
	UmpThirdBase  string = "ump3b"
)

type Umpire struct {
	umpireId string
	position string
}

type PlayedGame struct {
	startTime   time.Time
	umpires     []Umpire
	temperature int
	atBats      []PlayerAtBat
}

const (
	GameId          string = "id"
	GameVersion     string = "version"
	GameInfo        string = "info"
	StartingPlayers string = "start"
	Play            string = "play"
	Substitusion    string = "sub"
	Data            string = "data" // consider ignoring
)

// ParseFile parses a file from rotosheet and returns a list of games
// 		input - fileName: file on local drive containing rotosheet defined game
//	returns: list of files parsed from the doc
func ParseFile(fileName string) (parsedGames []PlayedGame, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open file " + fileName)
	}

	var singleGame [][]string
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		csv := strings.Split(currentLine, ",")
		if csv[0] == GameId {
			if len(singleGame) > 0 {
				currentGame, err := CsvToGame(singleGame)
			}
		}
	}

	return parsedGames, nil
}

// CsvToGame takes a rotosheet data doc and converts it into an object
func CsvToGame(lines [][]string) (game *PlayedGame, err error) {

	game = new(PlayedGame)
	for index, currentLine := range lines {
		if currentLine[0] == GameInfo {
			err = HandleInfoLine(currentLine, game)
		} else if currentLine[0] == GameInfo {

		} else if currentLine[0] == StartingPlayers {

		} else if currentLine[0] == Substitusion {

		} else if currentLine[0] == Data {

		} else {
			return nil, fmt.Errorf("Unexpected rotosheet opener on line [%d]: %q", index, currentLine[0])
		}
	}
	return game, nil
}

func HandleInfoLine(currentLine []string, game *PlayedGame) (err error) {
	if currentLine[1] == UmpHome {

	} else if currentLine[1] == UmpFirstBase {

	} else if currentLine[1] == UmpSecondBase {

	} else if currentLine[1] == UmpThirdBase {

	}
}
