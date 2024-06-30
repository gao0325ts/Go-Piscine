package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)

	// --- Additional Tests ---
	l = piscine.StrLen("こんにちは")
	fmt.Println(l)
}
