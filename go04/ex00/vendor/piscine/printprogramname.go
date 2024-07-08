package piscine

import "ft"

func PrintProgramName(name string) {
	for _, r := range name {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}