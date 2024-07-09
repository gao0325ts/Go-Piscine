package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.MakeRange(5, 10))
	fmt.Println(piscine.MakeRange(10, 5))

	// --- Additional Tests ---
	fmt.Println(piscine.MakeRange(5, 5))
	fmt.Println(piscine.MakeRange(5, 6))
}
