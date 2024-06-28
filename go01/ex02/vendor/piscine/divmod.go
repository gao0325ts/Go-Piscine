package piscine

import (
	_ "ft"
	_ "fmt"
)

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}