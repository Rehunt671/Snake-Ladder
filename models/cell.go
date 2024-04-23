package models

type Cell interface {
	AddPlayer(p Player)
	RemovePlayer(p Player)
	GetPlayers() []Player
	SetPlayers([]Player)
	HasPlayer() bool
	GetSymbols() []string
	SetSymbols([]string)
}

// TODO:
// 1.Change GetStandOn function name to GetPlayers
// 2.Change SetStandOn function name to SetPlayers
// 3.Change RemoveStandOn function name to RemovePlayer
// 4.Change AddStandOn function name to AddPlayer
type cellImpl struct {
	symbols           []string
	playersStandingOn []Player
}

func NewCell(symbols []string) Cell {
	return &cellImpl{
		symbols: symbols,
	}
}

func (c *cellImpl) AddPlayer(p Player) {
	c.playersStandingOn = append(c.playersStandingOn, p)
}

func (c *cellImpl) RemovePlayer(p Player) {

	for i, player := range c.playersStandingOn {
		if player == p {
			c.playersStandingOn = append(c.playersStandingOn[:i], c.playersStandingOn[i+1:]...)
			return
		}
	}
}

func (c *cellImpl) GetPlayers() []Player {
	return c.playersStandingOn
}

func (c *cellImpl) SetPlayers(players []Player) {
	c.playersStandingOn = players
}

func (c *cellImpl) GetSymbols() []string {
	return c.symbols
}

func (c *cellImpl) SetSymbols(symbols []string) {
	c.symbols = symbols
}

func (c *cellImpl) HasPlayer() bool {
	return len(c.playersStandingOn) != 0
}
