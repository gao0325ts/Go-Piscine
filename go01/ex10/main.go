package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Atoi("12345"))
	fmt.Println(piscine.Atoi("0000000012345"))
	fmt.Println(piscine.Atoi("012 345")) // 0
	fmt.Println(piscine.Atoi("Hello World!")) // 0
	fmt.Println(piscine.Atoi("+1234")) // 1234
	fmt.Println(piscine.Atoi("-1234")) // -1234
	fmt.Println(piscine.Atoi("++1234")) // 0
	fmt.Println(piscine.Atoi("--1234")) // 0
}