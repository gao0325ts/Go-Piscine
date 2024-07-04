package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RecursivePower(4, 3))

	// --- Additional Tests ---
	fmt.Println(piscine.RecursivePower(4, 0))
	fmt.Println(piscine.RecursivePower(4, 1))
	fmt.Println(piscine.RecursivePower(4, -1))
}
