package models

type Player interface{
	GetName() string
	GetPos() int
	SetPos(int) 
	SetWin(bool)
	GetWin() bool
}

type playerImpl struct {
	name string
	pos int
	// color int
	isWin bool
}

func NewPlayer(name string) Player {
	return &playerImpl{
		name : name,
		pos :  1,
	}
}

func (p *playerImpl) GetPos() int {
	return p.pos
}
func (p *playerImpl) GetName() string {
	return p.name
}

func (p *playerImpl) SetPos(newPos int)  {
	 p.pos = newPos
}

func (p *playerImpl) SetWin(isWin bool)  {
	 p.isWin = isWin
}
func (p *playerImpl) GetWin() bool  {
	return p.isWin
}