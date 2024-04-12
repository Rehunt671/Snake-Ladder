package utils

func NumToRowCol(pos int, size int) (int, int) {
	i := (pos - 1) / size
	j := (pos - 1) % size
	if i%2 != 0 {
		j = size - j - 1
	}
	return i, j
}