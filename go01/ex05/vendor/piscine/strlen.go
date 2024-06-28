package piscine

import (
	_ "ft"
	_ "fmt"
)

func StrLen(s string) int {
	runes := []rune(s)
	return len(runes)
}