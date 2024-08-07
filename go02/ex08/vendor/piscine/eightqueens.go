package piscine

import (
	"ft"
)

func isValid(pos []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if pos[i] == col || pos[i]-col == row-i || col-pos[i] == row-i {
			return false
		}
	}
	return true
}

func findPatterns(pos []int, row int) {
	if row == 8 {
		for i := 0; i < 8; i++ {
			ft.PrintRune(rune(pos[i] + '1'))
		}
		ft.PrintRune('\n')
		return
	}
	for col := 0; col < 8; col++ {
		if isValid(pos, row, col) {
			pos[row] = col
			findPatterns(pos, row+1)
		}
	}
}

func EightQueens() {
	var positions [8]int
	findPatterns(positions[:], 0)
}
