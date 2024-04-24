package models

import "math/rand"

type Dice interface {
	Roll() int
}

// FINISH: change max to faces
type diceImpl struct {
	faces int
}

func NewDice(faces int) Dice {
	return &diceImpl{
		faces: faces,
	}
}

func (d *diceImpl) Roll() int {
	return rand.Intn(d.faces) + 1
}
