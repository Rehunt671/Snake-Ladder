package models

type Cell interface {
	AddStandOn(p Player)
	RemoveStandOn(p Player)
	GetStandOn() []Player
	SetStandOn([]Player)
	GetSymbols() []string
	SetSymbols([]string)
}

type cellImpl struct {
	symbols []string
	standOn []Player
}

func NewCell(symbols []string) Cell {
	return &cellImpl{
		symbols: symbols,
	}
}

func (c *cellImpl) AddStandOn(p Player) {
	c.standOn = append(c.standOn, p)
}

func (c *cellImpl) RemoveStandOn(p Player) {

	for i, player := range c.standOn {
		if player == p {
			c.standOn = append(c.standOn[:i], c.standOn[i+1:]...)
			return
		}
	}
}
func (c *cellImpl) GetStandOn() []Player        { return c.standOn }
func (c *cellImpl) SetStandOn(players []Player) { c.standOn = players }
func (c *cellImpl) GetSymbols() []string        { return c.symbols }
func (c *cellImpl) SetSymbols(symbols []string) { c.symbols = symbols }
