package piscine

import (
	"ft"
	"os"
)

func RevParams() {
	params := os.Args
	var count int
	for range params {
		count++
	}
	for i := count - 1; i > 0; i-- {
		printStrln(params[i])
	}
}

func printStrln(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
