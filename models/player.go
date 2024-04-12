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
	}
}

func (p *Player) Move(roll int) {
	game := GetGameInstance()
	board := game.board
	boardSize := board.size*board.size
	
	newPos := p.pos + roll
	if newPos > boardSize {
		gap := newPos - boardSize
		newPos = boardSize - gap
		p.pos = newPos
	} else {
		p.pos = board.GetNewPosition(newPos)
		if board.IsDestination(p.pos) {
			fmt.Printf("%50s\n", "Won!!!")
			p.pos = 0;
			p.win = true
		}
	}
}
