package piscine

import (
	"ft"
)

func PrintStr(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		ft.PrintRune(runes[i])
	}
	ft.PrintRune('\n')
}