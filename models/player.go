package models

import "fmt"


type Player struct {
	name string
	pos int
	// color int
	win bool
}

func NewPlayer(name string) *Player {
	return &Player{
		name : name,
		pos :  1,
	}
}

func (p *Player) RollDice() int {
	game := GetGameInstance()
	dice := game.dice
	return dice.Roll()
}

func (p *Player) Move(roll int) {
	game := GetGameInstance()
	board := game.board
	boardSize := board.size*board.size


	board.RemoveStandOn(p);
	newPos := p.pos + roll
	
	if newPos > boardSize {
		gap := newPos - boardSize
		newPos = boardSize - gap
		p.pos = newPos
	} else {
		newPos = board.GetNewPosition(newPos)
		p.pos = newPos
	}

	board.AddStandOn(p);
	if board.IsDestination(p.pos) {
		fmt.Printf("%78s\n", "Won!!!")
		p.win = true
	}

}
