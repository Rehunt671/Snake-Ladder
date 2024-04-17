package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/snake-ladder/models"
)

func askNumber(message string) int {
	var num int
	fmt.Print(message)
	fmt.Scan(&num)
	bufio.NewReader(os.Stdin).ReadString('\n')
	return num
}

func askNumSnake() int {
	return askNumber("Type snake number: ")
}

func askNumLadder() int {
	return askNumber("Type ladder number: ")
}

func main() {

	numSnakes := askNumSnake()
	numLadders := askNumLadder()
	game := models.NewGame(numSnakes,numLadders,10)
	game.AddPlayer("red")
	game.AddPlayer("green")
	game.AddPlayer("blue")
	game.Play()
}