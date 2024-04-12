package models

type Region struct {
	symbol string
}

func NewRegion(symbol string) *Region {
	return &Region{
		symbol: symbol,
	}
}