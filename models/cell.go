package models

type Cell interface {
	AddPlayerStandingOn(p Player)
	RemovePlayerStandingOn(p Player)
	RecievePlayersStandingOn() []Player
	SetPlayersStandingOn([]Player)
	GetSymbols() []string
	SetSymbols([]string)
}

// TODO:
// 1.Change GetStandOn function name to RecievePlayersStandingOn
// 2.Change SetStandOn function name to SetPlayersStandingOn
// 3.Change RemoveStandOn function name to RemovePlayerStandingOn
// 4.Change AddStandOn function name to AddPlayerStandingOn
type cellImpl struct {
	symbols           []string
	playersStandingOn []Player
}

func NewCell(symbols []string) Cell {
	return &cellImpl{
		symbols: symbols,
	}
}

func (c *cellImpl) AddPlayerStandingOn(p Player) {
	c.playersStandingOn = append(c.playersStandingOn, p)
}

func (c *cellImpl) RemovePlayerStandingOn(p Player) {

	for i, player := range c.playersStandingOn {
		if player == p {
			c.playersStandingOn = append(c.playersStandingOn[:i], c.playersStandingOn[i+1:]...)
			return
		}
	}
}
func (c *cellImpl) RecievePlayersStandingOn() []Player    { return c.playersStandingOn }
func (c *cellImpl) SetPlayersStandingOn(players []Player) { c.playersStandingOn = players }
func (c *cellImpl) GetSymbols() []string                  { return c.symbols }
func (c *cellImpl) SetSymbols(symbols []string)           { c.symbols = symbols }
