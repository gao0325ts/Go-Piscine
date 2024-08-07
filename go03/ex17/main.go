package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))
	fmt.Println(piscine.AtoiBase("1111101", "01"))
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))

	// --- Additional Tests ---
	fmt.Println(piscine.AtoiBase("7", "7"))
	fmt.Println(piscine.AtoiBase("7D", "+0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("7D", "-0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("7D", "00123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("0", "00123456789ABCDEF"))
}
