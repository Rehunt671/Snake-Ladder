package models

import "math/rand"



type Dice struct {
	max int
}

func NewDice(max int) *Dice {
	return &Dice{
		max: max,
	}
}

func (d *Dice) Roll() int {
	return rand.Intn(d.max)
}
