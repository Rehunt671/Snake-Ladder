package models

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/snake-ladder/utils"
)


type Board interface{
	SetPosition(p Player , roll int)
	GetSize() int
	GetPath(int , int ) Path
}

type boardImpl struct {
	size    int
	paths [][]Path
	snakes  []Snake
	ladders []Ladder
}

func NewBoard(numSnakes int, numLadders int, size int) Board {
	b := &boardImpl{
    size:    size,
	}

	snakeLadderMap := make(map[string]bool)
	b.initSnakes( numLadders , size , &snakeLadderMap)
	b.initLadders( numLadders , size , &snakeLadderMap)
	b.initRegions();

	return b
}
func (b *boardImpl) SetPosition(p Player , newPos int)  {
	boardSize := b.size * b.size
	playerPos := p.GetPos()
	
	if newPos > boardSize {
		gap := newPos - boardSize
		newPos = boardSize - gap
	} else {
		newPos = b.getNewPosition(newPos)
	}

	row , col := utils.NumToRowCol(playerPos , b.size)
	b.paths[row][col].RemoveStandOn(p);

	row , col = utils.NumToRowCol(newPos , b.size) 
	b.paths[row][col].AddStandOn(p);

	p.SetPos(newPos)
	if b.isDestination(newPos) {
		fmt.Printf("%78s\n", "Won!!!")
		p.SetWin(true)
	}

}

func (b *boardImpl) GetSize() int{
	return b.size
}

func (b *boardImpl) GetPath(i int , j int) Path{
	return b.paths[i][j]
}


func (b *boardImpl) getNewPosition(pos int) int {
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

func (b *boardImpl) isDestination(pos int) bool {
	return pos == b.size*b.size
}

func (b *boardImpl) initSnakes(numSnakes int , size int , snakeLadderMap *map[string]bool){
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

func (b *boardImpl) initLadders(numLadders int , size int , snakeLadderMap *map[string]bool){
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

//////////////////Init Regions/////////////
func (b *boardImpl) initRegions(){
	b.createRegionsSlice()
	b.addNumberSymbol()
	b.addSnakesSymbol()
	b.addLadderSymbol()
}

func (b *boardImpl) createRegionsSlice(){
	b.paths = make([][]Path, b.size)
	for i := range b.paths {
    b.paths[i] = make([]Path, b.size)
	}
}

func (b *boardImpl) addNumberSymbol(){
	size := b.size

	for i := size - 1 ; i >= 0  ; i-- {
		num := 0;
		if i % 2 == 0 {
			for j := size - 1 ; j >= 0  ; j-- {
				num = size * i + j + 1
				b.paths[i][j] = NewPath([]string{strconv.Itoa(num)})
			}
		}else{
			for j := 0 ; j < size ; j++ {
				num = size * i + j + 1
				b.paths[i][size - j - 1] = NewPath([]string{strconv.Itoa(num)})
			}
		}
	}
}

func (b *boardImpl) addLadderSymbol(){
	
	for i, ladder := range b.ladders {
		
		start := ladder.GetStart()
		row , col := utils.NumToRowCol(start , b.size)
		path := b.paths[row][col]
		newSymbols := append(path.GetSymbols(),fmt.Sprintf("L%d", i+1))
		path.SetSymbols(newSymbols)

		end := ladder.GetEnd()
		row , col = utils.NumToRowCol(end , b.size) 
		path = b.paths[row][col]
		newSymbols = append(path.GetSymbols(),fmt.Sprintf("l%d", i+1))
		path.SetSymbols(newSymbols)

	}
}

func(b *boardImpl) addSnakesSymbol(){
	for i, snake := range b.snakes {
		start := snake.GetStart()
		row , col := utils.NumToRowCol(start , b.size)
		path := b.paths[row][col]
		newSymbols := append(path.GetSymbols(),fmt.Sprintf("S%d", i+1))
		path.SetSymbols(newSymbols)

		
		end := snake.GetEnd()
		row , col = utils.NumToRowCol(end , b.size) 
		path = b.paths[row][col]
		newSymbols = append(path.GetSymbols(),fmt.Sprintf("s%d", i+1))
		path.SetSymbols(newSymbols)
	}
}
//////////////////Init Regions/////////////