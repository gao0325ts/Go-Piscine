package piscine

import (
	"ft"
)

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 0)
	var start int
	var end int
	for i, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if start == i {
				result = append(result, "")
			} else {
				result = append(result, s[start:end])
			}
			start = i + 1
		}
		end++
	}
	if start <= end {
		result = append(result, s[start:])
	}
	return result
}

func PrintWordsTables(a []string) {
	for _, s := range a {
		printStrln(s)
	}
}

func printStrln(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
