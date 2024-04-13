package models

//package is a collection of source files in the same directory that are compiled together
import (
	"fmt"
	"strings"
)

var gameInstance *Game

type Game struct {
	dice    *Dice
	board   *Board
	players []*Player
	firstPlayer *Player
}

// Don't have Constructor in golang ,So we have to create Global func to setter for struct
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
func (g *Game) AddPlayer(name string){
	player := NewPlayer(name)
	g.players = append(g.players,player)
	g.board.AddStandOn(player)
	if len(g.players) == 1 {
		g.firstPlayer = g.players[0]
	}
}

func (g *Game) Play() {
	g.render()

	for {
		g.playRound()
		if g.isWinAll() {
			g.resetGame()
			g.render()
			continue
		}
		g.changeTurn()
		g.render()
	}
}


func (g *Game) playRound() {
	curPlayer := g.getCurrentPlayer()
	g.printInformation(curPlayer)
	
	roll := curPlayer.RollDice()
	g.printRoll(roll)

	curPlayer.Move(roll)
	g.printPlayerPosition(curPlayer)
}

func (g *Game) getCurrentPlayer() *Player {
	return g.players[0]
}

func (g *Game) resetBoard(){
	board := g.board
	size := board.size
	regions := board.regions

	for i := 0 ; i < size ; i++{
		for j := 0 ; j < size ; j++{
			regions[i][j].standOn = regions[i][j].standOn[:0]
		}
	}
	
}

func (g *Game) resetPlayersInfo(){
	board := g.board
	for _, player := range g.players {
		player.pos = 1
		player.win = false 
		board.AddStandOn(player)
	}
}

func (g *Game) resetQueue(){
	for g.players[0] != g.firstPlayer {
		g.players = append(g.players[1:], g.players[0])
	}
}

func (g *Game) resetGame() {
	fmt.Printf("%92s\n", "All Player is Winning Reset Game!!")
	g.resetBoard()	
	g.resetPlayersInfo()
	g.resetQueue()
}

func (g *Game) render() {
	board := g.board
	size := board.size

	g.printBoarder()

	for i := size - 1 ; i >= 0  ; i-- {
		for j := 0 ; j < size ; j++ {
			symbols := ""
			if len(board.regions[i][j].standOn) > 0 {

				names  := []string{}
				for _, player := range board.regions[i][j].standOn {
					names = append(names, player.name)
				}
				symbols = strings.Join(names, ",")
				
			}else{
				symbols = strings.Join(board.regions[i][j].symbols, ",")
			}

			fmt.Printf("%15s",symbols)
		}
		
		fmt.Println()
	}

	g.printBoarder()

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

func (g *Game) changeTurn() {
	g.players = append(g.players[1:], g.players[0])
	for g.players[0].win {
		g.players = append(g.players[1:], g.players[0])
	}
}


func (g *Game)printBoarder()	{
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------")
}

func (g *Game) printInformation(player *Player) {
	fmt.Printf("%80s %s\n", "Current Player =", player.name)
	fmt.Printf("%86s\n", "Press Enter to Roll!!!")
	fmt.Scanln()
}

func (g *Game) printRoll(roll int) {
	fmt.Printf("%77s %d\n", "Roll =", roll)
}

func (g *Game) printPlayerPosition(player *Player) {
	fmt.Printf("%70s %s %d\n", player.name, "Position =", player.pos)
}