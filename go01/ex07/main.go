package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
	s = "42東京"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
