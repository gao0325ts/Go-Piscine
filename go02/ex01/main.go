package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))

	// --- Additional Tests ---
	fmt.Println(piscine.RecursiveFactorial(-42)) // non possible value
	fmt.Println(piscine.RecursiveFactorial(-1))  // non possible value
	fmt.Println(piscine.RecursiveFactorial(0))
	fmt.Println(piscine.RecursiveFactorial(1))
	fmt.Println(piscine.RecursiveFactorial(2))
	fmt.Println(piscine.RecursiveFactorial(20))
	fmt.Println(piscine.RecursiveFactorial(42)) // overflow
}
