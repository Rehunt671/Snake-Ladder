package models

import (
	"fmt"
	"strings"
)

// TODO:
// 1. Change changeTurn algorithm
// 2. Change isWinAll function name => isGameEnd
// 3. Make render function more readable
type Game interface {
	AddPlayer(name string)
	Play()
}

type gameImpl struct {
	dice    Dice
	board   Board
	players []Player
	turnIdx int
}

func NewGame(numberOfSnakes int, numberOfLadders int, boardSize int) Game {
	return &gameImpl{
		dice:    NewDice(6),
		board:   NewBoard(numberOfSnakes, numberOfLadders, boardSize),
		players: []Player{},
	}
}

func (g *gameImpl) AddPlayer(name string) {
	player := NewPlayer(name)
	g.players = append(g.players, player)
	g.board.SetPosition(player, 1)
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
	curPlayer := g.getCurrentPlayer()
	g.printInformation(curPlayer)

	roll := g.dice.Roll()
	g.printRoll(roll)

	newPos := g.board.SetPosition(curPlayer, curPlayer.GetPos()+roll)
	g.printPlayerPosition(curPlayer)

	if g.board.IsDestination(newPos) {
		g.printWin()
		curPlayer.SetIsWin(true)
	}
}

func (g *gameImpl) getCurrentPlayer() Player {
	return g.players[g.turnIdx]
}

func (g *gameImpl) resetBoard() {
	board := g.board
	size := board.GetSize()
	boardSize := size * size

	for i := 0; i < boardSize; i++ {
		cell := board.GetCell(i)
		players := cell.RecievePlayersStandingOn()
		cell.SetPlayersStandingOn(players[:0])
	}
}

func (g *gameImpl) resetPlayersInfo() {
	board := g.board

	for _, player := range g.players {
		board.SetPosition(player, 1)
		player.SetIsWin(false)
	}
}

func (g *gameImpl) resetQueue() {
	g.turnIdx = 0
}

func (g *gameImpl) resetGame() {
	fmt.Printf("%92s\n", "All player are winning Reset Game!!")
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

func (g *gameImpl) printRegion(idx int) {
	board := g.board
	symbols := ""
	cell := board.GetCell(idx)

	if len(cell.RecievePlayersStandingOn()) > 0 {
		names := make([]string, len(cell.RecievePlayersStandingOn()))
		for i, player := range cell.RecievePlayersStandingOn() {
			names[i] = player.GetName()
		}
		symbols = strings.Join(names, ",")
	} else {
		symbols = strings.Join(cell.GetSymbols(), ",")
	}

	fmt.Printf("%15s", symbols)
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
	g.changeTurnIdx()
	curPlayer := g.getCurrentPlayer()
	for curPlayer.GetIsWin() {
		g.changeTurnIdx()
		curPlayer = g.getCurrentPlayer()
	}
}

func (g *gameImpl) changeTurnIdx() {
	g.turnIdx = (g.turnIdx + 1) % len(g.players)
}

func (g *gameImpl) printBorder() {
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------")
}

func (g *gameImpl) printInformation(p Player) {
	fmt.Printf("%80s %s\n", "Current Player =", p.GetName())
	fmt.Printf("%86s\n", "Press Enter to Roll!!!")
	fmt.Scanln()
}

func (g *gameImpl) printRoll(roll int) {
	fmt.Printf("%77s %d\n", "Roll =", roll)
}

func (g *gameImpl) printPlayerPosition(p Player) {
	fmt.Printf("%70s %s %d\n", p.GetName(), "Position =", p.GetPos())
}

func (g *gameImpl) printWin() {
	fmt.Printf("%78s\n", "Won!!!")
}
