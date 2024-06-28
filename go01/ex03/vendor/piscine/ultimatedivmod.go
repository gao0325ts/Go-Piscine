package piscine

import (
	_ "ft"
	_ "fmt"
)

func UltimateDivMod(a *int, b *int) {
	n := *a
	*a = *a / *b
	*b = n % *b
}