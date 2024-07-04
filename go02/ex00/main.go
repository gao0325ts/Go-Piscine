package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg))

	// --- Additional Tests ---
	fmt.Println(piscine.IterativeFactorial(-42)) // non possible value
	fmt.Println(piscine.IterativeFactorial(-1))  // non possible value
	fmt.Println(piscine.IterativeFactorial(0))
	fmt.Println(piscine.IterativeFactorial(1))
	fmt.Println(piscine.IterativeFactorial(2))
	fmt.Println(piscine.IterativeFactorial(20))
	fmt.Println(piscine.IterativeFactorial(42)) // overflow
}
