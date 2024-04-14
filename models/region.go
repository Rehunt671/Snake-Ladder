package models

type Region interface {
	AddStandOn(p Player)
	RemoveStandOn(p Player)
	GetStandOn() []Player
	SetStandOn([]Player)
	GetSymbols() []string
	SetSymbols([]string)
}

type regionImpl struct {
	symbols []string
	standOn []Player
}

func NewRegion(symbols []string) Region {
	return &regionImpl{
		symbols: symbols,
	}
}

func (r *regionImpl) AddStandOn(p Player) {
	r.standOn = append(r.standOn, p)
}

func (r *regionImpl) RemoveStandOn(p Player) {
	for i, player := range r.standOn {
		if player == p {
			r.standOn = append(r.standOn[:i], r.standOn[i+1:]...)
			return
		}
	}
}
func (r *regionImpl) GetStandOn() []Player        { return r.standOn }
func (r *regionImpl) SetStandOn(players []Player) { r.standOn = players }
func (r *regionImpl) GetSymbols() []string        { return r.symbols }
func (r *regionImpl) SetSymbols(symbols []string) { r.symbols = symbols }