package models

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Board interface {
	SetPosition(p Player, pos int) int
	IsDestination(pos int) bool
	GetSize() int
	GetCell(idx int) Cell
}

type boardImpl struct {
	size    int
	cells   []Cell
	snakes  []Snake
	ladders []Ladder
}

func NewBoard(numSnakes int, numLadders int, size int) Board {
	b := &boardImpl{
		size: size,
	}
	snakeLadderMap := make(map[string]bool)
	b.initSnakes(numLadders, size, &snakeLadderMap)
	b.initLadders(numLadders, size, &snakeLadderMap)
	b.initRegions()
	return b
}

func (b *boardImpl) SetPosition(p Player, newPos int) int {
	newPos = b.getValidPosition(newPos)
	b.setStanOn(p, newPos)
	p.SetPos(newPos)
	return newPos
}

func (b *boardImpl) IsDestination(pos int) bool {
	return pos == b.size*b.size
}

func (b *boardImpl) GetSize() int {
	return b.size
}

func (b *boardImpl) GetCell(idx int) Cell {
	return b.cells[idx]
}

func (b *boardImpl) setStanOn(p Player, newPos int) {
	b.cells[p.GetPos()-1].RemoveStandOn(p)
	b.cells[newPos-1].AddStandOn(p)
}

func (b *boardImpl) getValidPosition(pos int) int {
	boardSize := b.size * b.size

	if pos > boardSize {
		gap := pos - boardSize
		return boardSize - gap
	}

	if ok, val := b.isLadder(pos); ok {
		fmt.Printf("%82s %d, to: %d\n", "Climb ladder!!. Go up from:", pos, val)
		return val
	}

	if ok, val := b.isSnake(pos); ok {
		fmt.Printf("%88s %d, to: %d\n", "Oops got bitten by snake!!!. Down from:", pos, val)
		return val
	}

	return pos
}

func (b *boardImpl) isLadder(pos int) (bool, int) {

	for _, val := range b.ladders {
		if val.GetStart() == pos {
			return true, val.GetEnd()
		}
	}
	return false, -1
}

func (b *boardImpl) isSnake(pos int) (bool, int) {

	for _, val := range b.snakes {
		if val.GetStart() == pos {
			return true, val.GetEnd()
		}
	}
	return false, -1
}

func (b *boardImpl) initSnakes(numSnakes int, size int, snakeLadderMap *map[string]bool) {
	boardSize := size * size

	for i := 0; i < int(numSnakes); i++ {
		for {
			start := rand.Intn(boardSize) + 1
			end := rand.Intn(boardSize) + 1
			if end >= start || start == size {
				continue
			}
			if _, ok := (*snakeLadderMap)[fmt.Sprintf("%d:%d", start, end)]; !ok {
				b.snakes = append(b.snakes, NewSnake(start, end))
				(*snakeLadderMap)[fmt.Sprintf("%d:%d", start, end)] = true
				break
			}
		}
	}
}

func (b *boardImpl) initLadders(numLadders int, size int, snakeLadderMap *map[string]bool) {
	boardSize := size * size

	for i := 0; i < int(numLadders); i++ {
		for {
			start := rand.Intn(boardSize) + 1
			end := rand.Intn(boardSize) + 1
			if start >= end || start == 1 {
				continue
			}
			if _, ok := (*snakeLadderMap)[fmt.Sprintf("%d:%d", start, end)]; !ok {
				b.ladders = append(b.ladders, NewLadder(start, end))
				(*snakeLadderMap)[fmt.Sprintf("%d:%d", start, end)] = true
				break
			}
		}
	}
}

func (b *boardImpl) initRegions() {
	b.addNumberSymbol()
	b.addSnakesSymbol()
	b.addLadderSymbol()
}

func (b *boardImpl) addNumberSymbol() {
	size := b.size

	for i := 0; i < size*size; i++ {
		cell := NewCell([]string{strconv.Itoa(i + 1)})
		b.cells = append(b.cells, cell)
	}
}

func (b *boardImpl) addLadderSymbol() {

	for i, ladder := range b.ladders {
		start := ladder.GetStart()
		cell := b.cells[start-1]
		newSymbols := append(cell.GetSymbols(), fmt.Sprintf("L%d", i+1))
		cell.SetSymbols(newSymbols)

		end := ladder.GetEnd()
		cell = b.cells[end-1]
		newSymbols = append(cell.GetSymbols(), fmt.Sprintf("l%d", i+1))
		cell.SetSymbols(newSymbols)
	}
}

func (b *boardImpl) addSnakesSymbol() {
	for i, snake := range b.snakes {
		start := snake.GetStart()
		cell := b.cells[start-1]
		newSymbols := append(cell.GetSymbols(), fmt.Sprintf("S%d", i+1))
		cell.SetSymbols(newSymbols)

		end := snake.GetEnd()
		cell = b.cells[end-1]
		newSymbols = append(cell.GetSymbols(), fmt.Sprintf("s%d", i+1))
		cell.SetSymbols(newSymbols)
	}
}
