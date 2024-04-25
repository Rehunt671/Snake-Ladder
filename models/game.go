package models

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/snake-ladder/constants"
)

type Game interface {
	AddPlayer(name string)
	Play()
}

type gameImpl struct {
	dice      Dice
	board     Board
	players   []Player
	turnIndex int
}

func NewGame(numberOfSnakes int, numberOfLadders int, boardSize int) Game {
	return &gameImpl{
		dice:    NewDice(constants.MAX_DICE_FACES),
		board:   NewBoard(numberOfSnakes, numberOfLadders, boardSize),
		players: []Player{},
	}
}

func (g *gameImpl) AddPlayer(name string) {
	player := NewPlayer(name)
	g.players = append(g.players, player)
	g.setPosition(player, 1)
}

func (g *gameImpl) Play() {
	for {
		g.render()
		g.playRound()

		if g.isGameEnd() {
			g.resetGame()
			continue
		}

		g.changeTurn()
	}
}

func (g *gameImpl) playRound() {
	currentPlayer := g.getCurrentPlayer()
	g.printInformation(currentPlayer)

	faces := g.dice.Roll()
	g.printDiceFace(faces)

	newPos := g.setPosition(currentPlayer, currentPlayer.GetPos()+faces)
	g.printPlayerPosition(currentPlayer)

	if g.board.IsDestination(newPos) {
		g.printWin()
		currentPlayer.SetIsWin(true)
	}
}

func (g *gameImpl) setPosition(p Player, newPos int) int {
	newPos = g.getValidPosition(newPos)
	g.setStanOn(p, newPos)
	p.SetPos(newPos)

	return newPos
}

func (g *gameImpl) setStanOn(p Player, newPos int) {
	oldCell := g.board.GetCell(p.GetPos() - 1)
	newCell := g.board.GetCell(newPos - 1)
	oldCell.RemovePlayer(p)
	newCell.AddPlayer(p)
}

func (g *gameImpl) getValidPosition(pos int) int {
	size := g.board.GetSize()
	boardSize := size * size

	if pos > boardSize {
		gap := pos - boardSize
		return boardSize - gap
	}

	if ok, val := g.isLadder(pos); ok {
		g.printClimbLadder(pos, val)
		return val
	}

	if ok, val := g.isSnake(pos); ok {
		g.printGotBittenBySnake(pos, val)
		return val
	}

	return pos
}

func (g *gameImpl) isLadder(pos int) (bool, int) {
	board := g.board
	for _, val := range board.GetLadders() {
		if val.GetStart() == pos {
			return true, val.GetEnd()
		}
	}

	return false, -1
}

func (g *gameImpl) isSnake(pos int) (bool, int) {
	board := g.board

	for _, val := range board.GetSnakes() {
		if val.GetStart() == pos {
			return true, val.GetEnd()
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

	for i := 0; i < boardSize; i++ {
		cell := board.GetCell(i)
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

func (g *gameImpl) render() {
	board := g.board
	size := board.GetSize()
	g.printBorder()

	for i := size - 1; i >= 0; i-- {
		if i%2 == 0 {
			for j := 0; j < size; j++ {
				idx := size*i + j
				g.printRegion(idx)
			}
		} else {
			for j := size - 1; j >= 0; j-- {
				idx := size*i + j
				g.printRegion(idx)
			}
		}
		fmt.Println()
	}

	g.printBorder()
}

func (g *gameImpl) getPlayersName(players []Player) []string {
	names := lo.Map(players, func(player Player, idx int) string {
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

func (g *gameImpl) printRegion(idx int) {
	board := g.board
	symbol := ""
	cell := board.GetCell(idx)
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

func (g *gameImpl) printInformation(p Player) {
	fmt.Printf("%80s %s\n", "Current Player =", p.GetName())
	fmt.Printf("%86s\n", "Press Enter to Roll!!!")
	fmt.Scanln()
}

func (g *gameImpl) printDiceFace(roll int) {
	fmt.Printf("%77s %d\n", "Face =", roll)
}

func (g *gameImpl) printPlayerPosition(p Player) {
	fmt.Printf("%70s %s %d\n", p.GetName(), "Position =", p.GetPos())
}

func (g *gameImpl) printWin() {
	fmt.Printf("%78s\n", "Won!!!")
}

func (g *gameImpl) printGameEnd() {
	fmt.Printf("%92s\n", "All player are winning Reset Game!!")
}
