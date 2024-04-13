package models

type Snake interface {
	GetStart() int
	GetEnd() int
}

type snakeImpl struct {
	start int
	end   int
}

func NewSnake(start int, end int) Snake {
	return &snakeImpl{
		start: start,
		end:   end,
	}
}

func (s *snakeImpl) GetStart() int {
	return s.start
}
func (s *snakeImpl) GetEnd() int {
	return s.end
}