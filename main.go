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

func askSnakeNumber() int {
	return askNumber("Type snake number: ")
}

func askLadderNumber() int {
	return askNumber("Type ladder number: ")
}

func main() {
	snakeNumber := askSnakeNumber()
	ladderNumber := askLadderNumber()
	//TODO: fix infinite loop if snake > boardSize
	//TODO: shouldn't add player after game begin
	game := models.NewGame(snakeNumber, ladderNumber, constants.BOARD_SIZE)
	game.AddPlayer("red")
	game.AddPlayer("green")
	game.AddPlayer("blue")
	game.Play()
}
