package piscine

import (
	"ft"
)

func PrintStr(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
