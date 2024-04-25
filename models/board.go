package models

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Board interface {
	IsDestination(position int) bool
	GetSize() int
	GetCell(index int) Cell
	GetLadders() []Ladder
	GetSnakes() []Snake
}

type boardImpl struct {
	size    int
	cells   []Cell
	snakes  []Snake
	ladders []Ladder
}

func NewBoard(snakeNumber int, ladderNumber int, size int) Board {
	board := &boardImpl{
		size: size,
	}
	snakeLadderMap := make(map[int]int)
	board.initSnakes(snakeNumber, size, &snakeLadderMap)
	board.initLadders(ladderNumber, size, &snakeLadderMap)
	board.initCells()
	return board
}

func (b *boardImpl) IsDestination(position int) bool {
	return position == b.size*b.size
}

func (b *boardImpl) GetSize() int {
	return b.size
}

func (b *boardImpl) GetCell(index int) Cell {
	return b.cells[index]
}

func (b *boardImpl) GetLadders() []Ladder {
	return b.ladders
}

func (b *boardImpl) GetSnakes() []Snake {
	return b.snakes
}

func (b *boardImpl) initSnakes(snakeNumber int, size int, snakeLadderMap *map[int]int) {
	boardSize := size * size

	for snakeIndex := 0; snakeIndex < snakeNumber; snakeIndex++ {
		for {
			start := rand.Intn(boardSize) + 1
			end := rand.Intn(boardSize) + 1
			if end >= start || start == boardSize || (*snakeLadderMap)[start] > 0 {
				continue
			}
			b.snakes = append(b.snakes, NewSnake(start, end))
			(*snakeLadderMap)[start] = end
			break
		}
	}
}

func (b *boardImpl) initLadders(ladderNumber int, size int, snakeLadderMap *map[int]int) {
	boardSize := size * size

	for ladderIndex := 0; ladderIndex < ladderNumber; ladderIndex++ {
		for {
			start := rand.Intn(boardSize) + 1
			end := rand.Intn(boardSize) + 1
			if start >= end || start == 1 || (*snakeLadderMap)[start] > 0 {
				continue
			}
			b.ladders = append(b.ladders, NewLadder(start, end))
			(*snakeLadderMap)[start] = end
			break
		}
	}
}

func (b *boardImpl) initCells() {
	b.addNumberSymbol()
	b.addSnakesSymbol()
	b.addLadderSymbol()
}

func (b *boardImpl) addNumberSymbol() {
	size := b.size
	boardSize := size * size

	for boardIndex := 0; boardIndex < boardSize; boardIndex++ {
		cell := NewCell([]string{strconv.Itoa(boardIndex + 1)})
		b.cells = append(b.cells, cell)
	}
}

func (b *boardImpl) addLadderSymbol() {
	for ladderIndex, ladder := range b.ladders {
		start := ladder.GetStart()
		cell := b.cells[start-1]
		newSymbols := append(cell.GetSymbols(), fmt.Sprintf("L%d", ladderIndex+1))
		cell.SetSymbols(newSymbols)

		end := ladder.GetEnd()
		cell = b.cells[end-1]
		newSymbols = append(cell.GetSymbols(), fmt.Sprintf("l%d", ladderIndex+1))
		cell.SetSymbols(newSymbols)
	}
}

func (b *boardImpl) addSnakesSymbol() {
	for snakeIndex, snake := range b.snakes {
		start := snake.GetStart()
		cell := b.cells[start-1]
		newSymbols := append(cell.GetSymbols(), fmt.Sprintf("S%d", snakeIndex+1))
		cell.SetSymbols(newSymbols)

		end := snake.GetEnd()
		cell = b.cells[end-1]
		newSymbols = append(cell.GetSymbols(), fmt.Sprintf("s%d", snakeIndex+1))
		cell.SetSymbols(newSymbols)
	}
}
