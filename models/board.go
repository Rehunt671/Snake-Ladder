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
	b.initLadders(numLadders,size,&snakeLadderMap)
	b.initSnakes(numSnakes,size,&snakeLadderMap)
	b.initRegions();
	return b
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
		fmt.Printf("%55s %d, to: %d\n","OOps got bitten by snake!!!. Down from:", pos, val)
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
	size := b.size
	
	b.regions = make([][]*Region, size)
	for i := range b.regions {
    b.regions[i] = make([]*Region, size)
	}

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

	for i, snake := range b.snakes {
		start := snake.start
		end := snake.end

		row , col := utils.NumToRowCol(start , size) 
		b.regions[row][col].symbols = append(b.regions[row][col].symbols,fmt.Sprintf("S%d", i+1) )
		row , col = utils.NumToRowCol(end , size) 
		b.regions[row][col].symbols = append(b.regions[row][col].symbols,fmt.Sprintf("s%d", i+1) )

	}

	for i, ladder := range b.ladders {
		start := ladder.start
		row , col := utils.NumToRowCol(start , size) 
		b.regions[row][col].symbols = append(b.regions[row][col].symbols,fmt.Sprintf("L%d", i+1) )

		end := ladder.end
		row , col = utils.NumToRowCol(end , size) 
		b.regions[row][col].symbols = append(b.regions[row][col].symbols,fmt.Sprintf("l%d", i+1) )
	}

}

func (b *Board) initSnakes(numSnakes int , size int , snakeLadderMap  *map[string]bool){
	boardSize := size * size
	for i := 0; i < numSnakes; i++ {
		for {
			start := rand.Intn(boardSize) + 1		// 1 - 100
			end := rand.Intn(boardSize) + 1			// 1 - 100
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

func (b *Board) initLadders(numLadders int , size int , snakeLadderMap  *map[string]bool){
	boardSize := size * size
	for i := 0; i < numLadders; i++ {
		for {
			start := rand.Intn(boardSize) + 1		// 1 - 100
			end := rand.Intn(boardSize) + 1			// 1 - 100
			if start >=  end || start == size {
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

