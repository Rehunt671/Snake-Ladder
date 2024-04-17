package models

//package is a collection of source files in the same directory that are compiled together
import (
	"fmt"
	"strings"
)

type Game interface {
	AddPlayer(name string)
	Play()
}

type gameImpl struct {
	dice    Dice
	board   Board
	players []Player
	firstPlayer Player
}

func NewGame(numberOfSnakes int, numberOfLadders int, boardSize int) Game {
	return &gameImpl{
			dice:    NewDice(6),
			board:   NewBoard(numberOfSnakes, numberOfLadders, boardSize),
			players: []Player{},
	}
}

func (g *gameImpl) AddPlayer(name string){
	player := NewPlayer(name)
	g.players = append(g.players,player)
	g.board.SetPosition(player , 1)	
	if len(g.players) == 1 {
		g.firstPlayer = g.players[0]
	}
}

func (g *gameImpl) Play() {

	for {
		g.render()
		g.playRound()
		if g.isWinAll() {
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

	newPos := g.board.SetPosition(curPlayer , curPlayer.GetPos() + roll)
	g.printPlayerPosition(curPlayer)

	if g.board.IsDestination(newPos) {
		g.printWin()
		curPlayer.SetWin(true)
	}

}

func (g *gameImpl) getCurrentPlayer() Player {
	return g.players[0]
}

func (g *gameImpl) resetBoard(){
	board := g.board
	size := board.GetSize()

	for i := 0 ; i < size * size ; i++{
		cell := board.GetCell(i)
		standOn := cell.GetStandOn()
		cell.SetStandOn(standOn[:0])
	}
	
}

func (g *gameImpl) resetPlayersInfo(){
	board := g.board
	for _, player := range g.players {
		board.SetPosition(player , 1)
		player.SetWin(false) 
	}
}

func (g *gameImpl) resetQueue(){
	for g.players[0] != g.firstPlayer {
		g.players = append(g.players[1:], g.players[0])
	}
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

	for i := size - 1 ; i >= 0  ; i-- {
		if i % 2 != 0 {
			for j := size - 1 ; j >= 0  ; j-- {
				idx := size*i + j
				g.printRegion(idx)
			}
		}else{
			for j := 0 ; j < size ; j++ {
				idx := size*i + j
				g.printRegion(idx)
			}
		}
		fmt.Println()
	}

	g.printBorder()

}

func (g *gameImpl) printRegion(idx int ) {
	board := g.board

	symbols := ""
	cell := board.GetCell(idx)
	if len(cell.GetStandOn()) > 0 {
		names := make([]string, len(cell.GetStandOn()))
		for i, player := range cell.GetStandOn() {
			names[i] = player.GetName()
		}
		symbols = strings.Join(names, ",")
	} else {
		symbols = strings.Join(cell.GetSymbols(), ",")
	}

	fmt.Printf("%15s", symbols)
}

func (g *gameImpl) isWinAll() bool {
	winCount := 0
	for _, player := range g.players {
			if player.GetWin() {
					winCount++
			}
	}

	return winCount == len(g.players) - 1
}

func (g *gameImpl) changeTurn() {
	g.players = append(g.players[1:], g.players[0])
	for g.players[0].GetWin() {
		g.players = append(g.players[1:], g.players[0])
	}
}


func (g *gameImpl) printBorder()	{
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

func (g *gameImpl) printWin(){

	fmt.Printf("%78s\n", "Won!!!")

}