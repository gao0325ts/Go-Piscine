package piscine

import "ft"

func PrintParams(params []string) {
	for i, p := range params {
		if i >= 1 {
			printStr(p)
		}
	}
}

func printStr(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
