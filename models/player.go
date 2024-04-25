package models

type Player interface {
	GetName() string
	GetPosition() int
	SetPosition(int)
	SetIsWin(bool)
	GetIsWin() bool
}

type playerImpl struct {
	name     string
	position int
	isWin    bool
}

func NewPlayer(name string) Player {
	return &playerImpl{
		name:     name,
		position: 1,
	}
}

func (p *playerImpl) GetPosition() int {
	return p.position
}
func (p *playerImpl) GetName() string {
	return p.name
}

func (p *playerImpl) SetPosition(newPosition int) {
	p.position = newPosition
}

func (p *playerImpl) SetIsWin(isWin bool) {
	p.isWin = isWin
}

func (p *playerImpl) GetIsWin() bool {
	return p.isWin
}
