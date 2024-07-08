package piscine

import "ft"

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 0)
	var tmp string
	var i int
	for _, r := range s {
		if r == ' ' {
			result = append(result, tmp)
			i++
			tmp = ""
			continue
		}
		tmp += string(r)
	}
	result = append(result, tmp)
	return result
}

func PrintWordsTables(a []string) {
	for _, s := range a {
		printStr(s)
	}
}

func printStr(s string) {
	for _, r := range []rune(s) {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
