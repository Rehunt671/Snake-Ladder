package models

import (
	"fmt"
	"math/rand"
)

type Board struct {
	size    int
	regions [][]*Region
	snakes  []*Snake
	ladders []*Ladder
}

func NewBoard(numSnakes int, numLadders int, size int) *Board {
	b := &Board{
    size:    size,
    regions: make([][]*Region, size),
    snakes:  []*Snake{},
    ladders: []*Ladder{},
}

	boardSize := size * size
	for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
					b.regions[i] = append(b.regions[i], NewRegion("-")) 
			}
	}
	snakeLadderMap := make(map[string]bool)

	for i := 0; i < numSnakes; i++ {
		for {
			start := rand.Intn(boardSize) + 1		// 1 - 100
			end := rand.Intn(boardSize) + 1			// 1 - 100
			if end >= start || start == size {
				continue
			}
			if _, ok := snakeLadderMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
				b.snakes = append(b.snakes, NewSnake(start, end))
				snakeLadderMap[fmt.Sprintf("%d:%d", start, end)] = true
				break
			}
		}
	}

	for i := 0; i < numLadders; i++ {
		for {
			start := rand.Intn(boardSize) + 1		// 1 - 100
			end := rand.Intn(boardSize) + 1			// 1 - 100
			if start >=  end || start == size {
				continue
			}
			if _, ok := snakeLadderMap[fmt.Sprintf("%d:%d", start, end)]; !ok {
				b.ladders = append(b.ladders, NewLadder(start, end))
				snakeLadderMap[fmt.Sprintf("%d:%d", start, end)] = true
				break
			}
		}
	}

	return b
}

func (b *Board) isLadder(pos int) (bool, int) {
	for _, val := range b.ladders {
		if val.start == pos {
			return true, val.end
		}
	}
	return false, -1
}

func (b *Board) isSnake(pos int) (bool, int) {
	for _, val := range b.snakes {
		if val.start == pos {
			return true, val.end
		}
	}
	return false, -1
}

func (b *Board) IsDestination(pos int) bool {
	return pos == b.size*b.size
}

func (b *Board) GetNewPosition(pos int) int {
	if ok, val := b.isLadder(pos); ok {
		fmt.Printf("%55s %d, to: %d\n","Climb ladder. Go up from:", pos, val)
		return val
	}

	if ok, val := b.isSnake(pos); ok {
		fmt.Printf("%55s %d, to: %d\n","Found snake!!!. Down from:", pos, val)
		return val
	}

	return pos
}

