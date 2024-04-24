package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/snake-ladder/constants"
	"github.com/snake-ladder/models"
)

func askNumber(message string) int {
	var num int
	fmt.Print(message)
	fmt.Scan(&num)
	bufio.NewReader(os.Stdin).ReadString('\n')
	return num
}

// FINISH: change function name to askSnakeNumber
func askSnakeNumber() int {
	return askNumber("Type snake number: ")
}

// FINISH: change function name to askLadderNumber
func askLadderNumber() int {
	return askNumber("Type ladder number: ")
}

func main() {
	// FINISH: change numSnakes => snakeNumber, numLadders => ladderNumber
	snakeNumber := askSnakeNumber()
	ladderNumber := askLadderNumber()
	// FINISH: default value for size of board
	game := models.NewGame(snakeNumber, ladderNumber, constants.BOARD_SIZE)
	game.AddPlayer("red")
	game.AddPlayer("green")
	game.AddPlayer("blue")
	game.Play()
}
