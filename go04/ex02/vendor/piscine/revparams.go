package piscine

import "ft"

func RevParams(params []string) {
	var count int
	for range params {
		count++
	}
	for i := count - 1; i > 0; i-- {
		printStr(params[i])
	}
}

func printStr(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
