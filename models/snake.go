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
