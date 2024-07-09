package piscine

import (
	"ft"
	"os"
)

func PrintProgramName() {
	name := os.Args[0]
	for i, r := range name {
		// Skip the initial "./"
		if i <= 1 {
			continue
		}
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
