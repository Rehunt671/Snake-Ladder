package models

type Region struct {
	symbols []string
	standOn []*Player
}

func NewRegion(symbols []string) *Region {
	return &Region{
		symbols: symbols,
	}
}

func (r *Region) AddStandOn(p *Player) {
	r.standOn = append(r.standOn, p)
}

func (r *Region) RemoveStandOn(p *Player) {
	for i, player := range r.standOn {
		if player == p {
			r.standOn = append(r.standOn[:i], r.standOn[i+1:]...)
			return
		}
	}
}