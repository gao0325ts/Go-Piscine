package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
	l = piscine.StrLen("こんにちは")
	fmt.Println(l)
}
