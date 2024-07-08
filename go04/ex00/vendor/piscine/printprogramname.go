package piscine

import "ft"

func PrintProgramName(name string) {
	for i, r := range name {
		if i <= 1 {
			continue
		}
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}