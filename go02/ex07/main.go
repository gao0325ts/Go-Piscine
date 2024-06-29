package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(5))
	fmt.Println(piscine.FindNextPrime(4))
	fmt.Println(piscine.FindNextPrime(-42))
	fmt.Println(piscine.FindNextPrime(-1))
	fmt.Println(piscine.FindNextPrime(0))
	fmt.Println(piscine.FindNextPrime(1))
	fmt.Println(piscine.FindNextPrime(2))
	fmt.Println(piscine.FindNextPrime(3))
}
