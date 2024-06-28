package piscine

import (
	_ "ft"
	_ "fmt"
)

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}