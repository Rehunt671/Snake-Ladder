package models

type Ladder struct {
	start int
	end   int
}

func NewLadder(start int, end int) *Ladder {
	return &Ladder{
		start: start,
		end:   end,
	}
}
