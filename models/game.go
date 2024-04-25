package models

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/snake-ladder/constants"
)

// FINISH: p => player in all file except player file
// FINISH: pos => position
// FINISH: newPos => newPosition
// FINISH: change i variable in for loop
// FINISH: idx => index
// FINISH: ok , val => ... ?
type Game interface {
	Play()
}

type gameImpl struct {
	dice      Dice
	board     Board
	players   []Player
	turnIndex int
}

func NewGame(playerNames []string, numberOfSnakes int, numberOfLadders int, boardSize int) Game {
	g := &gameImpl{
		dice:  NewDice(constants.MAX_DICE_FACES),
		board: NewBoard(numberOfSnakes, numberOfLadders, boardSize),
	}
	g.initPlayers(playerNames)

	return g
}

func (g *gameImpl) Play() {
	for {
		g.renderBoard()
		g.playRound()

		if g.isGameEnd() {
			g.resetGame()
			continue
		}

		g.changeTurn()
	}
}

func (g *gameImpl) initPlayers(playerNames []string) {
	for _, playerName := range playerNames {
		player := NewPlayer(playerName)
		g.players = append(g.players, player)
		g.setPosition(player, 1)
	}
}

func (g *gameImpl) playRound() {
	currentPlayer := g.getCurrentPlayer()
	g.printInformation(currentPlayer)

	// FINISH: faces => step
	step := g.dice.Roll()
	g.printDiceFace(step)

	destination := currentPlayer.GetPosition() + step
	newPosition := g.setPosition(currentPlayer, destination)
	g.printPlayerPosition(currentPlayer)

	if g.board.IsDestination(newPosition) {
		g.printWin()
		currentPlayer.SetIsWin(true)
	}
}

// FINISH: getValidPosition function name => findValidPosition
func (g *gameImpl) setPosition(player Player, newPosition int) int {
	newPosition = g.findValidPosition(newPosition)
	g.setStanOn(player, newPosition)
	player.SetPosition(newPosition)

	return newPosition
}

func (g *gameImpl) setStanOn(player Player, newPosition int) {
	oldCell := g.board.GetCell(player.GetPosition() - 1)
	newCell := g.board.GetCell(newPosition - 1)
	oldCell.RemovePlayer(player)
	newCell.AddPlayer(player)
}

func (g *gameImpl) findValidPosition(position int) int {
	size := g.board.GetSize()
	boardSize := size * size

	if position > boardSize {
		gap := position - boardSize
		return boardSize - gap
	}

	if isLadder, end := g.isLadder(position); isLadder {
		g.printClimbLadder(position, end)
		return end
	}

	if isLadder, end := g.isSnake(position); isLadder {
		g.printGotBittenBySnake(position, end)
		return end
	}

	return position
}

func (g *gameImpl) isLadder(position int) (bool, int) {
	board := g.board
	for _, ladder := range board.GetLadders() {
		if ladder.GetStart() == position {
			return true, ladder.GetEnd()
		}
	}

	return false, -1
}

func (g *gameImpl) isSnake(position int) (bool, int) {
	board := g.board

	for _, snake := range board.GetSnakes() {
		if snake.GetStart() == position {
			return true, snake.GetEnd()
		}
	}

	return false, -1
}

func (g *gameImpl) getCurrentPlayer() Player {
	return g.players[g.turnIndex]
}

func (g *gameImpl) resetBoard() {
	board := g.board
	size := board.GetSize()
	boardSize := size * size

	for boardIndex := 0; boardIndex < boardSize; boardIndex++ {
		cell := board.GetCell(boardIndex)
		playersOnCell := cell.GetPlayers()
		cell.SetPlayers(playersOnCell[:0])
	}
}

func (g *gameImpl) resetPlayersInfo() {
	for _, player := range g.players {
		g.setPosition(player, 1)
		player.SetIsWin(false)
	}
}

func (g *gameImpl) resetQueue() {
	g.turnIndex = 0
}

func (g *gameImpl) resetGame() {
	g.printGameEnd()
	g.resetBoard()
	g.resetPlayersInfo()
	g.resetQueue()
}

// FINISH: render => renderBoard
func (g *gameImpl) renderBoard() {
	board := g.board
	size := board.GetSize()
	g.printBorder()

	for row := size - 1; row >= 0; row-- {
		if row%2 == 0 {
			for collumn := 0; collumn < size; collumn++ {
				regionIndex := size*row + collumn
				g.printRegion(regionIndex)
			}
		} else {
			for collumn := size - 1; collumn >= 0; collumn-- {
				regionIndex := size*row + collumn
				g.printRegion(regionIndex)
			}
		}
		fmt.Println()
	}

	g.printBorder()
}

func (g *gameImpl) getPlayersName(players []Player) []string {
	names := lo.Map(players, func(player Player, index int) string {
		return player.GetName()
	})

	return names
}

func (g *gameImpl) isGameEnd() bool {
	winCount := 0

	for _, player := range g.players {
		if player.GetIsWin() {
			winCount++
		}
	}

	return winCount == len(g.players)-1
}

func (g *gameImpl) changeTurn() {
	g.changeTurnIndex()
	currentPlayer := g.getCurrentPlayer()

	for currentPlayer.GetIsWin() {
		g.changeTurnIndex()
		currentPlayer = g.getCurrentPlayer()
	}
}

func (g *gameImpl) changeTurnIndex() {
	g.turnIndex = (g.turnIndex + 1) % len(g.players)
}

func (g *gameImpl) printRegion(index int) {
	board := g.board
	symbol := ""
	cell := board.GetCell(index)
	playersOnCell := cell.GetPlayers()

	if cell.HasPlayer() {
		names := g.getPlayersName(playersOnCell)
		symbol = strings.Join(names, ",")
	} else {
		symbol = strings.Join(cell.GetSymbols(), ",")
	}

	g.printSymbol(symbol)
}

func (g *gameImpl) printClimbLadder(source int, destination int) {
	fmt.Printf("%82s %d, to: %d\n", "Climb ladder!!. Go up from:", source, destination)
}

func (g *gameImpl) printGotBittenBySnake(source int, destination int) {
	fmt.Printf("%88s %d, to: %d\n", "Oops got bitten by snake!!!. Down from:", source, destination)
}

func (g *gameImpl) printSymbol(symbol string) {
	fmt.Printf("%15s", symbol)
}

func (g *gameImpl) printBorder() {
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------")
}

func (g *gameImpl) printInformation(player Player) {
	fmt.Printf("%80s %s\n", "Current Player =", player.GetName())
	fmt.Printf("%86s\n", "Press Enter to Roll!!!")
	fmt.Scanln()
}

func (g *gameImpl) printDiceFace(roll int) {
	fmt.Printf("%77s %d\n", "Face =", roll)
}

func (g *gameImpl) printPlayerPosition(player Player) {
	fmt.Printf("%70s %s %d\n", player.GetName(), "Position =", player.GetPosition())
}

func (g *gameImpl) printWin() {
	fmt.Printf("%78s\n", "Won!!!")
}

func (g *gameImpl) printGameEnd() {
	fmt.Printf("%92s\n", "All player are winning Reset Game!!")
}
