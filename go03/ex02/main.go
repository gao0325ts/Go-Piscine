package main

import (
	"ft"
	"piscine"
)

func main() {
	ft.PrintRune(piscine.LastRune("Hello!"))
	ft.PrintRune(piscine.LastRune("Salut!"))
	ft.PrintRune(piscine.LastRune("Ola!"))

	// --- Additional Tests ---
	ft.PrintRune(piscine.LastRune(""))
	ft.PrintRune('\n')
}
