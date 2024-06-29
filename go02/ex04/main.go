package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
	fmt.Println(piscine.Fibonacci(-42))
	fmt.Println(piscine.Fibonacci(-1))
	fmt.Println(piscine.Fibonacci(0))
	fmt.Println(piscine.Fibonacci(1))
	fmt.Println(piscine.Fibonacci(2))
	fmt.Println(piscine.Fibonacci(3))
}
