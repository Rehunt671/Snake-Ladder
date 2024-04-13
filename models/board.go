package models

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/snake-ladder/utils"
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
	}

	snakeLadderMap := make(map[string]bool)
	b.initSnakesAndLadders(numSnakes , numLadders , size , &snakeLadderMap)
	b.initRegions();
	return b
}

func (b *Board) IsDestination(pos int) bool {
	return pos == b.size*b.size
}

func (b *Board) GetNewPosition(pos int) int {
	if ok, val := b.isLadder(pos); ok {
		fmt.Printf("%82s %d, to: %d\n","Climb ladder!!. Go up from:", pos, val)
		return val
	}

	if ok, val := b.isSnake(pos); ok {
		fmt.Printf("%88s %d, to: %d\n","Oops got bitten by snake!!!. Down from:", pos, val)
		return val
	}

	return pos
}

func (b *Board) RemoveStandOn(p *Player)  {
	row , col := utils.NumToRowCol(p.pos , b.size)
	b.regions[row][col].RemoveStandOn(p);
}

func (b *Board) AddStandOn(p *Player)  {
	row , col := utils.NumToRowCol(p.pos , b.size) 
	b.regions[row][col].AddStandOn(p);
}


func (b *Board) initRegions(){
	b.createRegionsSlice()
	b.addNumberSymbol()
	b.addSnakesSymbol()
	b.addLadderSymbol()
}

func (b *Board) initSnakesAndLadders(numSnakes, numLadders, size int, snakeLadderMap *map[string]bool) {
	boardSize := size * size

	for i := 0; i < numSnakes + numLadders; i++ {
		for {
			start := rand.Intn(boardSize) + 1
			end := rand.Intn(boardSize) + 1
			if start == size || end == size || start == end {
				continue
			}
			key := fmt.Sprintf("%d:%d", start, end)
			if _, ok := (*snakeLadderMap)[key]; !ok {
				if start > end {
					b.snakes = append(b.snakes, NewSnake(start, end))
				} else {
					b.ladders = append(b.ladders, NewLadder(start, end))
				}
				(*snakeLadderMap)[key] = true
				break
			}
		}
	}
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

func (b *Board) createRegionsSlice(){
	b.regions = make([][]*Region, b.size)
	for i := range b.regions {
    b.regions[i] = make([]*Region, b.size)
	}
}

func (b *Board) addNumberSymbol(){
	size := b.size

	for i := size - 1 ; i >= 0  ; i-- {
		num := 0;
		if i % 2 == 0 {
			for j := size - 1 ; j >= 0  ; j-- {
				num = size * i + j + 1
				b.regions[i][j] = NewRegion([]string{strconv.Itoa(num)})
			}
		}else{
			for j := 0 ; j < size ; j++ {
				num = size * i + j + 1
				b.regions[i][size - j - 1] = NewRegion([]string{strconv.Itoa(num)})
			}
		}
	}
}

func (b *Board) addLadderSymbol(){
	
	for i, ladder := range b.ladders {
		start := ladder.start
		row , col := utils.NumToRowCol(start , b.size) 
		b.regions[row][col].symbols = append(b.regions[row][col].symbols,fmt.Sprintf("L%d", i+1) )

		end := ladder.end
		row , col = utils.NumToRowCol(end , b.size) 
		b.regions[row][col].symbols = append(b.regions[row][col].symbols,fmt.Sprintf("l%d", i+1) )
	}
}

func(b *Board) addSnakesSymbol(){
	for i, snake := range b.snakes {
		start := snake.start
		row , col := utils.NumToRowCol(start , b.size) 
		b.regions[row][col].symbols = append(b.regions[row][col].symbols,fmt.Sprintf("S%d", i+1) )
		
		end := snake.end
		row , col = utils.NumToRowCol(end , b.size) 
		b.regions[row][col].symbols = append(b.regions[row][col].symbols,fmt.Sprintf("s%d", i+1) )
	}
}
