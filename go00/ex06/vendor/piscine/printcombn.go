package piscine

import (
	"ft"
)

func PrintCombN(n int) {
	if n <= 0 || n >= 10 { return }

	var PrintCombNRecursive func(prefix []rune, n, start, cur_digit int)

	PrintCombNRecursive = func(prefix []rune, n, start, cur_digit int) {
		if cur_digit == n {
			for i := 0; i < n; i++ {
				ft.PrintRune(prefix[i])
			}
			if prefix[0] != rune((10 - n)+'0') {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
			return
		}

		for i := start; i <= 9; i++ {
			prefix[cur_digit] = rune(i + '0')
			PrintCombNRecursive(prefix, n, i + 1, cur_digit + 1)
		}
	}
	var prefix [10]rune
	PrintCombNRecursive(prefix[:], n, 0, 0)
	ft.PrintRune('\n')
}