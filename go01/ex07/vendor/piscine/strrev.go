package piscine

import (
	_ "ft"
	_ "fmt"
)

func StrRev(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		tmp := runes[i]
		runes[i] = runes[length-1-i]
		runes[length-1-i] = tmp
	}
	return string(runes)
}