package models

import (
	"fmt"
)


var gameInstance *Game

type Game struct {
	dice    *Dice
	board   *Board
	players []*Player
	firstPlayer *Player
}

// ไม่มี Constructor in golang ,So we have to create Global func to setter for struct
func NewGame(numberOfSnakes int, numberOfLadders int, boardSize int) *Game {
	gameInstance =  &Game{
			dice:    NewDice(6),
			board:   NewBoard(numberOfSnakes, numberOfLadders, boardSize),
			players: []*Player{},
	}
	return gameInstance
}

func GetGameInstance() *Game {
	if gameInstance == nil {
		gameInstance = NewGame(1,1,10);
	}
  return gameInstance;
}


//ตัว pointer receiver จะเป็นตัวบอกว่า method นั้นๆจะใช้ได้แค่กับ struct ของ pointer เท่านั้น
func (g *Game) 	AddPlayer(name string){
	g.players = append(g.players, NewPlayer(name))
	if len(g.players) == 1 {
		g.firstPlayer = g.players[0]
	}
}

func (g *Game) Play() {
	g.render()

	for {
		curPlayer := g.getCurrentPlayer()

		g.printCurrentPlayer(curPlayer)

		g.waitForRoll()

		roll := g.rollDice()
		g.printRoll(roll)

		curPlayer.Move(roll)
		g.printPlayerPosition(curPlayer)

		if g.isWinAll() {
			g.resetAndContinue()
			continue
		}

		g.changeTurn()
		g.render()
	}
}

func (g *Game) getCurrentPlayer() *Player {
	return g.players[0]
}

func (g *Game) printCurrentPlayer(player *Player) {
	fmt.Printf("%55s %s\n", "Current Player =", player.name)
}

func (g *Game) waitForRoll() {
	fmt.Printf("%61s\n", "Press Enter to Roll!!!")
	fmt.Scanln()
}

func (g *Game) rollDice() int {
	return g.dice.Roll()
}

func (g *Game) printRoll(roll int) {
	fmt.Printf("%50s %d\n", "Roll =", roll)
}

func (g *Game) printPlayerPosition(player *Player) {
	fmt.Printf("%45s %s %d\n", player.name, "Position =", player.pos)
}

func (g *Game) resetAndContinue() {
	fmt.Printf("%65s\n", "All Player is Winning Reset Game!!")
	fmt.Println("----------------------------------------------------------------------------------------------------------")
	g.resetGame()
}

func (g *Game) render() {
	board := g.board
	// snakes := board.snakes
	// ladders := board.ladders
	size := board.size

	fmt.Println("----------------------------------------------------------------------------------------------------------")
	
	for i := size ; i > 0  ; i-- {
		if i % 2 == 0 {
			for j := 0 ; j < size ; j++ {
				fmt.Printf("%10d", size * i - j) 
			}
		}else{
			for j := size - 1 ; j >= 0 ; j-- {
				fmt.Printf("%10d", size * i - j) 
			}
		}
		fmt.Println()
	}

	fmt.Println("----------------------------------------------------------------------------------------------------------")

}

func (g *Game) isWinAll() bool {
	winCount := 0
	for _, player := range g.players {
			if player.win {
					winCount++
			}
	}

	return winCount == len(g.players) - 1
}

func (g *Game) resetPlayers(){
	for _, player := range g.players {
		player.pos = 0
		player.win = false 
	}
}

func (g *Game) resetQueue(){
	for g.players[0] != g.firstPlayer {
		g.players = append(g.players[1:], g.players[0])
	}
}

func (g *Game) resetGame() {
	g.resetPlayers()
	g.resetQueue()
}

func (g *Game) changeTurn() {
	g.players = append(g.players[1:], g.players[0])
	for g.players[0].win {
		g.players = append(g.players[1:], g.players[0])
	}
}
