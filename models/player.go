package models

type Player interface {
	GetName() string
	GetPos() int
	SetPos(int)
	SetIsWin(bool)
	GetIsWin() bool
}

// FINISH:
// 1. Change SetWin function name to SetIsWin
// 2. Change GetWin function name to GetIsWin
type playerImpl struct {
	name string
	pos  int
	// color int
	isWin bool
}

func NewPlayer(name string) Player {
	return &playerImpl{
		name: name,
		pos:  1,
	}
}

func (p *playerImpl) GetPos() int {
	return p.pos
}
func (p *playerImpl) GetName() string {
	return p.name
}

func (p *playerImpl) SetPos(newPos int) {
	p.pos = newPos
}

func (p *playerImpl) SetIsWin(isWin bool) {
	p.isWin = isWin
}

func (p *playerImpl) GetIsWin() bool {
	return p.isWin
}
