package models

type Path interface {
	AddStandOn(p Player)
	RemoveStandOn(p Player)
	GetStandOn() []Player
	SetStandOn([]Player)
	GetSymbols() []string
	SetSymbols([]string)
}

type pathImpl struct {
	symbols []string
	standOn []Player
}

func NewPath(symbols []string) Path {
	return &pathImpl{
		symbols: symbols,
	}
}

func (r *pathImpl) AddStandOn(p Player) {
	r.standOn = append(r.standOn, p)
}

func (r *pathImpl) RemoveStandOn(p Player) {
	for i, player := range r.standOn {
		if player == p {
			r.standOn = append(r.standOn[:i], r.standOn[i+1:]...)
			return
		}
	}
}
func (r *pathImpl) GetStandOn() []Player        { return r.standOn }
func (r *pathImpl) SetStandOn(players []Player) { r.standOn = players }
func (r *pathImpl) GetSymbols() []string        { return r.symbols }
func (r *pathImpl) SetSymbols(symbols []string) { r.symbols = symbols }