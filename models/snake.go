package models

type Snake struct {
	start int
	end   int
}

func NewSnake(start int, end int) *Snake {
	return &Snake{
		start: start,
		end:   end,
	}
}

func (s *Snake) GetStart() int {
	return s.start
}
func (s *Snake) GetEnd() int {
	return s.end
}