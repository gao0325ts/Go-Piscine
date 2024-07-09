package main

import (
	"fmt"
	"piscine"
)

func main() {
	result := piscine.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)

	// --- Additional Tests ---
	fmt.Println(piscine.ConvertBase("256", "0123456789", "0123456789abcdef")) // 100
	fmt.Println(piscine.ConvertBase("0", "01", "ABCDE")) // A
}
