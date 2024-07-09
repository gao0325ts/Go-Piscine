package piscine

import (
	"ft"
	"os"
)

func compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func printStrln(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func SortParams() {
	params := os.Args
	var sorted bool

	// Sort params in ASCII order
	for !sorted {
		sorted = true
		for i := range params {
			if i <= 1 {
				continue
			}
			if compare(params[i-1], params[i]) > 0 {
				params[i-1], params[i] = params[i], params[i-1]
				sorted = false
			}
		}
	}

	// Print sorted params, skipping the program name
	for i, p := range params {
		if i == 0 {
			continue
		}
		printStrln(p)
	}
}
