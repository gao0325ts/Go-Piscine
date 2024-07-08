package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))

	// --- Additional Tests ---
	fmt.Println(piscine.IsPrintable(""))
}
