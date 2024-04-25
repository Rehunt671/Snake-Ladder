package main

import (
	"fmt"

	"github.com/snake-ladder/constants"
	"github.com/snake-ladder/models"
	"github.com/snake-ladder/utils.go"
)

func askSnakeNumber() int {
	maxSnakeNumber := constants.MAX_SNAKE
	question := fmt.Sprintf("Type snake number ( <= %d): ", maxSnakeNumber)

	return utils.AskNumber(question, maxSnakeNumber)
}

func askLadderNumber() int {
	maxLadderNumber := constants.MAX_LADDER
	question := fmt.Sprintf("Type ladder number ( <= %d): ", maxLadderNumber)

	return utils.AskNumber(question, maxLadderNumber)
}

func askPlayerNumber() int {
	maxPlayer := constants.MAX_PLAYER
	question := fmt.Sprintf("Type player number ( <= %d): ", maxPlayer)

	return utils.AskNumber(question, maxPlayer)
}

func askPlayerNames(playerNumber int) []string {
	playerNames := make([]string, 0)
	nameMap := make(map[string]bool)

	for playerIndex := 0; playerIndex < playerNumber; {
		question := fmt.Sprintf("Type P%d name:", playerIndex+1)
		playerName := utils.AskString(question)

		if _, isNameExist := nameMap[playerName]; !isNameExist {
			playerNames = append(playerNames, playerName)
			nameMap[playerName] = true
			playerIndex++
		} else {
			fmt.Println("Name already exists. Please enter a unique name.")
		}
	}

	return playerNames
}

func main() {
	snakeNumber := askSnakeNumber()
	ladderNumber := askLadderNumber()
	playerNumber := askPlayerNumber()
	playerNames := askPlayerNames(playerNumber)
	//FINISH: fix infinite loop if snake > boardSize
	//FINISH: shouldn't add player after game begin
	game := models.NewGame(playerNames, snakeNumber, ladderNumber, constants.BOARD_SIZE)
	game.Play()
}
