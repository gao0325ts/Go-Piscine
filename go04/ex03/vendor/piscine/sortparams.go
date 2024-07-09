package piscine

import "ft"

func compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func SortParams(params []string) {
	var sorted bool
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
	for i, p := range params {
		if i == 0 {
			continue
		}
		printStr(p)
	}
}

func printStr(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}