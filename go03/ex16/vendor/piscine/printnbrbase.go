package piscine

import "ft"

func strLen(s string) int {
	var length int
	for range s {
		length++
	}
	return length
}

func isValidBase(base string) bool {
	if strLen(base) < 2 {
		return false
	}
	var extChars string
	for _, r := range base {
		if r == '-' || r == '+' {
			return false
		}
		for _, c := range extChars {
			if r == c {
				return false
			}
		}
		extChars += string(r)
	}
	return true
}

func printResult(n, len int, base string) {
	if n >= len {
		printResult(n/len, len, base)
	}
	ft.PrintRune(rune(base[n%len]))
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	baseLen := strLen(base)
	if nbr < 0 {
		ft.PrintRune('-')
		nbr = -nbr
	}
	if nbr == 0 {
		ft.PrintRune(rune(base[0]))
		return
	}
	printResult(nbr, baseLen, base)
}
