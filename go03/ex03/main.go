package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))

	// --- Additional Tests ---
	fmt.Println(piscine.Index("nihao", "o"))
	fmt.Println(piscine.Index("Bonjour", "jou"))
}
