package piscine

import (
	"ft"
	"os"
)

func PrintParams() {
	params := os.Args
	for i, p := range params {
		// Skip the program name
		if i == 0 {
			continue
		}
		printStrln(p)
	}
}

func printStrln(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
