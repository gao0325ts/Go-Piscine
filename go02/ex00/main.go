package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg))
	fmt.Println(piscine.IterativeFactorial(-42))
	fmt.Println(piscine.IterativeFactorial(-1))
	fmt.Println(piscine.IterativeFactorial(0))
	fmt.Println(piscine.IterativeFactorial(1))
	fmt.Println(piscine.IterativeFactorial(2))
	fmt.Println(piscine.IterativeFactorial(42))
}