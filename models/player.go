package models

type Player interface {
	GetName() string
	GetPos() int
	SetPos(int)
	SetIsWin(bool)
	GetIsWin() bool
}

// TODO: pos => position
type playerImpl struct {
	name  string
	pos   int
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
