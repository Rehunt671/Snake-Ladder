package models

type Ladder interface {
	GetStart() int
	GetEnd() int
}

type ladderImpl struct {
	start int
	end   int
}

func NewLadder(start int, end int) Ladder {
	return &ladderImpl{
		start: start,
		end:   end,
	}
}

func (l *ladderImpl) GetStart() int {
	return l.start
}
func (l *ladderImpl) GetEnd() int {
	return l.end
}